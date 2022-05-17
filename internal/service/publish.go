package service

type PublishListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type VideoInfo struct {
	Id            uint     `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type PublishListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

func (svc *Service) PublishList(userId uint) (pubResp PublishListResponse, err error) {
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
	for i := range video {
		pubResp.VideoList[i] = VideoInfo{
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
	}
	return
}
