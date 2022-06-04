package message

import "encoding/json"

// Email 邮件消息
type Email struct {
	UserName []string `json:"user_name"`
	Password string   `json:"password"` // 验证登录时为空
	UserId   uint     `json:"user_id"`
	LoginIP  string   `json:"login_ip"`
	Token    string   `json:"token"`
	Type     int      `json:"type"` // 判断邮件类型：1：注册， 2： 验证登录
}

func (e Email) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}
