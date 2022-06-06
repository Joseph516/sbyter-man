package cronjob

import (
	"douyin_service/global"
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"strconv"
)

const (
	//"0 30 2 1/2 * ?"
	FlashFavorCnt         = "0 30 2 1/2 * ?" //每两日的凌晨2点半执行一次
	FlashUserFavoritedCnt = "0 30 3 1/2 * ?" //每两日的凌晨3点半执行一次
	FlashUserFavoriteCnt  = "0 0 4 1/2 * ?"
)

// FlashFavorCnt 定时把视频点赞数量刷新到mysql中，并在此时清除redis缓存
func (dc DouyinCron) FlashFavorCnt() {
	pattern := util.VIDEOFAVORCNTPRIFIX + "*"
	keys, err := dc.redis.GetCacheKeys(pattern)
	if err != nil {
		global.Logger.Error(err)
		return
	}

	for _, key := range keys {
		//eg. VC2截取获得videoId = 2
		vId, err := strconv.Atoi(key[len(util.VIDEOFAVORCNTPRIFIX):])
		videoId := uint(vId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Infof("正在获取%s并启动刷新", key)
		//不要判断exist，能从缓存中获取key说明肯定是存在的
		_, cnt, err := dc.redis.QueryFavorCnt(videoId)
		if err != nil {
			global.Logger.Error(err)
			return
		}

		var video model.Video
		video.ID = videoId
		video.FavoriteCount = cnt

		//写回数据库
		err = dc.dao.UpdateFavoriteCnt(video)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Infof("刷新成功，目前数据库值为%d", cnt)
		//写回后删除对应缓存
		//刷数据库后，再检查一次缓存是否被更新，如果被更新了，那么不能删除缓存，否则会丢失新的更新
		_, cnt2, err := dc.redis.QueryFavorCnt(videoId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		if cnt2 != cnt {
			global.Logger.Info("缓存被更新，未删除")
		} else {
			_, err = dc.redis.DeleteOneCache(key)
			if err != nil {
				global.Logger.Error(err)
				return
			}
			global.Logger.Info("删除对应缓存")
		}
	}
	global.Logger.Info("刷新完成")
}

// FlashUserFavoritedCnt 定时刷新用户获赞数量
// 总体逻辑：
// 1.获取缓存中需要刷新的key
// 2.将key对应的value刷新到数据库
// 3.检查此时缓存中的数据是否改变(两次缓存查询之间间隔了一个数据库操作)，如果改变了则缓存不能被删除！
func (dc DouyinCron) FlashUserFavoritedCnt() {
	pattern := util.USERTOTALFAVORITEDCNT + "*"
	favoritedKey, err := dc.redis.GetCacheKeys(pattern)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	for index, _ := range favoritedKey {
		uId, err := strconv.Atoi(favoritedKey[index][len(util.USERTOTALFAVORITEDCNT):])
		userId := uint(uId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Infof("正在获取%s并启动刷新", favoritedKey[index])
		//第一次查缓存
		_, favoritedCnt, err := dc.redis.QueryUserFavoritedCount(userId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		//写回数据库
		err = dc.dao.UpdateById(userId,
			map[string]interface{}{"total_favorited": favoritedCnt})
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Info("刷新成功")
		//第二次查缓存
		_, favoritedCnt2, err := dc.redis.QueryUserFavoritedCount(userId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		if favoritedCnt2 != favoritedCnt {
			global.Logger.Info("缓存被更新，未删除")
		} else {
			_, err = dc.redis.DeleteOneCache(favoritedKey[index])
			if err != nil {
				global.Logger.Error(err)
				return
			}
			global.Logger.Info("删除对应缓存")
		}
	}
	global.Logger.Info("FlashUserFavoritedCnt完成")
}

func (dc DouyinCron) FlashUserFavoriteCnt() {
	pattern := util.USERTOTALFAVORITECNT + "*"
	favoriteKey, err := dc.redis.GetCacheKeys(pattern)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	for index, _ := range favoriteKey {
		uId, err := strconv.Atoi(favoriteKey[index][len(util.USERTOTALFAVORITECNT):])
		userId := uint(uId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Infof("正在获取%s并启动刷新", favoriteKey[index])
		//第一次查缓存
		_, favoriteCnt, err := dc.redis.QueryUserFavoriteCount(userId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		//写回数据库
		err = dc.dao.UpdateById(userId,
			map[string]interface{}{"favorite_count": favoriteCnt})
		if err != nil {
			global.Logger.Error(err)
			return
		}
		global.Logger.Info("刷新成功")
		//第二次查缓存
		_, favoriteCnt2, err := dc.redis.QueryUserFavoriteCount(userId)
		if err != nil {
			global.Logger.Error(err)
			return
		}
		if favoriteCnt2 != favoriteCnt {
			global.Logger.Info("缓存被更新，未删除")
		} else {
			_, err = dc.redis.DeleteOneCache(favoriteKey[index])
			if err != nil {
				global.Logger.Error(err)
				return
			}
			global.Logger.Info("删除对应缓存")
		}
	}
	global.Logger.Info("FlashUserFavoriteCnt完成")
}
