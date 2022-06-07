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

		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := uint(0)
	if param.Token != "" {
		valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
		if !valid {
			global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
			response.ToErrorResponse(tokenErr)
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
		uid = uint(user_id)
	}

	// 不管用户是否登陆、有没有Token，下面都是查询全体用户的所有视频，最后根据latest_time返回20条视频

	// 先不考虑推荐算法，我们显然只需要查询全体用户的所有视频，根据latest_time返回20条结果

	// 获取视频发布列表
	//var resp service.FeedResponse
	svc := service.New(c.Request.Context())
	resp, err := svc.Feed(uid, param.LastTime)
	if err != nil {
		global.Logger.Errorf("svc.PublishFeed err: %v", err)
		response.ToErrorResponse(errcode.ErrorListPublishFail)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "获取用户的视频发布列表成功"
	response.ToResponse(resp)
}
