package o

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/io"
)

const (
	USER_TOKEN			= "token"
	USER_TICKET			= "ticket"
	TERMINAL			= "terminal"
	AllowHeaders 		= "Content-Type,TERMINAL,TOKEN,TICKET,VERSION"
)

//gateway
func Gateway(c *gin.Context) {

	//构造请求信息
	io.App.Context = c

	if origin := c.GetHeader("Access-Control-Allow-Origin");origin != ""{
		c.Header("Access-Control-Allow-Origin",origin)
	} else {
		c.Header("Access-Control-Allow-Origin","*")
	}
	c.Header("Access-Control-Allow-Credentials","true")
	c.Header("Access-Control-Allow-Headers",AllowHeaders)
	c.Header("Access-Control-Allow-Methods","*")
	c.Header("Access-Control-Allow-Options","*")
}

//no route
func NoRoute(c *gin.Context) {
	io.App.Response(C_ERROR,M_INVALID_ROUTER,nil)
	return
}

//no method
func NoMethod(c *gin.Context) {
	io.App.Response(C_ERROR,M_INVALID_FUNCTION,nil)
	return
}