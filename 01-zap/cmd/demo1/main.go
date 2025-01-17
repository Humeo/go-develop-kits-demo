package main

import (
	"01-zap/cmd/demo1/config"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// User 示例结构体
type User struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// 初始化日志
	config.InitLogger()
	defer config.Sync()

	// 1. 使用Logger（性能更好）
	config.Logger.Info("server starting...",
		zap.String("version", "1.0.0"),
		zap.Int("port", 8080),
	)

	// 2. 使用SugarLogger（更方便）
	config.SugarLogger.Infof("Server is running on port: %d", 8080)

	// 3. 不同级别的日志
	config.SugarLogger.Debug("这是一条调试日志")
	config.SugarLogger.Info("这是一条信息日志")
	config.SugarLogger.Warn("这是一条警告日志")
	config.SugarLogger.Error("这是一条错误日志")

	// 4. 结构化日志示例
	user := User{
		Name:    "张三",
		Age:     25,
		Address: "北京市",
	}

	// 使用Logger记录结构化数据
	config.Logger.Info("user info",
		zap.String("name", user.Name),
		zap.Int("age", user.Age),
		zap.String("address", user.Address),
	)

	// 5. 记录错误和堆栈信息
	if err := someFunction(); err != nil {
		config.Logger.Error("operation failed",
			zap.Error(err),
			zap.String("additional_info", "some context"),
		)
	}

	// 使用方式2
	if err := someFunction2(); err != nil {
		config.Logger.Error("operation failed",
			zap.Error(err),
			zap.String("error_code", "INTERNAL_ERROR"),
		)
	}

	// 使用方式3
	if err := someFunction3(); err != nil {
		customErr, ok := err.(*CustomError)
		if ok {
			config.Logger.Error("operation failed",
				zap.Error(err),
				zap.String("error_code", customErr.Code),
				zap.String("error_message", customErr.Message),
			)
		}
	}

	// 6. 使用with字段创建子logger
	userLogger := config.Logger.With(
		zap.String("module", "user"),
		zap.String("requestID", "123456"),
	)
	userLogger.Info("processing user request")

	// 7. 性能测试示例
	for i := 0; i < 100; i++ {
		config.Logger.Info("performance test",
			zap.Int("iteration", i),
			zap.Time("timestamp", time.Now()),
		)
	}
}

// 方式1：使用标准错误
func someFunction() error {
	return errors.New("INTERNAL_ERROR")
}

// 方式2：使用fmt.Errorf
func someFunction2() error {
	return fmt.Errorf("INTERNAL_ERROR: %s", "something went wrong")
}

// 方式3：自定义错误类型
type CustomError struct {
	Code    string
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func someFunction3() error {
	return &CustomError{
		Code:    "INTERNAL_ERROR",
		Message: "something went wrong",
	}
}
