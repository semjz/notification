package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notification/domain/notify"
	"notification/pkg/setup"
)

type SmsPayload struct {
	Sender    string   `json:"sender" validate:"required"`
	Recipient []string `json:"recipient" validate:"required,dive,phonenumber"`
	Message   string   `json:"message" validate:"required"`
	Priority  string   `json:"priority,omitempty"`
}

func (s *SmsPayload) GetRecipient() []string {
	return s.Recipient
}

func (s *SmsPayload) GetMessage() string {
	return s.Message
}

func (s *SmsPayload) GetPriority() string {
	return s.Priority
}

type SmsNotifier struct{}

func (en *SmsNotifier) Send(payload notify.NotifyPayload) error {

	smsPayload, ok := payload.(*SmsPayload)
	if !ok {
		return fmt.Errorf("payload does not support sender")
	}

	sender := smsPayload.Sender
	url := "https://api.sms.ir/v1/send/bulk"
	requiredSmsPayload := make(map[string]interface{})
	requiredSmsPayload["lineNumber"] = sender
	requiredSmsPayload["messageText"] = smsPayload.GetMessage()
	requiredSmsPayload["mobiles"] = smsPayload.GetRecipient()

	jsonData, _ := json.Marshal(requiredSmsPayload)
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

	setup.RegisterNotifier("sms", setup.NotifierEntry{
		NewPayload: func() notify.NotifyPayload { return &SmsPayload{} },
		Notifier:   &SmsNotifier{},
	})

}
