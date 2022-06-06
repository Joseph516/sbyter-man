package controller

import (
	"douyin_service/global"
	v1 "douyin_service/internal/controller/api/v1"
	"douyin_service/internal/middleware"
	"douyin_service/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second * 10,
	Capacity:     20,
	Quantum:      20,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	user := v1.NewUser()
	favorite := v1.NewFavorite()
	pub := v1.NewPublish()
	fol := v1.NewFollow()
	com := v1.NewComment()
	feed := v1.NewFeed()
	notify := v1.NewNotify()
	apiv1 := r.Group("/douyin/")
	apiv1.Use()
	{
		// user
		apiv1.POST("/user/login/", user.Login)
		apiv1.POST("/user/register/", user.Register)
		apiv1.GET("/user/", user.Get)
		apiv1.GET("/verifyRegister/", notify.VerifyRegister)
		apiv1.GET("/verifyLogin/", notify.VerifyLogin)

		// favor
		apiv1.GET("favorite/list/", favorite.FavoriteList)
		apiv1.POST("/favorite/action/", favorite.Action)

		// publish
		apiv1.GET("/publish/list/", pub.List)
		apiv1.POST("/publish/action/", pub.Action)

		//comment
		apiv1.GET("/comment/list/", com.List)
		apiv1.POST("/comment/action/", com.CommentAction)

		// feed
		apiv1.GET("/feed/", feed.Feed)

		//follow
		apiv1.POST("/relation/action/", fol.Action)
		apiv1.GET("/relation/follower/list/", fol.FollowerList)
		apiv1.GET("/relation/follow/list/", fol.FollowList)
	}


	return r
}
