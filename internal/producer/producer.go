package producer

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"github.com/Shopify/sarama"
	"log"
)

func NewSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{global.KafkaSetting.Host}, config)
	if err != nil {
		log.Fatalln("NewConsumer err: ", err)
		return nil, err
	}
	return producer, nil
}

// Producer 生产者方法
func Producer(svc service.Service, topic string, message string, limit int)  {
	for i := 0; i < limit; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(message),
		}
		partition, offset, err := svc.Kafka.SyncProducer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, message)
	}
}
