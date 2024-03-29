package config

import (
	"time"
)

var (
	JwtCfg      *JwtConfig
	DatabaseCfg *DatabaseConfig
)

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
