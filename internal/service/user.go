package service

import (
	"douyin_service/internal/model"
)

type CreateUserRequest struct {
	UserName string `form:"username" binding:"required,len=11"`
	Password string `form:"password" binding:"required,len>=8"`
}

type UserInfo struct {
	ID            uint    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type GetUserInfoRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type GetUserInfoResponse struct {
	ResponseCommon
	User *UserInfo `json:"user" binding:"required"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.UserName, param.Password)
}

func (svc *Service) GetUserById(param *GetUserInfoRequest) (model.User, error) {
	return svc.dao.GetUserById(param.UserId)
}
