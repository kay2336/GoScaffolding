package middleware

import (
	"awesomeProject/pkg/e"
	"awesomeProject/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT
// token鉴权
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		// 解析前端发送的请求中的token
		token := c.Request.Header.Get("token")

		// 无token
		if token == "" {
			code = http.StatusNotFound
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseTokenHs256(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}

		// 给请求重新赋值干嘛？？
		//c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		c.Next()
	}
}
