package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type BaseApi struct {
}

type LoginResponse struct {
	User      system.LyadminUsers `json:"user"`
	Access    string              `json:"access"`
	ExpiresAt int64               `json:"expiresAt"`
}

// Login
// @Tags      Base
// @Summary   用户登录
// @accept    application/json
// @Produce   application/json
// @Param    data  body      systemReq.LoginRequestParams 			true  "用户名, 密码, 验证码"
// @Success 2000 {object} response.StructResponse{data=LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router    /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var req systemReq.LoginRequestParams
	err := c.ShouldBind(&req)
	// ip := utils.GetRealClientIP(c)

	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	if capchaStore.Verify(req.CaptchaKey, req.Captcha, true) {
		var user system.LyadminUsers
		err = global.GL_DB.Where("username = ?", req.Username).First(&user).Error
		if err == nil {
			if ok := utils.CheckPassword(req.Password, user.Password); !ok {
				response.ErrorResponse("账号密码错误", c)
				return
			}
			if !user.IsActive {
				global.GL_LOG.Error("该用户被禁用，请联系管理员:" + req.Username)
				response.ErrorResponse("该用户被禁用，请联系管理员!", c)
				return
			}
			b.IssueJwtToken(c, user)
			return
		} else {
			global.GL_LOG.Error("登陆失败! 用户名不存在或密码错误!", zap.Error(err))
			response.ErrorResponse("用户名不存在或密码错误", c)
			return
		}
	} else {
		response.ErrorResponse("验证码错误", c)
		return
	}
}

// IssueJwtToken 登录以后签发jwt
func (b *BaseApi) IssueJwtToken(c *gin.Context, user system.LyadminUsers) {
	j := &utils.JWT{SecretKey: []byte(global.GL_CONFIG.JWT.SecretKey)} // 唯一签名
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:     user.UUID,
		ID:       uint(user.ID),
		Nickname: user.Nickname,
		Username: user.Username,
		Identity: user.Identity,
		DeptId:   user.DeptId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GL_LOG.Error("获取token失败!", zap.Error(err))
		response.ErrorResponse("获取token失败", c)
		return
	}

	if !global.GL_CONFIG.System.UseMultipoint {
		response.SuccessResponse(LoginResponse{
			User:      user,
			Access:    token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GL_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.ErrorResponse("设置登录状态失败", c)
			return
		}
		response.SuccessResponse(LoginResponse{
			User:      user,
			Access:    token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GL_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.ErrorResponse("设置登录状态失败", c)
	} else {
		var blackJWT system.LyadminJwtBlacklist
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
			Access:    token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// @Tags     User
// @Summary  创建用户
// @Produce   application/json
// @Param    data  body      systemReq.CreateUserRequestParams                                            true  "用户名, 密码"
// @Success  2000   {object}  response.StructResponse{data=systemReq.CreateUserRequestParams,msg=string}  "创建用户,返回包括用户信息"
// @Router   /system/user/user [post]
func (b *BaseApi) CreateUser(c *gin.Context) {
	var req systemReq.CreateUserRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	var i int64
	err = global.GL_DB.Model(&system.LyadminUsers{}).Where("username = ?", req.Username).Count(&i).Error
	if err != nil {
		global.GL_LOG.Error("创建用户失败!", zap.Error(err))
		response.ErrorResponse("创建用户失败", c)
		return
	}
	if i > 0 {
		response.ErrorResponse("用户名已存在！", c)
		return
	}
	user := &system.LyadminUsers{Username: req.Username, Nickname: req.Nickname, Password: req.Password, Avatar: req.Avatar, Mobile: req.Mobile, Gender: req.Gender, Email: req.Email}
	// 加密密码
	user.Password = utils.MakePassowrd(req.Password)
	err = global.GL_DB.Create(&user).Error

	if err != nil {
		global.GL_LOG.Error("创建用户失败!", zap.Error(err))
		response.ErrorResponse("创建用户失败!", c)
		return
	}
	response.SuccessResponse(user, "创建成功", c)
}

// @Tags      User
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.StructResponse{msg=string}  "用户修改密码"
// @Router    /system/user/change_password [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.LyadminUsers{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse("修改失败，原密码错误！", c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
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
// 	query := global.GL_DB.Table("lyadmin_users").Select("id", "name", "username")
// 	p := pagination.Page[ExampleService]{}
// 	p.PaginateQuery(query, c)
// 	response.PaginateResponse(p.Data, p, "获取成功", c)
// }
