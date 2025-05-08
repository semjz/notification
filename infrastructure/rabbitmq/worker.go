package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"notification/ent"
	"notification/pkg"
	"time"
)

func (r *RabbitMQ) Receiver(client *ent.Client) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if r.logError("Failed to connect to RabbitMQ", err) {
		return
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if r.logError("Failed to open a channel", err) {
		return
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    "", // default exchange
			"x-dead-letter-routing-key": "retry_queue",
			// arguments
		})

	if r.logError("Failed to declare a queue", err) {
		return
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if r.logError("Failed to set QoS", err) {
		return
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if r.logError("Failed to register a consumer", err) {
		return
	}

	// This channel will keep the main goroutine alive
	var forever chan struct{}

	type TypeHint struct {
		Type string `json:"type"`
	}

	// Start a goroutine to process incoming messages
	go func() {
		for d := range msgs {

			var hint TypeHint
			id := d.Headers["id"].(string)
			uuid, _ := uuid.Parse(id)
			err := json.Unmarshal(d.Body, &hint)
			if r.logError("Failed to unmarshal", err) {
				return
			}

			notifier, factoryFunc, err := pkg.GetNotifier(hint.Type)
			if r.logError("Failed to get notifier", err) {
				return
			}

			payload := factoryFunc()
			err = json.Unmarshal(d.Body, payload)

			if r.logError("Failed to unmarshal", err) {
				return
			}

			err = notifier.Send(payload)
			if r.logError("Failed to send", err) {
				DBMessages, _ := client.Message.Get(context.Background(), uuid)
				rc := DBMessages.Attempts + 1
				client.Message.UpdateOneID(uuid).
					SetAttempts(rc).
					SetNextRetryAt(time.Now().Add(5 * time.Minute)).
					Save(context.Background())
				d.Reject(false)
				continue
			}

			client.Message.UpdateOneID(uuid).
				SetStatus("sent").
				Save(context.Background())
			d.Ack(false)
		}
	}()

	// Block forever (waiting for messages)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
