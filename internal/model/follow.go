package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	FollowedId 	int64
	FollowerId  int64
}

func (f Follow) TableName() string {
	return "douyin_follow"
}

func (f Follow) Create(db *gorm.DB) error {
	return db.Create(&f).Error
}

func (f Follow) Delete(db *gorm.DB) error {
	return db.Where("followed_id = ? And follower_id = ?", f.FollowedId, f.FollowerId).Delete(&f).Error
}

func (f Follow) IsExist(db *gorm.DB ) (bool, error) {
	var follows []Follow
	db.Where("followed_id = ? and follower_id = ?", f.FollowedId, f.FollowerId).Find(&follows)
	fmt.Println(follows)
	if len(follows)==0{
		return false, nil
	}
	return true, nil
}