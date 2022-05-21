package dao

import (
	"douyin_service/internal/model"
	"gorm.io/gorm"
	"time"
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

func (d *Dao) QueryVideoById(videoId int64) (model.Video, error) {
	var video model.Video
	video, err := video.QueryVideoById(videoId, d.engine)
	if err != nil {
		return video, err
	}
	return video, nil
}

func (d *Dao) QueryBatchVideoById(favorList []int64) ([]model.Video, error) {
	var video model.Video
	videos, err := video.QueryBatchVdieoById(favorList, d.engine)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (d *Dao) UpdatesVideo(video model.Video) error {
	return video.UpdatesVideo(d.engine)
}
