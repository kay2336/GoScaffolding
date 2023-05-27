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

	// 查询是否有重复账号
	//user, err := userDao.FindUserByPhone(registerUserForm.Phone)

	// bcrypt密码加密存储
	hashedPwd, err := utils.HashPwd(user.Password)
	if err != nil {
		log.Println("hash pwd error")
		return
	}

	user.Password = hashedPwd
	if err = userDao.Register(user); err != nil {
		log.Println(err)
		return
	}
	//log.Println(user, "service")

	c.JSON(http.StatusOK, gin.H{
		"data": "serviceRegister OK",
		"user": user,
	})
	return
}

func (serv *UserServ) Login(c *gin.Context, user *table.User) (err error) {
	// 创建dao层对象，使用dao层方法
	userDao := dao.NewUserDao()
	_, err = userDao.FindUserByPhone(user.Phone)

	if err == gorm.ErrRecordNotFound {
		//err = errors.New("用户不存在")
		c.JSON(http.StatusOK, gin.H{
			"data": "no phoneNumber",
		})
		return
	}

	// 校验手机号和密码是否正确
	//if !user.CheckPassword(req.Password) {
	//	err = errors.New("账号/密码错误")
	//	utils.LogrusObj.Info(err)
	//	return
	//}

	// 获取token
	//token, err := utils.GenerateToken(user.ID, req.UserName, 0)
	//if err != nil {
	//	utils.LogrusObj.Info(err)
	//	return
	//}

	//给api层返回状态信息

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}
