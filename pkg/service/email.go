package service

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"notification/domain/notify"
	"notification/pkg/setup"
)

type EmailPayload struct {
	Sender      string   `json:"sender" validate:"required,email"`
	Recipient   []string `json:"recipient" validate:"required,dive,email"`
	Subject     string   `json:"subject" validate:"required"`
	Message     string   `json:"message" validate:"required"`
	Priority    string   `json:"priority,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}

func (e *EmailPayload) GetRecipient() []string {
	return e.Recipient
}

func (e *EmailPayload) GetMessage() string {
	return e.Message
}

func (e *EmailPayload) GetSubject() string {
	return e.Subject
}

func (e *EmailPayload) GetPriority() string {
	return e.Subject
}

type EmailNotifier struct{}

func (en *EmailNotifier) Send(payload notify.NotifyPayload) error {
	emailPayload, ok := payload.(*EmailPayload)
	if !ok {
		return fmt.Errorf("payload does not support sender")
	}
	sender := emailPayload.Sender
	apiKey := "re_6cCDxFbw_34HjduPxRD4jY5W4wtumstCa"
	client := resend.NewClient(apiKey)
	params := &resend.SendEmailRequest{
		From:    sender,
		To:      emailPayload.GetRecipient(),
		Subject: emailPayload.GetSubject(),
		Html:    emailPayload.GetMessage(),
	}

	_, err := client.Emails.Send(params)
	return err
}

func init() {
	setup.RegisterNotifier("email", setup.NotifierEntry{
		NewPayload: func() notify.NotifyPayload { return &EmailPayload{} },
		Notifier:   &EmailNotifier{},
	})
}
