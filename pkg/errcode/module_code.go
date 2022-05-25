package errcode

var (
	ErrorRegisterFail    = NewError(10010001, "注册失败")
	ErrorLoginFail       = NewError(10010002, "用户名或密码错误")
	ErrorLoginExpire     = NewError(10010003, "登录过期")
	ErrorLoginDanger     = NewError(10010004, "登录IP异常, 请验证邮箱邮件")
	ErrorGetUserInfoFail = NewError(10010005, "获取用户信息失败")
	ErrorVerifyExpire     = NewError(10010006, "链接已过期")
	// publish
	ErrorListPublishFail   = NewError(10020001, "获取用户的视频发布列表错误")
	ErrorActionPublishFail = NewError(10020002, "发布视频错误")

	ErrorActionFail      = NewError(10050001, "点赞失败")
	ErrorActionListFail  = NewError(10050002, "查询点赞作品失败")

	// Follow
	ErrorFollowActionFail = NewError(10040001, "关注用户失败")

	// comment
	ErrorListCommentFail   = NewError(10030001, "获取视频的评论列表错误")
	ErrorActionCommentFail = NewError(10030002, "评论操作错误")

)
