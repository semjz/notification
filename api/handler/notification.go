package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"notification/domain/notify"
	"notification/ent"
	"notification/internal"
	_ "notification/internal/service"
)

func SendNotification(broker notify.Broker, client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var meta struct {
			Type string `json:"type"`
		}

		if err := c.BodyParser(&meta); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, err.Error())
		}

		_, factory, err := internal.GetNotifier(meta.Type)
		if err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "unsupported notifier "+meta.Type)
		}

		payload := factory()

		decoder := internal.ValidatePayloadStructure(c.Body())

		if err := decoder.Decode(&payload); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "payload structure validation error")
		}

		if err := internal.ValidatePayloadFields(payload); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "payload field validation error")
		}
		DBTypePayload, err := internal.StructToMap(payload)
		message, err := client.Message.
			Create().
			SetType(meta.Type).
			SetPayload(DBTypePayload).
			Save(context.Background())
		fmt.Println(message)
		broker.Process(message)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Notification sent"})
	}
}

func respondWithError(c *fiber.Ctx, status int, errMsg string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": errMsg,
	})
}
