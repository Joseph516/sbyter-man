package service

import (
	"douyin_service/global"
	"douyin_service/internal/model/message"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestProducer(t *testing.T) {
	email := message.Email{
		UserName: []string{"2664006323@qq.com"},
		UserId:   1,
		LoginIP:  "127.0.0.1",
		Token:    "123",
	}
	svc := New(&gin.Context{})
	svc.Producer(global.KafkaSetting.TopicEmail, email.String(), 1)
}
