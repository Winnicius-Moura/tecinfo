package main

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type GalleryCardMessage struct {
	ContributorID string  `json:"contributor_id"`
	HtmlContent   string  `json:"html_content"`
	ApprovedAt    string  `json:"approved_at"`
	Percentage    float64 `json:"percentage"`
}

func startRabbitMQConsumer(hub *Hub) {
	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"approved_submissions_ex", // name
		"fanout",                  // type
		true,                      // durable
		false,                     // auto-deleted
		false,                     // internal
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	q, err := ch.QueueDeclare(
		"",    // name - let rabbitmq generate a random one
		false, // durable
		true,  // delete when unused (important for ephemeral workers)
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"approved_submissions_ex", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf(" [*] Waiting for approved submissions. To exit press CTRL+C")

	for d := range msgs {
		var card GalleryCardMessage
		if err := json.Unmarshal(d.Body, &card); err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		log.Printf("Received new approved submission from: %s", card.ContributorID)
		
		// Broadcast the actual message to connected websocket clients
		hub.broadcast <- d.Body
	}
}
