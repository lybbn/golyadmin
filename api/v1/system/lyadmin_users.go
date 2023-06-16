package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

type LoginResponse struct {
	User      system.LyadminUsers `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}

type LoginRequestParams struct {
	Username   string `form:"username" json:"username" binding:"required" msg:"用户名不能为空"` // 用户名
	Password   string `form:"password" json:"password" binding:"required" msg:"密码不能为空"`  // 密码
	Captcha    string `form:"captcha" json:"captcha" binding:"required" msg:"验证码不能为空"`   // 验证码
	CaptchaKey string `form:"captchaKey" json:"captchaKey"`                              // 验证码key
}

// Login
// @Tags      Base
// @Summary   用户登录
// @accept    application/json
// @Produce   application/json
// @Param    data  body      LoginRequestParams 			true  "用户名, 密码, 验证码"
// @Success 200 {object} response.StructResponse{data=LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router    /base/login [post]
func (l *BaseApi) Login(c *gin.Context) {
	var req LoginRequestParams
	err := c.ShouldBind(&req)
	ip := utils.GetRealClientIP(c)

	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	if capchaStore.Verify(req.CaptchaKey, req.Captcha, true) {

	} else {
		response.ErrorResponse("验证码错误", c)
		return
	}
	response.SuccessResponse(ip, "", c)
}

// 请求参数结构（分页）
// type ExampleQueryParmas struct {
// 	response.StructPageQueryParams
// 	Name string `json:"name" form:"name"` //查询参数
// }

// GetExaExampleList
// @Tags      Example
// @Summary   分页获取信息列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ExampleQueryParmas                                        true  "页码, 每页大小"
// @Success 200 {object} response.StructPageResponse{data=ExampleService}
// @Router    /example/exampleList [get]
// func (e *ExampleApi) GetExaExampleList(c *gin.Context) {
// 	//单独获取请求参数
// 	name := c.Query("name")
// 	fmt.Println("====== single By Query String ======")
// 	fmt.Println(name)

// 	//按结构体接收请求参数
// 	var pageParams ExampleQueryParmas
// 	err := c.ShouldBindQuery(&pageParams)
// 	fmt.Println("====== Only Bind By Query String ======")
// 	fmt.Println(pageParams)
// 	fmt.Println(pageParams.Name)
// 	if err != nil {
// 		response.ErrorResponse(err.Error(), c)
// 		return
// 	}

// 	//分页方法
// 	query := global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username")
// 	p := pagination.Page[ExampleService]{}
// 	p.PaginateQuery(query, c)
// 	response.PaginateResponse(p.Data, p, "获取成功", c)
// }
