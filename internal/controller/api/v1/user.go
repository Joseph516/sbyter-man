package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Get 获取用户信息
func (u User) Get(c *gin.Context) {
	param := service.GetUserInfoRequest{}
	response := app.NewResponse(c)
	var res service.GetUserInfoResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.GetUserById(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetUserById err: %v", err)
		response.ToResponse(errcode.ErrorGetUserInfoFail)
		return
	}
	exist, followCnt, err := svc.QueryFollowCntRedis(user.ID)
	if err !=nil{
		global.Logger.Errorf("svc.QueryFollowCntRedis err: %v", err)
		response.ToResponse(errcode.ErrorGetUserInfoFail)
		return
	}
	if exist{
		user.FollowCount = followCnt
	}
	exist, fanCnt, err := svc.QueryFanCntRedis(user.ID)
	if err !=nil{
		global.Logger.Errorf("svc.QueryFanCntRedis err: %v", err)
		response.ToResponse(errcode.ErrorGetUserInfoFail)
		return
	}
	if exist{
		user.FollowerCount = fanCnt
	}

	// 获取新增的点赞字段
	totalFavorited, err := svc.GetTotalFavoritedById(user.ID)
	if err ==nil {
		user.TotalFavorited = totalFavorited // 更新数据库次数
	} else {
		global.Logger.Errorf("svc.GetTotalFavoritedById err: %v", err)
	}

	favoriteCount, err := svc.GetFavoriteCountById(user.ID)
	if err ==nil {
		user.FavoriteCount = favoriteCount // 更新数据库次数
	} else {
		global.Logger.Errorf("svc.GetFavoriteCountById err: %v", err)
	}
	isFollowReq := service.IsFollowRequest{
		A: param.UserId,
		B: user.ID,
	}
	flag, err := svc.IsFollow(isFollowReq)
	if err != nil {
		global.Logger.Errorf("svc.IsFollow err: %v", err)
	}
	res = service.GetUserInfoResponse{
		User: &service.UserInfo{
			ID:              user.ID,
			Name:            user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        flag,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
			TotalFavorited: user.TotalFavorited,
			FavoriteCount: user.FavoriteCount,
		},
	}
	res.StatusCode = 0
	res.StatusMsg = "获取信息成功"
	response.ToResponse(res)

	// 更新redis数据到mysql

	if followCnt != user.FollowCount || fanCnt != user.FollowerCount || totalFavorited != user.TotalFavorited || favoriteCount == user.FavoriteCount {
		req := &service.UpdateByIdRequest{
			UserId: user.ID,
			Data:   map[string]interface{}{
				"total_favorited": user.TotalFavorited,
				"favorite_count": user.FavoriteCount,
				"follow_count": user.FollowCount,
				"follower_count": user.FollowerCount,
			},
		}
		err = svc.UpdateById(req)
		if err != nil {
			global.Logger.Errorf("svc.UpdateById err: %v", err)
		}
	}
	//return	//多余的return
}
