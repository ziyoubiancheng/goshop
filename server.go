package main

import (
	"github.com/ziyoubiancheng/drivers"
	"github.com/ziyoubiancheng/drivers/pkg/cache/redis"
	"github.com/ziyoubiancheng/drivers/pkg/database/mysql"
	"github.com/ziyoubiancheng/drivers/pkg/server/gin"
	"github.com/ziyoubiancheng/drivers/pkg/server/stat"
	"go.uber.org/zap"

	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/pkg/bootstrap"
	"github.com/ziyoubiancheng/goshop/router"
	"github.com/ziyoubiancheng/goshop/service"
)

func startFn() {
	if err := drivers.Container(
		bootstrap.Arg.CfgFile,
		mysql.Register,
		redis.Register,
		gin.Register,
		stat.Register,
	); err != nil {
		panic(err)
	}
	// 配置初始化
	if err := bootstrap.InitConfig(bootstrap.Arg.CfgFile); err != nil {
		model.Logger.Panic(err.Error())
	}

	model.Init()
	service.Init()
	service.InitGen()
	// 主服务器
	endless.DefaultReadTimeOut = gin.Config().Drivers.Server.Gin.ReadTimeout.Duration
	endless.DefaultWriteTimeOut = gin.Config().Drivers.Server.Gin.WriteTimeout.Duration
	endless.DefaultMaxHeaderBytes = 100000000000000
	server := endless.NewServer(gin.Config().Drivers.Server.Gin.Addr, router.InitApi())
	server.BeforeBegin = func(add string) {
		model.Logger.Info("started")
	}

	if err := server.ListenAndServe(); err != nil {
		model.Logger.Error("ServerApi err", zap.String("err", err.Error()))
	}
}
