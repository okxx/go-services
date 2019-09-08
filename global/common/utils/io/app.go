package io

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppServer struct {
	Context *gin.Context
}
var App = &AppServer{}

// response function
func (a *AppServer)Response(code int,msg string, data interface{}) {
	if data == nil { data = gin.H{} }
	a.Context.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}