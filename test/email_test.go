package test

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"notification/router"
	"testing"
)

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "notification",
	})
	router.SetUpRoutes(app)
	return app
}

func TestEmailSuccessful(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
		"recipient": "user@example.com",
		"subject": "Test subject",
		"message": "Test message"
	}`)
	payload := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/notification", payload)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Nil(t, err)

	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, `{"message":"Notification sent"}`, string(body))
}

func TestEmailWrongPayload(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
		"recipient": "user@example.com",
		"subject": "Test subject",
		"message": "Test message",
		"extarField"
	}`)
	payload := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/notification", payload)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Error(t, err)

	assert.Equal(t, 400, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, `{"error": "invalid payload"}`, string(body))
}
