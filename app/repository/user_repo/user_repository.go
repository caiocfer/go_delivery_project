package userrepo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/caiocfer/go_delivery_project/app/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &Users{db}
}

func (repository *Users) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := repository.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return exists, nil
}

func (repository *Users) CreateUser(user models.User) error {
	query := `
					INSERT INTO users (name, email, password_hash, phone)
					VALUES ($1, $2, $3, $4)
	`

	_, err := repository.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Phone,
	)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err) // Log do erro completo
		return fmt.Errorf("failed to insert user into database: %w", err)
	}

	log.Printf("User '%s' created successfully", user.Name)

	return nil
}
