package producer

import (
	"douyin_service/global"
	"log"

	"github.com/Shopify/sarama"
)

// Producer 生产者方法
func Producer(topic string, message string, limit int) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{global.KafkaSetting.Host}, config)
	if err != nil {
		log.Fatalln("NewSyncProducer err: ", err)
		return
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(message),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, message)
	}
}
