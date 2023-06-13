package notifiers

import (
	"fmt"
)

// notifiers paketi içinde SmsNotifier adında bir yapı tanımlanıyor.
type SmsNotifier struct {
}

// NewSmsNotifier, SmsNotifier yapısının yeni bir örneğini döndürür.
// Bu fonksiyon, SmsNotifier yapısından bir örnek oluşturarak geri döndürür.
func NewSmsNotifier() *SmsNotifier {
	return &SmsNotifier{}
}

// Bu fonksiyon, SmsNotifier yapısına ait bir metottur (yöntem).
// Fonksiyon, s *SmsNotifier alıcısıyla birlikte çağrılabilir.
func (s *SmsNotifier) Notify(message string) error {
	fmt.Println("SMS notification:", message)
	return nil
}
