
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

