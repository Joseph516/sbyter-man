package setting

import "time"

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

// OSSSettingS 阿里云OSS设置
type OSSSettingS struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

type DatabaseSettingS struct {
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

type KafkaSettings struct {
	Host             string
	TopicEmail       string // 发送邮件主题
	TopicComment     string //发送评论主题
	TopicCompression string
	DefaultPartition int
	ConsumerGroupID  string
}

type RedisSettingS struct {
	Addr     string // ip和端口号
	Password string
	DB       int
}

type JWTSettingS struct {
	Key    string
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
