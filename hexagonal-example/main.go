package main

import (
	"fmt"
	"hexagonal/internal/application"
	"hexagonal/internal/infrastructure"
)

func main() {
	userRepository := &infrastructure.InMemoryUserRepository{}

	userService := application.NewUserService(userRepository)

	user := userService.CreateUser("john_doe", "john.doe@example.com")
	fmt.Printf("Yeni kullanici oluşturuldu: %+v\n", user)

	users := userService.GetAllUsers()
	fmt.Println("Tüm kullanicilar:")
	for _, u := range users {
		fmt.Printf("- %+v\n", u)
	}
}
