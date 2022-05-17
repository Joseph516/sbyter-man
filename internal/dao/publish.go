package dao

import "douyin_service/internal/model"

// type PublishDao struct{}

// var publishDao *PublishDao

func (d *Dao) ListVideoByUserId(userId uint) ([]model.Video, error) {
	return model.Video{AuthorId: int64(userId)}.ListVideoByUserId(d.engine)
}
