package service

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"fmt"
	"sync"
)

type ActionRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	VideoId    int64  `json:"video_id" form:"video_id" binding:"required"`
	ActionType int    `json:"action_type" form:"action_type" binding:"required"`
}

type ActionResponse struct {
	ResponseCommon
}

type FavoriteListRequest struct {
	UserId int64  `json:"user_id" form:"user_id" binding:"required"`
	Token  string `json:"token" form:"token" binding:"required"`
}

type FavoriteListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

var lock sync.Mutex

// Action 点赞
func (svc *Service) Action(param *ActionRequest, userId int64) error {
	user := userId
	video := param.VideoId
	action := param.ActionType
	ok, err := svc.IsFavor(user, video)
	if err != nil {
		return err
	}
	//点赞
	if action == 1 && !ok {
		err := svc.redis.FavorAction(user, video)
		if err != nil {
			return err
		}
		err = svc.afterFavoriteAction(video, action)
		return err
	}
	//取消点赞
	if action == 2 && ok {
		err := svc.redis.CancelFavorAction(user, video)
		if err != nil {
			return err
		}
		err = svc.afterFavoriteAction(video, action)
		return err
	}
	return nil
}

// FavoriteList 用户点赞的列表
// 根据自身userId查询点赞的videoId
// 根据videoId查询点赞的video
// 根据video中的authorId查询对应的author
// 根据video和author拼接[]VideoInfo
func (svc *Service) FavoriteList(param *FavoriteListRequest) ([]VideoInfo, error) {
	var videoInfos []VideoInfo
	//查询videoId
	videosId, err := svc.redis.FavorList(param.UserId)
	if err != nil {
		return nil, err
	}
	//查询video
	//TODO：限制查询个数
	fmt.Println(videosId)
	videos, err := svc.QueryBatchVdieoById(videosId)
	//筛选video的authorId
	authorsId := make([]int64, 0)
	for _, video := range videos {
		authorsId = append(authorsId, video.AuthorId)
	}
	//查询author
	//TODO：限制查询个数
	authors, err := svc.GetUsersByIds(authorsId)
	if err != nil {
		return nil, err
	}
	//构建{authorId: author}映射
	authorMap := make(map[int64]UserInfo, 0)
	for _, author := range authors {
		//TODO：是否关注需要调用关注接口查询,先假设这是调用得到的结果。
		isFollow, err := svc.dao.IsFollow(param.UserId, int64(author.ID))
		if err != nil {
			return nil, err
		}
		authorMap[int64(author.ID)] = UserInfo{
			ID:              author.ID,
			Name:            author.UserName,
			FollowCount:     author.FollowCount,
			FollowerCount:   author.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          author.Avatar,
			Signature:       author.Signature,
			BackgroundImage: author.BackgroundImage,
		}
	}

	for _, video := range videos {
		//调用isFavor接口查询是否点赞了
		isFavorite, _ := svc.IsFavor(video.AuthorId, int64(video.ID))
		//video是从数据库查的，为了保证最新的点赞数量，应该先从缓存查
		favoriteCnt := video.FavoriteCount
		ok, cnt, err := svc.redis.QueryFavorCnt(int64(video.ID))
		if err != nil {
			return nil, err
		}
		if !ok {
			favoriteCnt = cnt
		}
		videoInfo := VideoInfo{
			Id:            video.ID,
			Author:        authorMap[video.AuthorId],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: favoriteCnt,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoInfos = append(videoInfos, videoInfo)
	}
	return videoInfos, nil
}

// IsFavor 查询是否点赞的功能
// 由于用户对哪些视频点赞使用bitmap存储到redis中，因此直接在redis查询。
func (svc *Service) IsFavor(userId int64, videoId int64) (bool, error) {
	return svc.redis.IsFavor(userId, videoId)
}

// QueryFavorCnt 获取video的点赞数量(先查缓存再查数据库)
func (svc *Service) QueryFavorCnt(videoId int64) (int64, error) {
	ok, cnt, err := svc.redis.QueryFavorCnt(videoId)
	if err != nil {
		return 0, err
	}
	if ok {
		return cnt, err
	}
	cnt, err = svc.dao.QueryFavorCntById(videoId)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// afterFavoriteAction 执行favoriteAction之后更改缓存和数据库中favorite_count的操作
func (svc *Service) afterFavoriteAction(videoId int64, action int) error {
	//先在缓存中尝试查找
	var (
		err   error
		exist bool
		cnt   int64
		key   string
	)
	key = util.VideoFavorCntKey(videoId)
	exist, cnt, err = svc.redis.QueryFavorCnt(videoId)
	if err != nil {
		return err
	}
	if !exist {
		//保证只有一个请求是数据库的
		lock.Lock()
		//再查一次缓存
		exist, cnt, err = svc.redis.QueryFavorCnt(videoId)
		if err != nil {
			return err
		}
		if !exist {
			//走数据库
			cnt, err = svc.dao.QueryFavorCntById(videoId)
			if err != nil {
				return err
			}
			if action == 1 {
				cnt++
			} else {
				cnt--
			}
			err = svc.redis.Set(key, cnt, 0)
			if err != nil {
				return err
			}
		} else {
			//走缓存
			if action == 1 {
				cnt = svc.redis.IncrFavorCnt(videoId)
			} else {
				cnt = svc.redis.DecrFavorCnt(videoId)
			}
		}
		lock.Unlock()
	} else {
		if action == 1 {
			cnt = svc.redis.IncrFavorCnt(videoId)
		} else {
			cnt = svc.redis.DecrFavorCnt(videoId)
		}
	}

	var newVideo model.Video
	newVideo.ID = uint(videoId)
	newVideo.FavoriteCount = cnt
	//更新数据库，已经交给定时任务

	return nil
}
