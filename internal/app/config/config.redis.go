package config

// Redis redis配置参数
type RedisConfig struct {
	Addr     string
	Password string
	Database int
}
