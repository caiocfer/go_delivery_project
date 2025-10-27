package userrepo

import "github.com/caiocfer/go_delivery_project/app/models"

type UserRepository interface {
	CreateUser(user models.User) error
	EmailExists(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
}
