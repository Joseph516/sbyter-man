package service

import (
	"douyin_service/internal/model"
)

type CreateUserRequest struct {
	UserName string `form:"username" binding:"required,len=11"`
	Password string `form:"password" binding:"required,len>=8"`
	LoginIP  string `form:"login_ip"`
}

type UserInfo struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`           // 头像
	Signature       string `json:"signature"`        // 个性签名
	BackgroundImage string `json:"background_image"` // 背景图片
	LoginIP         string `json:"login_ip"`         // 登录IP
}

type GetUserInfoRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type GetUserInfoResponse struct {
	ResponseCommon
	User *UserInfo `json:"user" binding:"required"`
}

type GetUserByEmailRequest struct {
	UserName string   `form:"user_name"  binding:"required"`
}

type UpdateIPRequest struct {
	ID      uint   `form:"id"`
	LoginIP string `form:"login_ip"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) (uint, error) {
	return svc.dao.CreateUser(param.UserName, param.Password, param.LoginIP)
}

func (svc *Service) UpdateUserLoginIP(param *UpdateIPRequest) error {
	return svc.dao.UpdateUserLoginIP(param.ID, param.LoginIP)
}

func (svc *Service) GetUserById(param *GetUserInfoRequest) (model.User, error) {
	return svc.dao.GetUserById(param.UserId)
}

func (svc *Service) GetUsersByIds(userIds []uint) ([]model.User, error) {
	return svc.dao.GetUsersByIds(userIds)
}

func (svc *Service) GetUserByEmail(param *GetUserByEmailRequest) (model.User, error) {
	return svc.dao.GetUserByEmail(param.UserName)
}
