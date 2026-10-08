package initdb

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
)

// 数据来源: backend/lyadmin_go.sql（REPLACE INTO `lyadmin_users`，共 3 行，password 为 bcrypt 哈希原样保留）
var seedUsers = []system.LyadminUsers{
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 1, CreatedAt: mustTime("2023-06-26 08:37:21.068"), UpdatedAt: mustTime("2023-07-01 00:25:47.008")},
		UUID: "10ac6b3495304b4c9ad01ec1a437c22d", Username: "superadmin", Password: "$2a$10$8bFSiG0THdgR2al7yfQBu.kPhS5NGKfZo/C5J2DId8KY5CmpPzjga",
		Name: "超级管理员", Nickname: "超级管理员", Mobile: "18000000000", Email: "", Avatar: "", Gender: "男",
		DeptId: 0, IsStaff: true, IsSuperuser: true, IsActive: true, Identity: 1,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 0, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 2, CreatedAt: mustTime("2023-06-26 10:44:48.690"), UpdatedAt: mustTime("2023-07-09 16:14:08.828")},
		UUID: "1401c17e70184764b327be8942a1ac80", Username: "admin", Password: "$2a$10$aWM7YczX8hq5htpe1yh2v.6TxIGbZhTsPOg1h4U4qS9.XJBiiPyU.",
		Name: "管理员", Nickname: "管理员", Mobile: "18000000000", Email: "", Avatar: "", Gender: "男",
		DeptId: 1, IsStaff: true, IsSuperuser: false, IsActive: true, Identity: 2,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 3, CreatedAt: mustTime("2023-07-09 19:56:41.488"), UpdatedAt: mustTime("2023-07-30 21:53:45.587")},
		UUID: "653505e254fe4a2eaf9179a82d80a695", Username: "test", Password: "$2a$10$HGpnhxNuLBwqt1HJ5gbWnOVof4b6tVIFRDN3/VSQw43hwnwmnIer2",
		Name: "", Nickname: "测试前端用户", Mobile: "18000000000", Email: "", Avatar: "", Gender: "男",
		DeptId: 0, IsStaff: true, IsSuperuser: false, IsActive: true, Identity: 3,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	// IsDelete 均为 false（SQL 中 is_delete=0）
}
