package main

import (
	"fmt"
	"log"
	"os"

	"crud-app/config"
	"crud-app/controller"
	"crud-app/database"
	"crud-app/logger"
	"crud-app/manager"
	"crud-app/mongoDatabase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var userController *controller.UserController

// var validate *validator.Validate

// func init() {
// 	config.Load()
// 	logger.InitLogger()
// 	client := database.Connect()
// 	userCollection := mongoDatabase.InitUserCollection(client)

//		userManager := manager.NewUserManager(userCollection)
//		validate := validator.New()
//		userController = controller.NewUserController(userManager, validate)
//	}
func init() {
	config.Load()
	logger.InitLogger()
	database.Connect()

	userCollection := mongoDatabase.InitUserCollection(database.MongoClient)

	userManager := manager.NewUserManager(userCollection)
	validate := validator.New()
	userController = controller.NewUserController(userManager, validate)
}

func main() {
	e := echo.New()

	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetAllUsers)
	e.GET("/users/:id", userController.GetUserByID)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running  http://localhost:%s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
