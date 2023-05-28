package dao

import (
	"awesomeProject/model/table"
	"awesomeProject/sql"
	"gorm.io/gorm"
)

//错误处理交给调用的一方

// UserDao
// 为什么不用写变量名？
type UserDao struct {
	Db *gorm.DB
}

// NewUserDao
// 通过新建实例userDao，实现service层调用dao层的方法
func NewUserDao() *UserDao {
	return &UserDao{Db: sql.GetMysqlDB()}
}

// Register
// 新增用户至mysql
func (dao *UserDao) Register(user *table.User) (err error) {
	err = sql.GetMysqlDB().Create(&user).Error
	return
}

// FindUserByPhone
// 通过手机号码查询用户
func (dao *UserDao) FindUserByPhone(phone string) (user *table.User, err error) {
	err = dao.Db.Where("phone = ?", phone).First(&user).Error
	return
}
