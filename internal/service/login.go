package service

type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) Login(param *LoginRequest) (uint, bool, error) {
	return svc.dao.CheckUser(param.UserName, param.Password)
}

func (svc Service) Register(param *RegisterRequest) (uint, bool, error) {
	createUserRequest := CreateUserRequest{
		UserName: param.UserName,
		Password: param.Password,
	}
	uid, err := svc.CreateUser(&createUserRequest)
	if err != nil {
		return uid, false, err
	}
	getUserInfoRequest := GetUserInfoRequest{
		UserId: uid,
		Token:  "",
	}
	user, err := svc.GetUserById(&getUserInfoRequest)
	if err != nil {
		return user.ID, false, err
	}
	return  user.ID, true, nil
}
