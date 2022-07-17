package repository

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"errors"
	"sync"
)

type Model struct {
	ID        uint  `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type User struct {
	*Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	// Avatar          string `json:"avatar"`           // 头像
	// Signature       string `json:"signature"`        // 个性签名
	// BackgroundImage string `json:"background_image"` // 背景图片
	// LoginIP         string `json:"login_ip"`         // 最近登陆的IP地址
	// TotalFavorited  int64  `json:"total_favorited"`  // 被赞的总次数
	// FavoriteCount   int64  `json:"favorite_count"`   // 喜欢总数量
}

func (User) TableName() string {
	return "douyin_user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(username, password, loginIP string) (uint, error) {
	user := User{
		UserName:      username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (*UserDao) CheckUser(username, password string) (uint, error) {
	user := model.User{
		UserName: username,
	}

	if err := db.Where("user_name = ?", user.UserName).First(&user).Error; err != nil {
		return 0, err
	}
	// 核实数据库密码
	if !util.CheckBcrypt(user.Password, password) {
		return 0, errors.New("CheckBcrypt failed")
	}
	return user.ID, nil
}

func (*UserDao) GetUserById(userId uint) (*User, error) {
	var user User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
