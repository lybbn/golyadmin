package system

// LyadminRoleMenu 是 LyadminRole 和 LyadminMenu 的连接表
type LyadminRoleMenu struct {
	LyadminRoleId uint `gorm:"column:lyadmin_role_id"`
	LyadminMenuId uint `gorm:"column:lyadmin_menu_id"`
}

func (LyadminRoleMenu) TableName() string {
	return "lyadmin_role_menu"
}
