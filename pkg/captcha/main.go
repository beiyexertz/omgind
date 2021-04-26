package captcha

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

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

var store = base64Captcha.DefaultMemStore

func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4,
		"234567890abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQSTUVWXYZ",
		&color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}
