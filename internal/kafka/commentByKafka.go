package kafka

import (
	"douyin_service/global"
	"douyin_service/internal/dao"
	"encoding/json"
)

type CommentActionRequest struct {
	UserId      int64  `form:"user_id" binding:"-"`
	Token       string `form:"token" binding:"required"`
	VideoId     int64  `form:"video_id"  binding:"required"`
	ActionType  int64  `form:"action_type"  binding:"required"`
	CommentText string `form:"comment_text"  binding:"-"`
	CommentId   uint   `form:"comment_id"  binding:"-"`
}

func (c CommentActionRequest) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

const PUBCOMMENT int64 = 1

//const DELCOMMENT int64 = 2

func (param *CommentActionRequest) CommentAction() error {
	dao := dao.New(global.DBEngine)
	if param.ActionType == PUBCOMMENT {
		return dao.PublishComment(param.VideoId, param.UserId, param.CommentText)
	}
	return dao.DeleteComment(param.CommentId)
}
