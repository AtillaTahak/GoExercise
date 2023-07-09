package notifiers

import (
	"fmt"
	"sync"
)

type PushNotifier struct {
}

var lock sync.Mutex
var pushNotifier *PushNotifier

func NewPushNotifier() *PushNotifier {
	if pushNotifier == nil {
		lock.Lock()
		if pushNotifier == nil {
			pushNotifier = &PushNotifier{}
		}
		lock.Unlock()
	}
	return pushNotifier
}

func (s *PushNotifier) Notify(message string) error {
	fmt.Println("Push notification:", message)
	return nil
}
