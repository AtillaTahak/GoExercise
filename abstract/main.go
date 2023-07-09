package main

import (
	"abstract/notifiers"
	"abstract/handlers"
)


// composition root ( Herşeyinizi Başlatılan yer)
func main() {
	// notifiers
	pushNotifier := notifiers.NewPushNotifier()
	pushNotifier2 := notifiers.NewPushNotifier()

	// repositories

	// handlers
	regHandler := handlers.NewRegistrationHandler(pushNotifier)
	regHandler2 := handlers.NewRegistrationHandler(pushNotifier2)
	regHandler2.Register("John2")
	regHandler.Register("John")
	println("regHandler", regHandler)
	println("regHandler", regHandler2)


	// server

	// start
}
