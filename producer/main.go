package main

import (
	"context"
	"strconv"

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
	for i := 0; i < 10; i++ {
		message := amqp091.Publishing{
			Headers: amqp091.Table{
				"sample": "value",
			},
			Body: []byte("Hello [" + strconv.Itoa(i) + "]"),
		}
		err := channel.PublishWithContext(ctx, "notification", "email", false, false, message)
		if err != nil {
			panic(err)
		}
	}
}