package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/uuid"
	"gorm.io/gorm"
)

/* 注意点 :
1. 	结构体里面的变量 (Name) 必须是首字符大写
gorm 指定类型
json 表示json接收时候的名称
binding required 表示必须传入
*/

type LyadminUsers struct {
	global.GL_BASE_MODEL
	UUID        string        `gorm:"<-:create;type:varchar(50);comment:uuid" form:"uuid" json:"uuid"`    // 允许读和创建
	Username    string        `json:"username" gorm:"type:varchar(50);not null;index;unique;comment:用户名"` //用户名
	Password    string        `json:"-"  gorm:"type:varchar(128);comment:密码"`                             //密码
	Name        string        `json:"name" gorm:"type:varchar(20);comment:姓名"`                            //姓名
	Nickname    string        `json:"nickname" gorm:"type:varchar(20);comment:昵称"`                        //昵称
	Mobile      string        `json:"mobile"  gorm:"type:char(25);comment:手机号"`                           //手机号
	Email       string        `json:"email"  gorm:"type:varchar(100);comment:邮箱"`                         //邮箱
	Avatar      string        `json:"avatar" gorm:"type:varchar(255);comment:头像"`                         //头像
	Gender      string        `json:"gender" gorm:"type:varchar(10);default:男;comment:性别"`                //性别（男、女）
	DeptId      uint          `json:"dept_id" gorm:"comment:部门"`                                          //部门ID
	Dept        LyadminDept   `json:"dept" gorm:"foreignKey:DeptId"`                                      //部门外键
	Post        []LyadminPost `json:"post" gorm:"many2many:lyadmin_users_post;"`                          //岗位
	Role        []LyadminRole `json:"role" gorm:"many2many:lyadmin_users_role;"`                          //角色
	IsStaff     bool          `json:"is_staff" gorm:"default:true;comment:是否可登录后台"`                       //是否可登录后台
	IsSuperuser bool          `json:"is_superuser" gorm:"default:false;comment:是否超管"`                     //是否超管
	IsActive    bool          `json:"is_active" gorm:"default:true;comment:状态(1正常、0冻结)"`                  //状态(1正常、0冻结)
	Identity    int           `json:"identity" gorm:"size:4;default:2;comment:身份(1 超级管理员 、2后台、3前台)"`      //身份（1 超级管理员,2后台用户、3前台用户）
}

func (LyadminUsers) TableName() string {
	return "lyadmin_users"
}

func (u *LyadminUsers) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.MakeUUID()
	return
}
