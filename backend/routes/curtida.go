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
	api.DELETE("/curtida", controllers.DelCurtida)

}
