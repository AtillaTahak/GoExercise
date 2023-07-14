package application

import "hexagonal/internal/domain"

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(username, email string) domain.User {
	user := domain.User{
		ID:       generateUserID(),
		Username: username,
		Email:    email,
	}

	s.userRepository.Store(user)
	return user
}

func (s *UserService) GetAllUsers() []domain.User {
	return s.userRepository.FindAll()
}

func generateUserID() int {
	return 0
}
