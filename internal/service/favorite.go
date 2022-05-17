package service

import (
	"douyin_service/global"
	"strconv"
)

type ActionRequest struct {
	UserId     uint   `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	VideoId    uint   `form:"video_id" binding:"required"`
	ActionType int    `form:"action_type" binding:"required"`
}

type ActionResponse struct {
	ResponseCommon
}

func (svc *Service) Action(param *ActionRequest) (bool, error) {
	//拼接用户id_videoId
	userId := strconv.Itoa(int(param.UserId))
	videoId := strconv.Itoa(int(param.VideoId))

	actionType := param.ActionType
	if actionType == 2 {
		actionType = 0
	}
	//TODO: 保证只有登录的用户才能点赞？
	//TODO: 两个操作可能需要原子操作？
	//修改用户中点赞的视频
	err := global.Rd.SetBit(userId, int64(param.VideoId), actionType).Err()
	if err != nil {
		return false, err
	}
	//修改视频中点赞的用户
	err = global.Rd.SetBit(videoId, int64(param.UserId), actionType).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
