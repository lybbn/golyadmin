package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type DeptRouter struct{}

func (m *MenuRouter) InitDeptRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("dept").Use(middleware.OperationLog())
	deptApi := v1.ApiGroupApp.SystemApiGroup.DeptApi
	{
		deptRouter.GET("dept", deptApi.GetDept)        // 获取部门全部列表
		deptRouter.POST("dept", deptApi.CreateDept)    // 新增部门
		deptRouter.PUT("dept/:id", deptApi.UpdateDept) // 编辑部门
		deptRouter.DELETE("dept", deptApi.DeleteDept)  // 删除部门
	}
}
