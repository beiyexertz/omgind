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
	if C.System.PrintConfig {
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
	System       SystemConfig `mapstructure:"System"  json:"system"`
	HTTP         HTTPConfig   `mapstructure:"HTTP" json:"http"`
	Menu         MenuConfig   `mapstructure:"menu" json:"menu"`
	Casbin       CasbinConfig `mapstructure:"Casbin" json:"casbin"`
	Log          LogConfig    `mapstructure:"Log" json:"log"`
	LogGormHook  LogGormHook
	LogMongoHook LogMongoHook
	Root         RootConfig        `mapstructure:"Root" json:"root"`
	JWTAuth      JWTAuthConfig     `mapstructure:"JWTAuth" json:"jwtAuth"`
	Monitor      MonitorConfig     `mapstructure:"Monitor" json:"monitor"`
	Captcha      CaptchaConfig     `mapstructure:"Captcha" json:"captcha"`
	RateLimiter  RateLimiterConfig `mapstructure:"RateLimiter" json:"rateLimiter"`
	CORS         CORSConfig        `mapstructure:"CORS" json:"CORS"`
	GZIP         GZIPConfig        `mapstructure:"GZIP" json:"GZIP"`
	Redis        RedisConfig       `mapstructure:"Redis" json:"redis"`
	Gorm         GormConfig        `mapstructure:"Gorm" json:"gorm"`
	MySQL        MySQLConfig       `mapstructure:"MySQL" json:"mysql"`
	Postgres     PostgresConfig    `mapstructure:"Postgres" json:"postgres"`
	Sqlite3      Sqlite3Config     `mapstructure:"Sqlite3" json:"sqlite3"`
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.System.RunMode == "debug"
}
