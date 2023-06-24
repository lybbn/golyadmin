package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type ButtonRouter struct{}

func (m *MenuRouter) InitMenuButtonRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("menu_button").Use(middleware.OperationLog())
	menuButtonApi := v1.ApiGroupApp.SystemApiGroup.MenuButtonApi
	{
		userRouter.GET("menu_button", menuButtonApi.GetMenuButton)        // 获取菜单全部列表
		userRouter.POST("menu_button", menuButtonApi.CreateMenuButton)    // 新增菜单
		userRouter.PUT("menu_button/:id", menuButtonApi.UpdateMenuButton) // 编辑菜单
		userRouter.DELETE("menu_button", menuButtonApi.DeleteMenuButton)  // 删除菜单
	}
}
