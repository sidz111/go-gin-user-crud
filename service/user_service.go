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

func (s *UserService) GetByID(id int) (*model.User, error) {
	if id < 1 {
		return nil, fmt.Errorf("Id is Invalid")
	}
	user, err := s.repo.GetByid(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeletebyId(id int) error {
	if id < 1 {
		return fmt.Errorf("Id is Invalid")
	}
	return s.DeletebyId(id)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUser(id int, user *model.User) error {
	if id < 1 {
		return fmt.Errorf("Id is Invalid")
	}

	u, err := s.repo.GetByid(id)
	if err != nil {
		return fmt.Errorf("User Not Found")
	}

	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Address != "" {
		u.Address = user.Address
	}
	return s.repo.UpdateUser(id, u)
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
