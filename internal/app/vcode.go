package app

import (
	"github.com/go-redis/redis"
	"github.com/wanhello/omgind/pkg/global"
	"github.com/wanhello/omgind/pkg/vcode"
)

func InitVcode(cli redis.Cmdable) *vcode.Vcode {

	cfg := global.CFG.Captcha
	vc := vcode.New(cli, cfg)
	return vc
}
