package main

import (
	"github.com/gofiber/fiber/v2"
	"notification/router"
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
	app.Listen(":3000")
}
