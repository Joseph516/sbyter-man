package dao

import (
	"douyin_service/internal/model"
	"time"

	"gorm.io/gorm"
)

// type PublishDao struct{}

// var publishDao *PublishDao

func (d *Dao) ListVideoByUserId(userId uint) ([]model.Video, error) {
	return model.Video{AuthorId: int64(userId)}.ListVideoByUserId(d.engine)
}

func (d *Dao) PublishVideo(authorId int64, playUrl, coverUrl, title string) error {
	now := time.Now()
	return model.Video{
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
	}.Create(d.engine)
}
