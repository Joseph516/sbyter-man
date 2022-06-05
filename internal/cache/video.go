package cache

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"encoding/json"
	"strconv"
)

//video 的缓存操作

// QueryVideoByVideoId 返回缓存中的video
// 返回（是否存在，缓存的key，key对应的值，err）
func (r *Redis) QueryVideoByVideoId(videoId uint) (bool, string, model.Video, error) {
	key := util.VideoKey(videoId)
	var video model.Video
	exist, err := r.IsExist(key)
	if err != nil {
		return false, "", video, err
	}
	if !exist {
		return false, key, video, nil
	}
	videoJson, err := r.redis.Get(key).Result()
	if err != nil {
		return false, "", video, err
	}
	err = json.Unmarshal([]byte(videoJson), &video)
	if err != nil {
		return false, "", video, err
	}
	return true, key, video, nil
}

func (r *Redis) QueryAuthorIdByVideoId(videoId uint) (bool, uint, error) {
	key := util.VideoAuthorKey(videoId)
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
	return true, uint(cnt), nil
}
