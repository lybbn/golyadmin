package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitAdminUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationLog())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("adminUser", userApi.CreateAdminUser)        // 创建管理员用户
		userRouter.DELETE("adminUser/:id", userApi.DeleteAdminUser)  // 删除管理员
		userRouter.POST("changePassword", userApi.ChangePassword)    // 用户修改密码
		userRouter.GET("getUserInfo", userApi.GetUserInfo)           // 获取用户信息
		userRouter.POST("setUserInfo", userApi.SetUserInfo)          // 设置用户信息
		userRouter.GET("getAdminUserList", userApi.GetAdminUserList) // 获取管理员用户列表（分页）
		userRouter.PUT("adminUser/:id", userApi.UpdateAdminUser)     // 编辑管理员
	}
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationLog())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("getUserList", userApi.GetUserList)  // 获取用户列表（分页）
		userRouter.DELETE("users/:id", userApi.DeleteUser)  // 删除用户
		userRouter.POST("disableuser", userApi.DisableUser) // 用户修改状态
		userRouter.PUT("users/:id", userApi.UpdateUser)     // 编辑用户
		userRouter.POST("users", userApi.CreateUser)        // 新增用户
	}
}
