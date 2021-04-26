package store

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type redisStore struct {
	cli        *redis.Client
	prefix     string
	out        Logger
	expiration time.Duration
}

// NewRedisStore create an instance of a redis store
func NewRedisStore(opts *redis.Options, expiration time.Duration, out Logger,
	prefix ...string) base64Captcha.Store {
	if opts == nil {
		panic("options cannot be nil")
	}
	return NewRedisStoreWithCli(
		redis.NewClient(opts),
		expiration,
		out,
		prefix...,
	)
}

// NewRedisStoreWithCli create an instance of a redis store
func NewRedisStoreWithCli(cli *redis.Client, expiration time.Duration, out Logger, prefix ...string) base64Captcha.Store {
	store := &redisStore{
		cli:        cli,
		expiration: expiration,
		out:        out,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

func (rs *redisStore) Set(id string, value string) {
	rs.cli.Set(rs.prefix+id, value, time.Minute*5)
}

func (rs *redisStore) Get(id string, clear bool) string {
	str := rs.cli.Get(rs.prefix + id).Val()
	if clear && str != "" {
		rs.cli.Del("captcha:" + id)
	}
	return str
}

func (rs *redisStore) Verify(id, answer string, clear bool) bool {
	vv := rs.Get(id, clear)
	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}

// Logger Define the log output interface
type Logger interface {
	Printf(format string, args ...interface{})
}
