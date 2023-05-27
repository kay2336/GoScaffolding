package user

type RegisterUserForm struct {
	Phone      string `varchar(32);comment:"手机号码"`
	Username   string `varchar(32);comment:"用户名"`
	Password   string `varchar(32);comment:"密码"`
	RePassword string `varchar(32);comment:"重复密码"`
}

type LoginUserForm struct {
	Phone    string `varchar(32);comment:"手机号码"`
	Password string `varchar(32);comment:"密码"`
}
