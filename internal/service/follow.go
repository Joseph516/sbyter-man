package service

import "douyin_service/pkg/errcode"

type FollowActionRequest struct {
	//UserId int64   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
	ToUserId int64 `form:"to_user_id" binding:"required"`
	ActionType int64 `form:"action_type" binding:"required"`

}

func (svc *Service) FollowAction(param *FollowActionRequest, user_id int64) error {
	switch param.ActionType{
	case 1:
		return svc.dao.CreateFollow(user_id, param.ToUserId)
	case 2:
		return svc.dao.CancelFollow(user_id, param.ToUserId)
	default:
		return errcode.InvalidParams
	}
}