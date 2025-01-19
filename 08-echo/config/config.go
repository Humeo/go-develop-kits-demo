package config

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		URL string
	}
}

func LoadConfig() (*Config, error) {
	// 加载配置的逻辑
	return &Config{}, nil
}
