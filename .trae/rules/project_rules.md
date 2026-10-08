# golyadmin 项目开发规则

## 项目概述

golyadmin 是一个基于 Go (Gin) + Vue3 (Element Plus) 的前后端分离后台管理系统，集成 JWT 鉴权、动态路由、动态菜单、RBAC 权限、数据级权限（基于部门）等功能。

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
│   ├── global.go        # 全局变量定义
│   ├── model.go         # 公共模型字段(GL_BASE_MODEL/GL_CONTROL_MODEL)
│   └── api_white.go     # 接口白名单(GL_API_WHILTELIST)
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
│   ├── middleware/      # 中间件(jwt/permission/operation_log/cors/error/ssl)
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
    global.GL_CONTROL_MODEL        // 控制字段(CreateBy/UpdateBy/BelongDept)
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

### 2. 分组聚合模式

后端采用分组聚合模式组织 ApiGroup / ServiceGroup / RouterGroup：

```go
// service/system/index.go
type ServiceGroup struct {
    JwtService
    UserService
    MenuService
    RoleService
    // ...
}
var ServiceGroupApp = new(ServiceGroup)

// api/v1/system/index.go
type ApiGroup struct {
    UserApi
    MenuApi
    RoleApi
    // ...
}

// router/system/index.go
type RouterGroup struct {
    UserRouter
    MenuRouter
    // ...
}
```

各 api 文件通过包级变量引用 service，避免每次手动实例化。

### 3. API层规范

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
// response.ErrorCodeResponse(code, msg, c) - 自定义错误码
// response.PaginateResponse(data, page, msg, c) - 分页
```

### 4. 响应规范

| Code | 含义 |
|------|------|
| 2000 | 成功 |
| 4000 | 通用错误 |
| 4001 | 认证失效(需重新登录) |

### 5. JWT认证规范

- 请求头格式: `Authorization: JWT {token}`
- Token过期前会自动刷新,通过响应头 `new-token` 返回新Token
- 支持单点/多点登录模式切换（`system.use-multipoint` 配置）

### 6. 路由注册规范

```go
// router/system/index.go
package system

type RouterGroup struct {
    UserRouter
    MenuRouter
    // ...
}

// 路由全局前缀: /api (由 config.yaml 的 system.router-prefix 控制)
// 实际接口路径示例: POST /api/system/user/adminUser
```

### 7. 路由分组机制（核心架构）

[initialize/router.go](file:///d:/laoyansoft/python/golyadmin/backend/initialize/router.go) 中将路由分为两组：

- **PublicGroup**（不需认证）：登录、验证码等基础接口，调用 `InitBaseRouter`
- **PrivateGroup**（需认证 + 权限）：挂载 `JWTAuthMiddleware()` + `PermissionMiddleware()` 中间件，再按模块分组：
  - `system` 分组：adminUser / operationLog / menu / role / dept / file / system 等
  - `user` 分组：前台用户管理

中间件链顺序：`CORS → GinRecovery(异常捕获) → JWTAuth → Permission → OperationLog(按路由挂载)`

### 8. 权限中间件与白名单

- 超级管理员(identity=1)跳过权限检查
- 接口白名单配置在 `global/api_white.go`，变量名为 `GL_API_WHILTELIST`（注意拼写）
- 白名单支持 `DataSource` 字段标识是否同时豁免数据权限
- 权限格式: `API路径:HTTP方法` 正则匹配，`:id` 会被转成 `([a-zA-Z0-9-]+)`

### 9. 数据级权限过滤

service 层广泛使用 `utils.DataLevelPermissionsFilter()` 实现基于部门的数据权限控制：

```go
db := tx.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model, id)
if db.RowsAffected == 0 {
    return errors.New("无该数据权限")
}
```

依赖 `GL_CONTROL_MODEL.BelongDept` 字段做数据归属判定，是 RBAC 之外的重要权限维度。

### 10. 用户身份标识（identity）

`LyadminUsers.Identity` 字段区分用户类型：

| Identity | 含义 |
|----------|------|
| 1 | 超级管理员（跳过权限检查） |
| 2 | 后台用户（管理端） |
| 3 | 前台用户（用户端） |

### 11. 操作日志中间件

`OperationLog()` 中间件用于记录操作日志（按路由挂载）：

- **默认不记录 GET 请求**，如需记录需注释 `operation_log.go` 中相关代码
- 对登录 `/api/base/login`、改密 `/api/system/user/changePassword` 等敏感接口做密码脱敏
- 文件上传接口 `/api/system/file/uploadFile` 不记录 body
- 文件下载/导出响应体超过 1024 字节会截断

### 12. 命令行工具

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
// ajaxGetDetailByID - 单例详情 GET /api/xxx/:id
// ajaxDownloadExcel - 下载excel文件流(blob)
// uploadImg  - 上传图片(FormData)
```

