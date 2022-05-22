package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) CreateFollow(follower, followed int64) error {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	if flag, err:= follow.IsExist(d.engine);flag{
		return err
	}
	err := follow.Create(d.engine)
	return err
}

func (d *Dao) CancelFollow(follower, followed int64) error {

	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	err := follow.Delete(d.engine)
	return err
}