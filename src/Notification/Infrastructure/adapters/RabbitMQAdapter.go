package adapters

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
}

type NotificationMessage struct {
	ClientID   int    `json:"client_id"`
	ClientName string `json:"client_name"`
	Content    string `json:"notification_content"`
}

func InitRabbitMQ() *RabbitMQAdapter {
	conn, err := amqp091.Dial("amqp://uriel:eduardo117@3.228.81.226:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		"paymanent", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}
}

func (r *RabbitMQAdapter) PublishMessage(message string) error {
	err := r.ch.PublishWithContext(
		context.Background(),
		"",          // exchange
		"paymanent", // routing key (queue name)
		false,       // mandatory
		false,       // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	log.Printf("Sent message: %s", message)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	r.ch.Close()
	r.conn.Close()
}
