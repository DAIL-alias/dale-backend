package services

import (
	"DALE/models"
	"DALE/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// Create a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *UserService) GetUserById(id int) (models.User, error) {
	return s.UserRepository.GetUserById(id)
}