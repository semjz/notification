package notify

type NotifyPayload interface {
	GetRecipient() []string
	GetMessage() string
	GetPriority() string
}
