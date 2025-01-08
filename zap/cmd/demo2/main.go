package main

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化logger
func initLogger() (*zap.Logger, error) {
	// 定义日志级别
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	// 定义日志编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
	}

	// 配置
	config := zap.Config{
		Level:         level,         // 日志级别
		Development:   true,          // 开发模式，堆栈跟踪
		Encoding:      "json",        // 输出格式 console 或 json
		EncoderConfig: encoderConfig, // 编码器配置
		InitialFields: map[string]interface{}{ // 初始化字段，如：添加一个服务器名称
			"serviceName": "demo-service",
		},
		OutputPaths:      []string{"stdout"}, // 输出到控制台和文件
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	return config.Build()
}

// 模拟的用户结构体
type User struct {
	Name    string
	Age     int
	Address string
}

// 模拟的业务函数
func processUser(logger *zap.Logger, user User) error {
	// 使用 With 创建带有上下文的 logger
	userLogger := logger.With(
		zap.String("user", user.Name),
		zap.Int("age", user.Age),
	)

	userLogger.Info("开始处理用户信息")

	// 模拟一些处理逻辑
	if user.Age < 0 {
		userLogger.Error("用户年龄不合法",
			zap.String("address", user.Address),
			zap.Error(errors.New("age cannot be negative")),
		)
		return errors.New("invalid age")
	}

	// 记录处理时间
	startTime := time.Now()
	time.Sleep(100 * time.Millisecond) // 模拟处理耗时
	userLogger.Info("用户信息处理完成",
		zap.Duration("processTime", time.Since(startTime)),
	)

	return nil
}

func main() {
	// 初始化logger
	logger, err := initLogger()
	if err != nil {
		panic("初始化logger失败: " + err.Error())
	}
	defer logger.Sync() // 退出前刷新缓冲区

	// 记录一条简单的日志
	logger.Info("系统启动")

	// 创建sugar logger
	sugar := logger.Sugar()
	sugar.Infof("Sugar Logger 示例: %s", "hello world")

	// 模拟处理几个用户
	users := []User{
		{Name: "张三", Age: 20, Address: "北京"},
		{Name: "李四", Age: -5, Address: "上海"}, // 这个会触发错误日志
		{Name: "王五", Age: 30, Address: "广州"},
	}

	for _, user := range users {
		err := processUser(logger, user)
		if err != nil {
			logger.Warn("处理用户信息失败",
				zap.String("user", user.Name),
				zap.Error(err),
			)
		}
	}

	// 记录不同级别的日志
	logger.Debug("这是一条调试日志")
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志",
		zap.Error(errors.New("发生了一个错误")),
	)

	// 使用 DPanic 在开发环境会 panic
	// logger.DPanic("这是一条会在开发环境触发 panic 的日志")

	// Fatal 和 Panic 级别的日志慎用
	// logger.Fatal("这是一条致命错误日志")  // 会导致程序退出
	// logger.Panic("这是一条会触发 panic 的日志") // 会触发 panic
}
