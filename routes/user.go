package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
	"github.com/guilherm5/middleware"
)

func RouterUser(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())
	c.POST("/new-user", controllers.NewUser)
	api.GET("/hello", controllers.Hello)
}

func Login(c *gin.Engine) {
	c.POST("login", controllers.LoginUser)
}
