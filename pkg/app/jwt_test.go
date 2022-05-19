package app

import (
	"fmt"
	"log"
	"testing"
)

func init()  {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("douyin", "douyin", "1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token, err := GenerateToken("douyin", "douyin", "2")
	if err != nil {
		t.Error(err)
	}
	claim, _ := ParseToken(token)
	if claim.Audience != "2" {
		t.Error("解析失败")
	}
}