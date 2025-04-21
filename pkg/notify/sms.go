package notify

import "fmt"

type SmsPayload struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
	Priority  string `json:"priority,omitempty"`
}

func (s *SmsPayload) GetRecipient() string {
	return s.Recipient
}

func (s *SmsPayload) GetMessage() string {
	return s.Message
}

func (s *SmsPayload) GetSubject() string {
	return s.Priority
}

func (e *SmsPayload) GetPriority() string {
	return e.Priority
}

type SmsNotifier struct{}

func (en *SmsNotifier) Send(payload *NotifyPayload) (bool, error) {
	fmt.Println("sending sms")
	fmt.Println("payload", payload)
	return true, nil
}

func init() {
	RegisterNotifier("sms", NotifierEntry{
		NewPayload: func() NotifyPayload { return &SmsPayload{} },
		Notifier:   &SmsNotifier{},
	})
}
