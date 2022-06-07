package main

import (
	"douyin_service/cronjob"
	"douyin_service/global"
	"douyin_service/internal/controller"
	"douyin_service/internal/kafka"
	"douyin_service/internal/model"
	"douyin_service/internal/service"
	"douyin_service/pkg/email"
	"douyin_service/pkg/logger"
	setting2 "douyin_service/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mattn/go-colorable"
	"github.com/robfig/cron/v3"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupKafka()
	if err != nil {
		log.Fatalf("init.setupKafka err: %v", err)
	}

	setupEmail()

	err = setupCron()
	if err != nil {
		log.Fatalf("init.setupCron err: %v", err)
	}

}

// @title 抖音平台
// @version 1.0
// @description 抖音后端服务
// @termsOfService https://*****
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := controller.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	svc := service.New(&gin.Context{})

	go svc.Kafka.ConsumeEmail()  // 开启一个协程监听kafka邮件消息
	go svc.Kafka.ConsumComment() // 开启一个协程监听kafka评论消息
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
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

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, // 600MB
		MaxAge:    10,  // 10天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupKafka() error {
	var err error
	global.Consumer, err = kafka.NewConsumer()

	if err != nil {
		return err
	}
	global.SyncProducer, err = kafka.NewSyncProducer()
	if err != nil {
		return err
	}
	return nil
}

func setupEmail() {
	smtpInfo := email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	}
	global.Email = email.NewEmail(&smtpInfo)
}

//设置定时任务
func setupCron() error {
	dc := cronjob.New()
	c := cron.New(cron.WithSeconds())
	//TODO:定时刷新用户点赞数量和被赞数量
	// 生成chain
	favorSkipChain1 := cronjob.SkipIfStillRunningChain()
	followSkipChain1 := cronjob.SkipIfStillRunningChain()

	// 生成job
	favorCntFlashJob := cronjob.GenerateJob(&favorSkipChain1, dc.FlashFavorCnt)
	followCntFlashJob := cronjob.GenerateJob(&followSkipChain1, dc.FlashFollowCnt)

	userFavoritedCntFlashJob := cronjob.GenerateJob(&favorSkipChain1, dc.FlashUserFavoritedCnt)
	userFavoriteCntFlashJob := cronjob.GenerateJob(&favorSkipChain1, dc.FlashUserFavoriteCnt)

	fanCntFlashJob := cronjob.GenerateJob(&followSkipChain1, dc.FlashFanCnt)
	global.Logger.Info("启动点赞数量定时刷新任务")

	//向cron注册经过对应chain修饰的job
	_, err := c.AddJob(cronjob.FlashFavorCnt, favorCntFlashJob)
	if err != nil {
		return err
	}
	_, err = c.AddJob(cronjob.FlashUserFavoritedCnt, userFavoritedCntFlashJob)
	if err != nil {
		return err
	}
	_, err = c.AddJob(cronjob.FlashUserFavoriteCnt, userFavoriteCntFlashJob)
	if err != nil {
		return err
	}
	_, err = c.AddJob(cronjob.FOLLOWCNTTIME, followCntFlashJob)
	if err != nil {
		return err
	}
	_, err = c.AddJob(cronjob.FOLLOWCNTTIME, fanCntFlashJob)
	if err != nil {
		return err
	}
	//开启
	c.Start()
	return nil
}
