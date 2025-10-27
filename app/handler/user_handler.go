package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caiocfer/go_delivery_project/app/models"
	"github.com/caiocfer/go_delivery_project/app/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req models.UserCreationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err.Error())})
		return
	}

	if err := req.PrepareField(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Validation error: %v", err.Error())})
		return
	}
	err := h.userService.CreateUser(req)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		statusCode := http.StatusInternalServerError
		errorMessage := "Failed to create user"

		if err.Error() == fmt.Sprintf("email '%s' is already registered", req.Email) {
			statusCode = http.StatusBadRequest
			errorMessage = err.Error()
		}

		c.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err.Error())})
		return
	}

	token, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
