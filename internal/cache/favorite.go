package cache

import (
	"douyin_service/pkg/util"
	"strconv"
	"time"
)

// 缓存的favorite-crud

// FavorAction userId给videoId点赞
func (r *Redis) FavorAction(userId uint, videoId uint) error {
	key := util.UserFavorKey(userId)
	err := r.redis.SetBit(key, int64(videoId), 1).Err()
	return err
}

// CancelFavorAction CancelFavor userId给videoId取消点赞
func (r *Redis) CancelFavorAction(userId uint, videoId uint) error {
	key := util.UserFavorKey(userId)
	err := r.redis.SetBit(key, int64(videoId), 0).Err()
	return err
}

// QueryFavorCnt 查询某个视频点赞数量
func (r *Redis) QueryFavorCnt(videoId uint) (bool, int64, error) {
	key := util.VideoFavorCntKey(videoId)
	exist, err := r.IsExist(key)
	if err != nil {
		return false, 0, err
	}
	if !exist {
		return false, 0, nil
	}
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return false, 0, err
	}
	cnt, _ := strconv.Atoi(result)
	return true, int64(cnt), nil
}

// QueryUserFavoritedCount 查询用户userId被赞总数量
func (r *Redis) QueryUserFavoritedCount(userId uint) (bool, int64, error) {
	key := util.UserFavoritedCntKey(userId)
	exist, err := r.IsExist(key)
	if err != nil {
		return false, 0, err
	}
	if !exist {
		return false, 0, nil
	}
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return false, 0, err
	}
	cnt, _ := strconv.Atoi(result)
	return true, int64(cnt), nil
}

// QueryUserFavoriteCount 查询用户userId点赞总数量
func (r *Redis) QueryUserFavoriteCount(userId uint) (bool, int64, error) {
	key := util.UserFavoriteCntKey(userId)
	exist, err := r.IsExist(key)
	if err != nil {
		return false, 0, err
	}
	if !exist {
		return false, 0, nil
	}
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return false, 0, err
	}
	cnt, _ := strconv.Atoi(result)
	return true, int64(cnt), nil
}

// IncrFavorCnt 增加video的点赞数量
func (r *Redis) IncrFavorCnt(videoId uint) int64 {
	key := util.VideoFavorCntKey(videoId)
	result := r.redis.Incr(key).Val()
	return result
}

// IncrUserFavoritedCnt 增加userId的被赞数量
func (r *Redis) IncrUserFavoritedCnt(userId uint) int64 {
	key := util.UserFavoritedCntKey(userId)
	result := r.redis.Incr(key).Val()
	return result
}

// IncrUserFavoriteCnt 增加userId的点赞数量
func (r *Redis) IncrUserFavoriteCnt(userId uint) int64 {
	key := util.UserFavoriteCntKey(userId)
	result := r.redis.Incr(key).Val()
	return result
}

// DecrFavorCnt 减少video的点赞数量
func (r *Redis) DecrFavorCnt(videoId uint) int64 {
	key := util.VideoFavorCntKey(videoId)
	result := r.redis.Decr(key).Val()
	return result
}

// DecrUserFavoritedCnt 减少userId的被赞数量
func (r *Redis) DecrUserFavoritedCnt(userId uint) int64 {
	key := util.UserFavoritedCntKey(userId)
	result := r.redis.Decr(key).Val()
	return result
}

// DecrUserFavoriteCnt 减少userId的点赞数量
func (r *Redis) DecrUserFavoriteCnt(userId uint) int64 {
	key := util.UserFavoriteCntKey(userId)
	result := r.redis.Decr(key).Val()
	return result
}

// IsFavor userId是否给videoId点赞
func (r *Redis) IsFavor(userId uint, videoId uint) (bool, error) {
	key := util.UserFavorKey(userId)
	result, err := r.redis.GetBit(key, int64(videoId)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// FavorList 获取userId点赞列表(userId点赞了什么)
func (r *Redis) FavorList(userId uint) ([]uint, error) {
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
func (r *Redis) FavoredList(videoId uint) ([]uint, error) {
	key := util.VideoFavoredKey(videoId)
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return nil, err
	}
	bitmap := []byte(result)
	favorList := util.GetListFromBitMap(bitmap)
	return favorList, nil
}

// ------一些常量------

const (
	VIDEO_EXPIRE = time.Hour * 48
)
