package model

import "gorm.io/gorm"

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