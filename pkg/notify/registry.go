package notify

import (
	"errors"
)

var notifyRegistry = make(map[string]NotifierEntry)

type NotifierEntry struct {
	NewPayload func() NotifyPayload
	Notifier   Notifier
}

func RegisterNotifier(notifyType string, entry NotifierEntry) {
	notifyRegistry[notifyType] = entry
}

func GetNotifier(notifyType string) (Notifier, func() NotifyPayload, error) {
	entry, exists := notifyRegistry[notifyType]
	if !exists {
		return nil, nil, errors.New("no notifier for " + notifyType)
	}
	return entry.Notifier, entry.NewPayload, nil
}
