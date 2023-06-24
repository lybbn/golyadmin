package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationLog())
	// userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("user", baseApi.CreateUser)                // 创建用户
		userRouter.POST("change_password", baseApi.ChangePassword) // 用户修改密码
	}
}
