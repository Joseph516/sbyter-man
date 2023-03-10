package rpc

import (
	"context"
	"douyin_service/pkg/util"
	"douyin_service/services/kitex_gen/user"
	"douyin_service/services/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

var userClient userservice.Client

func NewUserClient() {
	r, err := etcd.NewEtcdResolver([]string{etcdResolver})
	if err != nil {
		log.Fatalf("Init NewEtcdResolver failed: %v", err)
	}

	userClient, err = userservice.NewClient(UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
		// client.WithHostPorts(UserServiceHostPorts), // 直接使用etcd获取服务
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

func GetUserByToken(ctx context.Context, req *user.UserTokenRequest) (resp *user.UserResponse, err error) {
	return userClient.GetUserByToken(ctx, req)
}

func ValidToken(ctx context.Context, req *user.TokenRequest) (bool, error) {
	resp, err := userClient.CheckValidToken(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.IsValid, nil
}
