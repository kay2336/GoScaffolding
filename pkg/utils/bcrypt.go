package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd
// 加密明文密码，返回加密后的密码
func HashPwd(pwdStr string) (hashedPwd string, err error) {
	pwd := []byte(pwdStr)
	hashedPwdStr, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println("hash pwd error")
		return
	}
	hashedPwd = string(hashedPwdStr)
	return hashedPwd, err
}

// CmpPwd
// 将bcrypt加密后的密码和明文密码相比较，相同返回true，否则返回false
func CmpPwd(hashedPwdStr string, plainPwdStr string) bool {
	hashedPwd := []byte(hashedPwdStr)
	plainPwd := []byte(plainPwdStr)
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return false
	}
	return true
}
