package service

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/errcode"
	"gorm.io/gorm"
)

type TagListRequest struct {
	Tags []string `json:"tags"`
}

type GetTagByNameRequest struct {
	Name string `json:"name"`
}

// AddTagList 添加新标签到数据库
func (svc *Service) AddTagList(param TagListRequest) ([]uint, error) {
	var p GetTagByNameRequest
	ids := make([]uint, 0, len(param.Tags))
	for _, tag := range param.Tags {
		p.Name = tag
		t, err := svc.GetTagByName(p)
		if t.ID == errcode.ErrorTagID || err ==gorm.ErrRecordNotFound { // 表中无该数据
			id, err := svc.dao.CreateTag(tag) // 插入新标签
			if err != nil {
				continue
			}
			ids = append(ids, id)
		} else {
			ids = append(ids, t.ID)
		}
	}
	return ids, nil
}

func (svc *Service) GetTagByName(param GetTagByNameRequest) (model.Tag, error) {
	return svc.dao.GetTagByName(param.Name)
}