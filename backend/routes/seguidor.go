package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Seguidor(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.POST("/seguir", controllers.Seguir)
	api.DELETE("/seguir", controllers.Desseguir)
}
