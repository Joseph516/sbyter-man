package service

import (
	"douyin_service/global"
	"douyin_service/internal/model/message"
	"douyin_service/pkg/util"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

// NewConsumer 创建消费者实例
func NewConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{global.KafkaSetting.Host}, config)
	if err != nil {
		log.Fatalln("NewConsumer err: ", err)
		return nil, err
	}
	return consumer, nil
}


// ConsumeEmail 消费邮件
func (svc *Service) ConsumeEmail() {
	consumer := svc.Kafka.Consumer
	partitionConsumer, err := consumer.ConsumePartition(global.KafkaSetting.TopicEmail, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()
	for msg := range partitionConsumer.Messages() {
		data := message.Email{}
		_ = json.Unmarshal(msg.Value, &data)
		err := util.SendVerifiedEmail(data.UserName, data.UserId, data.LoginIP, data.Token) // 消费kafka消息
		if err != nil {
			global.Logger.Errorf("util.SendVerifiedEmail: %v", err)
		}
	}
}