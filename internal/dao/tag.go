package dao

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/errcode"
)

func (d *Dao) CreateTag(name string) (uint, error) {
	tag := model.Tag{
		Name: name,
	}
	err := tag.Create(d.engine)
	if err != nil {
		return errcode.ErrorTagID, err
	}
	return tag.ID, nil
}

func (d *Dao) GetTagByName(name string) (model.Tag, error) {
	tag := model.Tag{
		Name: name,
	}
	err := tag.GetTagByName(d.engine)
	if err != nil {
		return model.Tag{}, err
	}
	return tag, nil
}
