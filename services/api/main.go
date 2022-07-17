package main

import (
	"douyin_service/global"
	setting2 "douyin_service/pkg/setting"
	"douyin_service/services/api/handlers"
	"douyin_service/services/api/rpc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mattn/go-colorable"
	"log"
	"net/http"
	"time"
)

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("OSS", &global.OSSSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Kafka", &global.KafkaSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	global.Rd = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Addr,
		Password: global.RedisSetting.Password,
		DB:       0,
	})

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	if err = global.SetupLogger(); err != nil {
		return err
	}
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// init rpc
	rpc.NewUserClient()
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := handlers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// svc := service.New(&gin.Context{})

	// go svc.Kafka.ConsumeEmail()  // 开启一个协程监听kafka邮件消息
	// go svc.Kafka.ConsumComment() // 开启一个协程监听kafka评论消息
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
