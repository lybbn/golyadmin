package system

// LyadminUsersRole 是 LyadminUsers 和 LyadminRole 的连接表
type LyadminUsersRole struct {
	LyadminUsersId uint `gorm:"column:lyadmin_users_id;primaryKey"`
	LyadminRoleId  uint `gorm:"column:lyadmin_role_id;primaryKey"`
}

func (LyadminUsersRole) TableName() string {
	return "lyadmin_users_role"
}
