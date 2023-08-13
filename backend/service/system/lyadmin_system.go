package system

import (
	"gitee.com/lybbn/golyadmin/utils"
)

type SystemConfigService struct{}

// 获取服务器信息
func (s *SystemConfigService) GetSystemInfo() (utils.LySystemMonitorInfo, error) {
	var sv utils.LySystemMonitorInfo = utils.LyGetSystemAllInfo()
	return sv, nil
}
