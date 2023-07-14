package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func LInfoUser(c *gin.Engine) {
	//Adicionando groups para implementar middleware
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.PUT("/info-user", controllers.AfterCad) //Depois de criar o user, vamos rederecionar o user para esta pagina, onde vai pedir algumas infos
	api.GET("/info-user", controllers.GetInfoUser)
}

func PostUser(c *gin.Engine) {
	api := c.Group("api")
	api.POST("/new-user", controllers.NewUser)
}
