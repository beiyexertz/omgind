package vcode

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	config2 "github.com/wanhello/omgind/pkg/config"
	"github.com/wanhello/omgind/pkg/vcode/store"
)

type Captcha struct {
	Storer base64Captcha.Store
	cli    redis.Cmdable
}

func New(cli redis.Cmdable, cfg config2.CaptchaConfig) *Captcha {

	storer := store.NewRedisStore(cli, time.Minute*time.Duration(cfg.Duration), cfg.RedisPrefix)

	driver := base64Captcha.NewDriverString(cfg.Height, cfg.Width, cfg.NoiseCount, cfg.ShowLineOptions, cfg.Length, cfg.Source, cfg.BgColor, cfg.Fonts)

	base64Captcha.NewCaptcha(driver.ConvertFonts(), storer)
	return &Captcha{
		Storer: storer,
		cli:    cli,
	}
}

/*
//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverStringFunc(store base64Captcha.Store) (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4,
		"234567890abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQSTUVWXYZ",
		&color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	capt := base64Captcha.NewCaptcha(driver, store)
	return capt.Generate()
}

func DriverDigitFunc(store base64Captcha.Store) (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	capt := base64Captcha.NewCaptcha(driver, store)

	return capt.Generate()
}
*/
