package table

import "gorm.io/gorm"

type User struct {
	//Id string `gorm:"unique" ;comment:"主键"`
	//Email    string `varchar(32);comment:"邮箱"`
	gorm.Model
	Phone    string `varchar(32);comment:"手机号码" gorm:"phone,unique" json:"phone"`
	Username string `varchar(32);comment:"用户名" gorm:"username" json:"username"`
	Password string `varchar(32);comment:"密码" gorm:"password" json:"password"`
}
