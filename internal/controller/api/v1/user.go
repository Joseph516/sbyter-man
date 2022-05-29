package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Get 获取用户信息
func (u User) Get(c *gin.Context) {
	param := service.GetUserInfoRequest{}
	response := app.NewResponse(c)
	var res service.GetUserInfoResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	userStr := strconv.Itoa(int(param.UserId))
	valid, tokenErr := app.ValidToken(param.Token, userStr)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.GetUserById(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetUserById err: %v", err)
		response.ToResponse(errcode.ErrorGetUserInfoFail)
		return
	}

	res = service.GetUserInfoResponse{
		User: &service.UserInfo{
			ID:              user.ID,
			Name:            user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        false,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
		},
	}
	res.StatusCode = 0
	res.StatusMsg = "获取信息成功"
	response.ToResponse(res)
	//return	//多余的return
}
