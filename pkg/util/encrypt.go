package util

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// EncodeMD5
// encode string into md5
// 加密字符串为 md5码
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeBcrypt 使用Bcrypt加密字符串
func EncodeBcrypt(value string) (string, error)  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(bytes), err
}

//CheckBcrypt 对比密码是否和加密字符串对应
func CheckBcrypt(hashPassword, password string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}