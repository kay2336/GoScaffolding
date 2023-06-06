package service

import (
	"awesomeProject/dao"
	"awesomeProject/model/table"
	"awesomeProject/pkg/utils"
	"context"
	"gorm.io/gorm"
	"log"
)

type UserSrv struct {
}

// NewUserSrv
// 实例化UserSrv
func NewUserSrv() *UserSrv {
	return &UserSrv{}
}

func (serv *UserSrv) Register(ctx context.Context, user *table.User) (err error) {
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
	//ctx.JSON(http.StatusOK, gin.H{
	//	"data": "serviceRegister OK",
	//	"user": user,
	//})
	return
}

func (serv *UserSrv) Login(ctx context.Context, user *table.User) (err error) {
	// 创建dao层对象，使用dao层方法
	userDao := dao.NewUserDao()

	// mysql中无此手机号码的用户
	dbUser, err := userDao.FindUserByPhone(user.Phone)
	if err == gorm.ErrRecordNotFound {
		//c.JSON(http.StatusOK, gin.H{
		//	"data": "no phoneNumber",
		//})
		return
	}

	// 校验手机号和密码是否正确
	if !utils.CmpPwd(dbUser.Password, user.Password) {
		//c.JSON(http.StatusOK, gin.H{
		//	"data": "密码错误",
		//})
		return
	}

	//获取token
	token, err := utils.GenerateTokenUsingHS256(dbUser.Id, dbUser.Username, 0)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	//返回response

	//log.Println(user.Id)
	log.Println(token)
	return
}
