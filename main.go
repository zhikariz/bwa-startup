package main

import (
	"bwa-startup/auth"
	"bwa-startup/handler"
	"bwa-startup/helper"
	"bwa-startup/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	db := helper.SetupDb()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Repository
	userRepository := user.NewRepository(db)
	// Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	// Handler
	userHandler := handler.NewUserHandler(userService, authService)

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

}
