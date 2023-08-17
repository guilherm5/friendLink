package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Comentario(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.POST("/comentario", controllers.NewComment)
	api.POST("/list-comentario", controllers.GetComment)
	api.POST("/resp-comentario", controllers.RespComment)
	api.POST("/list-resp-comentario", controllers.ListRespComment)
}
