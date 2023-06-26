package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gorm.io/gorm"
)

type DeptService struct{}

// 分页获取部门列表
func (m *DeptService) GetLyadminDeptList(info systemReq.LyadminDeptSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminDept{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Method != "" {
		db = db.Where("method LIKE ?", "%"+info.Method+"%")
	}
	db = db.Order("id desc")
	return db
}

// 新增部门
func (m *DeptService) CreateDept(ReqData system.LyadminDept) error {
	return global.GL_DB.Create(&ReqData).Error
}

// 删除部门
func (m *DeptService) DeleteDept(id uint) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminDept{}).Error
	return err
}

// 编辑部门
func (m *DeptService) UpdateDept(ReqData system.LyadminDept) (err error) {
	var oldData system.LyadminDept
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = ReqData.ParentId
	upDateMap["name"] = ReqData.Name
	upDateMap["sort"] = ReqData.Sort
	upDateMap["status"] = ReqData.Status
	upDateMap["owner"] = ReqData.Owner
	upDateMap["phone"] = ReqData.Phone
	upDateMap["email"] = ReqData.Email
	upDateMap["update_by"] = ReqData.UpdateBy

	db := global.GL_DB.Where("id = ?", ReqData.ID).Find(&oldData)
	err = db.Updates(upDateMap).Error
	return err
}
