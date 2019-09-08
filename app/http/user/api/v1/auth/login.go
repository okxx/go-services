package auth

import (
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/io/o"
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/validation"
	"go-services/app/http/user/service/v1"
	"go-services/app/http/user/model/user"
	ticket2 "go-services/global/common/utils/auth/ticket"
	"go-services/global/common/utils/auth/token"
	"strconv"
	password2 "go-services/global/common/utils/auth/password"
	"go-services/global/common/utils/time/time"
	"go-services/global/common/utils/cookie"
)

/*
	app login interface
 */

type userLoginValidation struct {
	Account string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(255)"`
}

// @Summary 用户登录
// @Accept json
// @Tags 用户接口
// @Security Bearer
// @Produce  json
// @Param account query string true "account 手机号码"
// @Param password query string true "password 密码"
// @Resource Name
// @Router /user/auth/login [POST]
// @Success 200 {object} user.UserLoginResponse
func Login(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	valid := validation.Validation{}
	if ok,_ :=  valid.Valid(&userLoginValidation{Account:account,Password:password}); !ok {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}

	//获取用户是否存在
	userModel := v1.GetAttributes(map[string]interface{}{"account" : account})
	if userModel.Id < 1 {
		io.App.Response(o.C_ERROR,o.M_USER_NOT_FOUND,nil)
		return
	}
	if userModel.Status != 0 {
		io.App.Response(o.C_ERROR,o.M_USER_STATUS_INVALID,nil)
		return
	}

	userCreatedAt := time.DateToTimestamp(userModel.CreatedAt)
	buildPassword := password2.New(password,userModel.Salting,userCreatedAt)
	if buildPassword != userModel.Password {
		io.App.Response(o.C_ERROR,o.M_USER_PASSWORD_IS_INCORRECT,nil)
		return
	}

	//更新ticket
	ticket 		:= ticket2.New(account)
	attributes 	:= map[string]interface{}{"account" : account}
	values 		:= map[string]interface{}{"ticket" : ticket,"updated_at": time.CurrentTimeNow()}
	if err := v1.SetAttributes(attributes,values);err != nil {
		io.App.Response(o.C_ERROR,o.M_USER_LOGIN_FAILURE,nil)
		return
	}

	buildToken,err 	:= token.Build(strconv.Itoa(userModel.Id),userModel.Nickname,buildPassword)
	if err != nil {
		io.App.Response(o.C_ERROR,o.M_AUTHENTICATION_SYSYTEM_BUSY,nil)
		return
	}

	//set token and login ticket to cookie
	cookies := make(map[string]string)
	cookies["token"] = buildToken
	cookies["ticket"] = ticket
	cookie.Set(cookies)
	response := user.UserLoginResponse{
		Account:account,
		Ticket:ticket,
		Token: buildToken,
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,response)
}