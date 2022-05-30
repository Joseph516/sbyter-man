package middleware

import "github.com/Shopify/sarama"

type Kafka struct {
	Consumer      sarama.Consumer
	SyncProducer sarama.SyncProducer
}

func NewKafka(consumer sarama.Consumer, syncProducer sarama.SyncProducer) *Kafka {
	return &Kafka{Consumer: consumer, SyncProducer: syncProducer}
}
