package model

import "gorm.io/gorm"

type Tag struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}
func (t Tag) TableName() string {
	return "douyin_tag"
}
func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// GetTagByName 根据名称查询
func (t *Tag) GetTagByName(db *gorm.DB) error {
	if err := db.Where("name = ?", t.Name).Find(&t).Error; err != nil {
		return err
	}
	return nil
}