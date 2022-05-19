package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) ListCommentByVideoId(videoId int64) ([]model.Comment, error) {
	return model.Comment{VideoId: videoId}.ListCommentByVideoId(d.engine)
}
