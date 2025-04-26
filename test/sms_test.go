package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestSMSSuccessful(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "sms",
		"recipient": "+989121111111",
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

func TestSMSWrongNumberFormat(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "email",
		"recipient": "11111",
		"message": "Test message"
	}`)
	payload := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/notification", payload)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Equal(t, 400, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	assert.Contains(t, string(body), "Validation error")
}
