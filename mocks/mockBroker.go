package mocks

import (
	"notification/ent"
)

type MockBroker struct {
	calledwith *ent.Message
}

func (m *MockBroker) Process(message *ent.Message) {
	m.calledwith = message
}
