package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caiocfer/go_delivery_project/app/models"
	"github.com/caiocfer/go_delivery_project/app/service"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantService *service.RestaurantService
}

func NewRestaurantHandler(svc *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{
		svc,
	}
}

func (h *RestaurantHandler) CreateRestaurantHandler(c *gin.Context) {
	var req models.RestaurantCreationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err.Error())})
		return
	}

	if err := req.PrepareField(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Validation error: %v", err.Error())})
		return
	}
	_, err := h.restaurantService.CreateRestaurant(req)

	if err != nil {
		log.Printf("Error creating restaurant: %v", err)
		statusCode := http.StatusInternalServerError
		errorMessage := "Failed to create restaurant"

		if err.Error() == fmt.Sprintf("email '%s' is already registered", req.Email) {
			statusCode = http.StatusBadRequest
			errorMessage = err.Error()
		}

		c.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Restaurant created successfully",
	})
}

func (h *RestaurantHandler) LoginHandler(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err.Error())})
		return
	}

	token, err := h.restaurantService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
