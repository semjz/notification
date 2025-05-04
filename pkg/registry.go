package setup

import (
	"errors"
	"notification/domain/notify"
)

var NotifyRegistry = make(map[string]NotifierEntry)

type NotifierEntry struct {
	NewPayload func() notify.NotifyPayload
	Notifier   notify.Notifier
}

func RegisterNotifier(notifyType string, entry NotifierEntry) {
	NotifyRegistry[notifyType] = entry
}

func GetNotifier(notifyType string) (notify.Notifier, func() notify.NotifyPayload, error) {
	entry, exists := NotifyRegistry[notifyType]
	if !exists {
		return nil, nil, errors.New("no notifier for " + notifyType)
	}
	return entry.Notifier, entry.NewPayload, nil
}
