package notify

import "notification/ent"

type Broker interface {
	Process(message *ent.Message)
}
