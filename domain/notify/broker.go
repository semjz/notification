package notify

import "notification/domain/notify"

type Broker interface {
	Send(payload NotifyPayload)
}
