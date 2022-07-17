package rpc

import (
	"context"
	"douyin_service/pkg/util"
	"douyin_service/services/user/kitex_gen/user"
	"douyin_service/services/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

var userClient userservice.Client

type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP  string `form:"login_ip"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP  string `form:"login_ip"`
}

type RegisterResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	// Avatar          string `json:"avatar"`           // 头像
	// Signature       string `json:"signature"`        // 个性签名
	// BackgroundImage string `json:"background_image"` // 背景图片
	// LoginIP         string `json:"login_ip"`         // 登录IP
	// TotalFavorited  int64    `json:"total_favorited"`  // 被赞的总次数
	// FavoriteCount   int64    `json:"favorite_count"`   //喜欢总数量
}

type GetUserInfoRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type GetUserInfoResponse struct {
	ResponseCommon
	User *UserInfo `json:"user" binding:"required"`
}

func NewUserClient() {
	var err error
	userClient, err = userservice.NewClient(UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		// client.WithResolver(r),                            // resolver
		client.WithHostPorts("0.0.0.0:8888"),
	)

	if err != nil {
		log.Fatalf("Init NewUserCilent failed: %v", err)
	}
}

func Register(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	var err error
	req.Password, err = util.EncodeBcrypt(req.Password)
	if err != nil {
		return nil, err
	}

	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return userClient.Login(ctx, req)
}

func GetUser(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	return userClient.GetUser(ctx, req)
}
