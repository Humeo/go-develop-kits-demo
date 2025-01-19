
这样的配置系统允许：

* 从多个来源读取配置
* 支持配置热重载
* 类型安全的配置读取
* 环境变量覆盖
* 默认值支持
* 配置文件监听

相关函数的作用：
1. 配置文件相关函数
```go
// 设置配置文件路径
viper.SetConfigFile("/path/to/config.yaml")

// 设置配置文件名（不带扩展名）
viper.SetConfigName("config")  // 将搜索 config.yaml, config.json 等

// 设置配置文件类型
viper.SetConfigType("yaml")    // 支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

// 添加配置文件搜索路径（可添加多个）
viper.AddConfigPath("/etc/appname/")  // 先搜索这里
viper.AddConfigPath("$HOME/.appname") // 然后是这里
viper.AddConfigPath(".")             // 最后是当前目录

```

2. 环境变量相关函数

```go
// 启用环境变量读取
viper.AutomaticEnv()

// 设置环境变量前缀
// 例如：设置前缀为 "MYAPP"，则环境变量 MYAPP_DATABASE_HOST 将映射到配置项 database.host
viper.SetEnvPrefix("MYAPP")

// 绑定环境变量
viper.BindEnv("database.host")                    // 绑定单个配置项
viper.BindEnv("database.host", "DB_HOST")         // 绑定到指定环境变量名

```

3. 读取配置值的函数

```go
// 读取各种类型的配置值
host := viper.GetString("database.host")          // 获取字符串
port := viper.GetInt("database.port")             // 获取整数
enabled := viper.GetBool("feature.enabled")       // 获取布尔值
timeout := viper.GetDuration("server.timeout")    // 获取时间间隔
names := viper.GetStringSlice("users.names")      // 获取字符串切片
sub := viper.Sub("database")                      // 获取子配置
all := viper.AllSettings()                        // 获取所有配置

// 获取配置并设置默认值
viper.Get("database.host")                        // 如果不存在返回nil
viper.GetString("database.host")                  // 如果不存在返回空字符串

```

4. 设置默认值

```go
// 设置默认值
viper.SetDefault("database.host", "localhost")
viper.SetDefault("database.port", 5432)
viper.SetDefault("timeout", "30s")

```

5. 写入配置

```go
// 保存配置到文件
viper.WriteConfig()                               // 写入到已加载的配置文件
viper.SafeWriteConfig()                          // 当文件不存在时写入
viper.WriteConfigAs("/path/to/my/config.yaml")   // 写入到指定文件

```



