package api

import (
	"todo_list/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(500, err)
	}
}
