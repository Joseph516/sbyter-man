package dao

import (
	"douyin_service/internal/model"
	"time"

	"gorm.io/gorm"
)

func (d *Dao) ListCommentByVideoId(videoId int64) ([]model.Comment, error) {
	return model.Comment{VideoId: videoId}.ListCommentByVideoId(d.engine)
}

func (d *Dao) PublishComment(videoId, userId int64, content string) error {
	now := time.Now()
	return model.Comment{
		Model: gorm.Model{
			CreatedAt: now,
		},
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}.Create(d.engine)
}
func (d *Dao) DeleteComment(CommentId uint) error {
	return model.Comment{
		Model: gorm.Model{
			ID: CommentId,
		},
	}.Delete(d.engine)
}
