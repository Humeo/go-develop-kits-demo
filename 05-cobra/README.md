
```go
rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "A brief description",
    Long:  `A longer description...`,
}
```
* Use: 定义如何使用这个命令
* Short: 简短描述，在 mycli help 中显示
* Long: 详细描述，在 mycli help <command> 中显示

**标志定义**

```go
// 持久性标志示例
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")

// 本地标志示例
rootCmd.Flags().BoolP("toggle", "t", false, "Help message")
```

* PersistentFlags: 这些标志可用于当前命令及其所有子命令
* Flags: 这些标志只能用于当前命令

命令树

```
app (根命令)
├── server (子命令)
│   ├── start (server的子命令)
│   └── stop  (server的子命令)
└── config (子命令)
```

使用方式
```
# 查看根命令帮助
./app --help

# 执行 server 命令
./app server

# 执行 server start 子命令
./app server start

# 执行 server stop 子命令
./app server stop

# 执行 config 命令
./app config

```

实现代码
```go
package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "This is the root command",
    Long:  `This is a longer description of the root command`,
}

// 子命令：server
var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Server management commands",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Server command called")
    },
}

// server的子命令：start
var serverStartCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the server",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Starting server...")
    },
}

// server的子命令：stop
var serverStopCmd = &cobra.Command{
    Use:   "stop",
    Short: "Stop the server",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Stopping server...")
    },
}

// 另一个子命令：config
var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Configuration commands",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Config command called")
    },
}

func init() {
    // 将 server 命令添加为根命令的子命令
    rootCmd.AddCommand(serverCmd)
    
    // 将 start 和 stop 命令添加为 server 命令的子命令
    serverCmd.AddCommand(serverStartCmd)
    serverCmd.AddCommand(serverStopCmd)
    
    // 将 config 命令添加为根命令的子命令
    rootCmd.AddCommand(configCmd)
}

func main() {
    rootCmd.Execute()
}

```

重要的函数说明

**cmd.Flags**

`cmd.Flags()`是 Cobra 中用于获取命令的标志集合（flag set）的函数，它返回一个 *pflag.FlagSet 对象，用于管理该命令的所有命令行标志

1. 添加标志
```go
// 添加各种类型的标志
cmd.Flags().String("name", "", "Name to use")
cmd.Flags().Int("count", 0, "Count value")
cmd.Flags().Bool("verbose", false, "Verbose output")
cmd.Flags().Float64("ratio", 0.0, "Ratio value")
cmd.Flags().Duration("timeout", time.Second, "Timeout duration")

// 带短标志的版本
cmd.Flags().StringP("name", "n", "", "Name to use")
cmd.Flags().IntP("count", "c", 0, "Count value")
```

2. 获取标志值

```go
func(cmd *cobra.Command, args []string) {
    name, _ := cmd.Flags().GetString("name")
    count, _ := cmd.Flags().GetInt("count")
    verbose, _ := cmd.Flags().GetBool("verbose")
}
```

3. 标志验证和必需标志
```go
// 设置必需标志
cmd.Flags().StringVarP(&name, "name", "n", "", "Name (required)")
cmd.MarkFlagRequired("name")

// 自定义标志验证
cmd.Flags().IntVarP(&age, "age", "a", 0, "Age (must be positive)")
cmd.MarkFlagRequired("age")
if err := cmd.Flags().Set("age", "0"); err != nil {
    fmt.Println("Invalid age value")
}
```

4. 遍历所有标志

```go
cmd.Flags().VisitAll(func(flag *pflag.Flag) {
    fmt.Printf("Flag: %s, Value: %s\n", flag.Name, flag.Value)
})
```

```go
package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "time"
)

func main() {
    var (
        config  string
        timeout time.Duration
        verbose bool
    )

    var rootCmd = &cobra.Command{
        Use: "app",
        PersistentPreRun: func(cmd *cobra.Command, args []string) {
            // 在命令执行前检查标志
            if verbose {
                fmt.Println("Verbose mode enabled")
            }
        },
        Run: func(cmd *cobra.Command, args []string) {
            // 获取并使用标志值
            fmt.Printf("Config: %s\n", config)
            fmt.Printf("Timeout: %v\n", timeout)

            // 检查哪些标志被显式设置了
            flags := cmd.Flags()
            flags.VisitAll(func(flag *pflag.Flag) {
                if flag.Changed {
                    fmt.Printf("Flag %s was set with value: %s\n", 
                        flag.Name, flag.Value)
                }
            })
        },
    }

    // 设置标志
    flags := rootCmd.Flags()
    
    // 添加基本标志
    flags.StringVarP(&config, "config", "c", "default.conf", "Config file path")
    flags.DurationVarP(&timeout, "timeout", "t", 30*time.Second, "Timeout duration")
    flags.BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

    // 添加必需标志
    flags.StringP("required-flag", "r", "", "This flag is required")
    rootCmd.MarkFlagRequired("required-flag")

    // 自定义标志验证
    flags.Int("port", 8080, "Port number (1-65535)")
    flags.SetAnnotation("port", "validation", []string{"must be between 1 and 65535"})

    rootCmd.Execute()
}

```

```bash
# 使用各种标志
./app --config=custom.conf --timeout=1m --verbose

# 查看帮助
./app --help

# 使用短标志
./app -c custom.conf -t 1m -v

```

命令行标志

```go
greetCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet")
```

这行代码可以拆解为以下部分：

greetCmd.Flags() - 获取命令的标志集合
* StringVarP - 添加一个字符串类型的标志，带有短标志名
* 函数签名: StringVarP(p *string, name string, shorthand string, value string, usage string)
* 参数解释：
  * &name - 指向存储标志值的变量的指针
  * "name" - 长标志名（使用时需要 --name）
  * "n" - 短标志名（使用时需要 -n）
  * "" - 默认值（这里是空字符串）
  * "Name to greet" - 帮助信息描述

设置持久标志（对所有子命令都可用）：

```go
rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
```

