# golyadmin 项目开发规则

## 项目概述

golyadmin 是一个基于 Go (Gin) + Vue3 (Element Plus) 的前后端分离后台管理系统。

## 技术栈

### 后端 (backend/)
- **语言**: Go >= 1.26
- **Web框架**: Gin v1.12
- **ORM**: GORM v1.31
- **配置**: Viper (YAML格式)
- **日志**: Zap
- **数据库**: MySQL(默认) / SQLite / SQL Server / PostgreSQL / Oracle
- **缓存**: Redis (go-redis/v9)
- **鉴权**: JWT (golang-jwt/jwt/v5)
- **API文档**: Swagger (swaggo)
- **命令行**: Cobra
- **定时任务**: robfig/cron

### 前端 (web/)
- **框架**: Vue 3.5 + Element Plus
- **构建工具**: Vite 6
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP**: Axios
- **国际化**: vue-i18n
- **富文本**: TinyMCE
- **图表**: ECharts
- **样式**: SCSS

## 目录结构规范

```
backend/
├── api/v1/              # API层 - 处理HTTP请求响应
│   └── system/          # 系统模块API
├── config/              # 配置结构体定义
├── docs/                # Swagger文档
├── global/              # 全局变量(DB/Redis/Config/Log)
├── initialize/          # 初始化逻辑(GORM/Redis/Router)
├── model/               # 数据模型
│   ├── common/request/  # 通用请求结构
│   └── system/          # 系统模块模型
│       ├── request/     # 请求DTO
│       └── response/    # 响应DTO
├── router/              # 路由定义
├── service/             # 业务逻辑层
├── utils/               # 工具函数
│   ├── cmd/             # cobra命令
│   ├── core/            # viper/zap核心
│   ├── middleware/      # 中间件
│   ├── pagination/      # 分页
│   ├── response/        # 响应封装
│   └── databases/       # 多数据库驱动
├── config.yaml          # 配置文件
└── main.go              # 入口文件

web/
├── public/              # 静态资源
└── src/
    ├── api/             # API请求封装
    │   ├── api.js       # 接口定义
    │   ├── request.js   # axios封装
    │   └── url.js       # URL前缀
    ├── assets/          # 资源文件(css/img)
    ├── components/      # 公共组件
    ├── config/          # 前端配置
    ├── icons/           # SVG图标
    ├── locales/         # 国际化
    ├── mixins/          # 混入
    ├── router/          # 路由配置
    ├── store/           # Pinia状态管理
    ├── utils/           # 工具函数
    ├── views/           # 页面视图
    ├── App.vue          # 根组件
    └── main.js          # 入口文件
```

## 后端开发规范

### 1. 模型定义规范

```go
// model/system/lyadmin_xxx.go
package system

import "gitee.com/lybbn/golyadmin/global"

type LyadminXxx struct {
    global.GL_BASE_MODEL           // 基础字段(ID/CreatedAt/UpdatedAt)
    Name      string `json:"name" form:"name" gorm:"type:varchar(50);comment:名称"`
    global.GL_CONTROL_MODEL        // 控制字段(创建者/更新者)
}

func (LyadminXxx) TableName() string {
    return "lyadmin_xxx"
}
```

**要点**:
- 结构体字段必须首字母大写(可导出)
- 使用 `json:"-"` 忽略敏感字段(如密码)
- 使用 `global.GL_BASE_MODEL` 和 `global.GL_CONTROL_MODEL` 嵌入公共字段
- 必须实现 `TableName()` 方法

### 2. API层规范

```go
// api/v1/system/lyadmin_xxx.go
package system

import (
    "gitee.com/lybbn/golyadmin/utils/response"
    "github.com/gin-gonic/gin"
)

// 列表
func (s *XxxApi) GetList(c *gin.Context) {
    // 1. 参数绑定
    // 2. 调用service
    // 3. 返回结果
    response.SuccessResponse(data, "获取成功", c)
}

// 新增
func (s *XxxApi) Create(c *gin.Context) {
    response.SuccessResponse(nil, "创建成功", c)
}

// 统一的响应方式:
// response.SuccessResponse(data, msg, c)  - 成功
// response.ErrorResponse(msg, c)          - 错误
// response.PaginateResponse(data, page, msg, c) - 分页
```

### 3. 响应规范

| Code | 含义 |
|------|------|
| 2000 | 成功 |
| 4000 | 通用错误 |
| 4001 | 认证失效(需重新登录) |

### 4. JWT认证规范

- 请求头格式: `Authorization: JWT {token}`
- Token过期前会自动刷新,通过响应头 `new-token` 返回新Token
- 支持单点/多点登录模式切换

### 5. 路由注册规范

```go
// router/system/index.go
package system

type RouterGroup struct {
    UserRouter
    MenuRouter
    // ...
}

// 路由前缀: /api/v1/
```

### 6. 权限中间件

