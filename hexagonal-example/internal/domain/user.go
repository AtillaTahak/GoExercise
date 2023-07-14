package domain

type User struct {
	ID       int
	Username string
	Email    string
}

type UserRepository interface {
	Store(user User)
	FindAll() []User
}
