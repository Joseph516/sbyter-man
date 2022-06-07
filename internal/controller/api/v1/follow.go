package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Follow struct{}

func NewFollow() Follow {
	return Follow{}
}

func (f *Follow) Action(c *gin.Context) {
	// 解析入参
	param := service.FollowActionRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	res := service.ResponseCommon{}
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 验证token
	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}
	//// 从token中获取user_id
	//claims, err := app.ParseToken(param.Token)
	//if err != nil {
	//	global.Logger.Errorf("app.ParseToken: %v", err)
	//	response.ToErrorResponse(errcode.ErrorActionFail)
	//	return
	//}
	//userId, _ := strconv.Atoi(claims.Audience)
	userId := param.UserId

	//更新数据库
	svc := service.New(c.Request.Context())
	_, err := svc.FollowAction(&param, uint(userId))
	if err != nil {
		global.Logger.Errorf("svc.FollowAction errs: %v", err)
		response.ToErrorResponse(errcode.ErrorFollowActionFail)
		return
	}

	res.StatusCode = 0
	res.StatusMsg = "操作成功"
	response.ToResponse(res)
	//return	//编译器提示，多余的return

}

func (f Follow) FollowList(c *gin.Context) {
	param := service.FollowListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	res := service.FollowListResponse{}
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 验证token
	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}
	// 从token中获取user_id
	//claims, err := app.ParseToken(param.Token)
	//if err != nil {
	//	global.Logger.Errorf("app.ParseToken: %v", err)
	//	response.ToErrorResponse(errcode.ErrorActionFail)
	//	return
	//}
	//userId, _ := strconv.Atoi(claims.Audience)

	userId := param.UserId
	//访问数据库
	svc := service.New(c.Request.Context())
	res, err := svc.FollowList(uint(userId))
	if err != nil {
		global.Logger.Errorf("FollowList errs: %v", err.Error())
		res.StatusCode = -1
		res.StatusMsg = err.Error()
	}
	response.ToResponse(res)
}


func (f Follow) FollowerList(c *gin.Context) {
	param := service.FollowListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	res := service.FollowListResponse{}
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 验证token
	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}
	// 从token中获取user_id
	claims, err := app.ParseToken(param.Token)
	if err != nil {
		global.Logger.Errorf("app.ParseToken: %v", err)
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	userId, _ := strconv.Atoi(claims.Audience)
	//访问数据库
	svc := service.New(c.Request.Context())
	res, err = svc.FollowerList(uint(userId))
	if err != nil {
		global.Logger.Errorf("FollowList errs: %v", err.Error())
		res.StatusCode = -1
		res.StatusMsg = err.Error()
	}
	response.ToResponse(res)
}
