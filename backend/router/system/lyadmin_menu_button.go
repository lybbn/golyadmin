package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type MenuButtonRouter struct{}

func (m *MenuRouter) InitButtonRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("button").Use(middleware.OperationLog())
	menuButtonApi := v1.ApiGroupApp.SystemApiGroup.MenuButtonApi
	{
		userRouter.GET("button", menuButtonApi.GetMenuButton)        // 获取菜单全部列表
		userRouter.POST("button", menuButtonApi.CreateMenuButton)    // 新增菜单
		userRouter.PUT("button/:id", menuButtonApi.UpdateMenuButton) // 编辑菜单
		userRouter.DELETE("button", menuButtonApi.DeleteMenuButton)  // 删除菜单
	}
}
