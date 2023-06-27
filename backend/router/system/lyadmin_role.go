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
		userRouter.GET("role", roleApi.GetRoleList)    // 获取角色分页列表
		userRouter.POST("role", roleApi.CreateRole)    // 新增角色
		userRouter.PUT("role/:id", roleApi.UpdateRole) // 编辑角色
		userRouter.DELETE("role", roleApi.DeleteRole)  // 删除角色
	}
}
