package config

import (
	"os"
	"strings"
	"sync"

	"github.com/wanhello/omgind/pkg/helper/json"

	"github.com/koding/multiconfig"
)

var (
	// C 全局配置(需要先执行MustLoad，否则拿不到配置)
	C    = new(Config)
	once sync.Once
)

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// Config 配置参数
type Config struct {
	RunMode      string
	WWW          string
	Swagger      bool
	PrintConfig  bool
	HTTP         HTTP
	Menu         Menu
	Casbin       Casbin
	Log          Log
	LogGormHook  LogGormHook
	LogMongoHook LogMongoHook
	Root         Root
	JWTAuth      JWTAuth
	Monitor      Monitor
	Captcha      Captcha
	RateLimiter  RateLimiter
	CORS         CORS
	GZIP         GZIP
	Redis        Redis
	Gorm         Gorm
	MySQL        MySQL
	Postgres     Postgres
	Sqlite3      Sqlite3
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// Root root用户
type Root struct {
	UserName  string
	Password  string
	RealName  string
	FirstName string
	LastName  string
}

// Monitor 监控配置参数
type Monitor struct {
	Enable    bool
	Addr      string
	ConfigDir string
}
