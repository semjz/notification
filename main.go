package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"notification/infrastructure/rabbitmq"
	"notification/router"
	"os"
	"os/signal"
	"syscall"
)

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "notification",
	})
	router.SetUpRoutes(app)
	return app
}

func main() {
	app := SetUp()

	go rabbitmq.Receiver()

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
