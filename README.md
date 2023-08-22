## 基本介绍

### 项目介绍

> golyadmin是一个基于 golang、 [vue](https://vuejs.org) 和 [gin](https://gin-gonic.com) 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由，动态菜单，RBAC鉴权等功能，提供多种示例文件，让您的开发效率更高。前端admin部分依旧采用django-vue-lyadmin的前端，熟悉的权限设计和前端页面以及熟悉的代码。大大节省您多个不同语言项目切换时所花费的学习成本。

[ 预览 ](https://golyadmin.lybbn.cn) |[ 官方文档 ](https://doc.lybbn.cn/) | [django-vue-lyadmin](https://gitee.com/lybbn/django-vue-lyadmin)

## 在线体验

* 演示地址：[https://golyadmin.lybbn.cn](https://golyadmin.lybbn.cn) 账号：admin 密码：123456

![image-shouquan](https://gitee.com/lybbn/golyadmin/raw/master/web/src/assets/img/shouquan.png)

## 源码地址

* gitee地址(主推)：https://gitee.com/lybbn/golyadmin

## 内置功能

*  DashBoard： 数据分析查看
*  CRUD： 面向配置的crud功能
*  服务器监控面板（运维能力），支持windows和linux服务器的实时服务器资源状态监控
*  部门管理：配置系统组织机构（公司、部门、角色），树结构展现支持数据权限。
*  菜单管理：配置系统菜单，操作权限，按钮权限标识、后端接口权限等。
*  角色管理：角色菜单权限、数据权限、设置角色按部门进行数据范围权限划分。
*  权限管理：授权角色的权限范围。
*  管理员管理：主要管理系统管理员账号。
*  用户管理：主要管理前端用户。
*  个人中心：主要设置登录系统的个人昵称、密码等账号信息。
*  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。

特别鸣谢：本平台的一部分设计模式，部分参考[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

## 交流

- 开发者WX号：laoyanyj
- QQ群号：
* golyadmin交流01群：810799958 <a target="_blank" href="http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=ciRlLqTydhPej_GH_VQKgqyvWpvPftZo&authKey=JFbkoTwhq74OKbixC71VnxnIBqYY1WNce0bRI3954bl7YPPSzjBlRMsy0LH9Hm2k&noverify=0&group_code=810799958"><img border="0" src="//pub.idqqimg.com/wpa/images/group.png" alt="点击链接加入群聊" title="golyadmin交流01群"></a>

## 使用说明

```
- node版本 >= v16.19.1
- golang版本 >= v1.20
- IDE推荐：vscode
```

### 项目

使用 `vscode` 等编辑工具，打开项目目录

```bash

# 克隆项目
git clone https://gitee.com/lybbn/golyadmin.git
# 进入golyadmin/backend文件夹
cd golyadmin
cd backend
# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o golyadmin main.go (windows编译命令为go build -o golyadmin.exe main.go )

# 运行二进制
./golyadmin start [-c 配置文件] (windows运行命令为 golyadmin.exe start)

# 或直接非编译运行(开发环境)
go run main.go start [-c 配置文件]

```

其他命令(打包后go run main.go替换成golyadmin)

```bash
# migrate同步model表（需要同步的model写入modle/migrate.go中）

go run main.go migrate [-d 数据库别名] [-c 配置文件]

# 创建超级管理员用户

go run main.go createsuperuser -u superadmin -p 123456

# 修改用户密码

go run main.go changepassword -u superadmin -p 123456
```

### swagger自动化API文档

#### 安装 swagger

##### （1）可以访问外国网站

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn) 或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 如果嫌麻烦,可以使用go generate 编译前自动执行代码
cd golyadmin
cd backend
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 生成API文档

```bash
cd golyadmin
cd backend
swag init
```

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:9000/api/v1/swagger/index.html](http://localhost:9000/api/v1/swagger/index.html) 即可查看swagger文档

## lyadmin前端

#### 安装教程

```
cd frontend
npm install --registry=https://registry.npm.taobao.org
```

#### 使用说明

调试开发直接运行： 

```
npm start
```

默认监听端口：8090

#### 打包


```
npm run build
```

打包后静态文件在 dist 目录中

## 技术选型

- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 前端：用 [Vue3](https://vuejs.org) 和[ElementPlus](https://element-plus.org/zh-CN/) ，框架部分依然采用[django-vue-lyadmin](https://gitee.com/lybbn/django-vue-lyadmin)的前端部分。
- 数据库：采用`MySql` >= (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 线上部署

```
方法一、前后端分离部署：按正常分离模式部署即可
方法二、集成部署（不使用nginx）：前端执行打包命令 npm run build。把打包后的dist目录放入backend，然后在打开initialize/router.go下面相关集成部署注释即可
```

## 项目二开

```
- 默认集成mysql和sqlite驱动，如需sqlserver、oracle、postgresql需要到utils/databases/dbinitialize放开相应的驱动和方法注释
- 验证码captcha默认为mem内存模式(180s过期时间)，需要redis存储可修改api/v1/system/lyadmin_captcha.go
- 默认【操作日志】不记录GET请求，如需记录请修改utils/middleware/operation_log.go中注释掉说明方法
```

## 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者文件头部等信息声明。

## 捐赠该项目

开源不易，可使用支付宝、微信扫下面二维码打赏支持。您的支持是我不断创作的动力！！！

<table>
    <tr>
        <td><img src="https://gitee.com/lybbn/django-vue-lyadmin/raw/master/frontend/src/assets/img/alipay.jpg" height="300" width="400"/></td>
        <td><img src="https://gitee.com/lybbn/django-vue-lyadmin/raw/master/frontend/src/assets/img/wechat.jpg" height="300" width="400"/></td>
    </tr>
</table>
