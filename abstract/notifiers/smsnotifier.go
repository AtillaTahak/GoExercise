package notifiers

import (
	"fmt"
	"sync"
)

// notifiers paketi içinde SmsNotifier adında bir yapı tanımlanıyor.
type SmsNotifier struct {
}

var smsNotifier *SmsNotifier
var lock sync.Mutex

// NewSmsNotifier, SmsNotifier yapısının yeni bir örneğini döndürür.
// Bu fonksiyon, SmsNotifier yapısından bir örnek oluşturarak geri döndürür.
func NewSmsNotifier() *SmsNotifier {
	if smsNotifier == nil {
		lock.Lock()
		if smsNotifier == nil {
			smsNotifier = &SmsNotifier{}
		}
		lock.Unlock()
	}
	return smsNotifier
}

// Bu fonksiyon, SmsNotifier yapısına ait bir metottur (yöntem).
// Fonksiyon, s *SmsNotifier alıcısıyla birlikte çağrılabilir.
func (s *SmsNotifier) Notify(message string) error {
	fmt.Println("SMS notification:", message)
	return nil
}
