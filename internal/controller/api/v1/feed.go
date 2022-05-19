package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Feed struct{}

func NewFeed() Feed {
	return Feed{}
}

func (f Feed) Feed(c *gin.Context) {
	param := service.FeedRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	var resp service.FeedResponse
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
		return
	}
	user_id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		return
	}

	// 获取视频发布列表
	svc := service.New(c.Request.Context())
	resp, err = svc.Feed(uint(user_id), param.LastTime)
	if err != nil {
		global.Logger.Errorf("svc.PublishFeed err: %v", err)
		response.ToErrorResponse(errcode.ErrorListPublishFail)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "获取用户的视频发布列表成功"
	response.ToResponse(resp)
}
