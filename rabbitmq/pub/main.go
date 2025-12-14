package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://admin:secret@localhost:5672/shop-prod")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	fmt.Println("Hello, Go!")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	exchangeName := "orders"
	routingKey := "order.created"
	body := "Pera com banana"
	err = ch.PublishWithContext(ctx,
		exchangeName, // Exchange (Empty string = default exchange)
		routingKey,   // Routing Key (Target queue name)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			// Critical for Quorum Queues: Persist the message to disk/Raft log
			DeliveryMode: amqp.Persistent,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s to %s", body, routingKey)
}
