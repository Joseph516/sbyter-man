package v1

import (
	"douyin_service/global"
	"douyin_service/internal/model/message"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/email"
	"douyin_service/pkg/errcode"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register 注册用户
func (u User) Register(c *gin.Context) {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	if ok := email.VerifyEmailFormat(param.UserName); !ok {
		res := &service.RegisterResponse{
			UserID: errcode.ErrorUserID,
			Token:  "",
		}
		res.StatusCode = errcode.ErrorEmail.Code()
		res.StatusMsg = errcode.ErrorEmail.Msg()
		response.ToResponse(res)
		return
	}

	param.LoginIP = c.ClientIP() // 注册ip
	svc := service.New(c)

	user, err := svc.GetUserByEmail(&service.GetUserByEmailRequest{UserName: param.UserName})
	if err != nil {
		global.Logger.Errorf("svc.GetUserByEmail err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterFail)
		return
	}
	fmt.Println(user, err)
	if user.UserName != "" {
		global.Logger.Errorf("svc.GetUserByEmail err: %v", err)
		res := &service.RegisterResponse{
			UserID: errcode.ErrorUserID,
			Token:  "",
		}
		res.StatusCode = errcode.ErrorRegisterAgain.Code()
		res.StatusMsg = errcode.ErrorRegisterAgain.Msg()
		response.ToResponse(res)
		return
	}
	newEmail := message.Email{
		UserName: []string{param.UserName},
		Password: param.Password,
		LoginIP:  param.LoginIP,
		Type:     1,
	}
	svc.Kafka.Producer(global.KafkaSetting.TopicEmail, newEmail.String(), 1) // 向kafka生产一条消息
	res := &service.RegisterResponse{
		UserID: errcode.ErrorUserID,
		Token:  "",
	}
	res.StatusCode = errcode.ErrorRegisterVerify.Code()
	res.StatusMsg = errcode.ErrorRegisterVerify.Msg()
	response.ToResponse(res)
}

// Login 登录
func (u User) Login(c *gin.Context) {
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
		UserID: userId,
		Token:  "",
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
			if err != nil {
				global.Logger.Errorf("app.GenerateToken err: %v", err)
				response.ToErrorResponse(errcode.ErrorRegisterFail)
				return
			}
			email := message.Email{
				UserName: []string{param.UserName},
				UserId:   userId,
				LoginIP:  param.LoginIP,
				Token:    token,
				Type:     2,
			}
			svc.Kafka.Producer(global.KafkaSetting.TopicEmail, email.String(), 1) // 向kafka生产一条消息
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
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterFail)
		return
	}
	res = &service.LoginResponse{
		UserID: userId,
		Token:  token,
	}
	res.StatusCode = 0
	res.StatusMsg = "登录成功"
	response.ToResponse(res)
	//return	//多余的return
}
