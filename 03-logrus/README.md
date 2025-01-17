logrus使用
1. config下logger.go

配置初始化日志配置

```go
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

```
