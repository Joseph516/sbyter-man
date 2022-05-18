package service

type CommentListRequest struct {
	UserId  int64  `form:"user_id"  binding:"required"`
	VideoId int64  `form:"video_id"  binding:"required"`
	Token   string `form:"token" binding:"required"`
}

type CommentListResponse struct {
	ResponseCommon
	CommentList []CommentInfo `json:"comment_list"`
}

type CommentInfo struct {
	Id         uint     `json:"id"`
	Author     UserInfo `json:"author"` //评论的用户
	Content    string   `json:"content"`
	CreateDate int64    `json:"create_date"`
}

func (svc *Service) GetCommentList(video_id int64) (comResp CommentListResponse, err error) {
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
		comResp.CommentList[i] = CommentInfo{
			Id: comments[i].ID,
			Author: UserInfo{
				ID:          user.ID,
				Name:        user.UserName,
				FollowCount: user.FollowCount,
				IsFollow:    false,
			},
			Content:    comments[i].Content,
			CreateDate: comments[i].CreatedAt,
		}
	}
	return
}
