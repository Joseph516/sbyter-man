package model

import (
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	*Model
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	Avatar          string `json:"avatar"`           // 头像
	Signature       string `json:"signature"`        // 个性签名
	BackgroundImage string `json:"background_image"` // 背景图片
	LoginIP         string `json:"login_ip"`         // 最近登陆的IP地址
	TotalFavorited  int64  `json:"total_favorited"`  // 被赞的总次数
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢总数量
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

func (u User) UpdateById(db *gorm.DB, data map[string]interface{}) error {
	//修改多列的时候传入结构体会导致0值无法修改，利用map传参即可
	//Model里的参数中的主键会自动作为查询条件
	return db.Model(&u).Updates(data).Error
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
	if err = db.Where("user_name = ?", u.UserName).Find(&user).Error; err != nil {
		return 0, false, err
	}
	if (user == User{}) {
		return errcode.ErrorUserID, false, err
	}
	if !util.CheckBcrypt(user.Password, u.Password) { // 核实数据库密码
		return errcode.ErrorUserID, false, err
	}
	if u.LoginIP != user.LoginIP {
		return user.ID, false, errcode.ErrorLoginDanger
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

func (u User) GetUsersByIds(userIds []uint, db *gorm.DB) ([]User, error) {
	users := make([]User, 0)
	err := db.Where("id IN ?", userIds).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) GetUserByEmail(db *gorm.DB) (User, error) {
	var user User
	if err := db.Where("user_name = ?", u.UserName).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u User) UpdateIP(db *gorm.DB) error {
	err := db.Model(&User{}).Where("id = ?", u.ID).Update("login_ip", u.LoginIP).Error
	fmt.Println("err: ", err)
	return err
}
