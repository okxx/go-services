package auth

import (
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/io/o"
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/validation/v"
	"go-services/global/common/utils/sms"
	"go-services/global/common/utils/time/time"
	"go-services/global/common/utils/strings"
	"go-services/app/http/user/service/v1"
	password2 "go-services/global/common/utils/auth/password"
	user2 "go-services/app/http/user/model/user"
	"go-services/global/common/cache/redis"
)

type userRegisterValidation struct {
	Account 	string `valid:"Required;MaxSize(50)"`
	Code 		string `valid:"Required;MaxSize(6)"`
	Password 	string `valid:"Required;MaxSize(255)"`
}

// @Summary 用户注册
// @Accept json
// @Tags 用户接口
// @Security Bearer
// @Produce  json
// @Param code query string true "code 验证码"
// @Param account query string true "account 手机号码"
// @Param password query string true "password 密码"
// @Resource Name
// @Router /user/auth/register [POST]
// @Success 200 {object} user.UserResponse
func Register(c *gin.Context) {
	code 		:= c.Query("code")
	account 	:= c.Query("account")
	password 	:= c.Query("password")
	if !v.V(&userRegisterValidation{Code:code,Account:account,Password:password}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}

	//校验验证码
	key := sms.SMS_MODE_REGISTER+"_"+account
	if !sms.CheckCode(key,code) {
		io.App.Response(o.C_ERROR,o.M_INCORRECT_VERIFICATION_CODE,nil)
		return
	}

	//检测用户是否存在
	attributes := map[string]interface{}{"account" : account}
	userModel := v1.GetAttributes(attributes)
	if userModel.Id > 0 {
		io.App.Response(o.C_ERROR,o.M_USER_ALREADY_EXISTS,nil)
		return
	}

	user 		:= v1.AuthService{}
	salting		:= strings.Random(5,25)
	timestamp 	:= time.CurrentTimeNow()
	user.UserModel.Account 		= account
	user.UserModel.Password 	= password2.New(password,salting,time.DateToTimestamp(timestamp))
	user.UserModel.Salting 		= salting
	user.UserModel.Nickname		= strings.Random(5,25)
	user.UserModel.CreatedAt	= timestamp
	user.UserModel.UpdatedAt	= timestamp
	if err := user.Register(); err != nil {
		io.App.Response(o.C_ERROR,o.M_OPERATION_FAIL,nil)
		return
	}

	//删除验证码
	conn := redis.Server.Get()
	defer conn.Close()
	conn.Do("DEL",key)

	//构造返回值
	response := user2.UserResponse{
		Id:user.UserModel.Id,
		Nickname:user.UserModel.Nickname,
		Account:account,
		Status:user.UserModel.Status,
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,response)
}