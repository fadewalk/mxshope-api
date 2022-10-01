package main

import (
	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction()  // 生产环境，打印 json
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	url := "https://imooc.com"
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	"url", url,
	//	"attempt", 3)
	//sugar.Infof("Failed to fetch URL:%s", url)

	//logger.Info("failed to fetch URL") // 传递的是filed

	logger.Info("failed to fetch URL", // 需要这样去传参
		zap.String("url", url),
		zap.Int64("nums", 3))
}
