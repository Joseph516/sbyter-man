package oss

import (
	"douyin_service/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// UploadOSS 传输视频到OSS存储
func UploadOSS(localPath, ossPath string) error {
	endpoint := global.OSSSetting.Endpoint
	accessKeyID := global.OSSSetting.AccessKeyId
	accessKeySecret := global.OSSSetting.AccessKeySecret
	bucketName := global.OSSSetting.Bucket
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		global.Logger.Errorf("oss.New err: %v", err)
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		global.Logger.Errorf("client.Bucket err: %v", err)
		return err
	}

	fd, err := os.Open(localPath)
	if err != nil {
		global.Logger.Errorf("os.Open err: %v", err)
		return err
	}
	defer fd.Close()

	err = bucket.PutObject(ossPath, fd)
	if err != nil {
		global.Logger.Errorf("bucket.PutObject err: %v", err)
		return err
	}

	return nil
}
