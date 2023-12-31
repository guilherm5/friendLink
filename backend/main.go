package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	routes "github.com/guilherm5/routes"
)

func main() {
	api := gin.Default()

	api.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//Ligando package routes
	routes.LInfoUser(api)
	routes.Post(api)
	routes.Login(api)
	routes.PostUser(api)
	routes.Comentario(api)
	routes.Curtida(api)
	routes.Story(api)
	routes.Seguidor(api)

	go controllers.RunCronJobs()

	api.Run(":8080")

}
