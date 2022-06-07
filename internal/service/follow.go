package service

import (
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
	"strconv"
)

type FollowActionRequest struct {
	//UserId uint   `form:"user_id"  binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   uint   `form:"to_user_id" binding:"required"`
	ActionType int64  `form:"action_type" binding:"required"`
}

type FollowListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token string `form:"token" binding:"required"`
}

type FollowListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

// IsFollowRequest 判断A是否关注B
type IsFollowRequest struct {
	A uint
	B uint
}

func (svc *Service) LoadFanCount(userId uint) (int64, error) {
	upKey := util.FanCountKey(userId)
	if exist, _:= svc.redis.IsExist(upKey);!exist{
		//读入
		userInfo, err := svc.dao.GetUserById(userId)
		if err!=nil{
			return 0, err
		}
		svc.redis.SetFollowInfo(upKey, userInfo.FollowerCount)
		//return userInfo.FollowerCount, nil
	}
	ret, err := svc.redis.Get(upKey)
	if err!=nil{
		return 0, err
	}
	retInt, err := strconv.Atoi(ret)
	if err!=nil{
		return 0, err
	}
	return int64(retInt), nil
}


func (svc *Service) LoadFollowCount(userId uint) (int64, error) {
	fanKey := util.FollowCountKey(userId)
	if exist, _:= svc.redis.IsExist(fanKey);!exist{
		//读入
		userInfo, err := svc.dao.GetUserById(userId)
		if err!=nil{
			return 0, err
		}
		svc.redis.SetFollowInfo(fanKey, userInfo.FollowCount)

		//return userInfo.FollowerCount, nil
	}
	ret, err := svc.redis.Get(fanKey)
	if err!=nil{
		return 0, err
	}
	retInt, err := strconv.Atoi(ret)
	if err!=nil{
		return 0, err
	}
	return int64(retInt), nil
}

func (svc *Service) FollowAction(param *FollowActionRequest, userId uint) (flag bool, err error) {
	switch param.ActionType {
	case 1:
		flag, err = svc.dao.CreateFollow(userId, param.ToUserId)
		if err!=nil{
			return
		}
		if flag {
			// 创建了新的记录，更新缓存中的关注数量
			_,_ = svc.LoadFollowCount(userId)
			_,_ = svc.LoadFanCount(param.ToUserId)
			svc.redis.FollowAction(param.ToUserId, userId)
		}
	case 2:
		flag, err = svc.dao.CancelFollow(userId, param.ToUserId)
		if err!=nil{
			return
		}
		if flag {
			_,_ = svc.LoadFollowCount(userId)
			_,_ = svc.LoadFanCount(param.ToUserId)
			// 删除了记录，更新缓存中的关注数量
			svc.redis.CancelFollowAction(param.ToUserId, userId)
		}
	default:
		return false, errcode.InvalidParams
	}

	//svc.redis.Output()
	//_, followCount, _ := svc.redis.QueryFollowCnt(userId)
	//fmt.Println("关注人数:",followCount)
	//_, fanCount, _ := svc.redis.QueryFanCnt(param.ToUserId)
	//fmt.Println("粉丝人数:", fanCount)
	return
}

func (svc *Service) FollowList(userId uint) (res FollowListResponse, err error) {
	follows, err := svc.dao.FollowList(userId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowedId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		userInfo.IsFollow = true
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}

func (svc *Service) FollowerList(userId uint) (res FollowListResponse, err error) {
	follows, err := svc.dao.FollowerList(userId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowerId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
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

func (svc *Service) QueryFanCntRedis(userId uint) (bool, int64, error){
	return svc.redis.QueryFanCnt(userId)
}

func (svc *Service) QueryFollowCntRedis(userId uint) (bool, int64, error){
	return svc.redis.QueryFollowCnt(userId)
}

func (svc *Service) IsFollow(param IsFollowRequest) (bool, error) {
	return svc.dao.IsFollow(param.A, param.B)
}