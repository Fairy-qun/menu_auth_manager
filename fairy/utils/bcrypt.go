package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func Enbcrypt(pwd string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(hash),err
}

// 密码解密
func Debcrypt(hashPwd string,pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd),[]byte(pwd))
	if err != nil {
		return false;
	}
	return true;
}