package controller

import (
	"awesomeProject/model/table"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type UserController struct {
//}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user table.User
		if err := c.ShouldBind(&user); err != nil {
			// 返回给前端错误信息
			c.JSON(http.StatusOK, "bind error")

		} else {
			//log.Println(user, "controller")
			userServ := service.NewUserServ()
			// 校验参数是否正确（判断手机号码是否重复）

			// 调用service层的register
			err = userServ.Register(c, &user)

			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": "controllerRegister OK",
			})
		}
		return
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user table.User
		if err := c.ShouldBind(&user); err != nil {
			// 返回给前端错误信息
			c.JSON(http.StatusOK, "bind error")
		} else {
			userServ := service.NewUserServ()
			// 校验参数是否正确（判断手机号码是否重复）

			// 调用service层的register
			err = userServ.Login(c, &user)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": "controllerLogin OK",
			})
		}
	}
}
