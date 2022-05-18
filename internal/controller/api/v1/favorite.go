package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Favorite struct {
}

func NewFavorite() Favorite {
	return Favorite{}
}

// Action 登录用户对视频的点赞和取消点赞操作
func (f Favorite) Action(c *gin.Context) {
	//获取请求参数
	param := service.ActionRequest{}
	response := app.NewResponse(c)
	var res service.ActionResponse
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		res.StatusCode = errcode.ErrorLoginExpire.Code()
		res.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(res)
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.Action(&param)
	if err != nil {
		global.Logger.Errorf("svc.Action err: %v", err)
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "操作成功"
	response.ToResponse(res)
	return
}

// FavoriteList 登录用户点赞列表
func (f Favorite) FavoriteList(c *gin.Context) {
	param := service.FavoriteListRequest{}
	response := app.NewResponse(c)
	var res service.ActionResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	valid, err := app.ValidToken(param.Token)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		res.StatusCode = errcode.ErrorLoginExpire.Code()
		res.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(res)
		return
	}
	svc := service.New(c.Request.Context())
	favoriteList, err := svc.FavoriteList(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetUserById err: %v", err)
		response.ToResponse(errcode.ErrorActionListFail)
		return
	}

	//TODO:根据favoriteList查询相应的视频。。。
	fmt.Println(favoriteList)
}
