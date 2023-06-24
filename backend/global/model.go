package global

import (
	"time"
)

type GL_BASE_MODEL struct {
	ID        uint      `json:"id" form:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:主键"` //主键
	CreatedAt time.Time `json:"create_at" form:"create_at" gorm:"comment:创建时间"`                                //创建时间
	UpdatedAt time.Time `json:"update_at" form:"update_at" gorm:"comment:更新时间"`                                //更新时间
}

type GL_CONTROL_MODEL struct {
	CreateBy   uint `json:"create_by" form:"create_by" gorm:"index;comment:创建者"`        //创建者
	UpdateBy   uint `json:"update_by" form:"update_by" gorm:"index;comment:更新者"`        //更新着
	BelongDept uint `json:"belong_dept" form:"belong_dept" gorm:"index;comment:数据归属部门"` //数据归属部门
}
