package request

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LoginRequestParams struct {
	Username   string `form:"username" json:"username" binding:"required" msg:"用户名不能为空"` // 用户名
	Password   string `form:"password" json:"password" binding:"required" msg:"密码不能为空"`  // 密码
	Captcha    string `form:"captcha" json:"captcha" binding:"required" msg:"验证码不能为空"`   // 验证码
	CaptchaKey string `form:"captchaKey" json:"captchaKey"`                              // 验证码key
}

type ChangePasswordReq struct {
	ID          uint   `json:"-" form:"-"`                             // ID
	OldPassword string `json:"oldPassword" form:"oldPassword" bind:""` // 密码
	NewPassword string `json:"newPassword" form:"newPassword"`         // 新密码
}

type DisableUserReq struct {
	ID       uint `json:"id" form:"id" binding:"required" msg:"不能为空"` // ID
	IsActive bool `json:"is_active" form:"is_active"`                 //状态(1正常、0冻结)
}

type CreateUserRequestParams struct {
	Username string `form:"username" json:"username"  binding:"required" msg:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required" example:"密码" msg:"密码不能为空"`
	Name     string `json:"name" form:"name"` //姓名
	Nickname string `form:"nickname" json:"nickname" example:"昵称"`
	Mobile   string `form:"mobile" json:"mobile"  example:"电话号码"`
	Email    string `form:"email" json:"email"  example:"电子邮箱"`
	Avatar   string `form:"avatar" json:"avatar" example:"头像"`
	Gender   string `form:"gender" json:"gender" example:"性别"`
	DeptId   uint   `json:"dept_id" form:"dept_id"`     //部门ID
	RoleIds  []uint `json:"roleIds" form:"roleIds"`     //角色id数组
	IsActive bool   `json:"is_active" form:"is_active"` //状态(1正常、0冻结)
	Identity int    `json:"identity" form:"identity"`   //身份（1 超级管理员,2后台用户、3前台用户）
	global.GL_CONTROL_MODEL
}

type UpdateUsersRequestParams struct {
	ID       uint   `json:"id" form:"id"`               // ID
	UUID     string `form:"uuid" json:"uuid"`           // 允许读和创建
	Username string `json:"username" form:"username"`   //用户名
	Password string `json:"password" form:"password"`   //密码
	Name     string `json:"name" form:"name"`           //姓名
	Nickname string `json:"nickname" form:"nickname"`   //昵称
	Mobile   string `json:"mobile" form:"mobile"`       //手机号
	Email    string `json:"email"  form:"email"`        //邮箱
	Avatar   string `json:"avatar" form:"avatar"`       //头像
	Gender   string `json:"gender" form:"gender"`       //性别（男、女）
	DeptId   uint   `json:"dept_id" form:"dept_id"`     //部门ID
	RoleIds  []uint `json:"roleIds" form:"roleIds"`     //角色id数组
	IsActive bool   `json:"is_active" form:"is_active"` //状态(1正常、0冻结)
	Identity int    `json:"identity" form:"identity"`   //身份（1 超级管理员,2后台用户、3前台用户）
	global.GL_CONTROL_MODEL
}

type ChangeUserInfo struct {
	Name     string `json:"name" form:"name"`         //姓名
	Nickname string `json:"nickname" form:"nickname"` //昵称
	Mobile   string `json:"mobile" form:"mobile"`     //手机号
	Email    string `json:"email" form:"email"`       //邮箱
	Avatar   string `json:"avatar" form:"avatar"`     //头像
	Gender   string `json:"gender" form:"gender"`     //性别（男、女）
}

type LyadminUserSearch struct {
	Search   string `json:"search" form:"search"`       //搜索关键词
	BeginAt  string `json:"beginAt" form:"beginAt"`     //搜索开始时间
	EndAt    string `json:"endAt" form:"endAt"`         //搜索结束时间
	Name     string `json:"name" form:"name"`           //名称
	Mobile   string `json:"mobile" form:"mobile"`       //电话
	Username string `json:"username" form:"username"`   //账号
	IsActive string `json:"is_active" form:"is_active"` //状态
	response.StructPageQueryParams
}
