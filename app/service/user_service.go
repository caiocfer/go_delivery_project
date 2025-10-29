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
	exists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return errors.New("email already exists")
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return errors.New("error hashing password")
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

func (s *UserService) Login(req models.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	if err := security.VerifyPassword(user.Password, req.Password); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := security.GenerateJWT(user.Id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}

func (s *UserService) GetUserFromToken(tokenStr string) (*models.User, error) {
	claims, err := security.ValidateJWT(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}
