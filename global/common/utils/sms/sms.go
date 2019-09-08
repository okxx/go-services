package sms

import (
	"go-services/global/common/cache/redis"
)

const (

	//--短信有效期
	SMS_DEFAULT_EFFEXTIVE		= 1800

	//--短信模型
	SMS_MODE_REGISTER 			= "REGISTER"//注册
	SMS_MODE_LOGIN				= "LOGIN"//登录
	SMS_MODE_RESET_PASSWORD		= "RESET_PASSWORD"//找回密码
)


//检测动作是否有效
func CheckActions(action string) bool{
	result := false
	switch action {
	case SMS_MODE_REGISTER:
		result = true
	case SMS_MODE_LOGIN:
		result = true
	case SMS_MODE_RESET_PASSWORD:
		result = true
	default:
		result = false
	}
	return result
}

//验证码校验
func CheckCode(key,code string) bool {
	conn := redis.Server.Get()
	defer conn.Close()
	value,_ := redis.String(conn.Do("GET",key))
	if value != code {
		return false
	}
	return true
}