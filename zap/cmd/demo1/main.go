package main

import (
	"errors"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 2. 记录日志
	logger.Info("这是一条信息日志",
		zap.String("user", "张三"),
		zap.Int("age", 20),
	)

	logger.Error("这是一条错误日志",
		zap.Error(errors.New("出错了")),
	)

}
