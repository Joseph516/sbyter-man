package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

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
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	userId, _ := strconv.Atoi(claims.Audience)

	svc := service.New(c.Request.Context())
	err2 := svc.Action(&param, uint(userId))
	if err2 != nil {
		global.Logger.Errorf("svc.Action err: %v", err2)
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "操作成功"
	response.ToResponse(res)
}

// FavoriteList 登录用户点赞列表
func (f Favorite) FavoriteList(c *gin.Context) {
	param := service.FavoriteListRequest{}
	response := app.NewResponse(c)
	var res service.FavoriteListResponse

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	if len(param.Token) == 0 {
		//获取不到token说明未登录
		res.StatusCode = 0
		res.StatusMsg = "操作成功"
		res.VideoList = nil
		response.ToResponse(res)
		return
	}
	valid, err := app.ValidToken(param.Token, strconv.Itoa(int(param.UserId)))
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", err)
		res.StatusCode = errcode.ErrorLoginExpire.Code()
		res.StatusMsg = errcode.ErrorLoginExpire.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	favoriteList, err2 := svc.FavoriteList(&param)
	if err2 != nil {
		global.Logger.Errorf("svc.GetUserById err: %v", err2)
		response.ToResponse(errcode.ErrorActionListFail)
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "操作成功"
	res.VideoList = favoriteList
	response.ToResponse(res)
}
