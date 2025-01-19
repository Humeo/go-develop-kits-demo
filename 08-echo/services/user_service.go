package services

import (
	"08-echo/models"
	"errors"
)

type UserService struct {
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]models.User, 0),
	}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.users, nil
}

func (s *UserService) GetByID(id int) (*models.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	user.ID = len(s.users) + 1
	s.users = append(s.users, *user)
	return user, nil
}
