package dao

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
)

func (d *Dao) CreateUser(userName, password string) (uint, error) {
	sign, err := util.RandomSign()
	if err != nil {
		return errcode.ErrorUserID, err
	}
	img, err := util.RandomBackground()
	if err != nil {
		return errcode.ErrorUserID, err
	}
	user := model.User{
		UserName:      userName,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
		Avatar:          util.RandomAvatar(userName),
		Signature:       sign,
		BackgroundImage: img,
	}
	err = user.Create(d.engine)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// CheckUser 校验用户名和密码
func (d *Dao) CheckUser(username, password string) (uint, bool, error) {
	user := model.User{
		UserName: username,
		Password: password,
	}
	return user.CheckUser(d.engine)
}

func (d *Dao) GetUserById(userId uint) (model.User, error) {
	user := model.User{
		Model: &model.Model{
			ID: userId,
		},
	}
	return user.GetUserById(d.engine)
}

func (d *Dao) GetUsersByIds(userIds []uint) ([]model.User, error) {
	var user model.User
	users, err := user.GetUsersByIds(userIds, d.engine)
	if err != nil {
		return nil, err
	}
	return users, nil
}
