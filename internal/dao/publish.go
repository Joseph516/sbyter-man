package dao

import (
	"douyin_service/internal/model"
	"strconv"
	"strings"
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

func (d *Dao) PublishVideoWithTag(authorId uint, playUrl, coverUrl, title string, tags []string) error {
	now := time.Now()

	return d.engine.Transaction(func(tx *gorm.DB) error {
		// 创建标签记录项
		tagsModel := make([]model.Tag, 0, len(tags))
		for _, tag := range tags {
			tagsModel = append(tagsModel, model.Tag{Name: tag})
		}
		err := model.Tag{}.FirstOrCreateBatch(tx, tagsModel)
		if err != nil {
			return err
		}

		// 获取标签对应的编号id
		tagIds := []string{}
		for _, tag := range tagsModel {
			tagIds = append(tagIds, strconv.Itoa(int(tag.ID)))
		}

		// 创建视频记录项
		video := model.Video{
			Model: gorm.Model{
				CreatedAt: now,
			},
			AuthorId:      authorId,
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: 0,
			CommentCount:  0,
			Title:         title,
			Tags:          strings.Join(tagIds, ";"), // 逗号分隔标签
			PublishDate:   now,
		}
		if err := video.Create(tx); err != nil {
			return err
		}

		return nil
	})
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

// QueryFavorCntById 根据Id查询点赞数量
func (d *Dao) QueryFavorCntById(videoId uint) (int64, error) {
	var video model.Video
	video.ID = uint(videoId)
	return video.QueryFavorCntById(d.engine)
}
