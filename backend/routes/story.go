package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func Story(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())
	api.POST("/story", controllers.NewStory)
	api.POST("/get-story", controllers.GetStory)
	api.GET("/story", controllers.GetStoryUser)
}
