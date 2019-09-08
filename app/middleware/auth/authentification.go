package auth

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/io/o"
	token2 "go-services/global/common/utils/auth/token"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"go-services/app/http/user/service/v1"
	"go-services/global/common/server/app"
)

//checking user
func Authentification(c *gin.Context) {
	token := c.Request.Header.Get(o.USER_TOKEN)//get token in header
	//_ := c.Request.Header.Get(o.TERMINAL)//get terminal in header
	if token == "" {//if token in header is null. then get token in cookies
		if cookie,err := c.Request.Cookie(o.USER_TOKEN); err != nil {
			io.App.Response(o.C_ERROR,o.M_USER_NOT_LOGGED_OR_INVALID,nil);c.Abort();return
		} else {
			token = cookie.Value
		}
	}

	if claims,err := token2.Parse(token); err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			io.App.Response(o.C_ERROR,o.M_ERROR_AUTH_CHECK_TOKEN_TIMEOUT,nil);c.Abort();return
		default:
			io.App.Response(o.C_ERROR,o.M_USER_NOT_LOGGED_OR_INVALID,nil);c.Abort();return
		}
	} else {
		if claims == nil {
			io.App.Response(o.C_ERROR,o.M_USER_NOT_LOGGED_OR_INVALID,nil);c.Abort();return
		}
		if pk,err := strconv.Atoi(claims.Pk); err != nil || pk < 1 {
			io.App.Response(o.C_ERROR,o.M_USER_NOT_LOGGED_OR_INVALID,nil);c.Abort();return
		} else {
			user := v1.GetAttributes(map[string]interface{}{"id":pk})
			if user.Id < 1 {
				io.App.Response(o.C_ERROR,o.M_USER_NOT_LOGGED_OR_INVALID,nil);c.Abort();return
			}
			if user.Status != 0 {
				io.App.Response(o.C_ERROR,o.M_USER_STATUS_INVALID,nil);c.Abort();return
			}
			app.User = user
		}
	}

}