package middleware

import (
	"time"
	"todo_list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		code = 200
		//	var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400
			msg = "没有检查到登录状态,请先登录"
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 400
				msg = "登录超时,请重新登录"
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 400
				msg = "登录已过期,请重新登录"
			}

		}

		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    msg,
			})
			c.Abort()
			return
		}

		c.Next()
	}

}
