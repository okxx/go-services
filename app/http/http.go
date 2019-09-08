package http

import (
	_ "go-services/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-services/global/common/utils/io/o"
	"go-services/global/common/server/app"
	"go-services/app/routes/api"
)

func Routes () *gin.Engine {
	app.Router = gin.New()
	app.Router.NoRoute(o.NoRoute)
	app.Router.NoMethod(o.NoMethod)
	gin.SetMode(app.Server.RunMode)
	app.Router.Use(gin.Logger(),gin.Recovery(),o.Gateway)
	app.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//加载公共服务路由
	api.CommonInterface()
	//加载用户服务路由
	api.UserInterface()
	//加载宠物服务路由
	api.PetInterface()

	return app.Router
}

