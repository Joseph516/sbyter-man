package errcode

var (
	ErrorLoginFail       = NewError(10010001, "登录失败")
	ErrorGetUserInfoFail = NewError(10010003, "获取用户信息失败")
	ErrorActionFail      = NewError(10020001, "操作失败")
)
