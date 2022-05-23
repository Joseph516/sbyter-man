package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

func NewComment() Comment {
	return Comment{}
}

//查看视频的评论，直接列出改视频的所有评论
func (co Comment) List(c *gin.Context) {
	param := service.CommentListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		response.ToErrorResponse(tokenErr)
		return
	}
	//获取评论列表
	var resp service.CommentListResponse
	svc := service.New(c.Request.Context())
	resp, err := svc.GetCommentList(param.VideoId)
	if err != nil {
		global.Logger.Errorf("svc.GetCommentList err: %v", err)
		response.ToErrorResponse(errcode.ErrorListCommentFail)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "获取视频的评论列表成功"
	response.ToResponse(resp)
}

//评论操作接口，发布或者删除评论

func (co Comment) CommentAction(c *gin.Context) {
	param := service.CommentActionRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		response.ToErrorResponse(tokenErr)
		return
	}
	// 从token中获取user_id
	claims, err := app.ParseToken(param.Token)
	if err != nil {
		global.Logger.Errorf("app.ParseToken: %v", err)
		response.ToErrorResponse(errcode.ErrorActionPublishFail)
		return
	}
	userId, _ := strconv.Atoi(claims.Audience)
	param.UserId = int64(userId)
	//提交评论
	var resp service.PublishListResponse
	svc := service.New(c.Request.Context())
	err = svc.CommentAction(&param)
	if err != nil {
		global.Logger.Errorf("svc.CommentAction err: %v", err)
		response.ToErrorResponse(errcode.ErrorActionCommentFail)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作评论成功"
	response.ToResponse(resp)
}
