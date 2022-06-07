package dao

import (
	"douyin_service/internal/model"
	"time"

	"gorm.io/gorm"
)

// type PublishDao struct{}

// var publishDao *PublishDao

func (d *Dao) ListVideoByUserId(userId uint) ([]model.Video, error) {
	return model.Video{AuthorId: userId}.ListVideoByUserId(d.engine)
}

func (d *Dao) PublishVideo(authorId uint, playUrl, coverUrl, title string) error {
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

func (d *Dao) QueryVideoById(videoId uint) (model.Video, error) {
	var video model.Video
	video, err := video.QueryVideoById(videoId, d.engine)
	if err != nil {
		return video, err
	}
	return video, nil
}

func (d *Dao) QueryBatchVideoById(favorList []uint) ([]model.Video, error) {
	var video model.Video
	videos, err := video.QueryBatchVdieoById(favorList, d.engine)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// UpdateFavoriteCnt 更新video的facorite_count
func (d *Dao) UpdateFavoriteCnt(video model.Video) error {
	return video.UpdateFavoriteCnt(d.engine)
}

// UpdateCommentCnt 更新video的comment_count
func (d *Dao) UpdateCommentCnt(video model.Video) error {
	return video.UpdateCommentCnt(d.engine)
}

// QueryFavorCntById 根据Id查询点赞数量
func (d *Dao) QueryFavorCntById(videoId uint) (int64, error) {
	var video model.Video
	video.ID = videoId
	return video.QueryFavorCntById(d.engine)
}
