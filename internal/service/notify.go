package service

type VerifyRequest struct {
	ID    uint   `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
	IP    string `form:"ip" binding:"required"`
}

type VerifyResponse struct {
	ResponseCommon
}
