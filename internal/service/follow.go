package service

import (
	"douyin_service/pkg/errcode"
)

type FollowActionRequest struct {
	//UserId int64   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
	ToUserId int64 `form:"to_user_id" binding:"required"`
	ActionType int64 `form:"action_type" binding:"required"`
}

type FollowListRequest struct {
	//UserId int64   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FollowListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

func (svc *Service) FollowAction(param *FollowActionRequest, userId int64) error {
	switch param.ActionType{
	case 1:
		return svc.dao.CreateFollow(userId, param.ToUserId)
	case 2:
		return svc.dao.CancelFollow(userId, param.ToUserId)
	default:
		return errcode.InvalidParams
	}
}

func (svc *Service) FollowList(userId int64) (res FollowListResponse,err error) {
	follows, err := svc.dao.FollowList(userId)
	if err != nil{
		return
	}
	for i :=range follows{
		f := follows[i]
		id := f.FollowedId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr!=nil{
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}
