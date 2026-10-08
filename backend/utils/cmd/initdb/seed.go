package initdb

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// SeedAll 按依赖顺序写入全部种子数据（数据来源: backend/lyadmin_go.sql）
// 写入顺序: dept → post → role → users → menu → button → menu_button → users_role → role_menu → role_menubutton → users_post → role_dept
func SeedAll(db *gorm.DB) {
	if db == nil {
		fmt.Println("数据库连接为空，无法初始化种子数据")
		return
	}
	fmt.Println("开始写入种子数据")
	seedTable(db, "lyadmin_dept", seedDepts)
	seedTable(db, "lyadmin_post", seedPosts)
	seedTable(db, "lyadmin_role", seedRoles)
	seedTable(db, "lyadmin_users", seedUsers)
	seedTable(db, "lyadmin_menu", seedMenus)
	seedTable(db, "lyadmin_button", seedButtons)
	seedTable(db, "lyadmin_menu_button", seedMenuButtons)
	// many2many 关联表（model 中无独立结构体，使用原生 SQL，表名列名以 lyadmin_go.sql 建表语句为准）
	seedRelations(db, "lyadmin_users_role", "lyadmin_users_id", "lyadmin_role_id", seedUsersRoles)
	seedRelations(db, "lyadmin_role_menu", "lyadmin_role_id", "lyadmin_menu_id", seedRoleMenus)
	seedRelations(db, "lyadmin_role_menubutton", "lyadmin_role_id", "lyadmin_menu_button_id", seedRoleMenuButtons)
	seedRelations(db, "lyadmin_users_post", "lyadmin_users_id", "lyadmin_post_id", seedUsersPosts)
	seedRelations(db, "lyadmin_role_dept", "lyadmin_role_id", "lyadmin_dept_id", seedRoleDepts)
	fmt.Println("数据库初始化完成")
}

// mustTime 将 SQL 备份中的时间字符串（格式 "2006-01-02 15:04:05.000"）解析为本地时区 time.Time
func mustTime(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05.000", s, time.Local)
	if err != nil {
		panic("initdb 种子数据时间解析失败: " + s)
	}
	return t
}

// seedTable 幂等写入一张业务表：已有数据则跳过，否则批量写入种子切片
// 使用 SkipHooks 跳过模型钩子（LyadminMenu.AfterCreate 会自动补建默认菜单按钮、
// LyadminUsers.BeforeCreate 会覆盖 uuid），保证写入数据与 SQL 备份完全一致
func seedTable[T any](db *gorm.DB, name string, items []T) {
	var count int64
	if err := db.Model(new(T)).Count(&count).Error; err != nil {
		fmt.Printf("检查 %s 失败: %v\n", name, err)
		return
	}
	if count > 0 {
		fmt.Printf("跳过 %s（已有 %d 行）\n", name, count)
		return
	}
	if len(items) == 0 {
		fmt.Printf("跳过 %s（SQL 备份中无种子数据）\n", name)
		return
	}
	if err := db.Session(&gorm.Session{SkipHooks: true}).Create(&items).Error; err != nil {
		fmt.Printf("写入 %s 失败: %v\n", name, err)
		return
	}
	fmt.Printf("写入 %s %d 行\n", name, len(items))
}

// seedRelations 幂等写入 many2many 关联表（原生 SQL），pairs 中每个元素为 [左列值, 右列值]
func seedRelations(db *gorm.DB, table, leftCol, rightCol string, pairs [][2]uint) {
	var count int64
	if err := db.Table(table).Count(&count).Error; err != nil {
		fmt.Printf("检查 %s 失败: %v\n", table, err)
		return
	}
	if count > 0 {
		fmt.Printf("跳过 %s（已有 %d 行）\n", table, count)
		return
	}
	if len(pairs) == 0 {
		fmt.Printf("跳过 %s（SQL 备份中无种子数据）\n", table)
		return
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ", table, leftCol, rightCol)
	args := make([]interface{}, 0, len(pairs)*2)
	for i, p := range pairs {
		if i > 0 {
			sql += ","
		}
		sql += "(?, ?)"
		args = append(args, p[0], p[1])
	}
	if err := db.Exec(sql, args...).Error; err != nil {
		fmt.Printf("写入 %s 失败: %v\n", table, err)
		return
	}
	fmt.Printf("写入 %s %d 行\n", table, len(pairs))
}
