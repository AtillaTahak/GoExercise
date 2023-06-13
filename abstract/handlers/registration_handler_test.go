package handlers

import (
	"testing"
)
type MockNotifier struct {
}

func (m MockNotifier) Notify(message string) error {
	return nil
}

func TestRegister(t *testing.T) {
	// handler 
	registrationHandler := NewRegistrationHandler(MockNotifier{})
	err := registrationHandler.Register("John")
	if err != nil {
		t.Error("Error should be nil")
	}
}