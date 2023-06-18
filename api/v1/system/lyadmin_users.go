package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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
// @Success 2000 {object} response.StructResponse{data=LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router    /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var req LoginRequestParams
	err := c.ShouldBind(&req)
	// ip := utils.GetRealClientIP(c)

	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	if capchaStore.Verify(req.CaptchaKey, req.Captcha, true) {
		var user system.LyadminUsers
		err = global.GVLA_DB.Where("username = ?", req.Username).First(&user).Error
		if err == nil {
			if ok := utils.CheckPassword(req.Password, user.Password); !ok {
				response.ErrorResponse("账号密码错误", c)
				return
			}
			if !user.IsActive {
				global.GVLA_LOG.Error("该用户被禁用，请联系管理员:" + req.Username)
				response.ErrorResponse("该用户被禁用，请联系管理员!", c)
				return
			}
			b.GetJwtToken(c, user)
			return
		} else {
			global.GVLA_LOG.Error("登陆失败! 用户名不存在或密码错误!", zap.Error(err))
			response.ErrorResponse("用户名不存在或密码错误", c)
			return
		}
	} else {
		response.ErrorResponse("验证码错误", c)
		return
	}
}

// GetJwtToken 登录以后签发jwt
func (b *BaseApi) GetJwtToken(c *gin.Context, user system.LyadminUsers) {
	j := &utils.JWT{SecretKey: []byte(global.GVLA_CONFIG.JWT.SecretKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:     user.UUID,
		ID:       uint(user.ID),
		Nickname: user.Nickname,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVLA_LOG.Error("获取token失败!", zap.Error(err))
		response.ErrorResponse("获取token失败", c)
		return
	}

	if !global.GVLA_CONFIG.System.UseMultipoint {
		response.SuccessResponse(LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVLA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.ErrorResponse("设置登录状态失败", c)
			return
		}
		response.SuccessResponse(LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVLA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.ErrorResponse("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinBlacklist(blackJWT); err != nil {
			response.ErrorResponse("jwt黑名单作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.ErrorResponse("设置登录状态失败", c)
			return
		}
		response.SuccessResponse(LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

type CreateUserRequestParams struct {
	Username string `form:"username" json:"username"  binding:"required" msg:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required" example:"密码" msg:"密码不能为空"`
	Nickname string `form:"nickname" json:"nickname" example:"昵称"`
	Mobile   string `form:"mobile" json:"mobile"  example:"电话号码"`
	Email    string `form:"email" json:"email"  example:"电子邮箱"`
	Avatar   string `form:"avatar" json:"avatar" example:"头像"`
	Gender   string `form:"gender" json:"gender" example:"性别"`
	DeptId   int    `form:"dept_id" json:"dept_id" example:"int 部门id"`
	PostId   int    `form:"post_id" json:"post_id" example:"int 岗位id"`
	RoleId   int    `form:"role_id" json:"role_id" example:"int 角色id"`
}

// Create
// @Tags     User
// @Summary  创建用户
// @Produce   application/json
// @Param    data  body      CreateUserRequestParams                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  2000   {object}  response.StructResponse{data=CreateUserRequestParams,msg=string}  "创建用户,返回包括用户信息"
// @Router   /user/user [post]
func (b *BaseApi) CreateUser(c *gin.Context) {
	var req CreateUserRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	var i int64
	err = global.GVLA_DB.Model(&system.LyadminUsers{}).Where("username = ?", req.Username).Count(&i).Error
	if err != nil {
		global.GVLA_LOG.Error("创建用户失败!", zap.Error(err))
		response.ErrorResponse("创建用户失败", c)
		return
	}
	if i > 0 {
		response.ErrorResponse("用户名已存在！", c)
		return
	}
	user := &system.LyadminUsers{Username: req.Username, Nickname: req.Nickname, Password: req.Password, Avatar: req.Avatar, Mobile: req.Mobile, Gender: req.Gender, DeptId: req.DeptId, PostId: req.PostId, RoleId: req.RoleId, Email: req.Email}
	// 加密密码
	user.Password = utils.MakePassowrd(req.Password)
	err = global.GVLA_DB.Create(&user).Error

	if err != nil {
		global.GVLA_LOG.Error("创建用户失败!", zap.Error(err))
		response.ErrorResponse("创建用户失败!", c)
		return
	}
	response.SuccessResponse(user, "创建成功", c)
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
