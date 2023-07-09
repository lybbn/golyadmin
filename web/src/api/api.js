import axios from 'axios';
import {reqExpost,ajaxGet,ajaxPost,ajaxDelete,ajaxPut,ajaxPatch,uploadImg,ajaxGetDetailByID,ajaxDownloadExcel} from './request';
import {url} from './url';

// 获取登录页的信息
export const login = params => ajaxPost({url: `base/login`,params})
// 获取验证码
export const getCaptcha = params => ajaxGet({url: `base/captcha`,params})
// 获取菜单
export const apiSystemWebRouter = params => ajaxGet({url: `system/menu/web_router`,params})
// 获取系统所有api列表
export const getSystemLyapiList = params => ajaxGet({url: `lyapi.json`,params})

/**
*系统配置
 * */

// 系统配置
export const platformsettingsSysconfig = params => ajaxGet({url: `platformsettings/sysconfig/`,params})
// 系统配置 -- 新增
export const platformsettingsSysconfigAdd = params => ajaxPost({url: `platformsettings/sysconfig/`,params})
// 系统配置 -- 编辑
export const platformsettingsSysconfigEdit = params => ajaxPut({url: `platformsettings/sysconfig/`,params})
// 系统配置 -- 删除
export const platformsettingsSysconfigDelete = params => ajaxDelete({url: `platformsettings/sysconfig/`,params})
// 系统配置 -- 保存子项
export const platformsettingsSysconfigSavecontent = params => ajaxPut({url: `platformsettings/sysconfig/save_content/`,params})
// 系统配置 -- 获取所有models列表信息
export const platformsettingsSysconfigGetmodelsInfoList = params => ajaxGet({url: `platformsettings/sysconfig/get_models_info_list/`,params})


/**
*系统管理
 * */
// 部门管理列表
export const apiSystemDept = params => ajaxGet({url: `system/dept/dept`,params})
// 部门管理列表 -- 新增部门
export const apiSystemDeptAdd = params => ajaxPost({url: `system/dept/dept`,params})
// 部门管理列表 -- 编辑部门
export const apiSystemDeptEdit = params => ajaxPut({url: `system/dept/dept`,params})
// 部门管理列表 -- 删除部门
export const apiSystemDeptDelete = params => ajaxDelete({url: `system/dept/dept`,params})


// 菜单管理列表
export const apiSystemMenu = params => ajaxGet({url: `system/menu/menu`,params})
// 菜单管理列表 -- 新增菜单
export const apiSystemMenuAdd = params => ajaxPost({url: `system/menu/menu`,params})
// 菜单管理列表 -- 编辑菜单
export const apiSystemMenuEdit = params => ajaxPut({url: `system/menu/menu`,params})
// 菜单管理列表 -- 删除菜单
export const apiSystemMenuDelete = params => ajaxDelete({url: `system/menu/menu`,params})
//获取部门数据,获取菜单树
export const systemMenuTree = params => ajaxGet({url: `system/menu_tree/`,params})
// 菜单管理 》 按钮列表(全部)
export const systemButton = params => ajaxGet({url:`system/button/button`,params})
// 菜单管理 》 按钮 编辑
export const systemButtonEdit = params => ajaxPut({url:`system/button/button`,params})
// 菜单管理 》 按钮 新增
export const systemButtonAdd = params => ajaxPost({url:`system/button/button`,params})
// 菜单管理 》 按钮 删除
export const systemButtonDelete = params => ajaxDelete({url:`system/button/button`,params})
// 菜单管理 》 按钮配置列表（针对单个菜单）
export const systemMenuButton = params => ajaxGet({url:`system/menu_button/menu_button`,params})
// 菜单管理 》 按钮配置列表（针对单个菜单） 新增
export const systemMenuButtonAdd = params => ajaxPost({url:`system/menu_button/menu_button`,params})
// 菜单管理 》 按钮配置列表（针对单个菜单） 编辑
export const systemMenuButtonEdit = params => ajaxPut({url:`system/menu_button/menu_button`,params})
// 菜单管理 》 按钮配置列表（针对单个菜单） 删除
export const systemMenuButtonDelete = params => ajaxDelete({url:`system/menu_button/menu_button`,params})



// 角色管理列表
export const apiSystemRole = params => ajaxGet({url: `system/role/roleList`,params})
// 获取全部角色管理
export const apiSystemRoleAll = params => ajaxGet({url: `system/role/role`,params})
// 角色管理列表-修改
export const apiSystemRoleEdit = params => ajaxPut({url: `system/role/role`,params})
// 角色管理列表 -- 新增角色
export const apiSystemRoleAdd = params => ajaxPost({url: `system/role/role`,params})
// 角色管理列表 -- 新增角色
export const apiSystemRoleDelete = params => ajaxDelete({url: `system/role/role`,params})
//通过角色id,获取菜单数据
export const apiSystemRoleIdToMenu = params => ajaxGet({url: `system/role/role_id_to_menu`,params})
export const apiSystemRoleIdToMenuid = (id) => ajaxGet({url: `system/role/role_id_to_menu/`+id})

//权限管理
// 权限管理 -- 保存
export const apiPermissionSave = params => ajaxPut({url: `system/role/permission`,params})

