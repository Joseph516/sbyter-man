package global

import (
	"douyin_service/pkg/logger"
	"douyin_service/pkg/setting"
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
