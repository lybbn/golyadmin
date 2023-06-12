package system

import (
	v1 "gitee.com/lybbn/go-vue-lyadmin/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (e *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	customerRouterWithoutRecord := Router.Group("base")
	systemApi := v1.ApiGroupApp.SystemApiGroup.LyadminAdminUsersApi
	{
		customerRouterWithoutRecord.POST("login", systemApi.Login)
	}
}
