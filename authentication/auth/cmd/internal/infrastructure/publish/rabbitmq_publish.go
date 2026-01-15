package publish

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublish struct {
	ch *amqp.Channel
}

func NewRabbitMQPublish(ch *amqp.Channel) (*RabbitMQPublish, error) {

	if err := ch.Confirm(false); err != nil {
		return nil, fmt.Errorf("failed to put channel in confirm mode: %w", err)
	}
	return &RabbitMQPublish{ch: ch}, nil
}

func (p *RabbitMQPublish) publish(ctx context.Context, exchange string, routingKey string, event domain.Event) error {

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	confirmation, err := p.ch.PublishWithDeferredConfirmWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	ok := confirmation.Wait()
	if !ok {
		return fmt.Errorf("message was not acknowledged by RabbitMQ")
	}

	return nil
}

func (p *RabbitMQPublish) PasswordForgotEvent(ctx context.Context, event domain.Event) error {

	return p.publish(ctx, "auth", "passwordforgot", event)
}
