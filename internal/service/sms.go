package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notification/domain/notify"
	"notification/internal"
)

type SmsPayload struct {
	Sender    string   `json:"lineNumber" validate:"required"`
	Recipient []string `json:"mobiles" validate:"required,dive,phonenumber"`
	Message   string   `json:"messageText" validate:"required"`
	Priority  string   `json:"priority,omitempty"`
}

func (s SmsPayload) GetRecipient() []string { return s.Recipient }

func (s SmsPayload) GetMessage() string {
	return s.Message
}

func (s SmsPayload) GetPriority() string {
	return s.Priority
}

func (s SmsPayload) GetSender() string { return s.Sender }

func (s SmsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":        "sms",
		"lineNumber":  s.GetSender(),
		"mobiles":     s.GetRecipient(),
		"messageText": s.GetMessage(),
	})
}

type SmsNotifier struct{}

func (en *SmsNotifier) Send(payload notify.NotifyPayload) error {

	url := "https://api.sms.ir/v1/send/bulk"

	jsonData, _ := payload.MarshalJSON()
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "C0upfrvnr5gm2q0aMciYssv3mLQXcDiyYJExGHUdkqBFPlRUbOqqle3JJWU7Lgpw")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(resp)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func init() {

	internal.RegisterNotifier("sms", internal.NotifierEntry{
		NewPayload: func() notify.NotifyPayload { return &SmsPayload{} },
		Notifier:   &SmsNotifier{},
	})

}
