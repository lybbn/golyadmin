package initdb

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
)

// 数据来源: backend/lyadmin_go.sql（REPLACE INTO `lyadmin_dept`，共 2 行）
var seedDepts = []system.LyadminDept{
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 1, CreatedAt: mustTime("2023-06-26 11:06:52.000"), UpdatedAt: mustTime("2023-07-01 13:00:15.138")},
		ParentId: 0, Name: "golyadmin团队", Sort: 1, Status: true, Owner: "", Phone: "", Email: "",
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 2, CreatedAt: mustTime("2023-06-27 10:42:20.776"), UpdatedAt: mustTime("2023-07-01 12:59:38.480")},
		ParentId: 1, Name: "财务部门", Sort: 1, Status: true, Owner: "", Phone: "", Email: "",
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
}

// 数据来源: backend/lyadmin_go.sql（lyadmin_post 表在备份中无数据）
var seedPosts = []system.LyadminPost{}
