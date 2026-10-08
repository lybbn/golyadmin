package initdb

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
)

// 数据来源: backend/lyadmin_go.sql（REPLACE INTO `lyadmin_role`，共 1 行）
var seedRoles = []system.LyadminRole{
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 1, CreatedAt: mustTime("2023-06-26 10:51:03.000"), UpdatedAt: mustTime("2023-09-07 20:55:09.849")},
		Name: "管理员", Key: "admin", Sort: 1, Status: true, DataRange: 3,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
}
