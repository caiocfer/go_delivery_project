package models

import (
	"errors"
	"strings"
)

type Restaurant struct {
	Id             uint64 `json:"id,omitempty"`
	OwnerName      string `json:"owner_name,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"-"`
	Phone          string `json:"phone,omitempty"`
	RestaurantName string `json:"restaurant_name,omitempty"`
	CuisineType    string `json:"cuisine_type,omitempty"`
	Description    string `json:"description,omitempty"`
}

type RestaurantCreationRequest struct {
	OwnerName      string `json:"owner_name" validate:"required,min=3,max=50"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=6"`
	Phone          string `json:"phone,omitempty"`
	RestaurantName string `json:"restaurant_name,omitempty"`
	CuisineType    string `json:"cuisine_type,omitempty"`
	Description    string `json:"description,omitempty"`
}

func (restaurant *RestaurantCreationRequest) PrepareField() error {
	err := restaurant.validateFields()
	if err != nil {
		return err
	}

	restaurant.formatField()

	return nil
}

func (restaurant *RestaurantCreationRequest) validateFields() error {
	if restaurant.OwnerName == "" {
		return errors.New("owner name can't be empty")
	}

	if restaurant.Password == "" {
		return errors.New("password can't be empty")
	}

	if restaurant.Email == "" {
		return errors.New("email can't be empty")
	}

	if restaurant.Phone == "" {
		return errors.New("phone can't be empty")
	}

	if restaurant.RestaurantName == "" {
		return errors.New("restaurant name can't be empty")
	}

	return nil
}

func (restaurant *RestaurantCreationRequest) formatField() {
	restaurant.OwnerName = strings.TrimSpace(restaurant.OwnerName)
	restaurant.Password = strings.TrimSpace(restaurant.Password)
	restaurant.Email = strings.TrimSpace(restaurant.Email)
	restaurant.Phone = strings.TrimSpace(restaurant.Phone)
	restaurant.RestaurantName = strings.TrimSpace(restaurant.RestaurantName)
}
