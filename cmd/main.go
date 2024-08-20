package main

import (
	"DALE/config"
	"DALE/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig() // Load .env
    config.InitializeDB()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8000")
}
