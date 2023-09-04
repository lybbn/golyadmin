package system

import (
	"time"

	"gitee.com/lybbn/golyadmin/global"
)

type LyadminJob struct {
	global.GL_BASE_MODEL
	TaskName       string    `json:"task_name" form:"task_name" gorm:"comment:任务名称(任务简称，一般为英文)"`             //任务名称
	CronExpression string    `json:"cron_expression" form:"cron_expression" gorm:"size:255;comment:cron表达式"` // cron表达式
	Args           string    `json:"args" form:"args" gorm:"size:255;comment:目标参数"`                          // 目标参数
	EntryId        int       `json:"entry_id" form:"entry_id" gorm:"size:11;comment:job启动时返回的id"`            // job启动时返回的id
	Sort           int       `json:"sort" form:"sort" gorm:"default:1;comment:显示顺序"`                         //显示顺序
	Status         bool      `json:"status" form:"status" gorm:"default:true;comment:状态"`                    //状态
	TotalRunCount  int       `json:"total_run_count" form:"total_run_count" gorm:"default:0;comment:总运行次数"`
	LastRunAt      time.Time `json:"last_run_at" form:"last_run_at" gorm:"default:null;comment:最后运行时间"` //最后运行时间
	global.GL_CONTROL_MODEL
}

func (LyadminJob) TableName() string {
	return "lyadmin_job"
}
