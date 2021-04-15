package main

import (
	"bwa-startup/handler"
	"bwa-startup/helper"
	"bwa-startup/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := helper.SetupDb()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// User Endpoint
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
