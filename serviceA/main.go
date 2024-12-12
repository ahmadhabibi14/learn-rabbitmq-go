package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Response struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
}

func handleRequest(ch *amqp.Channel, msg amqp.Delivery) {
	// Proses request
	resp := &Response{
		Success: true,
		Error: ``,
	}

	response, error := json.Marshal(resp)
	if error != nil {
		panic(error)
	}

	// Kirim respons kembali ke client
	err := ch.Publish(
		"",          // Exchange
		msg.ReplyTo, // Reply-to queue
		false,       // Mandatory
		false,       // Immediate
		amqp.Publishing{
			CorrelationId: msg.CorrelationId,
			ContentType:   "application/json",
			Body:          []byte(response),
		},
	)
	if err != nil {
		log.Println(`failed to publish a message`, err)
	}
}

func main() {
	// Koneksi ke RabbitMQ
	conn, err := amqp.Dial("amqp://express:express1234@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	} else {
		log.Println(`connected to RabbitMQ`)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Mendeklarasikan antrean untuk menerima request
	queue, err := ch.QueueDeclare(
		"auth.user.login", // Nama antrean
		false,       // Durable
		false,       // Auto-delete
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)
	if err != nil {
		log.Println(`failed to declare a queue`, err)
	}

	// Mendengarkan antrean
	msgs, err := ch.Consume(
		queue.Name, // Nama queue
		"",         // Consumer tag
		true,       // Auto-ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		log.Println(`failed to register a consumer`, err)
	}

	// Proses request yang diterima
	for msg := range msgs {
		go handleRequest(ch, msg) // Proses permintaan di goroutine untuk menangani banyak request
	}
}
