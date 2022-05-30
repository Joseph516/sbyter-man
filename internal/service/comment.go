package service

type CommentListRequest struct {
	UserId  int64  `form:"user_id"  binding:"-"`
	VideoId int64  `form:"video_id"  binding:"required"`
	Token   string `form:"token" binding:"required"`
}

type CommentListResponse struct {
	ResponseCommon
	CommentList []CommentInfo `json:"comment_list"`
}

type CommentInfo struct {
	Id         uint     `json:"id"`
	Author     UserInfo `json:"user"` //评论的用户
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}

type CommentActionRequest struct {
	UserId      int64  `form:"user_id" binding:"-"`
	Token       string `form:"token" binding:"required"`
	VideoId     int64  `form:"video_id"  binding:"required"`
	ActionType  int64  `form:"action_type"  binding:"required"`
	CommentText string `form:"comment_text"  binding:"-"`
	CommentId   uint   `form:"comment_id"  binding:"-"`
}

const PUBCOMMENT int64 = 1

//const DELCOMMENT int64 = 2

func (svc *Service) CommentAction(param *CommentActionRequest) error {
	if param.ActionType == PUBCOMMENT {
		return svc.dao.PublishComment(param.VideoId, param.UserId, param.CommentText)
	}
	return svc.dao.DeleteComment(param.CommentId)
}

func (svc *Service) GetCommentList(video_id int64) (comResp CommentListResponse, err error) {
	//根据视频id获取作者id
	video, err := svc.dao.QueryVideoById(uint(video_id))
	if err != nil {
		return
	}
	//根据视频id获取评论列表
	comments, err := svc.dao.ListCommentByVideoId(video_id)
	if err != nil {
		return
	}
	//遍历赋值
	comResp.CommentList = make([]CommentInfo, len(comments))
	for i := range comments {
		//根据评论id找到用户id，以获取用户自身的信息
		user, err := svc.dao.GetUserById(uint(comments[i].UserId))
		if err != nil {
			//这样合适吗？
			return CommentListResponse{}, err
		}
		//根据用户id和视频作者id显查询是否关注
		isFollow, err := svc.dao.IsFollow(user.ID, video.AuthorId)
		if err != nil {
			return CommentListResponse{}, err
		}
		comResp.CommentList[i] = CommentInfo{
			Id: comments[i].ID,
			Author: UserInfo{
				ID:          user.ID,
				Name:        user.UserName,
				FollowCount: user.FollowCount,
				IsFollow:    isFollow,
			},
			Content:    comments[i].Content,
			CreateDate: comments[i].CreatedAt.Format("01-02"),
		}
	}
	return
}
