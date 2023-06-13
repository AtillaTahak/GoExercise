package main

import (
	"fmt"
)


type SmsNotifier struct {
}

func NewSmsNotifier() *SmsNotifier {
	return &SmsNotifier{}
}

func (s *SmsNotifier) Notify(message string) error {
	fmt.Println("SMS notification:", message)
	return nil
}