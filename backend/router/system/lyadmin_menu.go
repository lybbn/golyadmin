package system

import (
	v1 "gitee.com/lybbn/go-vue-lyadmin/api/v1"
	"gitee.com/lybbn/go-vue-lyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("menu").Use(middleware.OperationLog())
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		userRouter.GET("menu", menuApi.GetMenuList) // 分页获取菜单列表
	}
}
