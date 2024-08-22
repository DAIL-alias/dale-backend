package routes

import (
	"DALE/config"
	"DALE/handlers"
	"DALE/repositories"
	"DALE/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Create repositories, services, handlers for users
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Create repositories, services, handlers for aliases
	aliasRepo := repositories.NewAliasRepository(config.DB)
	aliasService := services.NewAliasService(aliasRepo)
	aliasHandler := handlers.NewAliasHandler(aliasService)

	r.GET("/ping", handlers.PingHandler)

	//user routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserById)

	//alias routes
	r.POST("/aliases/createalias", aliasHandler.CreateAlias)
	r.GET("/aliases", aliasHandler.GetAliases)
	r.GET("/aliases/:id", aliasHandler.GetAliasByID)
	r.GET("/aliases/getusersaliases/:userID", aliasHandler.GetUsersAliases)
	r.POST("/aliases/toggleactivestatus/:id", aliasHandler.ToggleActivateStatus)

	//auth routes
}