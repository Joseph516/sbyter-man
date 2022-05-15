package dao

import (
	"douyin_service/internal/model"
)

func (d *Dao) CreateUser(userName, password string) error {
	user := model.User{
		UserName:      userName,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	return user.Create(d.engine)
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