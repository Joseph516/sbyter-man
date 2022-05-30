package service

import (
	"douyin_service/global"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestConsumeEmail(t *testing.T) {
	var err error
	global.Consumer, err = NewConsumer()
	if err != nil {
		t.Fatal(err)
	}
	svc := New(&gin.Context{})
	svc.ConsumeEmail()
}
