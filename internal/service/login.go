package service

type LoginRequest struct {
	UserName string `form:"username" binding:"required,len=11"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint    `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) Login(param *LoginRequest) (uint, bool, error) {
	return svc.dao.CheckUser(param.UserName, param.Password)
}

