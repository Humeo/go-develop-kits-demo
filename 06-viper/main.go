package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	Server struct {
		Host    string
		Port    int
		Timeout string
	}
}

func main() {
	// 1. 基本配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	// 2. 设置默认值
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "localhost")

	// 3. 读取环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MYAPP")

	// 4. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// 5. 监听配置变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	// 6. 获取配置
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	// 7. 使用配置
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
}
