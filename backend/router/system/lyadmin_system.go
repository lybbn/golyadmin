package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"github.com/gin-gonic/gin"
)

type SystemRouter struct{}

func (m *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		Router.GET("getSystemInfo", systemApi.GetSystemInfo) // 获取服务器信息
	}
}
