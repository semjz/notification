package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	database "notification/cmd/migrate"
	"notification/ent"
	"notification/internal"
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
		nil,
	)

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

	// Start a goroutine to process incoming messages
	// Process messages
	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
			r.processMessage(d, ch, client)
		}
	}()

	// Block forever (waiting for messages)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQ) processMessage(d amqp.Delivery, ch *amqp.Channel, client *ent.Client) {
	type TypeHint struct {
		Type string `json:"type"`
	}
	var hint TypeHint
	id := d.Headers["id"].(string)
	uuid, _ := uuid.Parse(id)
	err := json.Unmarshal(d.Body, &hint)
	if r.logError("Failed to unmarshal", err) {
		return
	}

	notifier, factoryFunc, err := internal.GetNotifier(hint.Type)
	if r.logError("Failed to get notifier", err) {
		return
	}

	payload := factoryFunc()
	err = json.Unmarshal(d.Body, payload)
	fmt.Println(payload)
	if r.logError("Failed to unmarshal", err) {
		return
	}

	err = notifier.Send(payload)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Failed to send:", err)
		r.handleRetry(client, uuid, ch, d)
	} else {
		fmt.Println("Successfully sent")
		client.Message.UpdateOneID(uuid).
			SetStatus("sent").
			Save(context.Background())
		d.Ack(false)
	}
}

func (r *RabbitMQ) handleRetry(client *ent.Client, uuid uuid.UUID, ch *amqp.Channel, d amqp.Delivery) {
	retryMsg, err := database.QueryRetry(context.Background(), client, uuid)
	var retries int

	if err != nil && retryMsg == nil {
		client.Retry.Create().
			SetMessageUUID(uuid).
			SetAttempts(1).
			SetNextRetryAt(time.Now().Add(5 * time.Minute)).
			Save(context.Background())
	} else {
		retries = retryMsg.Attempts
		if retries >= 3 {
			client.Message.UpdateOneID(uuid).
				SetStatus("failed").
				Save(context.Background())

			retryMsg.Update().SetStatus("failed").Save(context.Background())
			d.Ack(false)
			return
		}
		rtr := retries + 1
		retryMsg.Update().
			SetAttempts(rtr).
			SetNextRetryAt(time.Now().Add(5 * time.Minute)).
			Save(context.Background())
	}

	// Acknowledge the original message
	d.Ack(false)

	// Start the goroutine to retry after 5 minutes
	go func() {
		<-time.After(5 * time.Minute)
		r.retryMessage(client, uuid, ch)
	}()
}

func (r *RabbitMQ) retryMessage(client *ent.Client, uuid uuid.UUID, ch *amqp.Channel) {
	// Fetch the original message
	DBMessage, err := client.Message.Get(context.Background(), uuid)

	if err != nil {
		fmt.Println("Failed to fetch message:", err)
		return
	}

	// Republish the message
	bytesMsg, _ := json.Marshal(DBMessage)
	err = ch.Publish(
		"",
		"task_queue",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         bytesMsg,
			Priority:     uint8(DBMessage.Edges.Retry.Attempts),
			Headers: amqp.Table{
				"id": uuid.String(),
			},
		},
	)

	if err != nil {
		fmt.Println("Failed to republish message:", err)
	} else {
		fmt.Println("Retry message republished successfully!")
	}
}
