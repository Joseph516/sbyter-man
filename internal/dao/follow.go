package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao)IsFollow(follower, followed int64) (flag bool,err error) {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	flag, err = follow.IsExist(d.engine)
	return
}

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

func (d *Dao) FollowList(userId int64) ([]model.Follow, error) {
	follows, err := model.Follow{}.QueryFollowList(d.engine, userId)
	return follows, err
}