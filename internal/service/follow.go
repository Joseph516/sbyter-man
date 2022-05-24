package service

import (
	"douyin_service/pkg/errcode"
	"fmt"
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

func (svc *Service) FollowAction(param *FollowActionRequest, userId int64)(flag bool,err error) {
	switch param.ActionType{
	case 1:
		flag, err = svc.dao.CreateFollow(userId, param.ToUserId)
		if flag{
			svc.redis.FollowAction(param.ToUserId, userId)
		}
	case 2:
		flag, err = svc.dao.CancelFollow(userId, param.ToUserId)
		if flag{
			svc.redis.CancelFollowAction(param.ToUserId, userId)
		}
	default:
		return false, errcode.InvalidParams
	}
	_, followCount, _ := svc.redis.QueryFollowCnt(userId)
	fmt.Println("关注人数:",followCount)
	_, fanCount, _ := svc.redis.QueryFanCnt(param.ToUserId)
	fmt.Println("粉丝人数:", fanCount)
	return
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
