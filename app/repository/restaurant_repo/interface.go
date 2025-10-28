package restaurant_repo

import "github.com/caiocfer/go_delivery_project/app/models"

type RestaurantRepository interface {
	Create(restaurant models.RestaurantCreationRequest) (int, error)
	FindByEmail(email string) (models.Restaurant, error)
}
