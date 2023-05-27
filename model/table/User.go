package table

import "gorm.io/gorm"

type User struct {
	//Id string `gorm:"unique" ;comment:"主键"`
	//UId string ``
	//Email    string `varchar(32);comment:"邮箱"`
	gorm.Model
	Phone    string `varchar(32);comment:"手机号码"`
	Username string `varchar(32);comment:"用户名"`
	Password string `varchar(32);comment:"密码"`
}
