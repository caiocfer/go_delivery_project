package api

import (
	"github.com/caiocfer/go_delivery_project/app/handler"
	"github.com/caiocfer/go_delivery_project/common"
	"github.com/gin-gonic/gin"
)

func SetupGin(userHandler *handler.UserHandler, restaurantHandler *handler.RestaurantHandler) {
	router := gin.Default()
	setupRoutes(router, userHandler, restaurantHandler)
	router.Run(common.SERVER_PORT)
}

func setupRoutes(router *gin.Engine, userHandler *handler.UserHandler, restaurantHandler *handler.RestaurantHandler) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(
			200, gin.H{"hello": "world"})
	})

	router.POST("/createuser", userHandler.CreateUserHandler)
	router.POST("/login", userHandler.LoginHandler)
	router.GET("/getuser", userHandler.GetUserHandler)

	router.POST("/createrestaurant", restaurantHandler.CreateRestaurantHandler)
	router.POST("/loginrestaurant", restaurantHandler.LoginHandler)
}
