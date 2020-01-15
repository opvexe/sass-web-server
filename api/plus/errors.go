package plus

// 系统错误定义
var (
	PE_ServerError = NewWrapResponse(500, "服务器发生错误", 500)
)

//表单提交错误
var (
	PE_NotFoundUser        = NewWrapResponse(450, "用户不存在", 450)
	PE_UserFormatError     = NewWrapResponse(451, "用户格式错误", 451)
	PE_PassWordError       = NewWrapResponse(452, "密码错误", 452)
	PE_EmailFormatError    = NewWrapResponse(453, "邮箱格式不正确", 453)
	PE_EmailHasOccupy      = NewWrapResponse(454, "邮箱被占有", 454)
	PE_UserNameOccupy      = NewWrapResponse(455, "用户名被占有", 455)
	PE_IsNotEmptyName      = NewWrapResponse(456, "用户名/邮箱不能为空", 456)
	PE_ThirdAccountError   = NewWrapResponse(457, "第三方授权登录失败", 457)
	PE_UserNameHasSet      = NewWrapResponse(458, "用户名已设置，请勿重复设置", 458)
	PE_UpdateUserNameError = NewWrapResponse(459, "更新用户名失败", 459)
)
