package v1

import (
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"douyin_service/services/api/rpc"
	"douyin_service/services/api/service"
	"douyin_service/services/kitex_gen/user"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Register 注册用户
func (u User) Register(c *gin.Context) {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	req := &user.UserRegisterRequest{
		Username: param.UserName,
		Password: param.Password,
	}
	res := &service.RegisterResponse{}
	resp, err := rpc.Register(c, req)
	if err != nil {
		res.StatusCode = errcode.ErrorRegisterFail.Code()
		res.StatusMsg = errcode.ErrorRegisterFail.WithDetails(err.Error()).Details()
		response.ToResponse(res)
		return
	}

	res = &service.RegisterResponse{
		UserID: uint(resp.UserId),
		Token:  resp.Token,
	}
	res.StatusCode = errcode.Success.Code()
	res.StatusMsg = "注册成功"
	response.ToResponse(res)
}

// Login 登录
func (u User) Login(c *gin.Context) {
	param := service.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	req := &user.UserLoginRequest{
		Username: param.UserName,
		Password: param.Password,
	}
	resp, err := rpc.Login(c, req)
	res := &service.LoginResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorLoginFail.Code()
		res.StatusMsg = errcode.ErrorLoginFail.WithDetails(err.Error()).Details()
		response.ToResponse(res)
		return
	}
	res = &service.LoginResponse{
		UserID: uint(resp.UserId),
		Token:  resp.Token,
	}
	res.StatusCode = errcode.Success.Code()
	res.StatusMsg = "登录成功"
	response.ToResponse(res)
}

// Get 获取用户信息
func (u User) Get(c *gin.Context) {
	param := service.GetUserInfoRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	req := &user.UserRequest{
		UserId: int64(param.UserId),
		Token:  param.Token,
	}
	resp, err := rpc.GetUser(c, req)
	res := &service.GetUserInfoResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorGetUserInfoFail.Code()
		res.StatusMsg = errcode.ErrorGetUserInfoFail.WithDetails(err.Error()).Details()
		response.ToResponse(res)
		return
	}
	followerCount := resp.User.GetFollowerCount()
	res = &service.GetUserInfoResponse{
		User: &service.UserInfo{
			ID:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.GetFollowCount(),
			FollowerCount: followerCount,
			IsFollow:      followerCount != 0,
		},
	}
	res.StatusCode = 0
	res.StatusMsg = "获取信息成功"
	response.ToResponse(res)
	return
}
