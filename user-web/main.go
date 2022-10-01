package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/initialize"
)

func main() {

	port := 8021
	//logger, _ := zap.NewProduction()   // 不能打印debug 级别的日志
	//logger, _ := zap.NewDevelopment()
	//zap.ReplaceGlobals(logger)

	// 1.初始化 logger
	initialize.InitLogger()

	//2.初始化 Router
	Router := initialize.Routers()
	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径 。 自己有加锁
	*/

	//logger, _ := zap.NewProduction()
	//defer logger.Sync()
	//sugar := logger.Sugar()
	//sugar.Info("")

	zap.S().Debugf("启动服务器，端口：%d", port)
	//zap.L().Info("启动服务器") // 第二种日志打开方式

	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
