package main

import (
	"log"

	"github.com/caiocfer/go_delivery_project/app/api"
	"github.com/caiocfer/go_delivery_project/app/db"
	"github.com/caiocfer/go_delivery_project/app/handler"
	userrepo "github.com/caiocfer/go_delivery_project/app/repository/user_repo"
	"github.com/caiocfer/go_delivery_project/app/service"
)

func main() {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Error connecting to db %v", err)
	}
	defer dbConn.Close()

	userRepository := userrepo.NewUserRepo(dbConn)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api.SetupGin(userHandler)
}
