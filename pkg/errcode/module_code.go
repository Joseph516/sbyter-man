package errcode

var (
	ErrorLoginFail       = NewError(10010001, "登录失败")
	ErrorLoginExpire     = NewError(10010002, "登录过期")
	ErrorRegisterFail    = NewError(10010003, "注册失败")
	ErrorGetUserInfoFail = NewError(10010004, "获取用户信息失败")
	ErrorActionFail      = NewError(10020001, "点赞失败")
	ErrorActionListFail  = NewError(10020002, "查询点赞作品失败")
)
