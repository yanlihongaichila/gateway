package main

import (
	"fmt"
	"gateway/api"
	"gateway/consts"
	"gateway/middleware"
	"github.com/JobNing/frameworkJ/app"
	"github.com/JobNing/frameworkJ/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitViper("config", "./config"); err != nil {
		panic(err)
	}

	err := app.Init(
		consts.ServiceName,
		viper.GetString("nacos.ip"),
		viper.GetString("nacos.port"),
	)
	if err != nil {
		panic(err)
		return
	}

	r := gin.Default()
	//跨域中间件
	r.Use(middleware.Cors())

	api.Register(r)
	r.Run(fmt.Sprintf(":%v", viper.GetString("app.port")))
}
