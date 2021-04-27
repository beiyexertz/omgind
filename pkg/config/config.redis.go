package config

// Redis redis配置参数
type RedisConfig struct {
	Enable   bool
	Addr     string
	Password string
	Database int
}
