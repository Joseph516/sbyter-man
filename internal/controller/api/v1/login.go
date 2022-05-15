package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (u User) Login(c *gin.Context)  {
	param := service.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(c.Request.Context())
	p := service.LoginRequest{UserName:param.UserName, Password: param.Password}
	userId, flag, err := svc.Login(&p)
	if err != nil {
		global.Logger.Errorf("svc.Login err: %v", err)
		response.ToErrorResponse(errcode.ErrorLoginFail)
		return
	}

	if !flag {
		global.Logger.Error("用户名/密码错误")
		response.ToErrorResponse(errcode.ErrorLoginFail)
		return
	}

	token, err := app.GenerateToken(global.JWTSetting.Key, global.JWTSetting.Secret)
	res := &service.LoginResponse{
		UserID:         userId,
		Token:          token,
	}
	res.StatusCode = 0
	res.StatusMsg = "登录成功"
	response.ToResponse(res)
	return
}
