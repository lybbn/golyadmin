package global

import (
	"time"
)

type GVLA_MODEL struct {
	ID        int       `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time `json:"create_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"update_at" gorm:"comment:更新时间"`
}
