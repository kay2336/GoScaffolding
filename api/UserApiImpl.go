package api

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
	return func(ctx *gin.Context) {
		var user table.User
		userSrv := service.NewUserSrv() // 实例化service层 dao层对象
		userDao := dao.NewUserDao()

		if err := ctx.ShouldBind(&user); err != nil {
			// 返回response
			ctx.JSON(http.StatusOK, "bind user error")
			return
		}

		// 判断手机号码是否重复
		dbUser, err := userDao.FindUserByPhone(user.Phone)
		if err != gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"err":  "already register",
				"data": dbUser,
			})
			return
		}

		// 调用service层的register
		err = userSrv.Register(ctx.Request.Context(), &user)
		if err != nil {
			//ctx.JSON(http.StatusOK, gin.H{
			//	"err": "register error",
			//})
		}

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user table.User
		userServ := service.NewUserSrv() // 调用service层的register

		if err := ctx.ShouldBind(&user); err != nil {
			// 返回给前端错误信息
			ctx.JSON(http.StatusOK, "bind user error")

		} else {
			err = userServ.Login(ctx.Request.Context(), &user)
			if err != nil {
				//ctx.JSON(http.StatusOK, gin.H{
				//	"err": "login error",
				//})
			}

		}
	}
}
