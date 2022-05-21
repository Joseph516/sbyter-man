package model

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64     `json:"author_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	Title         string    `json:"title"`
	PublishDate   time.Time `json:"publish_date"`
}

func (v Video) TableName() string { return "douyin_video" }

func (v Video) ListVideoByUserId(db *gorm.DB) ([]Video, error) {
	var video []Video
	err := db.Select("id, play_url, cover_url, favorite_count, comment_count, title").Where("author_id = ?",
		v.AuthorId).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v Video) Create(db *gorm.DB) error {
	return db.Create(&v).Error
}

// UpdatesVideo 修改Video信息
func (v Video) UpdatesVideo(db *gorm.DB) error {
	// gorm 不会更新零值，通过Select强制更新0值，避免点赞数量或者评论数量为0而不修改
	// Model 里的参数主键不为0，将默认产生一个where id = 主键的条件（update必须加条件,gorm禁止全局更新）
	return db.Model(&v).
		Select("AuthorId", "PlayUrl", "CoverUrl", "FavoriteCount", "CommentCount", "Title").
		Updates(v).Error
}

// QueryVideoById 根据videoId查询video信息
func (v Video) QueryVideoById(favor int64, db *gorm.DB) (Video, error) {
	var video Video
	err := db.Select("id, author_id,play_url, cover_url, favorite_count, comment_count, title").Where("id = ?", favor).
		Find(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

// QueryBatchVdieoById  根据videoId批量查询video信息
func (v Video) QueryBatchVdieoById(favorList []int64, db *gorm.DB) ([]Video, error) {
	var videos []Video
	err := db.Select("id, play_url, cover_url, favorite_count, comment_count, title").Where("id IN ?", favorList).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
