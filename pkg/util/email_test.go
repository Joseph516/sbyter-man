package util

import "testing"

func TestSendVerifiedEmail(t *testing.T) {
	err := SendVerifiedEmail([]string{"xxxx@qq.com"}, 1, "127.0.0.1", "shkkjssas")
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyEmail(t *testing.T) {
	flag, err := VerifyEmail("127.0.0.1", "shkkjssas")
	if err != nil {
		t.Fatal(err)
	}
	if !flag {
		t.Fatal(flag)
	}
}


