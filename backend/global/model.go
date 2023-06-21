package global

import (
	"time"
)

type GL_BASE_MODEL struct {
	ID        int64     ` json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:主键"` //主键
	CreatedAt time.Time `json:"create_at" gorm:"comment:创建时间"`                                        //创建时间
	UpdatedAt time.Time `json:"update_at" gorm:"comment:更新时间"`                                        //更新时间
}

type GL_CONTROL_MODEL struct {
	CreateBy int `json:"create_by" gorm:"index;comment:创建者"` //创建者
	UpdateBy int `json:"update_by" gorm:"index;comment:更新者"` //更新着
}
