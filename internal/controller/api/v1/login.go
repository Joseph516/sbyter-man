package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Register 注册用户
func (u User) Register(c *gin.Context)  {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	param.LoginIP = c.ClientIP() // 注册ip
	svc := service.New(c)
	userId, flag, err := svc.Register(&param)
	if err != nil {
		global.Logger.Errorf("svc.Login err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterFail)
		return
	}

	if !flag {
		global.Logger.Error("创建失败")
		response.ToErrorResponse(errcode.ErrorRegisterFail)
		return
	}
	idStr := strconv.Itoa(int(userId))
	token, err := app.GenerateToken(global.JWTSetting.Key, global.JWTSetting.Secret, idStr)
	res := &service.RegisterResponse{
		UserID:         userId,
		Token:          token,
	}
	res.StatusCode = 0
	res.StatusMsg = "注册成功"
	response.ToResponse(res)
	return
}

// Login 登录
func (u User) Login(c *gin.Context)  {
	param := service.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	param.LoginIP = c.ClientIP()
	svc := service.New(c.Request.Context())
	userId, flag, err := svc.Login(&param)
	res := &service.LoginResponse{
		UserID:         userId,
		Token:          "",
	}
	res.StatusCode = errcode.ErrorLoginFail.Code()
	res.StatusMsg = errcode.ErrorLoginFail.Msg()
	if err != nil {
		global.Logger.Errorf("svc.Login err: %v", err)
		if err == errcode.ErrorLoginDanger { // 登录IP异常
			res.StatusCode = errcode.ErrorLoginDanger.Code()
			res.StatusMsg = errcode.ErrorLoginDanger.Msg()
			// 此处将登录IP存入Token，保证不会被篡改Token
			token, err := app.GenerateToken(global.JWTSetting.Key, global.JWTSetting.Secret, param.LoginIP)
			err = util.SendVerifiedEmail([]string{param.UserName}, userId, param.LoginIP, token)
			if err != nil {
				global.Logger.Errorf("util.SendVerifiedEmail: %v", err)
			}
		}
		response.ToResponse(res)
		return
	}

	if !flag {
		global.Logger.Error("用户名/密码错误")
		response.ToResponse(res)
		return
	}
	idStr := strconv.Itoa(int(userId))
	token, err := app.GenerateToken(global.JWTSetting.Key, global.JWTSetting.Secret, idStr)
	res = &service.LoginResponse{
		UserID:         userId,
		Token:          token,
	}
	res.StatusCode = 0
	res.StatusMsg = "登录成功"
	response.ToResponse(res)
	return
}


