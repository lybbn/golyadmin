package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type MenuButtonRouter struct{}

func (m *MenuButtonRouter) InitMenuButtonRouter(Router *gin.RouterGroup) {
	menuButtonRouter := Router.Group("menu_button").Use(middleware.OperationLog())
	menuButtonApi := v1.ApiGroupApp.SystemApiGroup.MenuButtonApi
	{
		menuButtonRouter.GET("menu_button", menuButtonApi.GetMenuButton)        // 获取菜单按钮全部列表
		menuButtonRouter.POST("menu_button", menuButtonApi.CreateMenuButton)    // 新增菜单按钮
		menuButtonRouter.PUT("menu_button/:id", menuButtonApi.UpdateMenuButton) // 编辑菜单按钮
		menuButtonRouter.DELETE("menu_button", menuButtonApi.DeleteMenuButton)  // 删除菜单按钮
	}
}
