package request

type LoginRequestParams struct {
	Username   string `form:"username" json:"username" binding:"required" msg:"用户名不能为空"` // 用户名
	Password   string `form:"password" json:"password" binding:"required" msg:"密码不能为空"`  // 密码
	Captcha    string `form:"captcha" json:"captcha" binding:"required" msg:"验证码不能为空"`   // 验证码
	CaptchaKey string `form:"captchaKey" json:"captchaKey"`                              // 验证码key
}

type ChangePasswordReq struct {
	ID          uint   `json:"-" form:"-"`                       // ID
	Password    string `json:"password" form:"password" bind:""` // 密码
	NewPassword string `json:"newPassword" form:"newPassword"`   // 新密码
}

type CreateUserRequestParams struct {
	Username string `form:"username" json:"username"  binding:"required" msg:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required" example:"密码" msg:"密码不能为空"`
	Nickname string `form:"nickname" json:"nickname" example:"昵称"`
	Mobile   string `form:"mobile" json:"mobile"  example:"电话号码"`
	Email    string `form:"email" json:"email"  example:"电子邮箱"`
	Avatar   string `form:"avatar" json:"avatar" example:"头像"`
	Gender   string `form:"gender" json:"gender" example:"性别"`
}
