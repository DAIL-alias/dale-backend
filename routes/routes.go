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

	usersRoute := r.Group("/users")
	usersRoute.Use(middleware.RoleRequired(2))
	{
		r.POST("", userHandler.CreateUser)
		r.GET("", userHandler.GetUsers)
		r.GET("/:id", userHandler.GetUserById)
	}

	// alias routes
	aliasesGroup := r.Group("/aliases")
	aliasesGroup.Use(middleware.AuthRequired())
	{
		aliasesGroup.POST("/", aliasHandler.CreateAlias)
		aliasesGroup.POST("/toggle/:id", aliasHandler.ToggleActivateStatus)
		// One for getting a user's aliases given their sessionID => userID
		aliasesGroup.GET("/", aliasHandler.GetUsersAliasesProtected)

		protectedAliasesGroup := aliasesGroup.Group("/admin")
		protectedAliasesGroup.Use(middleware.RoleRequired(2))
		protectedAliasesGroup.GET("", aliasHandler.GetAliases)
		protectedAliasesGroup.GET("/:id", aliasHandler.GetAliasByID)
		protectedAliasesGroup.GET("/user/:userID", aliasHandler.GetUsersAliases)
	}

	// auth routes
	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/signup", authHandler.SignUp)
	r.POST("/auth/logout", authHandler.Logout)
}
