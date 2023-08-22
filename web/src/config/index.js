//API DOMAIN
const API_DOMAIN = process.env.NODE_ENV === 'development' ? "127.0.0.1:9000" : "golyadmin.lybbn.cn"
// 接口地址
const API_BASEURL = process.env.NODE_ENV === 'development' ? "http://"+ API_DOMAIN +"/api/" : "https://"+ API_DOMAIN +"/api/"
//版本号
const APP_VER = require('../../package.json').version
//是否开启代理
const VUE_APP_PROXY = false

//lyadmin系统配置
module.exports = {

    //API DOMAIN
    API_DOMAIN : API_DOMAIN,

	//接口地址
	API_BASEURL : API_BASEURL,

	//是否开启代理
	VUE_APP_PROXY : VUE_APP_PROXY,

    //接口地址(支持本地跨域开发)
    API_URL : process.env.NODE_ENV === 'development' && VUE_APP_PROXY ? "/api/" : API_BASEURL,

    //标题
    APP_TITLE : "golyadmin后台管理系统",

	//APP版本
    APP_VER : APP_VER,

	//应用名称
    APP_NAME : "golyadmin后台管理系统",

	//项目布局: simple、msimple
	PROGRAM_LAYOUT: 'msimple',

	//是否开启多标签
	ISMULTITABS: true,

	//语言 简体中文 zh-cn、 英文 en（此功能只是示例）
	LANG: 'zh-cn',

	// elementplus 组件大小： small、default、large
	ELEMENT_SIZE: 'default',

	// elementplus 组件 zIndex
	ELEMENT_ZINDEX: 3000,

	// elementplus button组件 autoInsertSpace 是否自动在两个中文字符之间插入空格
	ELEMENT_BUTTON: false,

	//菜单默认是否折叠
	MENU_IS_COLLAPSE: false,

	//左侧菜单默认宽度 默认 200
	MENU_WIDTH: 200,

	// 左侧菜单和顶部导航颜色 默认 #272E39
	MENU_HEADER_COLOR:'#272E39',

	//主题颜色
	COLOR: '#1966ff',

    //默认主题 'dark' 暗黑、'light' 正常
	THEME: 'light',

	//分页组件样式:white、backgroud
	PAGING_LAYOUT: 'white',

    //登录信息数据存储方式 localStorage、sessionStorage
	STORAGE_METHOD: 'localStorage',

	//请求超时
	TIMEOUT: 350000,

}