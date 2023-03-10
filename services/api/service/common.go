package service

type ResponseCommon struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
