package app

import (
	"douyin_service/global"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)
type Setting struct {
	vp *viper.Viper
}
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("../../configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

// SetupSetting 用于实现测试用例的全局变量初始化
func SetupSetting() error {
	setting, err := NewSetting()
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

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	global.Rd = redis.NewClient(&redis.Options{
		Addr:               global.RedisSetting.Addr,
		Password:           global.RedisSetting.Password,
		DB:                 0,
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
