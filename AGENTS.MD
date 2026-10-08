# golyadmin 项目 AGENTS 说明

**本文件是所有 AI Agent 编程工作的统一入口**。开始任何任务前先读完本文件，并按下面「AI 上下文地图」加载对应规则。

## AI 上下文地图（必读）

| 资源 | 路径 | 何时加载 |
|---|---|---|
| 项目开发规范（技术栈/分层/模型/路由/权限/数据库规范） | `.trae/rules/project_rules.md` | **任何后端或前端编码任务必读** |
| 技能 | `.trae/skills/`（技能目录，按需调用其中技能） | 技能skills |
| 种子数据变更询问规则 | 本文件「数据库初始化与种子数据」一节 | 涉及 model 字段/初始数据变更时 |

## 项目结构约定

- 后端：`backend/`（Go Gin + GORM），前端：`web/`（Vue3 + Element Plus）
- 详细开发规范见 `.trae/rules/project_rules.md`

## 数据库初始化与种子数据（重要）

种子数据由 `backend/utils/cmd/initdb/` 包维护（`initdb` 命令），**不再依赖 lyadmin_go.sql**：

```
backend/utils/cmd/initdb/
├── index.go        # cobra 命令：AutoMigrate 建表 + SeedAll
├── seed.go         # 幂等写入逻辑（每表 Count>0 则跳过；关联表原生 SQL）
├── data_dept.go    # 部门种子
├── data_role.go    # 角色种子
├── data_users.go   # 用户种子（含 bcrypt 密码哈希）
├── data_menu.go    # 菜单种子（树形 parent_id）
├── data_button.go  # 按钮种子
└── data_relations.go  # many2many 关联表种子（users_role/role_menu/role_menubutton 等）
```

## ⚠️ 变更询问规则（必须遵守）

当任务涉及以下任一变更时：

1. **修改 model 字段**（`model/system/` 下新增/删除/改名字段）
2. **修改种子数据相关模型**（LyadminUsers / LyadminRole / LyadminMenu / LyadminDept / LyadminButton / LyadminMenuButton）
3. **新增/调整初始业务数据**（如新项目默认菜单、默认角色、默认部门、默认配置）

在完成主任务前，**必须先询问用户**：

> 「该变更会影响到 initdb 种子数据，是否同步变更到 backend/utils/cmd/initdb？」

- 用户确认同步：修改 `initdb/data_*.go` 对应种子代码，保证新环境初始化与预期一致；若涉及 model 结构变化，同时核对 `seedTable` 的 SkipHooks 写入是否仍兼容（如字段默认值、必填列）
- 用户选择不同步：在完成报告中**明确提示**「initdb 种子数据未同步，新环境跑 initdb 初始化的数据与当前 model 不一致的风险」

## initdb 维护注意事项

- 种子写入使用 `SkipHooks: true`（跳过 LyadminMenu.AfterCreate 自动补建按钮、LyadminUsers.BeforeCreate 覆盖 uuid 等钩子副作用），手动加种子数据时注意与钩子行为的一致性
- 关联表（lyadmin_users_role 等）在 model 中是 many2many 标签、无独立结构体，种子走 `seedRelations` 原生 SQL，**列名以建表语句为准**（如 `lyadmin_users_id` 而非 `users_id`）
- 时间统一用 `mustTime("2006-01-02 15:04:05.000")` 解析为本地时区
- 种子数据变更后运行 `go build ./utils/cmd/initdb/` 验证编译
