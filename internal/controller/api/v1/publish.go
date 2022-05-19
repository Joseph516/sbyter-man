package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Publish struct{}

func NewPublish() Publish {
	return Publish{}
}

// List发布列表
// 登录用户的视频发布列表，直接列出用户所有投稿过的视频
// 参数名	位置	参数类型	必填	说明
// token	query		是	用户鉴权token
// user_id	query		是	用户id
func (p Publish) List(c *gin.Context) {
	param := service.PublishListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var resp service.PublishListResponse
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		resp.StatusCode = errcode.ErrorLoginExpire.Code()
		resp.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(resp)
		return
	}

	// 获取视频发布列表
	svc := service.New(c.Request.Context())
	resp, err = svc.PublishList(param.UserId)
	if err != nil {
		global.Logger.Errorf("svc.PublishList err: %v", err)
		response.ToErrorResponse(errcode.ErrorListPublishFail)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "获取用户的视频发布列表成功"
	response.ToResponse(resp)
}

// Action投稿接口, 登录用户选择视频上传
// 参数名	参数类型	必填	说明
// data	file	是	视频数据
// token	text	是	用户鉴权token
// title	text	是	视频标题
func (p Publish) Action(c *gin.Context) {
	data, _ := c.FormFile("data")
	param := service.PublishActionRequest{
		Data:  data,
		Token: c.PostForm("token"),
		Title: c.PostForm("title"),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var resp service.ResponseCommon
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		resp.StatusCode = errcode.ErrorLoginExpire.Code()
		resp.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(resp)
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

	// 发布视频
	svc := service.New(c.Request.Context())
	err = svc.PublishAction(param.Data, param.Token, param.Title, uint(userId))
	if err != nil {
		global.Logger.Errorf("svc.PublishAction err: %v", err)
		response.ToErrorResponse(errcode.ErrorActionPublishFail)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "视频发布成功"
	response.ToResponse(resp)
}
