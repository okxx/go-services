package sms

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/validation/v"
	"go-services/global/common/utils/io/o"
	"go-services/global/common/utils/sms"
	"go-services/global/common/utils/io"
)

// @Summary 校验验证码接口
// @Accept json
// @Tags 公共接口
// @Security Bearer
// @Produce  json
// @Param code query string true "code 验证码"
// @Param action query string true "action 短信的动作"
// @Param account query string true "account 手机号码"
// @Resource Name
// @Router /common/sms/check [POST]
// @Success 200 string string "ok"
func Check(c *gin.Context)  {
	code 	:= c.Query("code")
	action 	:= c.Query("action")
	account := c.Query("account")
	if !v.V(&struct {
		Code 	string `valid:"Required;MaxSize(6)"`
		Account string `valid:"Required;MaxSize(50)"`
	}{Code:code,Account:account}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}
	if !sms.CheckActions(action) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}
	key := action+"_"+account
	verify := sms.CheckCode(key,code)
	if !verify {
		io.App.Response(o.C_ERROR,o.M_INCORRECT_VERIFICATION_CODE,nil)
		return
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,nil)
}