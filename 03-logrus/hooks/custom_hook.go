package hooks

import (
	"github.com/sirupsen/logrus"
)

// CustomHook 自定义Hook
type CustomHook struct {
	// 可以添加一些配置字段
}

// Fire 实现Hook接口
func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	// 添加自定义字段
	entry.Data["app_name"] = "my_app"
	entry.Data["environment"] = "production"

	return nil
}

// Levels 定义Hook处理的日志级别
func (hook *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
