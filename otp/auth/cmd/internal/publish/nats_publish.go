package publish

import (
	"context"
	"encoding/json"

	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
)

type AuthPublish struct {
}

func NewAuthPublish() *AuthPublish {
	return &AuthPublish{}
}

func (p *AuthPublish) Publish(ctx context.Context, event domain.PasswordForgotEvent) error {

	_, err := json.Marshal(event)
	if err != nil {	
		return err
	}

	return nil
}
