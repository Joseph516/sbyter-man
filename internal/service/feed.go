package service

import (
	"time"
)

type FeedRequest struct {
	LastTime int64  `form:"last_time"`
	Token    string `form:"token"`
}

type FeedResponse struct {
	ResponseCommon
	NextTime  int64       `json:"next_time"`
	VideoList []VideoInfo `json:"video_list"`
}

func (svc *Service) Feed(lastTime int64) (pubResp FeedResponse, err error) {
	// 根据lastTime获取最新的20条视频, len <= 20
	videos, err := svc.dao.GetLatestVideos(lastTime)
	if err != nil {
		return
	}

	// 获取video中的authorId
	uids := make([]uint, len(videos))
	for i := range videos {
		uids[i] = videos[i].AuthorId
	}
	// 根据用户id切片获取用户自身信息
	users, err := svc.dao.GetUsersByIds(uids)
	if err != nil {
		return
	}

	// 建立用户id到用户信息的map映射
	map_user := make(map[uint]UserInfo)
	for _, user := range users {
		map_user[user.ID] = UserInfo{
			ID:            user.ID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
	}

	// 遍历赋值
	pubResp.VideoList = make([]VideoInfo, len(videos))
	nextTime := time.Now().Unix()
	for i, video := range videos {
		pubResp.VideoList[i] = VideoInfo{
			Id:            video.ID,
			Author:        map_user[video.AuthorId],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.FavoriteCount != 0,
			Title:         video.Title,
		}
		if video.PublishDate.Unix() < nextTime {
			nextTime = video.PublishDate.Unix()
		}
	}
	pubResp.NextTime = nextTime
	return
}

/* 这个函数的原始版本
func (svc *Service) Feed(userId uint, lastTime int64) (pubResp FeedResponse, err error) {
	// 根据用户id获取发布视频信息
	video, err := svc.dao.ListVideoByUserId(userId)
	if err != nil {
		return
	}

	// 根据用户id获取用户自身信息
	user, err := svc.dao.GetUserById(userId)
	if err != nil {
		return
	}

	// 遍历赋值
	pubResp.VideoList = make([]VideoInfo, len(video))
	nextTime := time.Now().Unix()
	for i := range video {
		// 视频时间限制
		if video[i].PublishDate.Unix() > lastTime {
			continue
		}

		pubResp.VideoList[i] = VideoInfo{
			Id: video[i].ID,
			Author: UserInfo{
				ID:            user.ID,
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video[i].PlayUrl,
			CoverUrl:      video[i].CoverUrl,
			FavoriteCount: video[i].FavoriteCount,
			CommentCount:  video[i].CommentCount,
			IsFavorite:    video[i].FavoriteCount != 0,
			Title:         video[i].Title,
		}
		if video[i].PublishDate.Unix() < nextTime {
			nextTime = video[i].PublishDate.Unix()
		}
	}
	pubResp.NextTime = nextTime
	return
}
*/
