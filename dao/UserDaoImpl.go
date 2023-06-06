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
	db *gorm.DB
}

// NewUserDao
// 通过新建实例userDao，实现service层调用dao层的方法
func NewUserDao() *UserDao {
	return &UserDao{sql.GetMysqlDB()}
}

// Register
// 新增用户至mysql
func (dao *UserDao) Register(user *table.User) (err error) {
	err = dao.db.Create(&user).Error
	return
}

// FindUserByPhone
// 通过手机号码查询用户
func (dao *UserDao) FindUserByPhone(phone string) (user *table.User, err error) {
	err = dao.db.Where("phone = ?", phone).First(&user).Error
	return
}

// FindUserById
// 根据用户id找到用户
func (dao *UserDao) FindUserById(id uint) (user *table.User, err error) {
	err = dao.db.Model(&table.User{}).Where("id=?", id).
		First(&user).Error

	return
}
