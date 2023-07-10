package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/guilherm5/controllers"
)

func Login(c *gin.Engine) {
	api := c.Group("api")
	api.POST("login", controllers.LoginUser)
}
