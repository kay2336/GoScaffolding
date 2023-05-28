package controller

import (
	"awesomeProject/dao"
	"awesomeProject/model/table"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//type UserController struct {
//}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user table.User
		if err := c.ShouldBind(&user); err != nil {
			// 返回response
			c.JSON(http.StatusOK, "bind user error")
			return
		}
		// 实例化service层 dao层对象
		userServ := service.NewUserServ()
		userDao := dao.NewUserDao()

		// 判断手机号码是否重复
		dbUser, err := userDao.FindUserByPhone(user.Phone)
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"err":  "already register",
				"data": dbUser,
			})
			return
		}

		// 调用service层的register
		err = userServ.Register(c, &user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": "register error",
			})
		}

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user table.User
		if err := c.ShouldBind(&user); err != nil {
			// 返回给前端错误信息
			c.JSON(http.StatusOK, "bind user error")
		} else {
			userServ := service.NewUserServ()

			// 调用service层的register
			err = userServ.Login(c, &user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": "login error",
				})
				return
			}
		}
	}
}
