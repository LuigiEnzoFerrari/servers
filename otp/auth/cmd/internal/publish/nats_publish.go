package publish

import (
	"context"
	"encoding/json"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
)

type AuthPublish struct {
	
}

func NewAuthPublish() *AuthPublish {
	return &AuthPublish{}
}

func (p *AuthPublish) Publish(ctx context.Context, event domain.PasswordForgotEvent) error {

	data, err := json.Marshal(event)
	if err != nil {
		log.Fatal("Error marshaling event: ", err)
	}

	return p.nc.Publish("auths.password_forgot", data)
}
