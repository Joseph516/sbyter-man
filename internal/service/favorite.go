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
	videosId, err := svc.dao.QueryFavoriteByUserId(param.UserId)
	if err != nil {
		return nil, err
	}
	videos, err := svc.QueryBatchVdieoById(videosId)
	authorsId := make([]int64, 0)
	for _, video := range videos {
		authorsId = append(authorsId, int64(video.ID))
	}
	authors, err := svc.GetUsersByIds(authorsId)
	if err != nil {
		return nil, err
	}

}
