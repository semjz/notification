package handler

import (
	"github.com/gofiber/fiber/v2"
	"notification/pkg/notify"
)

func SendNotification(c *fiber.Ctx) error {
	var meta struct {
		Type string `json:"type"`
	}

	if err := c.BodyParser(&meta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	notifier, factory, err := notify.GetNotifier(meta.Type)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unsupported notifier " + meta.Type})
	}

	payload := factory()
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	ok, err := notifier.Send(&payload)
	if err != nil || !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send notification"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Notification sent"})
}
