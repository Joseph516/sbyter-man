package util

import (
	"douyin_service/global"
	"douyin_service/pkg/errcode"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// SendVerifiedEmail 发送验证短信, 这里to可以是多个人
func SendVerifiedEmail(to []string, id uint, ip, token string) error {
	subject := "账号报警"
	// 下面链接实际上线需要修改为公网IP
	body := fmt.Sprintf(`<b>您的抖音账号登录异常，登录IP为%s，</b> <a tagert="_blank" href="http://127.0.0.1:8000/douyin/verify?id=%d&token=%s&ip=%s">
点击此处允许登录</a>`, ip, id, token, ip)
	err := global.Email.SendMail(to, subject, body)
	if err != nil {
		return err
	}
	global.Rd.Set(ip, token, 60 * 10 * time.Second) // 设置一个10分钟过期的token
	return nil
}

// VerifyEmail 验证短信是否真实
func VerifyEmail(ip, token string) (bool, error) {
	value := global.Rd.Get(ip)
	fmt.Println(value.Err(), value.Val())
	if value.Err() == redis.Nil || value.Val() != token { // 邮件已经过期
		return false, errcode.ErrorVerifyExpire
	}
	global.Rd.Del(token) // 删除该Token
	return true, nil
}
