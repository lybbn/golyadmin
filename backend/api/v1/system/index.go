package system

import "gitee.com/lybbn/golyadmin/service"

type ApiGroup struct {
	BaseApi
	OperationLogApi
	MenuApi
	RoleApi
}

var (
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	operationLogService = service.ServiceGroupApp.SystemServiceGroup.OperationLogService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	roleService         = service.ServiceGroupApp.SystemServiceGroup.RoleService
)
