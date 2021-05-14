package config

import (
	"fmt"
)

// Postgres postgres配置参数
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a PostgresConfig) DSN(orm string) string {
	var dsn string

	switch orm {
	case "gorm":
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
	case "ent":
		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", a.User, a.Password, a.Host, a.Port, a.DBName, a.SSLMode)
	}
	return dsn
}
