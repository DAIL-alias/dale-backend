package routes

import (
	"DALE/handlers"
	"DALE/repositories"
	"DALE/services"
	"DALE/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Create repositories, services, handlers
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

    r.GET("/ping", handlers.PingHandler)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserById)
}