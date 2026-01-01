package consumer

import (
	"context"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

)

type Handler func(ctx context.Context, body []byte) error

type ConsumerConfig struct {
	QueueName   string
	WorkerCount int
	Handler     Handler
}

type RabbitMQConsumer struct {
	conn *amqp.Connection
	configs []ConsumerConfig
	wg      *sync.WaitGroup
}

func NewRabbitMQConsumer(conn *amqp.Connection, configs []ConsumerConfig) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		conn: conn,
		configs: configs,
		wg: &sync.WaitGroup{},
	}
}

func (m *RabbitMQConsumer) Start(ctx context.Context) {
	for _, cfg := range m.configs {
		m.wg.Add(1)
		go m.startConsumer(ctx, cfg)
	}
}

func (m *RabbitMQConsumer) startConsumer(ctx context.Context, cfg ConsumerConfig) {
	defer m.wg.Done()
	ch, err := m.conn.Channel()
	if err != nil {
		log.Printf("Failed to open channel for %s: %v", cfg.QueueName, err)
		return
	}
	defer ch.Close()

	err = ch.Qos(cfg.WorkerCount, 0, false)
	if err != nil {
		log.Printf("Failed to set QoS for %s: %v", cfg.QueueName, err)
		return
	}

	msgs, err := ch.Consume(
		cfg.QueueName, "", false, false, false, false, nil,
	)
	if err != nil {
		log.Printf("Failed to register consumer for %s: %v", cfg.QueueName, err)
		return
	}

	log.Printf("Started consumer for queue: %s with %d workers", cfg.QueueName, cfg.WorkerCount)

	sem := make(chan struct{}, cfg.WorkerCount)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Stopping consumer for %s...", cfg.QueueName)
			return
		case d, ok := <-msgs:
			if !ok {
				return
			}
			log.Printf("Received message for %s: %s", cfg.QueueName, string(d.Body))

			sem <- struct{}{}
			
			m.wg.Add(1)
			go func(delivery amqp.Delivery) {
				defer m.wg.Done()
				defer func() { <-sem }()

				err := cfg.Handler(ctx, delivery.Body)

				if err != nil {
					log.Printf("Error processing %s: %v", cfg.QueueName, err)
					delivery.Nack(false, false)
				} else {
					delivery.Ack(false)
				}
			}(d)
		}
	}
}