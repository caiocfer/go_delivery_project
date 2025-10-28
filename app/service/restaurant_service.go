package service

import (
	"errors"

	"github.com/caiocfer/go_delivery_project/app/models"
	"github.com/caiocfer/go_delivery_project/app/models/security"
	"github.com/caiocfer/go_delivery_project/app/repository/restaurant_repo"
)

type RestaurantService struct {
	restaurantRepo restaurant_repo.RestaurantRepository
}

func NewRestaurantService(repo restaurant_repo.RestaurantRepository) *RestaurantService {
	return &RestaurantService{restaurantRepo: repo}
}

func (s *RestaurantService) CreateRestaurant(restaurant models.RestaurantCreationRequest) (int, error) {
	if err := restaurant.PrepareField(); err != nil {
		return 0, err
	}

	hashedPassword, err := security.HashPassword(restaurant.Password)
	if err != nil {
		return 0, errors.New("error hashing password")
	}

	restaurant.Password = hashedPassword

	return s.restaurantRepo.Create(restaurant)
}

func (s *RestaurantService) Login(email, password string) (string, error) {
	restaurant, err := s.restaurantRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := security.VerifyPassword(restaurant.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := security.GenerateJWT(restaurant.Id)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}
