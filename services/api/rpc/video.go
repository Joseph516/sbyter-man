package rpc

import (
	"context"
	"douyin_service/services/kitex_gen/video"
	"douyin_service/services/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

var videoClient videoservice.Client

func NewVideoClient() {
	r, err := etcd.NewEtcdResolver([]string{etcdResolver})
	if err != nil {
		log.Fatalf("Init NewEtcdResolver failed: %v", err)
	}
	videoClient, err = videoservice.NewClient(VideoServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(time.Minute),                // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
		// client.WithHostPorts(VideoServiceHostPorts), // 直接使用etcd获取服务
	)
	if err != nil {
		log.Fatalf("Init NewUserCilent failed: %v", err)
	}
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	return videoClient.PublishAction(ctx, req)
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	return videoClient.PublishList(ctx, req)
}

func Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	return videoClient.Feed(ctx, req)
}

func FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse, err error) {
	return videoClient.FavoriteAction(ctx, req)
}

func FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	return videoClient.FavoriteList(ctx, req)
}

func CommentAction(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse, err error) {
	return videoClient.CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *video.CommentListRequest) (resp *video.CommentListResponse, err error) {
	return videoClient.CommentList(ctx, req)
}
