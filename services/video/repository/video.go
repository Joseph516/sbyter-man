package repository

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId      uint      `json:"author_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	Title         string    `json:"title"`
	PublishDate   time.Time `json:"publish_date"`
}

func (v Video) TableName() string { return "douyin_video" }

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (d *VideoDao) ListVideoByUserId(userId uint) ([]Video, error) {
	var video []Video
	err := db.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title").Where("author_id = ?",
		userId).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (d *VideoDao) PublishVideo(authorId uint, playUrl, coverUrl, title string) error {
	now := time.Now()
	err := db.Create(&Video{
		Model: gorm.Model{
			CreatedAt: now,
		},
		AuthorId:      authorId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		PublishDate:   now,
	}).Error
	return err
}

// FeedVideo 返回距离上次查询的最新视频，不区分用户
func (d *VideoDao) FeedVideo(lastTime int64, cnt int) ([]Video, error) {
	var video []Video
	var tm time.Time
	if lastTime == 0 {
		tm = time.Now()
	} else {
		tm = time.Unix(lastTime, 0) // 时间戳 to 时间类对象
	}
	err := db.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title, publish_date").Where("publish_date <= ?",
		tm).Find(&video).Order("publish_date DESC").Limit(cnt).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}