### 2. 路由规范

- 路由模式: Hash模式 (`createWebHashHistory`)
- 动态路由: 自动扫描 `views/` 目录下的 `.vue` 文件生成路由（排除 components 子目录、index.vue、login.vue、lyterminal.vue）
- 路由守卫: 在 `router/index.js` 中统一处理认证和权限，白名单路由：buttonConfig / menuManage / lyterminal / buttonManage / lyFilePreview

### 3. 状态管理 (Pinia)

```javascript
// src/store/mutitabs.js - 多标签、登录token、用户信息
// src/store/siteTheme.js - 主题设置、颜色、菜单样式
// src/store/keepAlive.js - 页面缓存
```

### 4. 组件规范

- 布局组件: `components/layout/`
- CRUD组件: `components/lycrud.vue` / `components/lyFormTable.vue`
- 上传组件: `components/upload/`
- 表单弹窗: `components/dialog/`
- 富文本: `components/TEditor.vue` / `components/teditorjs/`
- 图表: `components/analysis/`

### 5. 样式规范

- 使用 SCSS, 全局样式在 `assets/css/common.scss`
- Element Plus 主题色通过 CSS 变量控制
- 支持暗黑模式 (`.dark` 类名切换)

## 数据库规范

### 表命名
- 前缀: `lyadmin_`
- 命名不统一，大部分为单数：`lyadmin_role`、`lyadmin_menu`、`lyadmin_dept`、`lyadmin_menu_button`、`lyadmin_button`、`lyadmin_post`、`lyadmin_operation_log`
- 少数复数：`lyadmin_users`
- 关联表：`lyadmin_users_role`、`lyadmin_role_menu`

### 公共字段
```go
// GL_BASE_MODEL (global/model.go)
ID        uint      `gorm:"type:bigint;primaryKey;autoIncrement"` // 主键
CreatedAt time.Time // 创建时间
UpdatedAt time.Time // 更新时间

// GL_CONTROL_MODEL (global/model.go)
CreateBy   uint `gorm:"index"` // 创建者
UpdateBy   uint `gorm:"index"` // 更新者
BelongDept uint `gorm:"index"` // 数据归属部门（数据权限关键字段）
```

### 删除策略
- **实际删除采用 GORM 硬删除**：`tx.Select(clause.Associations).Delete(&model)`（连同关联表一起删除）
- `IsDelete` 字段虽定义但未在删除逻辑中使用
- 状态控制通过 `IsActive` 字段（true=正常, false=冻结），用于禁用而非删除

