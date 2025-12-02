package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, _ := jetstream.New(nc)
	ctx := context.Background()

	consumer, err := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Durable:   "WarehouseWorker", 
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal("Consumer init error: ", err)
	}

	// 3. Start Processing
	log.Println("🏭 Warehouse Worker started. Waiting for orders...")
	
	cons, err := consumer.Consume(func(msg jetstream.Msg) {
		log.Printf("📦 Processing Order: %s\n", string(msg.Data()))
	})
	if err != nil {
		log.Fatal("Consumer init error: ", err)
	}
	defer cons.Stop()
	
	r := gin.Default()
	r.Run(":8081")
}
