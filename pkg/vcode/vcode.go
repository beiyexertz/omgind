package vcode

import (
	"fmt"
	"io"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/wanhello/omgind/pkg/config"
	"github.com/wanhello/omgind/pkg/helper/str"
	"github.com/wanhello/omgind/pkg/helper/uid/uuid"
	"github.com/wanhello/omgind/pkg/vcode/store"
)

type Vcode struct {
	cli    redis.Cmdable
	cpch   *base64Captcha.Captcha
	store  base64Captcha.Store
	driver base64Captcha.Driver
	source string
}

func New(cli redis.Cmdable, cfg config.CaptchaConfig) *Vcode {
	fmt.Printf(" captach config %+v", cfg)
	driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, cfg.Fonts)

	if cfg.Store == "redis" {
		storer := store.NewRedisStore(cli, time.Minute*time.Duration(cfg.Duration), cfg.RedisPrefix)

		cpch := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			cli:    cli,
			cpch:   cpch,
			driver: driver,
			store:  storer,
			source: cfg.Source,
		}

	} else {
		storer := base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, time.Minute*time.Duration(cfg.Duration))
		cpch := base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
		return &Vcode{
			cli:    nil,
			cpch:   cpch,
			driver: driver,
			store:  storer,
			source: cfg.Source,
		}
	}
}

func (vc *Vcode) NewLen(length int) (id string) {
	id = uuid.MustString()
	val := str.RandCustom(length, vc.source)
	vc.store.Set(id, val)
	return
}

func (vc *Vcode) GenerateImage(id string, w io.Writer) error {

	val := vc.store.Get(id, false)
	item, err := vc.driver.DrawCaptcha(val)
	if err != nil {
		return err
	}

	_, err = item.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

func (vc *Vcode) GenerateBase64(id string) (string, error) {

	val := vc.store.Get(id, false)
	item, err := vc.driver.DrawCaptcha(val)
	if err != nil {
		return "", err
	}
	bs64 := item.EncodeB64string()
	return bs64, nil
}

func (vc *Vcode) Verify(id, answer string, clear bool) bool {
	return vc.cpch.Verify(id, answer, clear)
}
