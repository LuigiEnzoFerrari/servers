package publish

import (
	"context"
	"encoding/json"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
	"github.com/nats-io/nats.go"
)

type AuthPublish struct {
	nc *nats.Conn
}

func NewAuthPublish(nc *nats.Conn) *AuthPublish {
	return &AuthPublish{nc: nc}
}

func (p *AuthPublish) Publish(ctx context.Context, event domain.PasswordForgotEvent) error {

	data, err := json.Marshal(event)
	if err != nil {
		log.Fatal("Error marshaling event: ", err)
	}

	return p.nc.Publish("auths.password_forgot", data)
}
