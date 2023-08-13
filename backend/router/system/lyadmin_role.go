package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (m *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role").Use(middleware.OperationLog())
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.GET("role", roleApi.GetRole)                        // 获取全部角色
		roleRouter.GET("roleList", roleApi.GetRoleList)                // 获取角色分页列表
		roleRouter.POST("role", roleApi.CreateRole)                    // 新增角色
		roleRouter.PUT("role/:id", roleApi.UpdateRole)                 // 编辑角色
		roleRouter.DELETE("role/:id", roleApi.DeleteRole)              // 删除角色
		roleRouter.GET("role_id_to_menu/:id", roleApi.GetRoleMenuById) // 获取所有菜单按钮
		roleRouter.PUT("permission/:id", roleApi.UpdateRolePremission) // 更新角色权限
	}
}
