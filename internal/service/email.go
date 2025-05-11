package service

import (
	"encoding/json"
	"fmt"
	"notification/config"
	"notification/domain/notify"
	"notification/infrastructure/email"
	"notification/internal"
)

type EmailPayload struct {
	Recipient   []string `json:"recipient" validate:"required,dive,email"`
	Subject     string   `json:"subject" validate:"required"`
	Message     string   `json:"message" validate:"required"`
	Priority    string   `json:"priority,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}

func (e EmailPayload) GetRecipient() []string {
	return e.Recipient
}

func (e EmailPayload) GetMessage() string {
	return e.Message
}

func (e EmailPayload) GetPriority() string {
	return e.Priority
}

func (e EmailPayload) GetSubject() string { return e.Subject }

func (e EmailPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":      "email",
		"recipient": e.GetRecipient(),
		"subject":   e.GetSubject(),
		"message":   e.GetMessage(),
		"priority":  e.GetPriority(),
	})
}

type EmailNotifier struct{}

func (en *EmailNotifier) Send(payload notify.NotifyPayload) error {
	emailPayload, ok := payload.(*EmailPayload)
	if !ok {
		return fmt.Errorf("payload does not support sender")
	}
	SMTPService := email.NewSMTPService(config.GetConfig())
	err := SMTPService.SendMail(emailPayload.GetRecipient(), emailPayload.GetSubject(), emailPayload.GetMessage())
	if err != nil {
		return err
	}
	return nil
}

func init() {
	internal.RegisterNotifier("email", internal.NotifierEntry{
		NewPayload: func() notify.NotifyPayload { return &EmailPayload{} },
		Notifier:   &EmailNotifier{},
	})
}
