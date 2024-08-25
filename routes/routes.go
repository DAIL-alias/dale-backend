package routes

import (
	"DALE/config"
	"DALE/handlers"
	"DALE/middleware"
	"DALE/repositories"
	"DALE/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// For users
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// For aliases
	aliasRepo := repositories.NewAliasRepository(config.DB)
	aliasService := services.NewAliasService(aliasRepo)
	aliasHandler := handlers.NewAliasHandler(aliasService)

	// For auth
	authService := services.NewAuthService(config.RedisClient, userRepo)
	authHandler := handlers.NewAuthHandler(authService, userService)

	r.GET("/ping", handlers.PingHandler)

	// user routes (ADMIN)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserById)

	// alias routes
	aliasesGroup := r.Group("/aliases")
	aliasesGroup.Use(middleware.AuthRequired())
	{
		aliasesGroup.POST("/createalias", aliasHandler.CreateAlias)
		aliasesGroup.GET("", aliasHandler.GetAliases)
		aliasesGroup.GET("/:id", aliasHandler.GetAliasByID)
		aliasesGroup.GET("/getusersaliases/:userID", aliasHandler.GetUsersAliases)
		aliasesGroup.POST("/toggleactivestatus/:id", aliasHandler.ToggleActivateStatus)
	}

	// auth routes
	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/signup", authHandler.SignUp)
	r.POST("/auth/logout", authHandler.Logout)
}
