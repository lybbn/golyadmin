package system

import (
	v1 "gitee.com/lybbn/golyadmin/api/v1"
	"gitee.com/lybbn/golyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationLog())
	fileApi := v1.ApiGroupApp.SystemApiGroup.FileApi
	{
		fileRouter.POST("uploadFile", fileApi.UploadFileLocal) // 文件上传（本地存储）
	}
}
