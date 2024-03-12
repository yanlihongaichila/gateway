package main

import (
	"getway/api"
	"getway/consts"
	"github.com/gin-gonic/gin"
	"github.com/yanlihongaichila/framework/app"
)

func main() {
	//通过nacos获取配置信息并且连接相关工具
	err := app.Init(consts.SERVICENAME)
	if err != nil {
		panic(err)
		return
	}
	r := gin.Default()
	api.Register(r)
	r.Run(":7787")
}
