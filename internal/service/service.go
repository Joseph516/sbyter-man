package service

import (
	"context"
	"douyin_service/global"
	"douyin_service/internal/cache"
	"douyin_service/internal/dao"
	"douyin_service/internal/kafka"
)

type Service struct {
	ctx   context.Context
	dao   *dao.Dao
	redis *cache.Redis
	Kafka *kafka.Kafka
}

type ResponseCommon struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx:   ctx,
		dao:   dao.New(global.DBEngine),
		redis: cache.NewRedis(global.Rd),
		Kafka: kafka.NewKafka(global.Consumer, global.SyncProducer),
	}
	return svc
}
