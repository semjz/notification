package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"notification/domain/notify"
	"notification/ent"
	"notification/pkg"
	_ "notification/pkg/service"
)

func SendNotification(broker notify.Broker, client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var meta struct {
			Type string `json:"type"`
		}

		if err := c.BodyParser(&meta); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, err.Error())
		}

		_, factory, err := pkg.GetNotifier(meta.Type)
		if err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "unsupported notifier "+meta.Type)
		}

		payload := factory()

		decoder := pkg.ValidatePayloadStructure(c.Body())

		if err := decoder.Decode(&payload); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "payload structure validation error")
		}

		if err := pkg.ValidatePayloadFields(payload); err != nil {
			return respondWithError(c, fiber.StatusBadRequest, "payload field validation error")
		}
		DBTypePayload, err := pkg.StructToMap(payload)
		message, err := client.Message.
			Create().
			SetType(meta.Type).
			SetPayload(DBTypePayload).
			SetAttempts(1).
			Save(context.Background())
		broker.Process(message)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Notification sent"})
	}
}

func respondWithError(c *fiber.Ctx, status int, errMsg string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": errMsg,
	})
}
