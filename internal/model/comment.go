package model

import "gorm.io/gorm"

type Comment struct {
	*Model
	UserId  int64  `json:"user_id"`
	VideoId int64  `json:"video_id"`
	Content string `json:"content"`
}

func (c Comment) TableName() string {
	return "douyin_user"
}

func (c Comment) ListCommentByVideoId(db *gorm.DB) ([]Comment, error) {
	var comments []Comment
	err := db.Select("user_id, content, created_at").Where("video_id = ?", c.VideoId).Find(&comments).Error
	return comments, err
}
