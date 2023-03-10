package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type OSSSettingS struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

// UploadOSS 传输视频到OSS存储
func UploadOSS(localPath, ossPath string, cfg OSSSettingS) error {
	endpoint := cfg.Endpoint
	accessKeyID := cfg.AccessKeyId
	accessKeySecret := cfg.AccessKeySecret
	bucketName := cfg.Bucket
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	fd, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer fd.Close()

	err = bucket.PutObject(ossPath, fd)
	if err != nil {
		return err
	}

	return nil
}
