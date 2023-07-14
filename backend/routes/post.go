package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Post(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())
	//post Post(conteudo no feed)
	api.POST("/post", controllers.Post)
	api.POST("/list-post", controllers.Pagination)
}
