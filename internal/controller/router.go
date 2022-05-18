package controller

import (
	"douyin_service/global"
	v1 "douyin_service/internal/controller/api/v1"
	"douyin_service/internal/middleware"
	"douyin_service/pkg/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	apiv1 := r.Group("/douyin/")
	apiv1.POST("/user/login/", user.Login)
	apiv1.POST("/user/register/", user.Register)
	apiv1.POST("/favorite/action", favorite.Action)
	apiv1.Use()
	{
		apiv1.GET("/user/", user.Get)
		apiv1.GET("favorite/list/", favorite.FavoriteList)
	}

	return r
}
