package system

import "gitee.com/lybbn/go-vue-lyadmin/service"

type ApiGroup struct {
	BaseApi
	OperationLogApi
}

var (
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	operationLogService = service.ServiceGroupApp.SystemServiceGroup.OperationLogService
)
