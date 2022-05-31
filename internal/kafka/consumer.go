package kafka

import (
	"douyin_service/global"
	"douyin_service/internal/model/message"
	"douyin_service/pkg/util"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
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
func (k *Kafka) ConsumeEmail() {
	consumer := k.Consumer
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

// ConsumComment 消费评论
func (k *Kafka) ConsumComment() {
	consumer := k.Consumer
	//OffsetNewest从最新的开始消费，即该 consumer 启动,之前产生的消息都无法被消费
	partitionConsumer, err := consumer.ConsumePartition(global.KafkaSetting.TopicComment, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("ConsumComment:ConsumePartition err: ", err)
		return
	}
	defer partitionConsumer.Close()
	for msg := range partitionConsumer.Messages() {
		param := CommentActionRequest{}
		_ = json.Unmarshal(msg.Value, &param)
		//创建一个空的 context，一般用于未确定时的声明使用
		err = param.CommentAction()
		if err != nil {
			global.Logger.Errorf("svc.CommentAction err: %v", err)
		}
	}
}
