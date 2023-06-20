package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/common/request"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	systemReq "gitee.com/lybbn/go-vue-lyadmin/model/system/request"
	"gorm.io/gorm"
)

type OperationLogService struct{}

// 创建操作日志
func (o *OperationLogService) CreateLyadminOperationLog(lyadminOperationLog system.LyadminOperationLog) (err error) {
	err = global.GVLA_DB.Create(&lyadminOperationLog).Error
	return err
}

// 批量删除操作日志
func (o *OperationLogService) DeleteLyadminOperationLogByIds(ids request.Ids) (err error) {
	err = global.GVLA_DB.Delete(&[]system.LyadminOperationLog{}, "id in (?)", ids.Ids).Error
	return err
}

// 删除操作日志
func (o *OperationLogService) DeleteLyadminOperationLog(id uint) (err error) {
	err = global.GVLA_DB.Where("id = ?", id).Delete(&system.LyadminOperationLog{}).Error
	return err
}

// 清空全部操作日志
func (o *OperationLogService) DeleteAllLyadminOperationLog() (err error) {
	err = global.GVLA_DB.Where("id > ?", 0).Delete(&system.LyadminOperationLog{}).Error
	return err
}

// 根据id获取单条操作记录
func (o *OperationLogService) GetLyadminOperationLogDetail(id uint) (lyadminOperationLog system.LyadminOperationLog, err error) {
	err = global.GVLA_DB.Where("id = ?", id).First(&lyadminOperationLog).Error
	return
}

// 分页获取操作记录列表
func (o *OperationLogService) GetLyadminOperationLogList(info systemReq.LyadminOperationLogSearch) *gorm.DB {
	// 创建db
	db := global.GVLA_DB.Model(&system.LyadminOperationLog{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Code != 0 {
		db = db.Where("code = ?", info.Code)
	}
	db = db.Order("id desc").Preload("User")
	return db
}
