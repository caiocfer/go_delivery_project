package restaurant_repo

import (
	"database/sql"

	"github.com/caiocfer/go_delivery_project/app/models"
)

type restaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) RestaurantRepository {
	return &restaurantRepository{db: db}
}

func (r *restaurantRepository) Create(restaurant models.RestaurantCreationRequest) (int, error) {
	query := `INSERT INTO restaurants (owner_name, email, password_hash, phone, restaurant_name, cuisine_type, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING restaurant_id`

	var restaurantID int

	err := r.db.QueryRow(query, restaurant.OwnerName, restaurant.Email, restaurant.Password, restaurant.Phone, restaurant.RestaurantName, restaurant.CuisineType, restaurant.Description).Scan(&restaurantID)
	if err != nil {
		return 0, err
	}

	return restaurantID, nil
}

func (r *restaurantRepository) FindByEmail(email string) (models.Restaurant, error) {
	query := `SELECT restaurant_id, owner_name, email, password_hash, phone, restaurant_name, cuisine_type, description FROM restaurants WHERE email = $1`

	var restaurant models.Restaurant

	err := r.db.QueryRow(query, email).Scan(&restaurant.Id, &restaurant.OwnerName, &restaurant.Email, &restaurant.Password, &restaurant.Phone, &restaurant.RestaurantName, &restaurant.CuisineType, &restaurant.Description)
	if err != nil {
		return models.Restaurant{}, err
	}

	return restaurant, nil
}
