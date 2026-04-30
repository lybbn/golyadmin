package system

import (
	"image/color"
	"time"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// base64Captcha参数可视化调试：https://captcha.mojotv.cn/

// DefaultMemStore适用单服务器场景，多服务器场景可使用redis共享存储验证码

// var store = base64Captcha.DefaultMemStore //存储的验证码为 10240 个，过期时间为 10分钟,

// 可自定义存储对象 设置存储的验证码为 10240个，过期时间为 3分钟

var capchaStore = base64Captcha.NewMemoryStore(10240, 180*time.Second)

// var capchaStore = base64Captcha.NewMemoryStore(10240, time.Duration(global.GL_CONFIG.Captcha.CaptchaTimeout)*time.Second)

// 使用redis存储

// var capchaStore = captcha.NewDefaultRedisStore()

// mathConfig 生成图形化算术验证码配置
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          global.GL_CONFIG.Captcha.ImgHeight,
		Width:           global.GL_CONFIG.Captcha.ImgWidth,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 1,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	return mathType.ConvertFonts()
}

// digitConfig 生成图形化数字验证码配置
func digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   global.GL_CONFIG.Captcha.ImgHeight,
		Width:    global.GL_CONFIG.Captcha.ImgWidth,
		Length:   global.GL_CONFIG.Captcha.KeyLength,
		MaxSkew:  0.45,
		DotCount: 25,
	}
	return digitType
}

// stringConfig 生成图形化字符串验证码配置
func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          global.GL_CONFIG.Captcha.ImgHeight,
		Width:           global.GL_CONFIG.Captcha.ImgWidth,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine, //base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine
		Length:          global.GL_CONFIG.Captcha.KeyLength,
		Source:          "1234567890qwertyuiopasdfghjklzxcvb", //1234567890qwertyuiopasdfghjklzxcvbQWERTYUIOPLKJHGFDSAZXCVBNM
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 1,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	return stringType.ConvertFonts()
}

// chineseConfig 生成图形化汉字验证码配置
func chineseConfig() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          global.GL_CONFIG.Captcha.ImgHeight,
		Width:           global.GL_CONFIG.Captcha.ImgWidth,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          global.GL_CONFIG.Captcha.KeyLength,
		Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值,但是,作者,可是,老眼,天空",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return chineseType
}

// autoConfig 生成图形化数字音频验证码配置
func audioConfig() *base64Captcha.DriverAudio {
	audioType := &base64Captcha.DriverAudio{
		Length:   global.GL_CONFIG.Captcha.KeyLength,
		Language: "zh",
	}
	return audioType
}

// configJsonBody json request body.
type captchaResponse struct {
	CaptchaKey    string `json:"captchaKey"`
	Captcha       string `json:"captcha"`
	CaptchaLength int    `json:"captchaLength"`
}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.StructResponse{data=captchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [get]
func (b *BaseApi) GetCaptcha(c *gin.Context) {

	var driver base64Captcha.Driver

	switch global.GL_CONFIG.Captcha.CaptchaType {
	case "audio":
		driver = audioConfig()
	case "string":
		driver = stringConfig()
	case "math":
		driver = mathConfig()
	case "chinese":
		driver = chineseConfig()
	case "digit":
		driver = digitConfig()
	default:
		driver = stringConfig()
	}

	cp := base64Captcha.NewCaptcha(driver, capchaStore)

	id, b64s, _, err := cp.Generate()

	if err != nil {
		global.GL_LOG.Error("验证码获取失败!", zap.Error(err))
		response.ErrorResponse("验证码获取失败", c)
		return
	}

	response.SuccessResponse(captchaResponse{
		CaptchaKey:    id,
		Captcha:       b64s,
		CaptchaLength: global.GL_CONFIG.Captcha.KeyLength,
	}, "验证码获取成功", c)
}
