package service

type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP  string `form:"login_ip"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP  string `form:"login_ip"`
}

type RegisterResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	// Avatar          string `json:"avatar"`           // 头像
	// Signature       string `json:"signature"`        // 个性签名
	// BackgroundImage string `json:"background_image"` // 背景图片
	// LoginIP         string `json:"login_ip"`         // 登录IP
	// TotalFavorited  int64    `json:"total_favorited"`  // 被赞的总次数
	// FavoriteCount   int64    `json:"favorite_count"`   //喜欢总数量
}

type GetUserInfoRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type GetUserInfoResponse struct {
	ResponseCommon
	User *UserInfo `json:"user" binding:"required"`
}
