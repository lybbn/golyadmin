package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// GetSystemInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.SuccessResponse{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /system/GetSystemInfo [post]
func (s *SystemApi) GetSystemInfo(c *gin.Context) {
	info, err := systemConfigService.GetSystemInfo()
	if err != nil {
		global.GL_LOG.Error("获取失败", zap.Error(err))
		response.ErrorResponse("获取失败", c)
		return
	}
	response.SuccessResponse(info, "获取成功", c)
}