- 超级管理员(identity=1)跳过权限检查
- 接口白名单配置在 `global/api_white.go`
- 权限格式: `API路径:HTTP方法` 正则匹配

### 7. 命令行工具

```bash
go run main.go start [-c 配置文件]    # 启动服务
go run main.go migrate [-d 数据库别名] # 数据库迁移
go run main.go createsuperuser -u xxx -p xxx  # 创建超管
go run main.go changepassword -u xxx -p xxx   # 修改密码
go run main.go version                # 查看版本
```

## 前端开发规范

### 1. API请求封装

```javascript
// src/api/api.js
// 命名规范: 模块名 + 操作名
export const apiSystemUser = params => ajaxGet({url: `system/user/getAdminUserList`, params})
export const apiSystemUserAdd = params => ajaxPost({url: `system/user/adminUser`, params})
export const apiSystemUserEdit = params => ajaxPut({url: `system/user/adminUser`, params})
export const apiSystemUserDelte = params => ajaxDelete({url: `system/user/adminUser`, params})

// 请求方法:
// ajaxGet    - GET请求
// ajaxPost   - POST请求
// ajaxPut    - PUT请求
// ajaxDelete - DELETE请求
// ajaxPatch  - PATCH请求
```

### 2. 路由规范

- 路由模式: Hash模式 (`createWebHashHistory`)
- 动态路由: 自动扫描 `views/` 目录下的 `.vue` 文件生成路由
- 路由守卫: 在 `router/index.js` 中统一处理认证和权限

### 3. 状态管理 (Pinia)

```javascript
// src/store/mutitabs.js - 多标签
// src/store/siteTheme.js - 主题设置
// src/store/keepAlive.js - 缓存
```

### 4. 组件规范

- 布局组件: `components/layout/`
- CRUD组件: `components/lycrud.vue` / `components/lyFormTable.vue`
- 上传组件: `components/upload/`
- 表单弹窗: `components/dialog/`

### 5. 样式规范

- 使用 SCSS, 全局样式在 `assets/css/common.scss`
- Element Plus 主题色通过 CSS 变量控制
- 支持暗黑模式 (`.dark` 类名切换)

## 数据库规范

### 表命名
- 前缀: `lyadmin_`
- 复数形式: `lyadmin_users`, `lyadmin_role`

### 公共字段
```go
// GL_BASE_MODEL
ID        uint      // 主键
CreatedAt time.Time // 创建时间
UpdatedAt time.Time // 更新时间

// GL_CONTROL_MODEL
CreatorID uint  // 创建者ID
UpdaterID uint  // 更新者ID
```

### 软删除
- 使用 `IsDelete` bool字段实现逻辑删除
- 状态字段: `IsActive` (true=正常, false=冻结)

## 配置规范

### 后端配置 (config.yaml)
```yaml
system:
  run-mode: debug          # debug/release
  http-port: 9000          # 服务端口
  db-type: mysql           # 数据库类型
  use-multipoint: true     # 多点登录
  router-prefix: "/api"    # 路由前缀
  is-demo: false           # 演示模式

jwt:
  secret-key: xxx          # 密钥
  expires-time: 7d         # 过期时间
  buffer-time: 1d          # 缓冲时间
```

### 前端配置 (src/config/index.js)
```javascript
API_BASEURL  // API基础地址
APP_VER      // 应用版本
LANG         // 默认语言 'zh-cn'
THEME        // 主题 'light'/'dark'
TIMEOUT      // 请求超时 350000ms
```

## 国际化

- 语言包路径: `web/src/locales/lang/`
- 支持: 中文(zh-cn) / 英文(en)

## 端口约定

- 后端服务: 9000
- 前端开发: 8090 (Vite dev server)
- 前端代理: `/api` → `http://127.0.0.1:9000/api/`

## 开发流程

### 后端新增模块
1. 在 `model/system/` 定义模型结构体(含request/response)
2. 在 `service/system/` 实现业务逻辑
3. 在 `api/v1/system/` 编写API接口
4. 在 `router/system/` 注册路由
5. 在 `model/migrate.go` 添加迁移模型

### 前端新增页面
1. 在 `api/api.js` 定义接口
2. 在 `views/` 创建页面组件
3. 自动路由会扫描并注册,或手动在 `router/index.js` 添加
4. 如需菜单管理,在后台配置菜单和权限

## 代码风格注释

- Go代码注释使用中文,说明函数用途、参数含义
- 模型字段使用 `comment:` 标签说明
- Vue组件使用 `<script setup>` 组合式API
- 关键逻辑添加行内注释

## 安全规范

- 密码使用 bcrypt 加密存储
- JWT Secret 生产环境必须修改
- SQL注入防护: 使用GORM参数绑定
- XSS防护: 前端过滤用户输入
- 演示模式下禁止写操作
- API白名单机制放行公开接口

## 启动命令

```bash
# 后端
cd backend
go mod tidy
go run main.go start

# 前端
cd web
npm install
npm run dev    # 开发模式
npm run build  # 生产构建
```
