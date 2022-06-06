package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	*Model
	FollowedId uint
	FollowerId uint
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

func (f Follow) IsExist(db *gorm.DB) (bool, error) {
	var follows []Follow
	db.Where("followed_id = ? and follower_id = ?", f.FollowedId, f.FollowerId).Find(&follows)
	if len(follows) == 0 {
		return false, nil
	}
	return true, nil
}

func (f Follow) QueryFollowList(db *gorm.DB, userId uint) (follows []Follow, err error) {
	err = db.Where("follower_id = ?", userId).Find(&follows).Error
	return
}

func (f Follow) QueryFollowerList(db *gorm.DB, userId uint) (follows []Follow, err error) {
	err = db.Where("followed_id = ?", userId).Find(&follows).Error
	return
}
