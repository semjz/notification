package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	database "notification/cmd/migrate"
	"notification/infrastructure/rabbitmq"
	"notification/router"
	"os"
	"os/signal"
	"syscall"
)

var client = database.DBConnect()

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "notification",
	})
	router.SetUpRoutes(app, client)
	return app
}

func main() {
	app := SetUp()

	rabbitmqService := rabbitmq.NewRabbitMQ()

	go rabbitmqService.Receiver(client)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := app.Listen(":3000")
		if err != nil {
			log.Fatal("Error starting server", err)
		}
	}()

	<-signalChan
	log.Println("Shutting down...")

}
