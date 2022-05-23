package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId  int64  `json:"user_id"`
	VideoId int64  `json:"video_id"`
	Content string `json:"content"`
}

func (c Comment) TableName() string {
	return "douyin_comment"
}

func (c Comment) ListCommentByVideoId(db *gorm.DB) ([]Comment, error) {
	var comments []Comment
	err := db.Where("video_id = ?", c.VideoId).Find(&comments).Error
	return comments, err
}

func (c Comment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
func (c Comment) Delete(db *gorm.DB) error {
	return db.Where("id = ?", c.ID).Delete(&c).Error
}
