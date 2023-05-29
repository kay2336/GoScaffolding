package service

import (
	"awesomeProject/dao"
	"awesomeProject/model/table"
	"awesomeProject/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type UserServ struct {
}

func NewUserServ() *UserServ {
	return &UserServ{}
}

func (serv *UserServ) Register(c *gin.Context, user *table.User) (err error) {
	// 创建dao层对象，使用dao层方法
	userDao := dao.NewUserDao()

	// bcrypt密码加密存储
	hashedPwd, err := utils.HashPwd(user.Password)
	if err != nil {
		log.Println("hash pwd error")
		return
	}

	// 向mysql中添加用户
	user.Password = hashedPwd
	if err = userDao.Register(user); err != nil {
		log.Println(err)
		return
	}

	// 返回response
	c.JSON(http.StatusOK, gin.H{
		"data": "serviceRegister OK",
		"user": user,
	})
	return
}

func (serv *UserServ) Login(c *gin.Context, user *table.User) (err error) {
	// 创建dao层对象，使用dao层方法
	userDao := dao.NewUserDao()

	// mysql中无此手机号码的用户
	dbUser, err := userDao.FindUserByPhone(user.Phone)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"data": "no phoneNumber",
		})
		return
	}

	// 校验手机号和密码是否正确
	if !utils.CmpPwd(dbUser.Password, user.Password) {
		c.JSON(http.StatusOK, gin.H{
			"data": "密码错误",
		})
		return
	}

	//获取token
	token, err := utils.GenerateTokenUsingHS256(user.ID, user.Username)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	//返回response
	c.JSON(http.StatusOK, gin.H{
		"data":  "login success",
		"token": token,
		"user":  user,
	})
	return
}
