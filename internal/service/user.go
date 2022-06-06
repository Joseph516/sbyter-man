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
	TotalFavorited  int64    `json:"total_favorited"`  // 被赞的总次数
	FavoriteCount   int64    `json:"favorite_count"`   //喜欢总数量
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
	UserName string `form:"user_name"  binding:"required"`
}


type UpdateIPRequest struct {
	ID      uint   `form:"id"`
	LoginIP string `form:"login_ip"`
}

// UpdateByIdRequest redis更新Mysql用户信息请求
type UpdateByIdRequest struct {
	UserId uint
	Data map[string]interface{}
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

func (svc *Service) UpdateById(param *UpdateByIdRequest) error {
	return svc.dao.UpdateById(param.UserId, param.Data)
}

// GetTotalFavoritedById 查询用户获赞数量接口
func (svc *Service) GetTotalFavoritedById(userId uint) (int64, error) {
	exist, cnt, err := svc.redis.QueryUserFavoritedCount(userId)
	if err != nil {
		return 0, err
	}
	if exist {
		return cnt, nil
	}
	user, err := svc.dao.GetUserById(userId)
	if err != nil {
		return 0, err
	}
	cnt = user.TotalFavorited
	return cnt, nil
}

// GetFavoriteCountById 查询用户点赞数量接口
func (svc *Service) GetFavoriteCountById(userId uint) (int64, error) {
	exist, cnt, err := svc.redis.QueryUserFavoriteCount(userId)
	if err != nil {
		return 0, err
	}
	if exist {
		return cnt, nil
	}
	user, err := svc.dao.GetUserById(userId)
	if err != nil {
		return 0, err
	}
	cnt = user.FavoriteCount
	return cnt, nil
}
