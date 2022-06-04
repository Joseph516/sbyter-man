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

type Notify struct {}

func NewNotify() Notify {
	return Notify{}
}

// VerifyLogin 生成验证登录通知
func (n Notify) VerifyLogin(c *gin.Context) {
	param := service.VerifyRequest{}
	response := app.NewResponse(c)
	var res service.VerifyResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	valid, tokenErr := app.ValidToken(param.Token, param.IP)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}
	flag, err := util.VerifyEmail(param.IP, param.Token)
	if !flag || err != nil {
		res.StatusCode = errcode.ErrorVerifyExpire.Code()
		res.StatusMsg = errcode.ErrorVerifyExpire.Msg()
		response.ToResponse(res)
		return
	}
	svc := service.New(c)
	p := service.UpdateIPRequest{
		ID:      param.ID,
		LoginIP: param.IP,
	}
	err = svc.UpdateUserLoginIP(&p)
	if err != nil {
		global.Logger.Errorf("svc.UpdateUserLoginIP: %v", err)
		response.ToResponse("校验失败")
		return
	}
	res.StatusMsg = "验证成功"
	response.ToResponse(res)
}

//VerifyRegister 验证注册
func (n Notify) VerifyRegister(c *gin.Context)  {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	var res service.VerifyResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	flag, err := util.VerifyEmail(param.LoginIP, "")
	if !flag || err != nil {
		res.StatusCode = errcode.ErrorVerifyExpire.Code()
		res.StatusMsg = errcode.ErrorVerifyExpire.Msg()
		response.ToResponse(res)
		return
	}
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
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterFail)
		return
	}
	resp := &service.RegisterResponse{
		UserID: userId,
		Token:  token,
	}
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	response.ToResponse(resp)
}