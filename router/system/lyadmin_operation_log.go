package system

import (
	v1 "gitee.com/lybbn/go-vue-lyadmin/api/v1"
	"github.com/gin-gonic/gin"
)

type OperationLogRouter struct{}

func (s *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	operationLogRouter := Router.Group("operationlog")
	operationLogApi := v1.ApiGroupApp.SystemApiGroup.OperationLogApi
	{
		operationLogRouter.DELETE("log", operationLogApi.DeleteLyadminOperationLog)
		operationLogRouter.DELETE("deletelogbyids", operationLogApi.DeleteLyadminOperationLogByIds)
		operationLogRouter.GET("log", operationLogApi.GetLyadminOperationLogDetail)
		operationLogRouter.GET("loglist", operationLogApi.GetLyadminOperationLogList)

	}
}
