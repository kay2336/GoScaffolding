package table

type User struct {
	Id       uint   `gorm:"primaryKey;unique;not null;autoIncrement"`
	Phone    string `varchar(32);comment:"手机号码" gorm:"phone,unique"`
	Username string `varchar(32);comment:"用户名"`
	Password string `varchar(32);comment:"密码"`
}
