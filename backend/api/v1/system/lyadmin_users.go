package system

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/common/request"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/pagination"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type BaseApi struct {
}

type UserApi struct {
}

type LoginResponse struct {
	User      system.LyadminUsers `json:"user"`
	Access    string              `json:"access"`
	ExpiresAt int64               `json:"expiresAt"`
}

// Login
// @Tags      Base
// @Summary   用户后台登录
// @accept    application/json
// @Produce   application/json
// @Param    data  body      systemReq.LoginRequestParams 			true  "用户名, 密码, 验证码"
// @Success 2000 {object} response.StructResponse{data=LoginResponse,msg=string} "用户后台登录，返回包括用户信息,token,过期时间"
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
		err = global.GL_DB.Where("username = ?", req.Username).Preload("Role").Preload("Role.Dept").First(&user).Error
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
			if !user.IsStaff {
				global.GL_LOG.Error("非后台用户，禁止登录:" + req.Username)
				response.ErrorResponse("非后台用户，禁止登录!", c)
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
	roleIds := []int{}
	roleDataScopes := []int{}
	roleDeptIds := []int{}
	for i, v := range user.Role {
		roleIds = append(roleIds, int(v.ID))
		roleDataScopes = append(roleDataScopes, v.DataRange)
		for _, vm := range user.Role[i].Dept {
			roleDeptIds = append(roleDeptIds, int(vm.ID))
		}
	}
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:           user.UUID,
		ID:             uint(user.ID),
		Nickname:       user.Nickname,
		Username:       user.Username,
		Identity:       user.Identity,
		DeptId:         user.DeptId,
		RoleIds:        roleIds,
		RoleDeptIds:    utils.RemoveDuplicatesArrInt(roleDeptIds),
		RoleDataScopes: utils.RemoveDuplicatesArrInt(roleDataScopes),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GL_LOG.Error("获取token失败!", zap.Error(err))
		response.ErrorResponse("获取token失败", c)
		return
	}

	if global.GL_CONFIG.System.UseMultipoint {
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

// @Tags      AdminUser
// @Summary   分页获取管理员用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.LyadminUserSearch                                        true  "页码, 每页大小等"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取管理员用户列表"
// @Router    /system/user/getAdminUserList [get]
func (u *UserApi) GetAdminUserList(c *gin.Context) {
	var pageInfo systemReq.LyadminUserSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := userService.GetAdminUserInfoList(pageInfo).Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c))
	p := pagination.Page[system.LyadminUsers]{}
	err = p.PaginateQuery(query, c)
	if err != nil {
		global.GL_LOG.Error("获取管理员用户列表失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	data := []interface{}{}
	for _, v := range p.Data {
		m := utils.ConvertStruct2MapJsonReflect(v)
		var roleIds []uint
		for _, v1 := range v.Role {
			roleIds = append(roleIds, v1.ID)
		}
		m["roleIds"] = roleIds
		// m["created_at"] = utils.DateTimeFormat(v.CreatedAt) //时间日期格式化
		// m["updated_at"] = utils.DateTimeFormat(v.UpdatedAt)
		data = append(data, m)
	}
	response.PaginateResponse(data, p, "获取成功", c)
}

// @Tags     AdminUser
// @Summary  创建管理员用户
// @Produce   application/json
// @Param    data  body      systemReq.CreateUserRequestParams                                            true  "用户名, 密码,部门,姓名"
// @Success  2000   {object}  response.StructResponse{data=system.LyadminUsers,msg=string}  "创建管理员用户,返回包括用户信息"
// @Router   /system/user/adminUser [post]
func (u *UserApi) CreateAdminUser(c *gin.Context) {
	var req systemReq.CreateUserRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptIdDB(c)
	err = userService.CreateAdminUser(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		msg := err.Error()
		response.ErrorResponse(msg, c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      AdminUser
// @Summary   编辑管理员
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       systemReq.UpdateUsersRequestParams true "model UpdateUsersRequestParams"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑管理员"
// @Router    /system/user/adminUser/:id [put]
func (u *UserApi) UpdateAdminUser(c *gin.Context) {
	var req systemReq.UpdateUsersRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = userService.UpdateAdminUser(req, c)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}

// @Tags      AdminUser
// @Summary   根据ID删除管理员
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "ID"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除管理员"
// @Router    /system/user/adminUser [delete]
func (u *UserApi) DeleteAdminUser(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = userService.DeleteAdminUser(uint(req.Id), c)
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      AdminUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /system/user/getUserInfo [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	uinfo := utils.GetUserInfoDB(c)
	var userinfo = systemReq.ChangeUserInfo{
		Name:     uinfo.Name,
		Nickname: uinfo.Nickname,
		Mobile:   uinfo.Mobile,
		Email:    uinfo.Email,
		Gender:   uinfo.Gender,
	}
	response.SuccessResponse(userinfo, "获取成功", c)
}

// @Tags      AdminUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /system/user/setUserInfo [post]
func (u *UserApi) SetUserInfo(c *gin.Context) {
	var req systemReq.ChangeUserInfo
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = userService.SetUserInfo(req, utils.GetUserID(c))
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse("修改失败!", c)
		return
	}
	response.SuccessResponse(nil, "设置成功", c)
}

// @Tags      AdminUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.StructResponse{msg=string}  "用户修改密码"
// @Router    /system/user/changePassword [post]
func (u *UserApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	uid := utils.GetUserID(c)
	user := &system.LyadminUsers{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: uid}, Password: req.OldPassword}
	_, err = userService.ChangePassword(user, req.NewPassword, c)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse("修改失败，原密码错误！", c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}

// @Tags      User
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.LyadminUserSearch                                        true  "页码, 每页大小等"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取用户列表"
// @Router    /user/user/getUserList [get]
func (u *UserApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.LyadminUserSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := userService.GetUserInfoList(pageInfo).Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c))
	p := pagination.Page[system.LyadminUsers]{}
	err = p.PaginateQuery(query, c)
	if err != nil {
		global.GL_LOG.Error("获取用户列表失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.PaginateResponse(p.Data, p, "获取成功", c)
}

// @Tags      User
// @Summary   根据ID删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "ID"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除用户"
// @Router    /user/user/users [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = userService.DeleteUser(uint(req.Id), c)
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      User
// @Summary   修改用户状态
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.DisableUserReq    true  "ID,状态"
// @Success   200   {object}  response.StructResponse{msg=string}  "修改用户状态"
// @Router    /user/user/disableuser [post]
func (u *UserApi) DisableUser(c *gin.Context) {
	var req systemReq.DisableUserReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Println(err.Error())
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = userService.DisableUser(req, c)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse("修改失败", c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}

// @Tags     User
// @Summary  创建用户
// @Produce   application/json
// @Param    data  body      systemReq.CreateUserRequestParams                                            true  "用户名, 密码,部门,姓名"
// @Success  2000   {object}  response.StructResponse{data=system.LyadminUsers,msg=string}  "创建用户,返回包括用户信息"
// @Router   /system/user/adminUser [post]
func (u *UserApi) CreateUser(c *gin.Context) {
	var req systemReq.CreateUserRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptIdDB(c)
	err = userService.CreateUser(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		msg := err.Error()
		response.ErrorResponse(msg, c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      User
// @Summary   编辑用户
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       systemReq.UpdateUsersRequestParams true "model UpdateUsersRequestParams"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑用户"
// @Router    /system/user/adminUser/:id [put]
func (u *UserApi) UpdateUser(c *gin.Context) {
	var req systemReq.UpdateUsersRequestParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = userService.UpdateUser(req, c)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}
