package dao

import "douyin_service/internal/model"

func (d *Dao) FollowAction(follower, followed int64) error {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}

	err := follow.Create(d.engine)
	return err
}
