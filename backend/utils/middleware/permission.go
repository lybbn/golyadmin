package middleware

import (
	"strings"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

func hasPermission(method string, path string, btnlist []system.LyadminMenuButton) bool {
	if len(btnlist) < 1 {
		return false
	}
	newapi := path + ":" + method
	var new_api_list []string
	for _, vm := range btnlist {
		if vm.Api != "" {
			new_api_list = append(new_api_list, strings.Replace(vm.Api, ":id", "([a-zA-Z0-9-]+)", 1)+":"+vm.Method+"$")
		}
	}
	for _, item := range new_api_list {
		if utils.RegexpMatch(item, newapi) {
			return true
		}
	}
	return false
}

func isApiWhiteList(method string, path string) bool {
	newapi := path + ":" + method
	var api_white_list []string
	for _, vm := range global.GL_API_WHILTELIST {
		api_white_list = append(api_white_list, strings.Replace(vm.Api, ":id", "([a-zA-Z0-9-]+)", 1)+":"+vm.Method+"$")
	}
	for _, item := range api_white_list {
		if utils.RegexpMatch(item, newapi) {
			return true
		}
	}
	return false
}

// PermissionMiddleware 判断用户接口权限
func PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uinfo := utils.GetUserInfo(c)
		identity := uinfo.BaseClaims.Identity
		// 演示模式判断
		allow_methodname_list := []string{"GET", "OPTIONS"}
		if global.GL_CONFIG.System.IsDemo && !utils.IsContainStr(allow_methodname_list, c.Request.Method) {
			response.ErrorResponse("演示模式，不允许操作!", c)
			c.Abort()
			return
		}
		//超级管理员跳过
		if identity == 1 {
			c.Next()
		} else {
			method := c.Request.Method
			path := c.Request.URL.Path
			if isApiWhiteList(method, path) {
				c.Next()
			} else {
				userinfo := utils.GetUserInfoDB(c)
				roleIds := []int{}
				for _, v := range userinfo.Role {
					roleIds = append(roleIds, int(v.ID))
				}
				if len(roleIds) < 1 {
					response.ErrorResponse("暂无访问该接口权限:无角色", c)
					c.Abort()
					return
				}
				rolelist := []system.LyadminRole{}
				err := global.GL_DB.Model(&system.LyadminRole{}).Where("status = ? and id in (?)", 1, roleIds).Preload("Permission").Find(&rolelist).Error
				if err != nil {
					global.GL_LOG.Error("查询数据库失败：" + err.Error())
					response.ErrorResponse("系统繁忙，请稍后再试", c)
					c.Abort()
					return
				}
				menubtnlist := []system.LyadminMenuButton{}
				for _, rl := range rolelist {
					for _, mt := range rl.Permission {
						menubtnlist = append(menubtnlist, mt)
					}
				}
				if !hasPermission(method, path, menubtnlist) {
					response.ErrorResponse("暂无访问该接口权限", c)
					c.Abort()
					return
				} else {
					c.Next()
				}
			}

		}
	}
}