//管理员管理
export const apiSystemUser = params => ajaxGet({url: `system/user/getAdminUserList`,params})
//管理员管理-新增
export const apiSystemUserAdd = params => ajaxPost({url: `system/user/adminUser`,params})
//管理员管理-修改
export const apiSystemUserEdit = params => ajaxPut({url: `system/user/adminUser`,params})
//管理员管理-删除
export const apiSystemUserDelte = params => ajaxDelete({url: `system/user/adminUser`,params})

/**
 *日志管理
 * */
// 日志管理 查询
export const systemOperationlog= params => ajaxGet({url: `system/operation_log/loglist`,params})
// 日志管理 删除
export const systemOperationlogDelete= params => ajaxDelete({url: `system/operation_log/log`,params})
// 日志管理 清空全部日志
export const systemOperationlogDeletealllogsDelete= params => ajaxDelete({url: `system/operation_log/deletealllogs`,params})

/**
 *个人中心
 * */
// 获取当前个人用户信息
export const systemUserUserInfo= params => ajaxGet({url: `system/user/getUserInfo`,params})
// 更新修改当前个人用户信息
export const systemUserUserInfoEdit= params => ajaxPost({url: `system/user/setUserInfo`,params})
// 用户重置个人密码
export const systemUserChangePassword= params => ajaxPost({url: `system/user/changePassword`,params})


/**
 *消息中心
 * */

//消息公告
export const messagesMessagenotice = params => ajaxGet({url: `messages/messagenotice/`,params})
//消息公告-新增
export const messagesMessagenoticeAdd = params => ajaxPost({url: `messages/messagenotice/`,params})
//消息公告-修改
export const messagesMessagenoticeEdit = params => ajaxPut({url: `messages/messagenotice/`,params})
//消息公告-删除
export const messagesMessagenoticeDelete = params => ajaxDelete({url: `messages/messagenotice/`,params})



/**
 *省市区选择
 * */
// 省市区选择  获取省
export const getProvince= params => ajaxGet({url: `areas/`,params})
// 省市区选择  获取市/区
export const getCityDistrictByID= params => ajaxGetDetailByID({url: `areas/`,params})

// 省市区选择  根据详细地址获取经纬度
export const getAddressaccuracy= params => ajaxGet({url: `getaddressaccuracy/`,params})

// 省市区选择  递归获取所有省市区数据
export const getAllAreasList= params => ajaxGet({url: `getallareaslist/`,params})


/**
 *地区管理
 * */

// 地区管理列表
export const addressArea = params => ajaxGet({url: `address/area/`,params})
// 地区管理列表 获取根地区
export const addressAreaRoot = params => ajaxGet({url: `address/area/area_root/`,params})
// 地区管理列表 -- 新增
export const addressAreaAdd = params => ajaxPost({url: `address/area/`,params})
// 地区管理列表 -- 编辑
export const addressAreaEdit = params => ajaxPut({url: `address/area/`,params})
// 地区管理列表 -- 删除
export const addressAreaDelete = params => ajaxDelete({url: `address/area/`,params})

/**
 *用户管理
 * */

// 用户管理 列表
export const UsersUsers= params => ajaxGet({url: `user/user/getUserList`,params})
// 用户管理 新增
export const UsersUsersAdd= params => ajaxPost({url: `user/user/users`,params})
// 用户管理 编辑
export const UsersUsersEdit= params => ajaxPut({url: `user/user/users`,params})
// 用户管理 删除
export const UsersUsersDelete= params => ajaxDelete({url: `user/user/users`,params})
// 用户管理 禁用用户
export const UsersUsersdisableEdit= params => ajaxPost({url: `user/user/disableuser`,params})
// 用户管理 导出
export const UsersUsersExportexecl= params => ajaxGet({url: `user/user/exportexecl/`,params})

/**
*平台设置
*/
//轮播图列表
export const platformsettingsLunboimg= params => ajaxGet({url: `platformsettings/lunboimg/`,params})
// 轮播图列表 新增
export const platformsettingsLunboimgAdd= params => ajaxPost({url: `platformsettings/lunboimg/`,params})
// 轮播图列表 编辑
export const platformsettingsLunboimgEdit= params => ajaxPut({url: `platformsettings/lunboimg/`,params})
// 轮播图列表 删除
export const platformsettingsLunboimgDelete= params => ajaxDelete({url: `platformsettings/lunboimg/`,params})


//其他设置
export const platformsettingsOther= params => ajaxGet({url: `platformsettings/other/`,params})
// 其他设置 新增
export const platformsettingsOtherAdd= params => ajaxPost({url: `platformsettings/other/`,params})
// 其他设置 编辑
export const platformsettingsOtherEdit= params => ajaxPut({url: `platformsettings/other/`,params})
// 其他设置 删除
export const platformsettingsOtherDelete= params => ajaxDelete({url: `platformsettings/other/`,params})

// 平台设置 图片上传
export const platformsettingsUploadPlatformImg= params => uploadImg({url: `platformsettings/uploadplatformimg/`,params})

//前端访问操作 获取
export const superOerateGet= params => ajaxGet({url: `super/operate/`,params})
//前端访问操作 设置
export const superOerateSet= params => ajaxPost({url: `super/operate/`,params})