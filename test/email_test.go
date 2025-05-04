package test

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"notification/api/handler"
	"notification/mocks"
	"notification/router"
	"testing"
)

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "notification",
	})
	mockBroker := mocks.MockBroker{}
	app.Post("/notification", handler.SendNotification(&mockBroker))
	router.SetUpRoutes(app)
	return app
}

func TestEmailSuccessful(t *testing.T) {
	app := SetUp()

	jsonBody := []byte(`{
		"type": "email",
		"recipient": ["Test@gmail.com"],
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

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t,
		200, res.StatusCode, "Expected 200, got %d. Response: %s",
		res.StatusCode, string(body))

}

func TestEmailWrongEmailFormat(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
		"recipient": ["user"],
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

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t,
		400, res.StatusCode, "Expected 400, got %d. Response: %s",
		res.StatusCode, string(body))

	assert.Contains(t, string(body), "payload field validation error")
}

func TestEmailMissingField(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
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

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t,
		400, res.StatusCode, "Expected 400, got %d. Response: %s",
		res.StatusCode, string(body))

	assert.Contains(t, string(body), "payload field validation error")
}

func TestEmailExtraField(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
		"recipient": "user@yahoo.com",
		"subject": "Test subject",
		"message": "Test message",
		"extra": "extra field"
	}`)
	payload := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/notification", payload)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t,
		400, res.StatusCode, "Expected 400, got %d. Response: %s",
		res.StatusCode, string(body))

	assert.Contains(t, string(body), "payload structure validation error")
}
