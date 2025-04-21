package notify

import (
	"fmt"
)

type EmailPayload struct {
	Recipient   string   `json:"recipient"`
	Subject     string   `json:"subject"`
	Message     string   `json:"message"`
	Priority    string   `json:"priority,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}

func (e *EmailPayload) GetRecipient() string {
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

func (en *EmailNotifier) Send(payload *NotifyPayload) (bool, error) {
	fmt.Println("sending email")
	fmt.Println("payload", payload)
	return true, nil
}

func init() {
	RegisterNotifier("email", NotifierEntry{
		NewPayload: func() NotifyPayload { return &EmailPayload{} },
		Notifier:   &EmailNotifier{},
	})
}
