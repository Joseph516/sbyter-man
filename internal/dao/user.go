package dao

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
)

func (d *Dao) CreateUser(userName, password, loginIP string) (uint, error) {
	sign, err := util.RandomSign()
	if err != nil {
		return errcode.ErrorUserID, err
	}
	img, err := util.RandomBackground()
	if err != nil {
		return errcode.ErrorUserID, err
	}
	user := model.User{
		UserName:        userName,
		Password:        password,
		FollowCount:     0,
		FollowerCount:   0,
		Avatar:          util.RandomAvatar(userName),
		Signature:       sign,
		BackgroundImage: img,
		LoginIP:         loginIP,
	}
	err = user.Create(d.engine)
	if err != nil {
		return errcode.ErrorUserID, err
	}
	return user.ID, nil
}

// CheckUser 校验用户名和密码
func (d *Dao) CheckUser(username, password, loginIP string) (uint, bool, error) {
	user := model.User{
		UserName: username,
		Password: password,
		LoginIP:  loginIP,
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

func (d *Dao) GetUserByEmail(username string) (model.User, error) {
	user := model.User{
		UserName: username,
	}
	return user.GetUserByEmail(d.engine)
}

func (d Dao) UpdateUserLoginIP(userId uint, loginIP string) error {
	user := model.User{
		Model: &model.Model{
			ID: userId,
		},
		LoginIP: loginIP,
	}
	return user.UpdateIP(d.engine)
}

func (d Dao) UpdateById(userId uint, data map[string]interface{}) error {
	user := model.User{
		Model: &model.Model{
			ID: userId,
		},
	}
	return user.UpdateById(d.engine, data)
}
