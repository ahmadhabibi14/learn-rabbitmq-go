package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://habi:habi1234@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	fmt.Println(`connected to RabbitMQ`)
	
	// ctx := context.Background()
	// emailConsumer, err := channel.ConsumeWithContext(ctx,
	// 	"email", "consumer-email",
	// 	true, false, false, false, nil,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// for message := range emailConsumer {
	// 	fmt.Println("Routing key:", message.RoutingKey)
	// 	fmt.Println("Body:", string(message.Body))
	// }
}