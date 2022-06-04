package errcode

var (
	ErrorEmail           = NewError(10010001, "请输入合法邮箱")
	ErrorRegisterVerify  = NewError(10010002, "请查看邮箱邮件，完成注册！")
	ErrorRegisterAgain   = NewError(10010003, "请勿重复注册")
	ErrorRegisterFail    = NewError(10010004, "注册失败")
	ErrorLoginFail       = NewError(10010005, "用户名或密码错误")
	ErrorLoginExpire     = NewError(10010006, "登录过期")
	ErrorLoginDanger     = NewError(10010007, "登录IP异常, 请验证邮箱邮件")
	ErrorGetUserInfoFail = NewError(10010008, "获取用户信息失败")
	ErrorVerifyExpire    = NewError(10010009, "链接已过期")
	// publish
	ErrorListPublishFail   = NewError(10020001, "获取用户的视频发布列表错误")
	ErrorActionPublishFail = NewError(10020002, "发布视频错误")

	ErrorActionFail     = NewError(10050001, "点赞失败")
	ErrorActionListFail = NewError(10050002, "查询点赞作品失败")

	// Follow
	ErrorFollowActionFail = NewError(10040001, "关注用户失败")

	// comment
	ErrorListCommentFail   = NewError(10030001, "获取视频的评论列表错误")
	ErrorActionCommentFail = NewError(10030002, "评论操作错误")
)
