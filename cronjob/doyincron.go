package cronjob

import (
	"douyin_service/global"
	"douyin_service/internal/cache"
	"douyin_service/internal/dao"
)

type DouyinCron struct {
	dao   *dao.Dao
	redis *cache.Redis
}

func New() DouyinCron {
	return DouyinCron{
		dao:   dao.New(global.DBEngine),
		redis: cache.NewRedis(global.Rd),
	}
}
