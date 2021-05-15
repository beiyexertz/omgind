package app

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"time"

	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/pkg/global"
)

// InitGormDB 初始化gorm存储
func InitEntClient() (*ent.Client, func(), error) {
	cfg := global.CFG.Ent
	cli, cleanFunc, err := NewEntClient()
	if err != nil {
		return nil, cleanFunc, err
	}

	if cfg.EnableAutoMigrate {
		err := cli.Schema.Create(
			context.Background(),
			// FIXME::  sql/schema: pq: column "id" of relation "xxxx_table" is not an identity column
			//migrate.WithGlobalUniqueID(true),
		)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return cli, cleanFunc, nil
}

// NewGormDB 创建DB实例
func NewEntClient() (*ent.Client, func(), error) {
	cfg := global.CFG
	var dsn string

	switch cfg.Ent.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, nil, errors.New("unknown db")
	}

	db, err := sql.Open(cfg.Ent.DBType, dsn)
	if err != nil {
		return nil, func() {}, err
	}

	drv := entsql.OpenDB(cfg.Ent.DBType, db)

	cleanFunc := func() {
		drv.Close()
	}

	db.SetMaxIdleConns(cfg.Ent.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Ent.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Ent.MaxLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, cleanFunc, err
	}

	client := ent.NewClient(ent.Driver(drv))

	return client, cleanFunc, nil
}
