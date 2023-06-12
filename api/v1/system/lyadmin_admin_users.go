package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

type LyadminAdminUsersApi struct {
}

type LoginResponse struct {
	User      system.LyadminAdminUsers `json:"user"`
	Token     string                   `json:"token"`
	ExpiresAt int64                    `json:"expiresAt"`
}

type LoginRequestParams struct {
	Username   string `json:"username" binding:"required"` // 用户名
	Password   string `json:"password" binding:"required"` // 密码
	Captcha    string `json:"captcha"`                     // 验证码
	CaptchaKey string `json:"captchaKey"`                  // 验证码key
}

// Login
// @Tags      Base
// @Summary   用户登录
// @accept    application/json
// @Produce   application/json
// @Param    data  body      LoginRequestParams 			true  "用户名, 密码, 验证码"
// @Success 200 {object} response.StructResponse{data=LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router    /base/login [post]
func (l *LyadminAdminUsersApi) Login(c *gin.Context) {
	var req LoginRequestParams
	err := c.ShouldBindJSON(&req)
	ip := c.ClientIP()

	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(ip, "", c)

}
