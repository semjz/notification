package notify

type Notifier interface {
	Send(payload NotifyPayload) error
}
