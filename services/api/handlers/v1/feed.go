package v1

import (
	"context"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"douyin_service/services/api/rpc"
	"douyin_service/services/api/service"
	"douyin_service/services/kitex_gen/user"
	"douyin_service/services/kitex_gen/video"
	"github.com/gin-gonic/gin"
	"time"
)

type Feed struct {
}

func NewFeed() Feed {
	return Feed{}
}

func (f Feed) Feed(c *gin.Context) {
	// 根据lastTime获取最新的20条视频, len <= 20;
	// 此版本对任意uid都是返回同样的结果
	param := service.FeedRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, err := rpc.ValidToken(context.Background(), &user.TokenRequest{Token: param.Token})
	if !valid {
		response.ToErrorResponse(errcode.UnauthorizedTokenError.WithDetails(err.Error()))
		return
	}

	// 不管用户是否登陆、有没有Token，下面都是查询全体用户的所有视频，最后根据latest_time返回20条视频
	resp, err := rpc.Feed(context.Background(), &video.FeedRequest{
		LatestTime: &param.LastTime,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorFeedFail.WithDetails(err.Error()))
		return
	}

	var res service.FeedResponse
	res.StatusCode = 0
	res.StatusMsg = "获取视频列表成功"
	// TODO: 需要优化
	res.VideoList, res.NextTime, err = packFeed(resp.VideoList)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorFeedFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(res)
}

func packFeed(rpcVideo []*video.Video) (ginVideo []service.VideoInfo, nextTime int64, err error) {
	// TODO: 增加评论、点赞信息
	ginVideo = make([]service.VideoInfo, len(rpcVideo))
	nextTime = time.Now().Unix()
	for i := range rpcVideo {
		ginVideo[i] = service.VideoInfo{
			Id: rpcVideo[i].Id,
			Author: service.UserInfo{
				ID:   rpcVideo[i].Author.Id,
				Name: rpcVideo[i].Author.Name,
				// FollowCount:   *rpcVideo[i].Author.FollowCount,
				// FollowerCount: *rpcVideo[i].Author.FollowerCount,
				IsFollow: rpcVideo[i].Author.IsFollow,
			},
			PlayUrl:       rpcVideo[i].PlayUrl,
			CoverUrl:      rpcVideo[i].CoverUrl,
			FavoriteCount: rpcVideo[i].FavoriteCount,
			CommentCount:  rpcVideo[i].CommentCount,
			IsFavorite:    rpcVideo[i].IsFavorite,
			Title:         rpcVideo[i].Title,
		}
		if rpcVideo[i].PublishDate < nextTime {
			nextTime = rpcVideo[i].PublishDate
		}
	}
	return
}
