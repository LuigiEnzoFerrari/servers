package consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"sync"
	"time"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventHandler interface {
	SendOTPEmail(ctx context.Context, event domain.Event) error
}

type RabbitMQConsumer struct {
	ch      *amqp.Channel
	handler EventHandler
	workers int
}

func NewRabbitMQConsumer(ch *amqp.Channel, handler EventHandler, workers int) (*RabbitMQConsumer, error) {
	if err := ch.Qos(workers*2, 0, false); err != nil {
		return nil, err
	}
	
	return &RabbitMQConsumer{
		ch:      ch,
		handler: handler,
		workers: workers,
	}, nil
}

func (c *RabbitMQConsumer) Start(ctx context.Context, queueName string) error {
	msgs, err := c.ch.Consume(
		queueName,
		"my-consumer-id",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for i := 0; i < c.workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			c.worker(ctx, msgs, workerID)
		}(i)
	}

	<-ctx.Done()
	wg.Wait()

	return nil
}

func (c *RabbitMQConsumer) worker(ctx context.Context, msgs <-chan amqp.Delivery, id int) {
	logger, _ := ctx.Value("logger").(*slog.Logger)
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			var event domain.Event
			
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				msg.Nack(false, false)
				continue
			}
			
			rabbitMQLogger := logger.With(
				slog.String("trace_id", event.TraceID),
				slog.Int("worker_id", id),
				slog.String("event", event.Type),
			)

			ctx = context.WithValue(ctx, "logger", rabbitMQLogger)
			rabbitMQLogger.Info("Processing event", "type", event.Type)

			msgCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

			err := c.handler.SendOTPEmail(msgCtx, event)
			cancel()

			if err != nil {
				rabbitMQLogger.Error("Error processing event", "error", err)
				msg.Nack(false, false)
			} else {
				msg.Ack(false)
			}
		}
	}
}