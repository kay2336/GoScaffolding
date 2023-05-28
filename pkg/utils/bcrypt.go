package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd
// 加密明文密码，返回加密后的密码
func HashPwd(pwd string) (hashedPwdStr string, err error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println("hash pwd error")
		return
	}
	hashedPwdStr = string(hashedPwd)
	return hashedPwdStr, err
}

// CmpPwd
// 将bcrypt加密后的密码和明文密码相比较，相同返回true，否则返回false
func CmpPwd(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}
