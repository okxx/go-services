package api

import (
	"go-services/app/http/common/api/v1/sms"
	"go-services/global/common/server/app"
)

//公共路由
func CommonInterface() {
	common := app.Router.Group("common")//公共api
	{
		cs := common.Group("sms")//短信
		{
			cs.GET("send",sms.Send)//发送验证码
			cs.GET("check",sms.Check)//校验验证码
		}
	}
}