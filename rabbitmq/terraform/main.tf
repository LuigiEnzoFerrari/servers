terraform {
  required_providers {
    rabbitmq = {
      source = "cyrilgdn/rabbitmq"
      version = "1.8.0"
    }
  }
}

variable "rabbitmq_endpoint" {
  type    = string
  default = "http://localhost:15672"
}

provider "rabbitmq" {
  endpoint = var.rabbitmq_endpoint
  username = "admin"
  password = "secret"
}

# ==============================================================================
# 1. VHOST (Best Practice: Isolate environments/apps)
# ==============================================================================
resource "rabbitmq_vhost" "shop_prod" {
  name = "shop-prod"
}

resource "rabbitmq_user" "guest" {
  name     = "guest"
  password = "guest"
  tags     = []
}

resource "rabbitmq_permissions" "guest_shop" {
  user  = rabbitmq_user.guest.name
  vhost = rabbitmq_vhost.shop_prod.name
  permissions {
    configure = ".*"
    write     = ".*"
    read      = ".*"
  }
}

# ==============================================================================
# SCENARIO 1: E-COMMERCE (Topic Exchange)
# Pattern: Source.Entity.Action -> Exchange -> Queue
# ==============================================================================

# 1. The Exchange (Source of Truth)
resource "rabbitmq_exchange" "orders" {
  name  = "orders"
  vhost = rabbitmq_vhost.shop_prod.name

  settings {
    type    = "topic"
    durable = true
  }
}

# 2. The Queues (Consumers)
# Naming: Service.WorkType

resource "rabbitmq_queue" "shipping_label" {
  name  = "shipping.create-label"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    durable = true
  }
}

resource "rabbitmq_queue" "notify_receipt" {
  name  = "notifications.send-receipt"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    durable = true
  }
}

# 3. The Bindings (Routing Logic)

resource "rabbitmq_binding" "shipping_bind" {
  source           = rabbitmq_exchange.orders.name
  vhost            = rabbitmq_vhost.shop_prod.name
  destination      = rabbitmq_queue.shipping_label.name
  destination_type = "queue"
  routing_key      = "order.created"
}

resource "rabbitmq_binding" "notify_bind" {
  source           = rabbitmq_exchange.orders.name
  vhost            = rabbitmq_vhost.shop_prod.name
  destination      = rabbitmq_queue.notify_receipt.name
  destination_type = "queue"
  # Notifications wants created AND updated orders
  # Matches: order.created, order.updated
  routing_key      = "order.*" 
}


# ==============================================================================
# SCENARIO 2: IOT TELEMETRY (Deep Wildcards)
# ==============================================================================

resource "rabbitmq_exchange" "telemetry" {
  name  = "telemetry"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    type    = "topic"
    durable = true
  }
}

resource "rabbitmq_queue" "dashboard_stream" {
  name  = "dashboard.stream-all"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    durable     = false     # Dashboards might be transient
    auto_delete = true      # Clean up when consumer disconnects
  }
}

resource "rabbitmq_queue" "alert_critical" {
  name  = "alerts.trigger-pager"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    durable = true
  }
}

resource "rabbitmq_binding" "dashboard_bind" {
  source           = rabbitmq_exchange.telemetry.name
  vhost            = rabbitmq_vhost.shop_prod.name
  destination      = rabbitmq_queue.dashboard_stream.name
  destination_type = "queue"
  # Grab everything: sensor.temp.normal, sensor.humidity.critical
  routing_key      = "sensor.#"
}

resource "rabbitmq_binding" "alert_bind" {
  source           = rabbitmq_exchange.telemetry.name
  vhost            = rabbitmq_vhost.shop_prod.name
  destination      = rabbitmq_queue.alert_critical.name
  destination_type = "queue"
  # Grab only criticals from ANY sensor
  routing_key      = "*.critical"
}


# ==============================================================================
# SCENARIO 3: DEAD LETTER QUEUE (Advanced Reliability)
# Naming: OriginalQueue.dlq
# ==============================================================================

# 1. The DLQ itself (Where failed messages go)
resource "rabbitmq_queue" "payments_dlq" {
  name  = "payments.charge-card.dlq"
  vhost = rabbitmq_vhost.shop_prod.name
  settings {
    durable = true
  }
}

# 2. The Main Queue (With DLX Arguments)
resource "rabbitmq_queue" "payments_work" {
  name  = "payments.charge-card"
  vhost = rabbitmq_vhost.shop_prod.name
  
  settings {
    durable = true
    # "arguments_json" is preferred for passing complex x-args
    arguments_json = <<EOF
    {
      "x-dead-letter-exchange": "",
      "x-dead-letter-routing-key": "${rabbitmq_queue.payments_dlq.name}"
    }
    EOF
  }
}

# Note: We mapped DLX to "" (default exchange) so we can route directly 
# to the queue name. In complex setups, you might map to a "dlx" exchange.


# ==============================================================================
# SCENARIO 4: POLICIES (Global Rules)
# ==============================================================================

# Example: Enforce TTL (Time To Live) on all temporary queues
resource "rabbitmq_policy" "temp_cleanup" {
  name  = "enforce-ttl-temp"
  vhost = rabbitmq_vhost.shop_prod.name

  policy {
    pattern  = "^temp\\."   # Regex matching queues starting with "temp."
    priority = 10
    apply_to = "queues"
    definition = {
      "message-ttl" = 60000 # 60 seconds
      "expires"     = 120000 # Queue deleted after 2 mins idle
    }
  }
}