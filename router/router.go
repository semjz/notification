package router

import (
	"github.com/gofiber/fiber/v2"
	"notification/api/handler"
	"notification/ent"
	"notification/infrastructure/rabbitmq"
)

func SetUpRoutes(app *fiber.App, client *ent.Client) {
	app.Post("/notification", handler.SendNotification(rabbitmq.NewRabbitMQ(), client))
}
