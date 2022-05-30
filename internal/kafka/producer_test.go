package kafka

import (
	"douyin_service/global"
	"douyin_service/internal/model/message"
	"douyin_service/internal/service"
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
	kafka := NewKafka(global.Consumer, global.SyncProducer)
	kafka.Producer(global.KafkaSetting.TopicEmail, email.String(), 1)
}
