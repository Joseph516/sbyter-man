package cronjob

import (
	"douyin_service/global"
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"strconv"
)

const (
	FAVORCNTTIME = "0 30 2 1/2 * ?" //每两日的凌晨2点半执行一次
)

// FlashFavorCnt 定时把视频点赞数量刷新到mysql中，并在此时清除redis缓存
func (dc DouyinCron) FlashFavorCnt() {
	pattern := util.VIDEOFAVORCNTPRIFIX + "*"
	keys, err := dc.redis.GetCacheKeys(pattern)
	if err != nil {
		global.Logger.Error(err)
	}

	for _, key := range keys {
		//eg. VC2截取获得videoId = 2
		videoId, err := strconv.Atoi(key[2:])
		if err != nil {
			global.Logger.Error(err)
		}
		global.Logger.Infof("正在获取%s并启动刷新", key)
		_, cnt, err := dc.redis.QueryFavorCnt(int64(videoId))
		if err != nil {
			global.Logger.Error(err)
		}

		var video model.Video
		video.ID = uint(videoId)
		video.FavoriteCount = cnt

		//写回数据库
		err = dc.dao.UpdateFavoriteCnt(video)
		if err != nil {
			global.Logger.Error(err)
		}
		global.Logger.Infof("刷新成功，目前数据库值为%d", cnt)
		//写回后删除对应缓存
		//刷数据库后，再检查一次缓存是否被更新，如果被更新了，那么不能删除缓存，否则会丢失新的更新
		_, cnt2, err := dc.redis.QueryFavorCnt(int64(videoId))
		if err != nil {
			global.Logger.Error(err)
		}
		if cnt2 != cnt {
			global.Logger.Info("缓存被更新，未删除")
			return
		}
		_, err = dc.redis.DeleteOneCache(key)
		if err != nil {
			global.Logger.Error(err)
		}
		global.Logger.Info("删除对应缓存")
	}
	global.Logger.Info("刷新完成")
}
