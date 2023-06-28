package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gorm.io/gorm"
)

type LyadminMenu struct {
	global.GL_BASE_MODEL
	ParentId      uint                `json:"parent_id" gorm:"comment:父菜单ID"`                 //父菜单ID
	Name          string              `json:"name" gorm:"comment:菜单名称"`                       //菜单名称
	Icon          string              `json:"icon" gorm:"comment:菜单图标"`                       //菜单图标
	WebPath       string              `json:"web_path" gorm:"comment:路由地址"`                   //路由地址
	IsLink        bool                `json:"is_link" gorm:"default:false;comment:是否外链"`      // 是否外链
	Visible       bool                `json:"visible" gorm:"default:true;comment:是否显示菜单"`     //是否显示菜单
	Component     string              `json:"component" gorm:"comment:对应前端文件路径"`              //对应前端文件路径
	ComponentName string              `json:"component_name" gorm:"comment:对应前端文件名称"`         //对应前端文件名称
	Sort          int                 `json:"sort" gorm:"default:1;comment:显示顺序"`             //显示顺序
	IsCatalog     bool                `json:"is_catalog" gorm:"default:false;comment:是否目录"`   //是否目录
	KeepAlive     bool                `json:"keep_alive" gorm:"default:false;comment:是否缓存页面"` //是否缓存页面
	Status        bool                `json:"status" gorm:"default:true;comment:状态"`          //状态
	MenuButtons   []LyadminMenuButton `json:"menuButtons" gorm:"foreignKey:MenuID;"`
	global.GL_CONTROL_MODEL
}

func (LyadminMenu) TableName() string {
	return "lyadmin_menu"
}

// 批量新增菜单按钮
func createMutiMenuButton(ReqData []LyadminMenuButton) error {
	return global.GL_DB.Create(&ReqData).Error
}

func (m *LyadminMenu) AfterCreate(tx *gorm.DB) (err error) {
	if !m.IsCatalog {
		mbs := []LyadminMenuButton{
			{MenuID: m.ID, Name: "新增", Value: "Create", Api: "", Method: "POST", GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: m.CreateBy, BelongDept: m.BelongDept}},
			{MenuID: m.ID, Name: "删除", Value: "Delete", Api: "", Method: "DELETE", GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: m.CreateBy, BelongDept: m.BelongDept}},
			{MenuID: m.ID, Name: "编辑", Value: "Update", Api: "", Method: "PUT", GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: m.CreateBy, BelongDept: m.BelongDept}},
			{MenuID: m.ID, Name: "查询", Value: "Search", Api: "", Method: "GET", GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: m.CreateBy, BelongDept: m.BelongDept}},
			{MenuID: m.ID, Name: "详情", Value: "Detail", Api: "", Method: "GET", GL_CONTROL_MODEL: global.GL_CONTROL_MODEL{CreateBy: m.CreateBy, BelongDept: m.BelongDept}},
		}
		err := createMutiMenuButton(mbs)
		if err != nil {
			global.GL_LOG.Error("自动创建菜单按钮失败:" + err.Error())
		}
	}
	return
}
