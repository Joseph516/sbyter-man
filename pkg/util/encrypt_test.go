package util

import (
	"fmt"
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	s := "test"
	md5Str := EncodeMD5(s)
	fmt.Println(md5Str)
}

func TestEncodeBcrypt(t *testing.T) {
	s := "test"
	bcryptStr, err := EncodeBcrypt(s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bcryptStr)
}

func TestCheckBcrypt(t *testing.T) {
	s := "test"
	hash := "$2a$14$oj14a3UJ98ZX6VcBdF18euqbX36nJ59QbUxhWnhTl9Q5s2XGEXk5C"
	ok := CheckBcrypt(hash, s)
	if !ok {
		t.Fatal("密码和密文不匹配")
	}
	fmt.Println(ok)
}