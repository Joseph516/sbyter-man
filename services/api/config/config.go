package config

import "time"

var (
	ServerCfg *ServerSettingS
	AppCfg    *AppSettingS
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadZipsPath       string
	UploadServerUrl      string
	UploadServerZipUrl   string
	UploadVideoMaxSize   int
	UploadVideoAllowExts []string
}
