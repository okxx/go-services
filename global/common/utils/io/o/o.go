package o

/*
	todo  0 		~ 1000 系统相关
	todo  1000 		~ 1999 用户相关
	todo  2000 		~ 2999 暂定
	todo  3000 		~ 3999 文件相关
 */
const (
	//TODO 系统错误码
	C_SUCCESS								= 200
	C_ERROR									= 400
	C_SERVER_ERROR							= 500
	//TODO 用户
	C_USER_NOT_LOGGED_OR_INVALID			= 4000//您处于未登陆状态，请先登录！
)


const (

	// 系统相关
	M_SUCCESS								= "成功"
	M_ERROR									= "未知错误"
	M_NO_PERMISSION							= "没有操作权限"
	M_OPERATION_FAIL						= "操作失败"
	M_SERVER_ERROR							= "服务器内部错误"
	M_INVALID_PARAMS						= "无效的请求参数"
	M_INVALID_ROUTER						= "未知的路由"
	M_INVALID_FUNCTION						= "未知的方法"

	// 用户相关
	M_USER_NOT_FOUND						= "用户不存在"
	M_USER_STATUS_INVALID					= "用户处于无效状态，无权操作，请联系客服"
	M_USER_REGISTER_FAIL					= "用户注册失败"
	M_USER_ALREADY_EXISTS					= "账号已存在，请登录"
	M_USER_TOKEN_INVALID					= "token无效"
	M_ERROR_AUTH_CHECK_TOKEN_TIMEOUT		= "token超时"
	M_USER_PASSWORD_IS_INCORRECT			= "账号密码不正确"
	M_USER_LOGIN_FAILURE					= "用户登录失败"
	M_USER_NOT_LOGGED_OR_INVALID			= "您处于未登陆状态，请先登录！"
	M_VERIFICATION_CODE_HAS_BEEN_SEND		= "验证码已发送,请勿重复发送。"
	M_VERIFICATION_CODE_SEND_FAIL			= "验证码发送失败，请稍后重试"
	M_VERIFICATION_CODE_ACTION_INVALID		= "发送验证码动作无效"
	M_INCORRECT_VERIFICATION_CODE			= "验证码不正确。"
	M_AUTHENTICATION_SYSYTEM_BUSY			= "The authentication system is busy, please try again later"//认证系统繁忙，请稍后重试。

	// 文件相关
	M_ERROR_UPLOAD_SAVE_IMAGE_FAIL    		= "保存图片失败"
	M_ERROR_UPLOAD_CHECK_IMAGE_FAIL   		= "检查图片失败"
	M_ERROR_UPLOAD_CHECK_IMAGE_FORMAT 		= "校验图片错误，图片格式或大小有问题"
	M_ERROR_IMAGE_DISTINGUISH_FAIL			= "图片分析失败，请稍后重试。"
	M_READ_FILE_FAIL						= "读取文件内容失败，请稍后重试"


	// 宠物相关
	M_IS_NOT_ANIMAL							= "这不是动物，请提供有效的宠物!"
	M_PET_EXISTED							= "宠物已存在，请勿宠物添加"
	M_PET_NOT_EXIST							= "宠物不存在"
	M_REGISTER_PET_ERROR					= "添加宠物失败。"
	M_EDIT_PET_ERROR						= "编辑宠物失败。"
)