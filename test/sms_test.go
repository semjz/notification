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
		"lineNumber": "30007732904278",
		"mobiles": ["+989121111111"],
		"messageText": "Test message"
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

func TestSMSWrongNumberFormat(t *testing.T) {
	app := SetUp()
	jsonBody := []byte(`{
		"type": "sms",
		"lineNumber": "30007732904278",
		"mobiles": ["11111"],
		"messageText": "Test message"
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
