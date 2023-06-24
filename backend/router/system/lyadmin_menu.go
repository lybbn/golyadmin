package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("menu").Use(middleware.OperationLog())
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		userRouter.GET("menu", menuApi.GetMenu) // 获取菜单全部列表
		// userRouter.GET("menulist", menuApi.GetMenuList) // 分页获取菜单列表
		userRouter.POST("menu", menuApi.CreateMenu)    // 新增菜单
		userRouter.PUT("menu/:id", menuApi.CreateMenu) // 编辑菜单
		userRouter.DELETE("menu", menuApi.DeleteMenu)  // 删除菜单
	}
}
