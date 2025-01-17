package main

import (
	"03-logrus/config"
	"03-logrus/hooks"
	"fmt"

	"time"

	"github.com/sirupsen/logrus"
)

// User 示例结构体
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// 初始化日志
	config.InitLogger()

	// 添加自定义Hook
	config.Log.AddHook(&hooks.CustomHook{})

	// 1. 基本日志示例
	config.Log.Info("server starting...")

	// 2. 带字段的日志
	config.Log.WithFields(logrus.Fields{
		"version": "1.0.0",
		"port":    8080,
	}).Info("server configuration")

	// 3. 结构化日志
	user := User{
		Name:    "张三",
		Age:     25,
		Address: "北京市",
	}

	config.Log.WithFields(logrus.Fields{
		"user": user,
	}).Info("user info")

	// 4. 错误日志
	if err := someFunction(); err != nil {
		config.Log.WithError(err).Error("operation failed")
	}

	// 5. 使用不同日志级别
	config.Log.Debug("这是一条调试日志")
	config.Log.Info("这是一条信息日志")
	config.Log.Warn("这是一条警告日志")
	config.Log.Error("这是一条错误日志")

	// 6. 创建子logger
	contextLogger := config.Log.WithFields(logrus.Fields{
		"module":    "user_service",
		"requestID": "123456",
	})

	contextLogger.Info("处理用户请求")

	// 7. 性能记录示例
	startTime := time.Now()
	// ... 执行一些操作
	config.Log.WithFields(logrus.Fields{
		"duration": time.Since(startTime),
	}).Info("operation completed")
}

func someFunction() error {
	return fmt.Errorf("something went wrong")
}
