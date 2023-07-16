package global

type urlInfo struct {
	Api        string
	Method     string
	DataSource bool //数据权限是否也是白名单 true 是白名单、false 否
}

// 权限验证需要排除的路由列表 GET\POST\PUT\DELETE
var GL_API_WHILTELIST = []urlInfo{
	{Api: "/api/system/menu/web_router", Method: "GET", DataSource: false},
	{Api: "/api/system/file/uploadFile", Method: "POST", DataSource: true},
}
