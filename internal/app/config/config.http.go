package config

// JWTAuth 用户认证
type JWTAuth struct {
	Enable        bool
	SigningMethod string
	SigningKey    string
	Expired       int
	Store         string
	FilePath      string
	RedisDB       int
	RedisPrefix   string
}

// HTTP http配置参数
type HTTP struct {
	Host             string
	Port             int
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	MaxLoggerLength  int `default:"4096"`
}

// Captcha 图形验证码配置参数
type Captcha struct {
	Store       string
	Length      int
	Width       int
	Height      int
	RedisDB     int
	RedisPrefix string
}

// RateLimiter 请求频率限制配置参数
type RateLimiter struct {
	Enable  bool
	Count   int64
	RedisDB int
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// GZIP gzip压缩
type GZIP struct {
	Enable             bool
	ExcludedExtentions []string
	ExcludedPaths      []string
}
