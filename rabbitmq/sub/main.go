package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 1. Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://admin:secret@localhost:5672/shop-prod")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "notifications.send-receipt"
	// Make sure these boolean flags match your existing queue configuration (Durable, AutoDelete, etc).

	// 3. Register a consumer

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer tag (empty lets library generate one)
		true,      // auto-ack (true means message is removed from queue immediately upon delivery)
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 4. Read from the channel
	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// If auto-ack was false, you would call d.Ack(false) here
		}
	}()

	log.Printf(" [*] Waiting for messages on queue: %s. To exit press CTRL+C", queueName)
	<-forever
}
