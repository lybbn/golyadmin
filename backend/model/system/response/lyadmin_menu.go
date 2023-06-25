package response

import (
	"gitee.com/lybbn/golyadmin/model/system"
)

type LyadminWebRouterResponse struct {
	system.LyadminMenu
	MenuPermission []string `json:"menuPermission"`
}
