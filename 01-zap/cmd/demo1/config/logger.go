package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// InitLogger 初始化Logger
func InitLogger() {
	// 设置日志文件路径
	logFile := filepath.Join("logs", "app.log")

	// 设置输出到文件
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile, // 日志文件路径
		MaxSize:    10,      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		MaxAge:     7,       // 文件最多保存多少天
		Compress:   true,    // 是否压缩
	})

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	// 设置开发模式
	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 创建核心配置
	core := zapcore.NewTee(
		// 同时输出到控制台和文件
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(developmentCfg),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), fileWriter),
			atomicLevel,
		),
	)

	// 创建logger
	Logger = zap.New(core,
		zap.AddCaller(), // 添加调用者信息
		// zap.AddCallerSkip(1),                  // 跳过一层调用堆栈
		zap.AddStacktrace(zapcore.ErrorLevel), // Error级别及以上添加堆栈信息
	)

	// 创建SugarLogger
	SugarLogger = Logger.Sugar()
}

// Sync 同步日志
func Sync() {
	Logger.Sync()
	SugarLogger.Sync()
}
