package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (m *MenuRouter) InitRoleRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("role").Use(middleware.OperationLog())
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		userRouter.GET("role", roleApi.GetRole)        // 获取角色全部列表
		userRouter.POST("role", roleApi.CreateRole)    // 新增角色
		userRouter.PUT("role/:id", roleApi.CreateRole) // 编辑角色
		userRouter.DELETE("role", roleApi.DeleteRole)  // 删除角色
	}
}
