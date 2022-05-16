package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) CreateUser(userName, password string) (uint, error) {
	user := model.User{
		UserName:      userName,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := user.Create(d.engine)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// CheckUser 校验用户名和密码
func (d *Dao) CheckUser(username, password string) (uint, bool, error) {
	user := model.User{
		UserName:      username,
		Password:      password,
	}
	return user.CheckUser(d.engine)
}

func (d *Dao) GetUserById(userId uint) (model.User, error) {
	user := model.User{
		Model: &model.Model{
			ID:   userId     ,
		},
	}
	return user.GetUserById(d.engine)
}