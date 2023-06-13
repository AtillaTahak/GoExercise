package main

type Notifier interface {
	Notify(message string) error
}

type RegistrationHandler struct {
	Notifier Notifier
}

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

func main() {
	regHandler := NewRegistrationHandler(NewPushNotifier())

	regHandler.Register("John")
}