### 数据库迁移
需同步的 model 在 [model/migrate.go](file:///d:/laoyansoft/python/golyadmin/backend/model/migrate.go) 的 `MigrateModelList` 中注册，执行 `go run main.go migrate` 同步表结构。

## 配置规范

### 后端配置 (config.yaml)
```yaml
system:
  run-mode: debug          # debug/release
  host: 0.0.0.0
  http-port: 9000          # 服务端口
  is-cors: true            # 是否启用跨域
  db-type: mysql           # 数据库类型
  use-multipoint: true     # 多点登录
  router-prefix: "/api"    # 路由全局前缀
  upload-dir: "media/uploadfile/"  # 文件上传目录
  url-prefix: http://127.0.0.1:9000  # 文件访问地址前缀
  is-demo: false           # 演示模式(只读)
  is-swagger: true         # 是否注册swagger
  gorm-log-mode: ""        # GORM日志等级 silent/error/warn/info

jwt:
  secret-key: xxx          # 密钥(生产环境必须修改)
  expires-time: 7d         # 过期时间(d/h/m)
  buffer-time: 1d          # 缓冲时间(临近过期签发新token)

cors:
  mode: allow-all          # allow-all / whitelist / strict-whitelist
  whitelist: [...]         # 白名单模式下的域名配置
```

### 前端配置 (src/config/index.js)
```javascript
API_DOMAIN   // API域名
API_BASEURL  // API基础地址(含/api/后缀)
VITE_APP_PROXY // 是否走Vite代理(默认false,直接请求后端)
APP_VER      // 应用版本
LANG         // 默认语言 'zh-cn'
THEME        // 主题 'light'/'dark'
TIMEOUT      // 请求超时 350000ms
```

## 部署与访问

### 端口约定
- 后端服务: 9000
- 前端开发: 8090 (Vite dev server)

### 前端代理
- 默认 `VITE_APP_PROXY=false`，前端直接请求 `http://127.0.0.1:9000/api/`
- 设为 `true` 时走 Vite 代理：`/api` → `http://127.0.0.1:9000/api/`（rewrite 去掉 `/api` 前缀再拼接到 target）

### 静态文件服务
- 后端提供 `/media` 静态文件服务（文件上传目录 `media/uploadfile/`）
- 文件访问 URL 前缀由 `system.url-prefix` 配置

### Swagger 文档
- 访问地址：`http://localhost:9000/api/swagger/index.html`
- 生成命令：`swag init`（在 backend 目录执行）

### 部署方式
- **方式一 前后端分离部署**：前端 `npm run build` 后由 nginx 托管，后端独立运行
- **方式二 集成部署（不使用 nginx）**：前端打包后 dist 放入 backend，放开 `initialize/router.go` 中 `LoadHTMLGlob` / `Static` 相关注释

## 其他重要机制

### 验证码
- 默认 mem 内存模式（180s 过期）
- 支持 redis 存储，需修改 `api/v1/system/lyadmin_captcha.go`
- 类型支持：math(计算) / digit / string / chinese / audit

### 多数据库支持
- 默认仅集成 mysql 和 sqlite 驱动
- sqlserver / oracle / postgresql 需到 `utils/databases/dbinitialize/` 放开相应驱动和方法注释
- 支持多数据源，通过 `alias-name` 区分，访问使用 `global.GL_DATABASES[别名]`

### 错误处理
- `GinRecovery` 中间件统一捕获 panic
- 404 路由统一返回 `{code:404, msg:"404 not found"}`

## 国际化

- 语言包路径: `web/src/locales/lang/`
- 支持: 中文(zh-cn) / 英文(en)

## 开发流程

### 后端新增模块
1. 在 `model/system/` 定义模型结构体(含request/response)
2. 在 `model/migrate.go` 的 `MigrateModelList` 注册迁移模型
3. 在 `service/system/` 实现业务逻辑，并在 `service/system/index.go` 的 `ServiceGroup` 中聚合
4. 在 `api/v1/system/` 编写API接口，并在 `api/v1/system/index.go` 的 `ApiGroup` 中聚合
5. 在 `router/system/` 注册路由，并在 `router/system/index.go` 的 `RouterGroup` 中聚合
6. 在 `initialize/router.go` 的对应分组中调用 `InitXxxRouter`

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

- 密码使用 `utils.MakePassowrd` 加密存储（bcrypt）
- JWT Secret 生产环境必须修改
- SQL注入防护: 使用GORM参数绑定
- XSS防护: 前端过滤用户输入
- 演示模式下禁止写操作（`is-demo: true` 时仅允许 GET/OPTIONS）
- API白名单机制放行公开接口
- 敏感接口（登录、改密）操作日志做密码脱敏

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
