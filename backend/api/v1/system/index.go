package system

import "gitee.com/lybbn/golyadmin/service"

type ApiGroup struct {
	UserApi
	BaseApi
	OperationLogApi
	MenuApi
	RoleApi
	MenuButtonApi
	ButtonApi
	DeptApi
}

var (
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	operationLogService = service.ServiceGroupApp.SystemServiceGroup.OperationLogService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	roleService         = service.ServiceGroupApp.SystemServiceGroup.RoleService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
	menuButtonService   = service.ServiceGroupApp.SystemServiceGroup.MenuButtonService
	buttonService       = service.ServiceGroupApp.SystemServiceGroup.ButtonService
	deptService         = service.ServiceGroupApp.SystemServiceGroup.DeptService
)
