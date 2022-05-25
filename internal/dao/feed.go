package dao

import (
	"douyin_service/internal/model"
	"time"
)

// 根据latest_time，获取最新的20条视频
func (d *Dao) GetLatestVideos(lastTime int64) ([]model.Video, error) {
	var video []model.Video
	var tm time.Time
	if lastTime == 0 {
		tm = time.Now()
	} else {
		tm = time.Unix(lastTime, 0) // 时间戳 to 时间类对象
	}
	err := d.engine.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title").Where("publish_date <= ?",
		tm).Find(&video).Order("publish_date DESC").Limit(20).Error

	if err != nil {
		return nil, err
	}
	return video, nil
}
