package main

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	
	ctx := context.Background()
	emailConsumer, err := channel.ConsumeWithContext(ctx,
		"email", "consumer-email",
		true, false, false, false, nil,
	)
	if err != nil {
		panic(err)
	}

	for message := range emailConsumer {
		fmt.Println("Routing key:", message.RoutingKey)
		fmt.Println("Body:", string(message.Body))
	}
}