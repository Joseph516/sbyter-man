package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tag struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

func (t Tag) TableName() string { return "douyin_tag" }

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// CreateOnConflictBatch当记录存在时不能获取id
func (t Tag) CreateOnConflictBatch(db *gorm.DB, tagsModel []Tag) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&tagsModel).Error
}

// FirstOrCreateBatch获取第一个记录，如果不存在则创建
func (t Tag) FirstOrCreateBatch(db *gorm.DB, tagsModel []Tag) error {
	for i := range tagsModel {
		if err := db.Where(tagsModel[i]).FirstOrCreate(&tagsModel[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
