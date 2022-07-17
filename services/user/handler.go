package main

import (
	"context"
	"douyin_service/services/user/kitex_gen/user"
	"douyin_service/services/user/pkg/jwt"
	"douyin_service/services/user/repository"
	"errors"
	"strconv"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	if req.Username == "" || req.Password == "" {
		resp = &user.UserRegisterResponse{
			UserId: -1,
			Token:  "",
		}
		err = errors.New("用户或者密码错误")
		return
	}

	// 保存至数据库
	userId, err := repository.NewUserDaoInstance().CreateUser(req.Username, req.Password, "")
	if err != nil {
		resp = &user.UserRegisterResponse{
			UserId: -1,
			Token:  "",
		}
		return
	}

	// 生成token
	idStr := strconv.Itoa(int(userId))
	token, err := jwt.GenerateToken(idStr)

	resp = &user.UserRegisterResponse{
		UserId: int64(userId),
		Token:  token,
	}
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	userId, err := repository.NewUserDaoInstance().CheckUser(req.Username, req.Password)
	if err != nil {
		return
	}

	// 生成token
	idStr := strconv.Itoa(int(userId))
	token, err := jwt.GenerateToken(idStr)
	if err != nil {
		return
	}

	resp = &user.UserLoginResponse{
		UserId: int64(userId),
		Token:  token,
	}
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// 校验token
	var ok bool
	if ok, err = jwt.CheckToken(req.Token, strconv.FormatInt(req.UserId, 10)); !ok {
		return
	}

	// 读取数据库
	userInfo, err := repository.NewUserDaoInstance().GetUserById(uint(req.UserId))
	if err != nil {
		return
	}
	resp = &user.UserResponse{
		User: &user.User{
			Id:            int64(userInfo.ID),
			Name:          userInfo.UserName,
			FollowCount:   &userInfo.FollowCount,
			FollowerCount: &userInfo.FollowerCount,
			IsFollow:      userInfo.FollowerCount != 0,
		},
	}
	return
}
