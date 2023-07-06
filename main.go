package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "github.com/guilherm5/routes"
)

func main() {
	api := gin.Default()

	//Liberando cors para o front end
	api.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//Ligando package routes
	routes.RouterUser(api)
	routes.Post(api)
	routes.Login(api)

	api.Run(":8080")
}
