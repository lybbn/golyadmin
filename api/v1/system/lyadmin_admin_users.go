package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"github.com/gin-gonic/gin"
)

type LyadminAdminUsersApi struct {
}

type LoginResponse struct {
	User      system.LyadminAdminUsers `json:"user"`
	Token     string                   `json:"token"`
	ExpiresAt int64                    `json:"expiresAt"`
}

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Login
// @Tags      Base
// @Summary   用户登录
// @accept    application/json
// @Produce   application/json
// @Success 2000 {object} response.SuccessResponse{LoginResponse,string}
// @Router    /login [post]
func (e *LyadminAdminUsersApi) Login(c *gin.Context) {

}
