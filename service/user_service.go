package service

import (
	"fmt"

	"github.com/sidz111/user-management-crud/model"
	"github.com/sidz111/user-management-crud/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *model.User) (*model.User, error) {
	if err := ValidateUser(user); err != nil {
		return nil, err
	}
	if err := s.repo.Create(*user); err != nil {
		return nil, err
	}
	return user, nil
}

func ValidateUser(user *model.User) error {
	if user.Name == "" {
		return fmt.Errorf("Name is Required")
	}
	if user.Address == "" {
		return fmt.Errorf("Address is Required")
	}
	return nil
}
