package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type ButtonRouter struct{}

func (m *ButtonRouter) InitButtonRouter(Router *gin.RouterGroup) {
	buttonRouter := Router.Group("button").Use(middleware.OperationLog())
	buttonApi := v1.ApiGroupApp.SystemApiGroup.ButtonApi
	{
		buttonRouter.GET("button", buttonApi.GetButton)           // 获取按钮全部列表
		buttonRouter.POST("button", buttonApi.CreateButton)       // 新增按钮
		buttonRouter.PUT("button/:id", buttonApi.UpdateButton)    // 编辑按钮
		buttonRouter.DELETE("button/:id", buttonApi.DeleteButton) // 删除按钮
	}
}
