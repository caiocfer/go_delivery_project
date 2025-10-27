package models

import (
	"errors"
	"strings"
)

type User struct {
	Id       uint64 `json:"user_id,omitempty"`
	Name     string `json:"username,omitempty"`
	Password string `json:"-"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type UserCreationRequest struct {
	Name     string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone,omitempty"`
}

func (user *UserCreationRequest) PrepareField() error {
	err := user.validateFields()
	if err != nil {
		return err
	}

	user.formatField()

	return nil
}

func (user *UserCreationRequest) validateFields() error {
	if user.Name == "" {
		return errors.New("user can't be empty")
	}

	if user.Password == "" {
		return errors.New("password can't be empty")
	}

	if user.Email == "" {
		return errors.New("email can't be empty")
	}

	if user.Phone == "" {
		return errors.New("phone can't be empty")
	}

	return nil
}

func (user *UserCreationRequest) formatField() {
	user.Name = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = strings.TrimSpace(user.Phone)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
