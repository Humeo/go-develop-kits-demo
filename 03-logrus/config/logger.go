package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger 初始化日志配置
func InitLogger() {
	Log = logrus.New()

	// 设置日志格式
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		DataKey:           "data",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	// 设置输出
	Log.SetOutput(os.Stdout)

	// 设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	// 开启调用者信息
	Log.SetReportCaller(true)

	// 设置日志轮转
	configLocalFilesystemLogger(
		"logs/app.log",
		7*24*time.Hour,
		time.Hour*24,
	)
}

// 配置日志轮转
func configLocalFilesystemLogger(filename string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(filename)
	writer, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", err)
	}
	Log.SetOutput(writer)
}
