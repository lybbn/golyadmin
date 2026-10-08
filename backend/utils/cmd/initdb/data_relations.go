package initdb

// 数据来源: backend/lyadmin_go.sql 中各 many2many 关联表的 REPLACE INTO 语句
// 每个元素为 [左列值, 右列值]，表名与列名以 SQL 文件建表语句为准

// lyadmin_users_role (lyadmin_users_id, lyadmin_role_id)，共 1 行
var seedUsersRoles = [][2]uint{
	{2, 1},
}

// lyadmin_role_menu (lyadmin_role_id, lyadmin_menu_id)，共 14 行
var seedRoleMenus = [][2]uint{
	{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7},
	{1, 8}, {1, 9}, {1, 10}, {1, 11}, {1, 12}, {1, 13},
	{1, 14}, {1, 15},
}

// lyadmin_role_menubutton (lyadmin_role_id, lyadmin_menu_button_id)，共 51 行
var seedRoleMenuButtons = [][2]uint{
	{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}, {1, 10},
	{1, 11}, {1, 12}, {1, 13}, {1, 14}, {1, 15}, {1, 16}, {1, 17}, {1, 18}, {1, 19}, {1, 20},
	{1, 21}, {1, 22}, {1, 23}, {1, 24}, {1, 25}, {1, 26}, {1, 27}, {1, 28}, {1, 29}, {1, 30},
	{1, 31}, {1, 35}, {1, 36}, {1, 37}, {1, 38}, {1, 39}, {1, 40}, {1, 41}, {1, 42}, {1, 43},
	{1, 44}, {1, 48}, {1, 49}, {1, 50}, {1, 51}, {1, 52}, {1, 53}, {1, 54}, {1, 55}, {1, 56},
	{1, 57}, {1, 58},
}

// lyadmin_users_post (lyadmin_users_id, lyadmin_post_id)，SQL 备份中无数据
var seedUsersPosts = [][2]uint{}

// lyadmin_role_dept (lyadmin_role_id, lyadmin_dept_id)，SQL 备份中无数据
var seedRoleDepts = [][2]uint{}
