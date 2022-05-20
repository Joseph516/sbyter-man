package cache

import (
	"douyin_service/pkg/util"
	"strconv"
	"time"
)

// 缓存的favorite-crud

// Favor userId给videoId点赞
func (r *Redis) Favor(userId int64, videoId int64) (bool, error) {
	key := util.UserFavorKey(userId)
	err := r.redis.SetBit(key, int64(videoId), 1).Err()
	if err != nil {
		return false, err
	}
	key = util.VideoFavoredKey(videoId)
	err = r.redis.SetBit(key, int64(userId), 1).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// CancelFavor userId给videoId取消点赞
func (r *Redis) CancelFavor(userId int64, videoId int64) (bool, error) {
	key := util.UserFavorKey(userId)
	err := r.redis.SetBit(key, int64(videoId), 0).Err()
	if err != nil {
		return false, err
	}
	key = util.VideoFavoredKey(videoId)
	err = r.redis.SetBit(key, int64(userId), 0).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryFavorCnt 查询videoId的点赞数量
func (r *Redis) QueryFavorCnt(videoId int64) (int64, error) {
	key := util.VideoFavoredCntKey(videoId)
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return 0, err
	}
	cnt, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return int64(cnt), err
}

// IncrFavorCnt 增加userId的点赞数量
func (r *Redis) IncrFavorCnt(userId int64) (int64, error) {
	key := util.UserFavorCntKey(userId)
	//expire会更新key的过期时间，并且如果不存在的话 ‘不会’ 创建这个key
	err := r.redis.Expire(key, time.Hour).Err()
	if err != nil {
		return 0, err
	}
	cnt, err := r.redis.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	return int64(cnt), nil
}

// DecrFavorCnt  减少userId的点赞数量
func (r *Redis) DecrFavorCnt(userId int64) (int64, error) {
	key := util.UserFavorCntKey(userId)
	err := r.redis.Expire(key, time.Hour).Err()
	if err != nil {
		return 0, err
	}
	cnt, err := r.redis.Decr(key).Result()
	if err != nil {
		return 0, err
	}
	return int64(cnt), nil
}

// IncrFavoredCnt 增加videoId的被点赞数量
func (r *Redis) IncrFavoredCnt(videoId int64) (int64, error) {
	key := util.VideoFavoredCntKey(videoId)
	err := r.redis.Expire(key, time.Hour).Err()
	if err != nil {
		return 0, err
	}
	cnt, err := r.redis.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	return int64(cnt), nil
}

// DecrFavoredCnt 减少videoId的被点赞数量
func (r *Redis) DecrFavoredCnt(videoId int64) (int64, error) {
	key := util.VideoFavoredCntKey(videoId)
	err := r.redis.Expire(key, time.Hour).Err()
	if err != nil {
		return 0, err
	}
	cnt, err := r.redis.Decr(key).Result()
	if err != nil {
		return 0, err
	}
	return int64(cnt), nil
}

// IsFavor userId是否给videoId点赞
func (r *Redis) IsFavor(userId int64, videoId int64) (bool, error) {
	key := util.UserFavorKey(userId)
	result, err := r.redis.GetBit(key, int64(videoId)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// DeleteOneCache 清除一条指定key的缓存
func (r *Redis) DeleteOneCache(key string) (bool, error) {
	err := r.redis.Del(key).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// FavorList 获取userId点赞列表(userId点赞了什么)
func (r *Redis) FavorList(userId int64) ([]int64, error) {
	key := util.UserFavorKey(userId)
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return nil, err
	}
	bitmap := []byte(result)
	favorList := util.GetListFromBitMap(bitmap)
	return favorList, nil
}

// FavoredList 获取点赞videoId列表(谁点赞了videoId)
func (r *Redis) FavoredList(videoId int64) ([]int64, error) {
	key := util.VideoFavoredKey(videoId)
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return nil, err
	}
	bitmap := []byte(result)
	favorList := util.GetListFromBitMap(bitmap)
	return favorList, nil
}

// IsExist 是否存在名为key的缓存
func (r *Redis) IsExist(key string) (bool, error) {
	result, err := r.redis.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}
