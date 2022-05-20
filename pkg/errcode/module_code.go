package errcode

var (
	ErrorLoginFail       = NewError(10010001, "用户名或密码错误")
	ErrorLoginExpire     = NewError(10010002, "登录过期")
	ErrorRegisterFail    = NewError(10010003, "注册失败")
	ErrorGetUserInfoFail = NewError(10010004, "获取用户信息失败")
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
