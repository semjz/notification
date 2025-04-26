package router

import (
	"github.com/gofiber/fiber/v2"
	"notification/api/handler"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/notification", handler.SendNotification)
}
