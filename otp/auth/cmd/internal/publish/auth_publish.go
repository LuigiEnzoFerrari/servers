package publish

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type AuthPublish struct {
	
}

func NewAuthPublish() *AuthPublish {
	return &AuthPublish{}
}

func (p *AuthPublish) Publish() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, _ := jetstream.New(nc)
	ctx := context.Background()

	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.>"},
		Storage:  jetstream.FileStorage,
	})
	if err != nil {
		log.Fatal("Stream init error: ", err)
	}

	for i := 1; i <= 5; i++ {
		orderID := fmt.Sprintf("ORD-%d", i)
		
		_, err := js.Publish(ctx, "orders.created", []byte(orderID))
		if err != nil {
			log.Printf("Failed to publish %s: %v", orderID, err)
		} else {
			log.Printf("Sent: %s", orderID)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
