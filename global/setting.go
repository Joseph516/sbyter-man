package global

import (
	"douyin_service/pkg/logger"
	"douyin_service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	OSSSetting      *setting.OSSSettingS
	DatabaseSetting *setting.DatabaseSettingS
	KafkaSetting    *setting.KafkaSettings
	RedisSetting    *setting.RedisSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)

func SetupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt,
		MaxSize:   600, // 600MB
		MaxAge:    10,  // 10天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
