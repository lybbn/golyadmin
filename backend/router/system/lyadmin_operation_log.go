package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type OperationLogRouter struct{}

func (s *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	operationLogRouter := Router.Group("operationlog")
	operationLogRouterRecode := Router.Group("operationlog").Use(middleware.OperationLog())
	operationLogApi := v1.ApiGroupApp.SystemApiGroup.OperationLogApi
	{
		operationLogRouter.DELETE("log", operationLogApi.DeleteLyadminOperationLog)
		// operationLogRouter.DELETE("deletelogbyids", operationLogApi.DeleteLyadminOperationLogByIds)
		operationLogRouter.GET("log", operationLogApi.GetLyadminOperationLogDetail)
		operationLogRouter.GET("loglist", operationLogApi.GetLyadminOperationLogList)

	}
	{
		operationLogRouterRecode.DELETE("deletealllogs", operationLogApi.DeleteAllLyadminOperationLog)
	}
}
