package main

import (
	"fmt"
)

type PushNotifier struct {
}

func NewPushNotifier() *PushNotifier {
	return &PushNotifier{}
}

func (s *PushNotifier) Notify(message string) error {
	fmt.Println("Push notification:", message)
	return nil
}