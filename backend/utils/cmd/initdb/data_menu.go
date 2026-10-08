package initdb

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
)

// 数据来源: backend/lyadmin_go.sql（REPLACE INTO `lyadmin_menu`，共 14 行，id 从 2 开始）
var seedMenus = []system.LyadminMenu{
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 2, CreatedAt: mustTime("2023-06-26 15:57:36.000"), UpdatedAt: mustTime("2023-07-09 16:36:27.176")},
		ParentId: 0, Name: "管理员管理", Icon: "avatar", WebPath: "adminManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 20, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 3, CreatedAt: mustTime("2023-06-26 15:57:40.000"), UpdatedAt: mustTime("2023-07-09 16:37:04.440")},
		ParentId: 0, Name: "用户管理", Icon: "UserFilled", WebPath: "userManageCrud", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 30, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 4, CreatedAt: mustTime("2023-06-26 15:57:40.000"), UpdatedAt: mustTime("2023-06-29 23:53:56.118")},
		ParentId: 0, Name: "系统管理", Icon: "tools", WebPath: "", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 990, IsCatalog: true, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 5, CreatedAt: mustTime("2023-06-26 15:57:40.000"), UpdatedAt: mustTime("2023-07-09 22:16:40.987")},
		ParentId: 4, Name: "菜单管理", Icon: "Menu", WebPath: "menuManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 2, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 6, CreatedAt: mustTime("2023-06-26 15:57:40.000"), UpdatedAt: mustTime("2023-07-09 22:18:34.306")},
		ParentId: 4, Name: "部门管理", Icon: "Collection", WebPath: "departmentManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 1, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 0, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 7, CreatedAt: mustTime("2023-06-27 12:09:22.049"), UpdatedAt: mustTime("2023-06-27 12:09:22.049")},
		ParentId: 4, Name: "操作日志", Icon: "InfoFilled", WebPath: "journalManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 99, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 0, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 8, CreatedAt: mustTime("2023-06-27 12:12:34.511"), UpdatedAt: mustTime("2023-07-09 22:16:17.860")},
		ParentId: 4, Name: "角色管理", Icon: "Key", WebPath: "roleManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 5, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 9, CreatedAt: mustTime("2023-06-28 21:53:41.125"), UpdatedAt: mustTime("2023-07-09 22:19:11.712")},
		ParentId: 4, Name: "权限管理", Icon: "Lock", WebPath: "authorityManage", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 7, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 10, CreatedAt: mustTime("2023-06-29 21:47:17.157"), UpdatedAt: mustTime("2023-07-09 16:38:50.064")},
		ParentId: 0, Name: "个人中心", Icon: "Place", WebPath: "personalCenter", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 866, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 2, UpdateBy: 1, BelongDept: 1}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 11, CreatedAt: mustTime("2023-07-09 16:35:41.147"), UpdatedAt: mustTime("2023-07-09 16:36:15.662")},
		ParentId: 0, Name: "DashBoard", Icon: "DataLine", WebPath: "analysis", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 1, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 12, CreatedAt: mustTime("2023-08-13 20:24:28.260"), UpdatedAt: mustTime("2023-08-13 20:24:38.798")},
		ParentId: 0, Name: "系统监控", Icon: "TrendCharts", WebPath: "", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 888, IsCatalog: true, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 13, CreatedAt: mustTime("2023-08-13 20:25:31.231"), UpdatedAt: mustTime("2023-08-13 20:25:31.231")},
		ParentId: 12, Name: "服务监控", Icon: "Stopwatch", WebPath: "server", IsLink: false, Visible: true, Component: "", ComponentName: "", Sort: 1, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 0, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 14, CreatedAt: mustTime("2023-09-07 20:45:00.416"), UpdatedAt: mustTime("2023-09-07 20:45:06.686")},
		ParentId: 4, Name: "按钮管理", Icon: "", WebPath: "buttonManage", IsLink: false, Visible: false, Component: "", ComponentName: "", Sort: 12, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
	{GL_BASE_MODEL: global.GL_BASE_MODEL{ID: 15, CreatedAt: mustTime("2023-09-07 20:46:53.827"), UpdatedAt: mustTime("2023-09-07 20:47:04.014")},
		ParentId: 4, Name: "按钮配置", Icon: "", WebPath: "buttonConfig", IsLink: false, Visible: false, Component: "", ComponentName: "", Sort: 15, IsCatalog: false, KeepAlive: false, Status: true,
		GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: 1, UpdateBy: 1, BelongDept: 0}},
}
