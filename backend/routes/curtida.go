package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Curtida(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.POST("/curtida", controllers.NewCurtida)
	api.POST("/curte-comentario", controllers.CurteComment)
	api.POST("/curte-resp-comentario", controllers.CurtRespComment)

	//DELETE CURTIDAS
	api.DELETE("/curtida", controllers.DelCurtida)
	api.DELETE("/del-curtida-comentario", controllers.DelCurtidaComentario)
	api.DELETE("/del-curtida-resp-comentario", controllers.DelCurtidaRespComentario)
}
