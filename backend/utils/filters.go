package utils

import (
	"fmt"
	"strings"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func isApiWhiteList(method string, path string) bool {
	newapi := path + ":" + method
	var api_white_list []string
	for _, vm := range global.GL_API_WHILTELIST {
		if vm.DataSource {
			api_white_list = append(api_white_list, strings.Replace(vm.Api, ":id", "([a-zA-Z0-9-]+)", 1)+":"+vm.Method+"$")
		}
	}
	for _, item := range api_white_list {
		if RegexpMatch(item, newapi) {
			return true
		}
	}
	return false
}

type lyadminDeptSelect struct {
	ID       uint `json:"id" form:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;comment:主键"` //主键
	ParentId uint `json:"parent_id" gorm:"comment:上级部门"`
}

// 递归获取部门的所有下级部门
func getDept(deptId int, tempDeptList []lyadminDeptSelect) []int {
	var dept_list []int
	for _, v := range tempDeptList {
		if v.ParentId == uint(deptId) {
			dept_list = append(dept_list, int(v.ID))
			getDept(int(v.ID), tempDeptList)
		}

	}
	return dept_list
}

/*
		数据 级权限过滤器(白名单和超级管理员直接返回全部)
	    0. 获取用户的部门id，没有部门则返回空
	    1. 判断过滤的数据是否有创建人所在部门 "belong_dept" 字段,没有则返回全部
	    2. 如果用户没有关联角色则返回本部门数据
	    3. 根据角色的最大权限进行数据过滤(会有多个角色，进行去重取最大权限)
	    4. 只为仅本人数据权限时只返回过滤本人数据，并且部门为自己本部门(考虑到用户会变部门，只能看当前用户所在的部门数据)
	    5. 自定数据权限 获取部门，根据部门过滤
*/
func DataLevelPermissionsFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	//白名单
	method := c.Request.Method
	path := c.Request.URL.Path
	if isApiWhiteList(method, path) {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	uinfo := GetUserInfo(c)
	identity := uinfo.BaseClaims.Identity
	//超级管理员跳过
	if identity == 1 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	} else {
		userinfo := GetUserInfoDB(c)
		var roleIds []int
		var roleDataScopes []int
		var roleDeptIds []int
		for i, v := range userinfo.Role {
			roleIds = append(roleIds, int(v.ID))
			roleDataScopes = append(roleDataScopes, v.DataRange)
			for _, vm := range userinfo.Role[i].Dept {
				roleDeptIds = append(roleDeptIds, int(vm.ID))
			}
		}
		// 0. 获取用户的部门id，没有部门则返回空
		dept_id := userinfo.DeptId
		if dept_id < 1 {
			return func(db *gorm.DB) *gorm.DB {
				return db.Where("id < ?", 0)
			}
		}
		// 1. 判断过滤的数据是否有创建人所在部门 "belong_dept" 字段
		// 2. 如果用户没有关联角色则返回本部门数据
		role_list := roleIds
		if len(role_list) < 1 {
			return func(db *gorm.DB) *gorm.DB {
				return db.Where("belong_dept = ?", dept_id)
			}
		}
		// 3. 根据所有角色 获取所有权限范围
		dataScope_list := roleDataScopes
		if IsContainInt(dataScope_list, 3) {
			return func(db *gorm.DB) *gorm.DB {
				return db
			}
		}
		// 4. 只为仅本人数据权限时只返回过滤本人数据，并且部门为自己本部门(考虑到用户会变部门，只能看当前用户所在的部门数据)
		if IsContainInt(dataScope_list, 0) {
			return func(db *gorm.DB) *gorm.DB {
				return db.Where("create_by = ? and belong_dept = ?", uinfo.BaseClaims.ID, dept_id)
			}
		}
		// 5. 自定数据权限 获取部门，根据部门过滤
		var dept_list []int
		for _, v := range dataScope_list {
			//自定义数据权限
			if v == 4 {
				dept_list = append(dept_list, roleDeptIds...)
			} else if v == 2 {
				//本部门及以下数据权限
				dept_list = append(dept_list, int(dept_id))
				var tempDeptList []lyadminDeptSelect
				global.GL_DB.Model(&system.LyadminDept{}).Select("id", "parent_id").Find(&tempDeptList)
				dept_list = append(dept_list, getDept(int(dept_id), tempDeptList)...)
				fmt.Println(dept_list)
			} else if v == 1 {
				//本部门数据权限
				dept_list = append(dept_list, int(dept_id))
			}
		}
		dept_list = RemoveDuplicatesArrInt(dept_list)
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("belong_dept in (?)", dept_list)
		}
	}
}
