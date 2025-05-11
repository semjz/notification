package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"notification/ent"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	logger *log.Logger
}

func NewRabbitMQ() *RabbitMQ {
	logFile, err := os.OpenFile("notification.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Could not open log file %v", err)
	}
	logger := log.New(logFile, "RABBITMQ ERROR: ", log.LstdFlags|log.Lshortfile)
	return &RabbitMQ{logger: logger}
}

func (r *RabbitMQ) logError(context string, err error) bool {
	if err != nil {
		r.logger.Printf("%s: %s", context, err)
		return true
	}
	return false
}

func (r *RabbitMQ) Process(message *ent.Message) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if r.logError("Failed to connect to RabbitMQ", err) {
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	defer ch.Close()

	if r.logError("Failed to open a channel", err) {
		return
	}

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,
	)

	if r.logError("Failed to declare a queue", err) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	body, err := json.Marshal(message.Payload)

	if r.logError("Failed to marshal payload", err) {
		return
	}

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
			Priority:     0, // initial priority
			Headers:      amqp.Table{"id": message.ID.String()},
		})

	if r.logError("Failed to publish a message", err) {
		return
	}

	log.Printf(" [x] Sent %s", message.Payload)
}
