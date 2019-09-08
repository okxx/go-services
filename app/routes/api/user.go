package api

import (
	"go-services/global/common/server/app"
	"go-services/app/http/user/api/v1/auth"
)

//user interface
func UserInterface() {
	u := app.Router.Group("user")
	{
		ua := u.Group("auth")
		{
			ua.Any("login",auth.Login)//登录
			ua.Any("register",auth.Register)//注册
		}
	}
}