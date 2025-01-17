使用步骤：
1. 创建配置文件
在config下创建logger.go
logger.go 下面初始化logger
实现InitLogger

```go
package config

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// InitLogger 初始化Logger
func InitLogger() {
    // 设置日志文件路径
    logFile := filepath.Join("logs", "app.log")
    
    // 设置输出到文件
    fileWriter := zapcore.AddSync(&lumberjack.Logger{
        Filename:   logFile,      // 日志文件路径
        MaxSize:    10,           // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: 30,           // 日志文件最多保存多少个备份
        MaxAge:     7,            // 文件最多保存多少天
        Compress:   true,         // 是否压缩
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
        zap.AddCaller(),                   // 添加调用者信息
        zap.AddCallerSkip(1),              // 跳过一层调用堆栈
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


```

步骤：
1. 设置文件路径
2. 设置输出到文件选项
3. 设置日志级别
4. 设置开发模式
5. 创建logger，error级别以上添加堆栈信息
6. 创建SuggarLogger

初始化日志
```go
config.InitLogger()

config.Logger.Info("s")

// 3. 不同级别的日志
config.SugarLogger.Debug("这是一条调试日志")
config.SugarLogger.Info("这是一条信息日志")
config.SugarLogger.Warn("这是一条警告日志")
config.SugarLogger.Error("这是一条错误日志")

// 使用Logger记录结构化数据
config.Logger.Info("user info",
zap.String("name", user.Name),
zap.Int("age", user.Age),
zap.String("address", user.Address),
)
// 2025-01-17 15:12:40.879	INFO	demo1/main.go:47	user info	{"name": "张三", "age": 25, "address": "北京市"}

// 记录错误和堆栈信息
if err := someFunction(); err != nil {
config.Logger.Error("operation failed",
zap.Error(err),
zap.String("additional_info", "some context"),
)
}
/*
   2025-01-17 15:12:40.879	ERROR	demo1/main.go:55	operation failed	{"error": "INTERNAL_ERROR", "additional_info": "some context"}
   main.main
   	D:/go-develop-kits-demo/01-zap/cmd/demo1/main.go:55
   runtime.main
   	D:/BaiduSyncdisk/Go/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.4.windows-amd64/src/runtime/proc.go:272
*/

// 6. 使用with字段创建子logger
userLogger := config.Logger.With(
zap.String("module", "user"),
zap.String("requestID", "123456"),
)
userLogger.Info("processing user request")
/*
   2025-01-17 15:12:40.879	INFO	demo1/main.go:86	processing user request	{"module": "user", "requestID": "123456"}
*/

// 7. 性能测试示例
for i := 0; i < 100; i++ {
config.Logger.Info("performance test",
zap.Int("iteration", i),
zap.Time("timestamp", time.Now()),
)
}
/*
2025-01-17 15:12:40.879	INFO	demo1/main.go:90	performance test	{"iteration": 0, "timestamp": "2025-01-17 15:12:40.879"}
2025-01-17 15:12:40.879	INFO	demo1/main.go:90	performance test	{"iteration": 1, "timestamp": "2025-01-17 15:12:40.879"}
*/


```



