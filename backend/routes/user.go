package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func RInfoUser(c *gin.Engine) {
	//Adicionando groups para implementar middleware
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.PUT("/info-user", controllers.InfoUser) //Depois de criar o user, vamos rederecionar o user para esta pagina, onde vai pedir algumas infos
}

func PostUser(c *gin.Engine) {
	api := c.Group("api")
	api.POST("/new-user", controllers.NewUser)
}
