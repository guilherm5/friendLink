package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/guilherm5/routes"
)

func main() {
	api := gin.Default()

	routes.RouterUser(api)
	routes.Login(api)

	api.Run(":8080")
}
