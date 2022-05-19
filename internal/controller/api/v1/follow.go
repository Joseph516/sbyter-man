package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Follow struct{}

func NewFollow() Follow {
	return Follow{}
}

func (f *Follow)Action(c *gin.Context)  {
	// 解析入参
	param := service.FollowActionRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	res :=service.ResponseCommon{}
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 验证token
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}

	//更新数据库
	svc := service.New(c.Request.Context())
	err = svc.FollowAction(&param)
	if err!=nil{
		global.Logger.Errorf("svc.FollowAction errs: %v", err)
		response.ToErrorResponse(errcode.ErrorFollowActionFail)
		return
	}

	res.StatusCode = 0
	res.StatusMsg = "关注成功"
	response.ToResponse(res)
	return

}