package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

func NewComment() Comment {
	return Comment{}
}

//List评论列表
//查看视频的评论，直接列出改视频的所有评论
// 参数名	位置	参数类型	必填	说明

func (co Comment) List(c *gin.Context) {
	param := service.CommentListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	var resp service.CommentListResponse
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		resp.StatusCode = errcode.ErrorLoginExpire.Code()
		resp.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(resp)
		return
	}
	//获取评论列表
	svc := service.New(c.Request.Context())
	resp, err = svc.GetCommentList(param.VideoId)
	if err != nil {
		global.Logger.Errorf("svc.GetCommentList err: %v", err)
		response.ToErrorResponse(errcode.ErrorListPublishFail)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "获取视频的评论列表成功"
	response.ToResponse(resp)
}
