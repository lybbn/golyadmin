# golyadmin

基于 Go (Gin + GORM) + Vue3 (Element Plus) 的前后端分离后台管理系统，集成 JWT 鉴权、动态路由、动态菜单、RBAC 权限、数据级权限（基于部门）等功能。

[ 预览 ](https://golyadmin.lybbn.cn) | [ 官方文档 ](https://doc.lybbn.cn/) | [ gitee 地址 ](https://gitee.com/lybbn/golyadmin) | [ django-vue-lyadmin ](https://gitee.com/lybbn/django-vue-lyadmin)

**在线体验**：https://golyadmin.lybbn.cn 账号：admin 密码：123456

## 内置功能

- DashBoard 数据分析、面向配置的 CRUD、服务器实时监控面板（Windows/Linux）
- 部门管理（树结构 + 数据权限）、菜单管理（按钮/接口权限）、角色管理（菜单/数据权限）、权限管理
- 管理员管理、用户管理、个人中心、操作日志

特别鸣谢：部分设计模式参考 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

## 环境要求

- golang >= v1.26，node >= v16.19.1，MySQL >= 5.7（默认，亦支持 SQLite/SQLServer/PostgreSQL/Oracle）

## 快速开始

```bash
# 1. 后端
cd backend
go mod tidy

# 初始化数据库（建表 + 写入菜单/角色/用户等种子数据，无需 SQL 文件）
go run main.go initdb [-c 配置文件] [-d 数据库别名]

# 启动服务（默认端口 9000）
go run main.go start [-c 配置文件]

# 2. 前端
cd ../web
npm install
npm run dev    # 开发模式，默认端口 8090；生产构建 npm run build
```

访问 http://localhost:8090 ，使用 superadmin / 123456 登录。

### 其他命令

```bash
go run main.go migrate [-d 数据库别名]                      # 同步 model 表结构（model 注册到 model/migrate.go）
go run main.go createsuperuser -u superadmin -p 123456     # 创建超级管理员
go run main.go changepassword -u superadmin -p 123456      # 修改用户密码
go run main.go version                                     # 查看版本
```

## Swagger 文档

```bash
cd backend && swag init    # 生成文档，访问 http://localhost:9000/api/swagger/index.html
```

## 技术选型

- 后端：Gin + GORM + Viper + Zap + JWT + Redis（多点登录限制）+ Swagger
- 前端：Vue3 + Element Plus + Vite + Pinia + vue-i18n

## 线上部署

- **前后端分离**：前端 `npm run build` 后由 nginx 托管 dist，后端独立运行（linux `sh restart.sh` / windows `start.bat`）
- **集成部署（无 nginx）**：前端打包后的 dist 放入 backend，放开 `initialize/router.go` 中集成部署注释。后端默认端口 9000

## 项目二开

- 默认集成 mysql 和 sqlite 驱动，其他数据库到 `utils/databases/dbinitialize/` 放开相应驱动注释
- 验证码默认 mem 内存模式，需 redis 存储可修改 `api/v1/system/lyadmin_captcha.go`
- 操作日志默认不记录 GET 请求，如需记录修改 `utils/middleware/operation_log.go`

## 商用注意事项

商用请遵守 Apache2.0 协议并保留作者文件头部等信息声明。

## 交流

- 开发者WX号：laoyanyj
- QQ群 django-vue-lyadmin交流02群：877020250
