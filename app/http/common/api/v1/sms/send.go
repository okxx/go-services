package sms

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/io/o"
	"go-services/global/common/utils/validation/v"
	"go-services/global/common/utils/strings"
	"go-services/app/http/common/service/v1/sms"
	sms2 "go-services/global/common/utils/sms"
	"time"
	"go-services/global/common/cache/redis"
	sms3 "go-services/app/http/common/model/sms"
)

//发送验证码接口
type sendSmsValidation struct {
	Account 	string `valid:"Required;MaxSize(20)"`
}

// @Summary 发送短信验证码
// @Accept json
// @Tags 公共接口
// @Security Bearer
// @Produce  json
// @Param account query string true "account 用户手机号码"
// @Param action query string true "action 发送某种短信的动作"
// @Resource Name
// @Router /common/sms/send [POST]
// @Success 200 {object} sms.SmsSendResponse
func Send(c *gin.Context) {
	account := c.Query("account")
	action := c.Query("action")
	if !v.V(sendSmsValidation{Account:account}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}
	if ok := sms2.CheckActions(action); !ok {
		io.App.Response(o.C_ERROR,o.M_VERIFICATION_CODE_ACTION_INVALID,nil)
		return
	}

	key 	:= action+"_"+account
	smsStr	:= strings.Random(1,6)

	conn := redis.Server.Get()
	defer conn.Close()
	if ok,_ := redis.Bool(conn.Do("EXISTS",key)); ok {
		io.App.Response(o.C_ERROR,o.M_VERIFICATION_CODE_HAS_BEEN_SEND,nil)
		return
	}

	//将表中所有关于当前用户指定动作的验证码全部失效
	attributes 	:= map[string]interface{}{"account" : account, "mode" : action}
	values 		:= map[string]interface{}{"status" : 1, "updated_at" : time.Now()}
	if err := sms.Update(attributes,values);err != nil {
		io.App.Response(o.C_ERROR,o.M_VERIFICATION_CODE_SEND_FAIL,nil)
		return
	}

	service 	:= sms.SmsService{}
	service.Model.Account 	= account
	service.Model.Sms 		= smsStr
	service.Model.Mode 		= sms2.SMS_MODE_REGISTER//注册
	service.Model.Effective	= sms2.SMS_DEFAULT_EFFEXTIVE
	service.Model.CreatedAt	= time.Now()
	service.Model.UpdatedAt	= time.Now()
	if err := service.Create(); err != nil {
		io.App.Response(o.C_ERROR,o.M_VERIFICATION_CODE_SEND_FAIL,nil)
		return
	}

	if _,err := conn.Do("SET",key,smsStr,"EX",sms2.SMS_DEFAULT_EFFEXTIVE,"NX"); err != nil {
		io.App.Response(o.C_ERROR,o.M_VERIFICATION_CODE_SEND_FAIL,nil)
		return
	}

	smsSendResponse := sms3.SmsSendResponse{
		Sms		: smsStr,
		Hint	: "验证码已发送，有效期30分钟。",
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,smsSendResponse)
}