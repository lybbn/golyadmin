package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gorm.io/gorm"
)

type ButtonService struct{}

// 分页获取按钮列表
func (m *ButtonService) GetLyadminButtonList(info systemReq.LyadminButtonSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminButton{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Value != "" {
		db = db.Where("value LIKE ?", "%"+info.Value+"%")
	}
	db = db.Order("id desc")
	return db
}

// 新增按钮
func (m *ButtonService) CreateButton(ReqData system.LyadminButton) error {
	return global.GL_DB.Create(&ReqData).Error
}

// 删除按钮
func (m *ButtonService) DeleteButton(id uint) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminButton{}).Error
	return err
}

// 编辑按钮
func (m *ButtonService) UpdateButton(ReqData system.LyadminButton) (err error) {
	var oldData system.LyadminButton
	upDateMap := make(map[string]interface{})
	upDateMap["name"] = ReqData.Name
	upDateMap["value"] = ReqData.Value
	upDateMap["update_by"] = ReqData.UpdateBy

	db := global.GL_DB.Where("id = ?", ReqData.ID).Find(&oldData)
	err = db.Updates(upDateMap).Error
	return err
}
