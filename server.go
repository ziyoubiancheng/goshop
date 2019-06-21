package main

import (
	"fmt"

	"github.com/ziyoubiancheng/drivers"
	"github.com/ziyoubiancheng/drivers/pkg/cache/redis"
	"github.com/ziyoubiancheng/drivers/pkg/database/mysql"
	"github.com/ziyoubiancheng/drivers/pkg/server/gin"
	"github.com/ziyoubiancheng/drivers/pkg/server/stat"
	//"go.uber.org/zap"

	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/pkg/bootstrap"
	"github.com/ziyoubiancheng/goshop/router"
	"github.com/ziyoubiancheng/goshop/service"
)

func startFn() {
	// 配置初始化
	bootstrap.Arg.CfgFile = "conf/conf.toml"
	if err := drivers.Container(
		bootstrap.Arg.CfgFile,
		mysql.Register,
		redis.Register,
		gin.Register,
		stat.Register,
	); err != nil {
		panic(err)
	}
	if err := bootstrap.InitConfig(bootstrap.Arg.CfgFile); err != nil {
		model.Logger.Panic(err.Error())
	}

	//init
	model.Init()
	service.Init()
	service.InitGen()

	//服务器
	fmt.Println("start")
	router.InitApi().Run(gin.Config().Drivers.Server.Gin.Addr)
}
