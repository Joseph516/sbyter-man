package service

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/util"
	"sync"
)

type ActionRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	VideoId    uint   `json:"video_id" form:"video_id" binding:"required"`
	ActionType int    `json:"action_type" form:"action_type" binding:"required"`
}

type ActionResponse struct {
	ResponseCommon
}

type FavoriteListRequest struct {
	UserId uint   `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type FavoriteListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

var lock sync.Mutex //用于双重检查

// Action 点赞
// 总体逻辑：
// 1.点赞/取消点赞
// 2.修改视频的点赞数量
// 3.修改视频作者的获赞数量
// 4.修改用户的点赞数量
func (svc *Service) Action(param *ActionRequest, userId uint) error {
	user := userId
	videoId := param.VideoId
	//查询video的作者id
	authorId, err := svc.QueryAuthorIdByVideoId(videoId)
	if err != nil {
		return err
	}
	action := param.ActionType
	ok, err := svc.IsFavor(user, videoId)
	if err != nil {
		return err
	}
	//点赞
	if action == 1 && !ok {
		err := svc.redis.FavorAction(user, videoId)
		if err != nil {
			return err
		}
		err = svc.changeVideoFavoriteCount(videoId, action)
		if err != nil {
			return err
		}
		err = svc.changeUserFavoritedCount(authorId, action)
		if err != nil {
			return err
		}
		err = svc.changeUserFavoriteCount(userId, action)
		return err
	}
	//取消点赞
	if action == 2 && ok {
		err := svc.redis.CancelFavorAction(user, videoId)
		if err != nil {
			return err
		}
		err = svc.changeVideoFavoriteCount(videoId, action)
		if err != nil {
			return err
		}
		err = svc.changeUserFavoritedCount(authorId, action)
		if err != nil {
			return err
		}
		err = svc.changeUserFavoriteCount(userId, action)
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
	videos, err := svc.QueryBatchVdieoById(videosId)
	//筛选video的authorId
	authorsId := make([]uint, 0)
	for _, video := range videos {
		authorsId = append(authorsId, video.AuthorId)
	}
	//查询author
	authors, err := svc.GetUsersByIds(authorsId)
	if err != nil {
		return nil, err
	}
	//构建{authorId: author}映射
	authorMap := make(map[uint]UserInfo, 0)
	for _, author := range authors {
		isFollow, err := svc.dao.IsFollow(param.UserId, author.ID)
		if err != nil {
			return nil, err
		}
		authorMap[author.ID] = UserInfo{
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
		isFavorite, _ := svc.IsFavor(video.AuthorId, video.ID)
		//video是从数据库查的，为了保证最新的点赞数量，应该先从缓存查
		favoriteCnt := video.FavoriteCount
		ok, cnt, err := svc.redis.QueryFavorCnt(video.ID)
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
func (svc *Service) IsFavor(userId uint, videoId uint) (bool, error) {
	return svc.redis.IsFavor(userId, videoId)
}

// QueryFavorCnt 获取video的点赞数量
func (svc *Service) QueryFavorCnt(videoId uint) (int64, error) {
	ok, cnt, err := svc.redis.QueryFavorCnt(videoId)
	if err != nil {
		return 0, err
	}
	if ok {
		return cnt, nil
	}
	cnt, err = svc.dao.QueryFavorCntById(videoId)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// changeVideoFavoriteCount 执行favoriteAction之后更改缓存和数据库中视频的favorite_count的操作
func (svc *Service) changeVideoFavoriteCount(videoId uint, action int) error {
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
		defer lock.Unlock()
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
	} else {
		if action == 1 {
			cnt = svc.redis.IncrFavorCnt(videoId)
		} else {
			cnt = svc.redis.DecrFavorCnt(videoId)
		}
	}
	return nil
}

// changeUserFavoritedCount 执行favoriteAction之后更改缓存和数据库中用户的total_favorited的操作
func (svc *Service) changeUserFavoritedCount(userId uint, action int) error {
	//先在缓存中尝试查找
	var (
		err   error
		exist bool
		user  model.User
		cnt   int64
		key   string
	)
	key = util.UserFavoritedCntKey(userId)
	exist, cnt, err = svc.redis.QueryUserFavoritedCount(userId)
	if err != nil {
		return err
	}
	if !exist {
		//保证只有一个请求是数据库的
		lock.Lock()
		//再查一次缓存
		exist, cnt, err = svc.redis.QueryUserFavoritedCount(userId)
		if err != nil {
			return err
		}
		if !exist {
			//走数据库
			user, err = svc.dao.GetUserById(userId)
			if err != nil {
				return err
			}
			cnt = user.TotalFavorited
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
				cnt = svc.redis.IncrUserFavoritedCnt(userId)
			} else {
				cnt = svc.redis.DecrUserFavoritedCnt(userId)
			}
		}
		lock.Unlock()
	} else {
		if action == 1 {
			cnt = svc.redis.IncrUserFavoritedCnt(userId)
		} else {
			cnt = svc.redis.DecrUserFavoritedCnt(userId)
		}
	}
	//缓存刷新数据库交给定时完成
	return nil
}

// changeUserFavoriteCount 执行favoriteAction之后更改缓存和数据库中用户的favorite_count的操作
func (svc *Service) changeUserFavoriteCount(userId uint, action int) error {
	//先在缓存中尝试查找
	var (
		err   error
		exist bool
		user  model.User
		cnt   int64
		key   string
	)
	key = util.UserFavoriteCntKey(userId)
	exist, cnt, err = svc.redis.QueryUserFavoriteCount(userId)
	if err != nil {
		return err
	}
	if !exist {
		//保证只有一个请求是数据库的
		lock.Lock()
		//再查一次缓存
		exist, cnt, err = svc.redis.QueryUserFavoriteCount(userId)
		if err != nil {
			return err
		}
		if !exist {
			//走数据库
			user, err = svc.dao.GetUserById(userId)
			if err != nil {
				return err
			}
			cnt = user.TotalFavorited
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
				cnt = svc.redis.IncrUserFavoriteCnt(userId)
			} else {
				cnt = svc.redis.DecrUserFavoriteCnt(userId)
			}
		}
		lock.Unlock()
	} else {
		if action == 1 {
			cnt = svc.redis.IncrUserFavoriteCnt(userId)
		} else {
			cnt = svc.redis.DecrUserFavoriteCnt(userId)
		}
	}
	//缓存刷新数据库交给定时完成
	return nil
}
