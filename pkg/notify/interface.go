package notify

type Notifier interface {
	Send(payload *NotifyPayload) (bool, error)
}

type NotifyPayload interface {
	GetRecipient() string
	GetMessage() string
	GetSubject() string
	GetPriority() string
}
