package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"go-services/global/common/database/mysql"
	"go-services/global/common/cache/redis"
	"go-services/global/common/third/qiniu"
	"go-services/global/common/server/app"
	"go-services/global/common/third/baidu/Identification/animal"
	"time"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", app.App)
	mapTo("server", app.Server)
	mapTo("database", mysql.Driver)
	mapTo("redis", redis.Driver)
	mapTo("qiniu", qiniu.Driver)
	mapTo("baidu_identification_animal",animal.Driver)
	app.App.ImageMaxSize 			= app.App.ImageMaxSize * 1024 * 1024
	app.Server.ReadTimeout 			= app.Server.ReadTimeout * time.Second
	app.Server.WriteTimeout 		= app.Server.ReadTimeout * time.Second
	redis.Driver.IdleTimeout 		= redis.Driver.IdleTimeout * time.Second
}

//设置环境
func Setup() {
	mysql.Setup()
	redis.Setup()
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}