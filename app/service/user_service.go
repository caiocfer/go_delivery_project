package service

import (
	"errors"
	"fmt"

	"github.com/caiocfer/go_delivery_project/app/models"
	"github.com/caiocfer/go_delivery_project/app/models/security"
	userrepo "github.com/caiocfer/go_delivery_project/app/repository/user_repo"
)

type UserService struct {
	userRepo userrepo.UserRepository
}

func NewUserService(repo userrepo.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(req models.UserCreationRequest) error {
	// TODO Verify if email already exists

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		errors.New("error hashing password")
	}

	userToCreate := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
	}

	err = s.userRepo.CreateUser(userToCreate)
	if err != nil {
		return fmt.Errorf("failed to create user in repo: %w", err)
	}

	return nil
}
