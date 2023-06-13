package main

import (
	"abstract/notifiers"
	"abstract/handlers"
)
// composition root ( Herşeyinizi Başlatılan yer)
func main() {
	// notifiers
	pushNotifier := notifiers.NewPushNotifier()

	// repositories

	// handlers
	regHandler := handlers.NewRegistrationHandler(pushNotifier)
	regHandler.Register("John")

	// server

	// start
}
