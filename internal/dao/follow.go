package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) IsFollow(follower, followed uint) (flag bool, err error) {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	flag, err = follow.IsExist(d.engine)
	return
}

func (d *Dao) CreateFollow(follower, followed uint) (flag bool, err error) {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	//如果已经关注了，返回false
	if flag, err = follow.IsExist(d.engine); flag || err != nil {
		flag = false
		return
	}
	//如果还没有关注，添加本条记录，返回true
	err = follow.Create(d.engine)
	if err != nil {
		return
	}
	flag = true
	return
}

func (d *Dao) CancelFollow(follower, followed uint) (bool, error) {

	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	if flag, err := follow.IsExist(d.engine); !flag || err != nil {
		return false, err
	}
	err := follow.Delete(d.engine)
	if err != nil {
		return false, err
	}
	return true, err
}

func (d *Dao) FollowList(userId uint) ([]model.Follow, error) {
	follows, err := model.Follow{}.QueryFollowList(d.engine, userId)
	return follows, err
}


func (d *Dao) FollowerList(userId uint) ([]model.Follow, error) {
	follows, err := model.Follow{}.QueryFollowerList(d.engine, userId)
	return follows, err
}