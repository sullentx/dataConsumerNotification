package main

import (
	infraestructure "api-consumer/Notification/Infrastructure"
	"api-consumer/Notification/Infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	infraestructure.Init()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	router.Use(cors.New(config))

	routes.SetRoutes(router, infraestructure.PostNotificationHandler, infraestructure.GetAllNotificationHandler)
	router.Run(":8082")
}
