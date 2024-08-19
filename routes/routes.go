package routes

import (
	"DALE/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", controllers.PingHandler)
}