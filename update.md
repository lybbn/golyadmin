# v2.0.1（2026-10-08）更新日志

## 新增

- `initdb` 命令：一键初始化数据库表结构和种子数据，不依赖 SQL 文件，可重复执行
- 生产环境配置文件 `config.pro.yaml`，通过 `GOLYADMIN_ENV=prod` 或 `-c` 指定
- `AGENTS.md`：AI Agent 编程统一入口
- `golyadmin-module-scaffold` 技能：AI 配合描述直接生成业务模块五层代码（model/service/api/router + 前端页面 + 菜单按钮种子）
- 部署文档

## 优化

- 前端整体升级 v4 玻璃拟态主题：三层设计 token、极光背景、玻璃卡片、暗色模式适配
- ECharts 统一注册 lyv4 主题，图表风格与全局主色一致
- 表格页高度计算常量校准（178→184），消除页面空滚动条
- 首页系统状态卡按占用率三档变色（<60% 绿 / 60~84% 黄 / ≥85% 红），整体状态标签联动
- 移除顶栏版权文字（登录页保留版权信息）
- 菜单「用户管理CRUD」更名为「用户管理」
- README 精简重写

## 修复

- Element Plus 2.13 兼容：el-radio/el-checkbox label 废弃迁移、el-tag 空 type 校验警告
- 面包屑分隔符图标不显示（ElementUI 2.x 图标类不兼容）
- 按钮配置页路由 404（缺少前导斜杠）
- 多标签栏与顶栏「双白带」、暗色模式侧栏残留亮色
- ECharts 主题注册 TypeError（ES 模块命名空间冻结）
- 布局溢出导致的右侧空滚动条（首页/权限管理/服务监控/地区管理骨架常量联动）
- 表单提示样式统一（el-alert 替换为轻量 .ly-form-tip）
