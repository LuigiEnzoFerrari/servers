package consumer

import (
	"encoding/json"
	"log"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/nats-io/nats.go"
)

type OtpConsumer struct {
	nc *nats.Conn
	otpService domain.OptService	
}

func NewOtpConsumer(nc *nats.Conn, otpService domain.OptService) *OtpConsumer {
	return &OtpConsumer{
		nc: nc,
		otpService: otpService,
	}
}

func (c *OtpConsumer) Start() {
	_, err := c.nc.QueueSubscribe(
		"delivery.otp",
		"otp-group",
		c.passwordForgotEventHandler,
	)

    if err != nil {
        log.Fatal(err)
    }
}

func (c *OtpConsumer) passwordForgotEventHandler(msg *nats.Msg) {
	data := msg.Data

	passwordForgotEvent := domain.PasswordForgotEvent{}
	if err := json.Unmarshal(data, &passwordForgotEvent); err != nil {
		log.Println(err)
		return
	}

	_, err := c.otpService.GenerateOTP(passwordForgotEvent.Email)

	if err != nil {
		log.Println(err)
		return
	}
	msg.Ack()
}
