package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "notification/pkg/service"
	"notification/pkg/setup"
)

var validate = setup.Validate

func SendNotification(c *fiber.Ctx) error {
	var meta struct {
		Type string `json:"type"`
	}

	if err := c.BodyParser(&meta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	notifier, factory, err := setup.GetNotifier(meta.Type)
	fmt.Println(setup.NotifyRegistry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "unsupported notifier " + meta.Type})
	}

	payload := factory()

	var raw map[string]json.RawMessage
	json.Unmarshal(c.Body(), &raw)
	delete(raw, "type")

	body, _ := json.Marshal(raw)

	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Validation error": err.Error()})
	}

	if err := validate.Struct(payload); err != nil {
		errs := make([]string, 0)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, err := range err.(validator.ValidationErrors) {
				errs = append(errs, err.Field()+" is "+err.Tag())
			}
		} else {
			errs = append(errs, err.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Validation error": errs})
	}

	err = notifier.Send(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Notification sent"})
}
