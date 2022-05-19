package errcode

var (
	ErrorLoginFail       = NewError(10010001, "用户名或密码错误")
	ErrorLoginExpire     = NewError(10010002, "登录过期")
	ErrorRegisterFail    = NewError(10010003, "注册失败")
	ErrorGetUserInfoFail = NewError(10010004, "获取用户信息失败")
	// publish
	ErrorListPublishFail   = NewError(10020001, "获取用户的视频发布列表错误")
	ErrorActionPublishFail = NewError(10020002, "发布视频错误")

	ErrorActionFail      = NewError(10030001, "点赞失败")
	ErrorActionListFail  = NewError(10030002, "查询点赞作品失败")
)
