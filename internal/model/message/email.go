package message

import "encoding/json"

// Email 邮件消息
type Email struct {
	UserName []string `json:"user_name"`
	UserId   uint     `json:"user_id"`
	LoginIP  string   `json:"login_ip"`
	Token    string   `json:"token"`
}

func (e Email) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}
