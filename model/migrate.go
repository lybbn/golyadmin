package model

import "gitee.com/lybbn/go-vue-lyadmin/model/system"

//需要migrate同步的model表
var MigrateModelList = []interface{}{
	system.LyadminAdminUsers{},
}
