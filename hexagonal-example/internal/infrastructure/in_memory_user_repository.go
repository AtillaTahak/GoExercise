package infrastructure

import "hexagonal/internal/domain"

type InMemoryUserRepository struct {
	users []domain.User
}

func (r *InMemoryUserRepository) Store(user domain.User) {
	r.users = append(r.users, user)
}

func (r *InMemoryUserRepository) FindAll() []domain.User {
	return r.users
}
