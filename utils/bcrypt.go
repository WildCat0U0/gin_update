//package utils
//
//import (
//	"golang.org/x/crypto/bcrypt"
//	"log"
//)
//
//// BcryptMake ---加密密码
//func BcryptMake(pwd []byte) string {
//	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost) //加密处理
//	if err != nil {
//		log.Println(err)
//	}
//	return string(hash)
//}
//
//// BcryptMakeCheck ---验证密码
//func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
//	byteHash := []byte(hashedPwd) //将字符串转换成字节数组
//	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
//	if err != nil {
//		log.Println(err)
//		return false
//	} else {
//		return true
//	}
//}

package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
