package service

import (
	"context"
	"douyin_service/global"
	"douyin_service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

type ResponseCommon struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
	return svc
}
