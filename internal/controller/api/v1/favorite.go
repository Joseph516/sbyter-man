package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Favorite struct {
}

func NewFavorite() Favorite {
	return Favorite{}
}

// Action 用户点赞/
func (f Favorite) Action(c *gin.Context) {
	//获取请求参数
	param := service.ActionRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}
	svc := service.New(c.Request.Context())
	p := service.ActionRequest{UserId: param.UserId,
		Token:      param.Token,
		VideoId:    param.VideoId,
		ActionType: param.ActionType,
	}
	_, err := svc.Action(&p)
	if err != nil {
		global.Logger.Errorf("svc.Action err: %v", err)
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	res := &service.ActionResponse{}
	res.StatusCode = 0
	res.StatusMsg = "操作成功"
	response.ToResponse(res)
	return
}
