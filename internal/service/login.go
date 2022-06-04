package service

import (
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
)

type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP string `form:"login_ip"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP string  `form:"login_ip"`
}

type RegisterResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) Login(param *LoginRequest) (uint, bool, error) {
	return svc.dao.CheckUser(param.UserName, param.Password, param.LoginIP)
}

func (svc Service) Register(param *RegisterRequest) (uint, bool, error) {
	hashPassword, err := util.EncodeBcrypt(param.Password)
	if err != nil { // 加密失败
		return errcode.ErrorUserID, false, err
	}
	createUserRequest := CreateUserRequest{
		UserName: param.UserName,
		Password: hashPassword,
		LoginIP: param.LoginIP,
	}

	uid, err := svc.CreateUser(&createUserRequest)
	if err != nil {
		return uid, false, err
	}
	getUserInfoRequest := GetUserInfoRequest{
		UserId: uid,
		Token:  "",
	}
	user, err := svc.GetUserById(&getUserInfoRequest)
	if err != nil {
		return user.ID, false, err
	}
	return user.ID, true, nil
}
