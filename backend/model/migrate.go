package model

import "gitee.com/lybbn/golyadmin/model/system"

//需要migrate同步的model表
var MigrateModelList = []interface{}{
	system.LyadminDept{},
	system.LyadminMenu{},
	system.LyadminMenuButton{},
	system.LyadminRole{},
	system.LyadminPost{},
	system.LyadminUsers{},
	system.LyadminOperationLog{},
	system.LyadminJwtBlacklist{},
}
