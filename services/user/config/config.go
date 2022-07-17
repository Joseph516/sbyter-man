package config

import (
	"github.com/spf13/viper"
	"time"
)

var (
	JwtCfg      *JwtConfig
	DatabaseCfg *DatabaseConfig
)

type ViperConfig struct {
	vp *viper.Viper
}

func NewViperConfig() (*ViperConfig, error) {
	vp := viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &ViperConfig{vp}, nil
}

func (s *ViperConfig) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

type JwtConfig struct {
	Key    string
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseConfig struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
