package publish

import (
	"encoding/json"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublish struct {
	conn *amqp.Connection
}

func NewRabbitMQPublish(conn *amqp.Connection) *RabbitMQPublish {
	return &RabbitMQPublish{conn: conn}
}

func (p *RabbitMQPublish) Publish(routingKey string, event interface{}) error {

	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	body, err := json.Marshal(event); if err != nil {	
		return err
	}

	return ch.Publish(
		"auth",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (p *RabbitMQPublish) PublishPasswordForgotEvent(event domain.PasswordForgotEvent) error {
	log.Println("Publishing password forgot event")
	return p.Publish("passwordforgot", event)
}
