package main

import (
	"context"
	"douyin_service/services/kitex_gen/user"
	"douyin_service/services/kitex_gen/video"
	"douyin_service/services/video/repository"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	err = repository.NewVideoDaoInstance().PublishVideo(uint(req.UserId), req.PlayUrl, req.CoverUrl, req.Title)
	resp = new(video.PublishActionResponse)
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	videos, err := repository.NewVideoDaoInstance().ListVideoByUserId(uint(req.UserId))
	if err != nil {
		resp = &video.PublishListResponse{
			VideoList: nil,
		}
		return
	}
	resp = new(video.PublishListResponse)
	resp.VideoList = make([]*video.Video, len(videos))
	for i := range videos {
		resp.VideoList[i] = new(video.Video)
		resp.VideoList[i].Id = int64(videos[i].ID)
		resp.VideoList[i].Author = &user.User{
			Id: int64(videos[i].AuthorId),
		}
		resp.VideoList[i].Title = videos[i].Title
		resp.VideoList[i].PlayUrl = videos[i].PlayUrl
		resp.VideoList[i].CoverUrl = videos[i].CoverUrl
		resp.VideoList[i].CommentCount = videos[i].CommentCount
		resp.VideoList[i].FavoriteCount = videos[i].FavoriteCount
	}
	return
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	videos, err := repository.NewVideoDaoInstance().FeedVideo(*req.LatestTime, 20)
	if err != nil {
		resp = &video.FeedResponse{
			VideoList: nil,
		}
		return
	}
	resp = new(video.FeedResponse)
	resp.VideoList = make([]*video.Video, len(videos))
	for i := range videos {
		resp.VideoList[i] = new(video.Video)
		resp.VideoList[i].Id = int64(videos[i].ID)
		resp.VideoList[i].Author = &user.User{
			Id: int64(videos[i].AuthorId),
		}
		resp.VideoList[i].Title = videos[i].Title
		resp.VideoList[i].PlayUrl = videos[i].PlayUrl
		resp.VideoList[i].CoverUrl = videos[i].CoverUrl
		resp.VideoList[i].CommentCount = videos[i].CommentCount
		resp.VideoList[i].FavoriteCount = videos[i].FavoriteCount
		resp.VideoList[i].PublishDate = videos[i].PublishDate.Unix()
	}
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *video.FavoriteActionRequest) (resp *video.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *video.CommentListRequest) (resp *video.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
