package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func RouterUser(c *gin.Engine) {
	//Adicionando groups para implementar middleware
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	c.POST("/new-user", controllers.NewUser)
	api.PUT("/info-user", controllers.InfoUser) //Depois de criar o user, vamos rederecionar o user para esta pagina, onde vai pedir algumas infos
}
