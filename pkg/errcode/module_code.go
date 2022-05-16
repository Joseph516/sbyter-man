package errcode

var (
	ErrorLoginFail       = NewError(10010001, "登录失败")
	ErrorLoginExpire     = NewError(10010002, "登录过期")
	ErrorRegisterFail    = NewError(10010003, "注册失败")
	ErrorGetUserInfoFail = NewError(10010004, "获取用户信息失败")
)
