package handlers

type RegistrationHandler struct {
	Notifier Notifier
}
// NewRegistrationHandler, RegistrationHandler yapısının yeni bir örneğini döndürür.
func NewRegistrationHandler(n Notifier) *RegistrationHandler {
	return &RegistrationHandler{
		Notifier: n,
	}
}

func (r *RegistrationHandler) Register(name string) error {
	err := r.Notifier.Notify("New user registered: " + name)
	if err != nil {
		return err
	}
	return nil
}
/*

Notifier arayüzü, bildirim göndermek için kullanılacak fonksiyonun belirli bir imzasını temsil eder. Bu arayüzü kullanarak farklı bildirim yöntemlerini destekleyen nesneleri kullanabiliriz.
RegistrationHandler yapısı, kayıt işlemini yönetir. Notifier adında bir alan içerir, bu da bildirim göndermek için kullanılacak bir nesneyi temsil eder.
NewRegistrationHandler fonksiyonu, RegistrationHandler yapısının yeni bir örneğini döndürür. Bu örneği oluştururken, Notifier nesnesi parametre olarak alınır ve RegistrationHandler yapısının ilgili alanına atanır.
Register metodu, belirtilen isimle yeni bir kullanıcı kaydeder. Bu metot, Notifier nesnesini kullanarak bir bildirim gönderir. Eğer bildirim gönderimi sırasında bir hata oluşursa, bu hata geri döndürülür.

*/
