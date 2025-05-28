package api

import (
	"github.com/caiocfer/go_delivery_project/common"
	"github.com/gin-gonic/gin"
)

func SetupGin() {
	router := gin.Default()
	setupRoutes(router)
	router.Run(common.SERVER_PORT)
}

func setupRoutes(router *gin.Engine) {
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(
			200, gin.H{"hello": "world"})
	})
}
