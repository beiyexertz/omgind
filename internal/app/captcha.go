package app

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/wanhello/omgind/pkg/global"
	"github.com/wanhello/omgind/pkg/vcode"
)

func InitCaptcha(cli redis.Cmdable) *base64Captcha.Captcha {

	cfg := global.CFG.Captcha

	if cfg.Store == "redis" {
		capt := vcode.New(cli, cfg)
		driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, cfg.Fonts)

		return base64Captcha.NewCaptcha(driver.ConvertFonts(), capt.Storer)
	} else {
		driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, cfg.Fonts)

		storer := base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, time.Duration(cfg.Duration)*time.Minute)

		return base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
	}
}
