package service

type FollowActionRequest struct {
	UserId int64   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
	ToUserId int64 `form:"to_user_id" binding:"required"`
	ActionType int64 `form:"action_type" binding:"required"`

}

func (svc *Service) FollowAction(param *FollowActionRequest) error {
	return svc.dao.FollowAction(param.UserId, param.ToUserId)
}