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
}

func Post(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())
	api.POST("/post", controllers.Post)
}

func Login(c *gin.Engine) {
	c.POST("login", controllers.LoginUser)
}
