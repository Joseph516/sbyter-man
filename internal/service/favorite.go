package service

type ActionRequest struct {
	UserId     int64  `json:"user_id" form:"user_id" binding:"required"`
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

func (svc *Service) Action(param *ActionRequest) error {
	user := param.UserId
	video := param.VideoId
	action := param.ActionType
	ok, err := svc.dao.IsFavor(user, video)
	if err != nil {
		return err
	}
	//点赞
	if action == 1 && !ok {
		err := svc.dao.FavorAction(user, video)
		return err
	}
	//取消点赞
	if action == 2 && ok {
		err := svc.dao.CancelFavorAction(user, video)
		return err
	}
	return nil

}

//FavoriteList 用户点赞的列表
// 根据自身userId查询点赞的videoId
// 根据videoId查询点赞的video
// 根据video中的authorId查询对应的author
// 根据video和author拼接[]VideoInfo
func (svc *Service) FavoriteList(param *FavoriteListRequest) ([]VideoInfo, error) {
	var videoInfos []VideoInfo
	//查询videoId
	videosId, err := svc.dao.QueryFavoriteByUserId(param.UserId)
	if err != nil {
		return nil, err
	}
	//查询video
	videos, err := svc.QueryBatchVdieoById(videosId)
	//筛选video的authorId
	authorsId := make([]int64, 0)
	for _, video := range videos {
		authorsId = append(authorsId, int64(video.ID))
	}
	//查询author
	authors, err := svc.GetUsersByIds(authorsId)
	if err != nil {
		return nil, err
	}
	//构建{authorId: author}映射
	authorMap := make(map[int64]*UserInfo, 0)
	for _, author := range authors {
		//TODO：是否关注需要调用关注接口查询,先假设这是调用得到的结果。
		isFollow := false
		authorMap[int64(author.ID)] = &UserInfo{
			ID:            author.ID,
			Name:          author.UserName,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      isFollow,
		}
	}

	for _, video := range videos {
		//调用isFavor接口查询是否点赞了
		isFavorite, _ := svc.IsFavor(video.AuthorId, int64(video.ID))
		videoInfo := VideoInfo{
			Id:            video.ID,
			Author:        *authorMap[int64(video.ID)],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoInfos = append(videoInfos, videoInfo)
	}
	return videoInfos, nil
}

// IsFavor 查询是否点赞的功能，暂定只走数据库
func (svc *Service) IsFavor(userId int64, videoId int64) (bool, error) {
	return svc.dao.IsFavor(userId, videoId)
}
