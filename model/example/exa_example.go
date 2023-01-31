package example

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
)

/* 注意点 :
1. 	结构体里面的变量 (Name) 必须是首字符大写
gorm 指定类型
json 表示json接收时候的名称
binding required 表示必须传入
*/

type ExaExample struct {
	global.GVLA_MODEL
	Name   string `gorm:"type:varchar(20); not null; comment:姓名" json:"name" form:"name" binding:"required"` // 姓名
	Mobile string `gorm:"comment:手机号" json:"mobile" form:"mobile"`                                           // 手机号
}
