package service

import "mime/multipart"

type PublishListRequest struct {
	UserId uint   `form:"user_id"` // 考虑到客户端在未登录下仍会尝试调用publish/list,因此将UserId和Token的required标签取消了
	Token  string `form:"token"`
}

type PublishListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type PublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data" binding:"required"`
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
}

type FeedRequest struct {
	LastTime int64  `form:"last_time"`
	Token    string `form:"token" binding:"required"`
}

type FeedResponse struct {
	ResponseCommon
	NextTime  int64       `json:"next_time"`
	VideoList []VideoInfo `json:"video_list"`
}
