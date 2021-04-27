package vcode

import (
	"io"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/wanhello/omgind/pkg/config"
	"github.com/wanhello/omgind/pkg/vcode/store"
)

type Vcode struct {
	cli     redis.Cmdable
	captcha *base64Captcha.Captcha
	store   base64Captcha.Store
	driver  base64Captcha.Driver
}

func New(cli redis.Cmdable, cfg config.CaptchaConfig) *Vcode {

	driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, cfg.Fonts)

	if cfg.Store == "redis" {
		storer := store.NewRedisStore(cli, time.Minute*time.Duration(cfg.Duration), cfg.RedisPrefix)

		cpc := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			cli:     cli,
			captcha: cpc,
			driver:  driver,
			store:   storer,
		}

	} else {
		storer := base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, time.Minute*time.Duration(cfg.Duration))
		cpc := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			cli:     nil,
			captcha: cpc,
			driver:  driver,
			store:   storer,
		}
	}
}

func (vc *Vcode) GenerateImage(w io.Writer) (id string, err error) {
	id, content, answer := vc.driver.GenerateIdQuestionAnswer()
	item, err := vc.driver.DrawCaptcha(content)
	if err != nil {
		return "", err
	}
	vc.captcha.Store.Set(id, answer)
	item.WriteTo(w)
	return
}

func (vc *Vcode) GenerateBase64() (id, b64s string, err error) {
	return vc.captcha.Generate()
}
