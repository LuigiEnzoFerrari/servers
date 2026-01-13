terraform {
  required_providers {
    rabbitmq = {
      source = "cyrilgdn/rabbitmq"
      version = "1.10.1"
    }
  }
}

variable "rabbitmq_user" {
  type      = string
  sensitive = true
}

variable "rabbitmq_password" {
  type      = string
  sensitive = true
}

provider "rabbitmq" {
  endpoint = "http://rabbitmq:15672"
  username = var.rabbitmq_user
  password = var.rabbitmq_password
}

resource "rabbitmq_vhost" "dev" {
  name = "dev"
}

resource "rabbitmq_exchange" "auth" {
  name  = "auth"
  vhost = rabbitmq_vhost.dev.name

  settings {
    type        = "topic"
    durable     = true
    auto_delete = false
  }
}

resource "rabbitmq_queue" "otp_passwordforgot" {
  name  = "otp.passwordforgot"
  vhost = rabbitmq_vhost.dev.name

  settings {
    durable     = true
    auto_delete = false
    
    arguments = {
      "x-queue-type" = "quorum"
    }
  }
}

resource "rabbitmq_binding" "auth_otp_binding" {
  source           = rabbitmq_exchange.auth.name
  vhost            = rabbitmq_vhost.dev.name
  destination      = rabbitmq_queue.otp_passwordforgot.name
  destination_type = "queue"
  routing_key      = "passwordforgot"
}