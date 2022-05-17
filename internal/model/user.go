package model

import (
	"douyin_service/pkg/app"
	"gorm.io/gorm"
)

type User struct {
	*Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

type UserSwagger struct {
	List  []*User
	Pager *app.Pager
}

func (u User) TableName() string {
	return "douyin_user"
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("user_name = ?", u.UserName).Updates(u).Error
}

func (u User) Delete(db *gorm.DB) error {
	return db.Where("user_name = ?", u.UserName).Delete(&u).Error
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) CheckUser(db *gorm.DB) (uint, bool, error) {
	var user User
	var err error
	if err = db.Where("user_name = ? AND password = ?", u.UserName, u.Password).Find(&user).Error; err != nil {
		return 0, false, err
	}
	if (user == User{}) {
		return 0, false, err
	}
	return user.ID, true, nil
}

func (u User) GetUserById(db *gorm.DB) (User, error) {
	var user User
	if err := db.Where("id = ?", u.ID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}