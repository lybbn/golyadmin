<template>
    <div class="lyauthcontainer">
        <div class="authority-workspace">
            <aside class="role-panel">
                <div class="role-panel-heading">
                    <div>
                        <h2>角色列表</h2>
                        <p>选择角色配置权限</p>
                    </div>
                    <el-tag effect="light" round>{{ data.length }}</el-tag>
                </div>
                <el-input
                    v-model="filterText"
                    class="role-search"
                    clearable
                    placeholder="搜索角色">
                    <template #prefix><el-icon><Search /></el-icon></template>
                </el-input>
                <el-scrollbar class="role-list-scroll">
                    <el-tree
                        class="filter-tree role-tree"
                        :data="data"
                        :highlight-current="true"
                        :props="{ label: 'name' }"
                        default-expand-all
                        :filter-node-method="filterNode"
                        :expand-on-click-node="false"
                        @node-click="nodeClick"
                        node-key="node_id"
                        ref="tree" />
                </el-scrollbar>
                <div class="role-panel-hint">
                    <el-icon><InfoFilled /></el-icon>
                    保存后将更新该角色的菜单与数据范围。
                </div>
            </aside>

            <main class="permission-panel">
                <header class="permission-header">
                    <div>
                        <span class="permission-kicker">ROLE ACCESS</span>
                        <h1>权限配置</h1>
                        <p>管理角色可访问的菜单和业务数据范围</p>
                    </div>
                    <div class="permission-header-actions">
                        <el-tag v-if="roleObj.name" effect="light" round>{{ roleObj.name }}</el-tag>
                        <el-button type="primary" icon="Check" :disabled="!roleObj.name" @click="submitPermisson">
                            保存更改
                        </el-button>
                    </div>
                </header>

                <el-tabs v-if="roleObj.name" v-model="activeAuthTab" class="permission-tabs">
                    <el-tab-pane name="menus">
                        <template #label><span class="permission-tab-label"><el-icon><Menu /></el-icon>菜单权限</span></template>
                        <div class="tab-intro">
                            <div>
                                <h2>菜单与操作权限</h2>
                                <p>勾选角色可见的菜单，以及各菜单允许执行的操作。</p>
                            </div>
                        </div>
                        <el-scrollbar class="permission-tree-scroll">
                            <el-tree
                                class="lymenupermisson permission-tree"
                                ref="menuTree"
                                :key="roleObj.id"
                                :data="menuOptions"
                                node-key="id"
                                default-expand-all
                                show-checkbox
                                :expand-on-click-node="false"
                                :default-checked-keys="menuCheckedKeys"
                                :check-on-click-node="false"
                                :check-strictly="true"
                                empty-text="暂无可配置菜单"
                                @check-change="handleCheckClick">
                                <template #default="{ node, data }">
                                    <div class="menu-data">
                                        <div class="menu-name">{{ data.name }}</div>
                                        <div class="menu-buttons" @click.stop>
                                            <el-checkbox
                                                v-for="item in data.menuButtons"
                                                :key="item.id"
                                                v-model="item.checked">
                                                {{ item.name }}
                                            </el-checkbox>
                                        </div>
                                    </div>
                                </template>
                            </el-tree>
                        </el-scrollbar>
                    </el-tab-pane>

                    <el-tab-pane name="data">
                        <template #label><span class="permission-tab-label"><el-icon><OfficeBuilding /></el-icon>数据范围</span></template>
                        <div class="data-scope-layout">
                            <section class="scope-card">
                                <div class="tab-intro">
                                    <div>
                                        <h2>数据授权范围</h2>
                                        <p>设置该角色可以查看和操作的数据范围。</p>
                                    </div>
                                    <el-tooltip content="授权用户可操作的数据范围" placement="top">
                                        <el-icon class="scope-help"><QuestionFilled /></el-icon>
                                    </el-tooltip>
                                </div>
                                <el-select
                                    v-model="roleObj.data_range"
                                    class="data-scope-select"
                                    placeholder="请选择数据范围"
                                    @change="dataScopeSelectChange">
                                    <el-option
                                        v-for="item in dataScopeOptions"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value" />
                                </el-select>
                                <div v-if="roleObj.data_range !== 4" class="scope-note">
                                    <el-icon><InfoFilled /></el-icon>
                                    当前范围无需指定具体部门。
                                </div>
                            </section>
                            <section v-if="roleObj.data_range === 4" class="department-scope-card">
                                <div class="department-scope-heading">
                                    <h2>授权部门</h2>
                                    <span>选择该角色可访问的部门</span>
                                </div>
                                <el-scrollbar class="department-tree-scroll">
                                    <el-tree
                                        class="tree-border department-tree"
                                        :data="deptOptions"
                                        :key="roleObj.id"
                                        show-checkbox
                                        default-expand-all
                                        :default-checked-keys="deptCheckedKeys"
                                        ref="dept"
                                        node-key="id"
                                        :check-strictly="true"
                                        :props="{ label: 'name', children: 'children', disabled: 'disabled' }" />
                                </el-scrollbar>
                            </section>
                            <section v-else class="scope-preview-card">
                                <div class="scope-preview-icon"><el-icon><OfficeBuilding /></el-icon></div>
                                <div>
                                    <span>当前数据范围</span>
                                    <h2>{{ selectedDataScope?.label || '尚未设置' }}</h2>
                                    <p>{{ selectedDataScope?.description || '请选择数据范围。' }}</p>
                                </div>
                            </section>
                        </div>
                    </el-tab-pane>
                </el-tabs>

                <div v-else class="permission-empty-state">
                    <div class="empty-icon"><el-icon><Lock /></el-icon></div>
                    <h2>选择一个角色开始配置</h2>
                    <p>从左侧角色列表选择角色，即可设置菜单权限和数据范围。</p>
                </div>
            </main>
        </div>
    </div>
</template>

<script>
    import {apiSystemRoleAll,apiSystemRoleIdToMenuid, apiSystemDept,apiPermissionSave} from '@/api/api'
    import XEUtils from 'xe-utils'
    import {Search} from '@element-plus/icons-vue'
    export default {
        name: "authorityManage",
        components: {Search},
        data() {
            return {
                filterText: '',
                activeAuthTab:'menus',
                menuOptions: [],
                menuRequestId: 0,
                menuCheckedKeys: [], // 菜单默认选中的节点
                deptOptions:[],
                deptCheckedKeys:[],
                data: [],
                roleObj: {
                    name: null,
                    data_range: null
                },
                loadingPage:false,
                optionsData:[],
                optionsDataAll:[],
                dataScopeOptions: [
                    {
                        value: 0,
                        label: '仅本人数据权限',
                        description: '仅能访问当前账号关联的数据记录。'
                    },
                    {
                        value: 1,
                        label: '本部门数据权限',
                        description: '仅能访问当前所在部门的数据。'
                    },
                    {
                        value: 2,
                        label: '本部门及以下数据权限',
                        description: '可访问本部门及下属部门的数据。'
                    },
                    {
                        value: 3,
                        label: '全部数据权限',
                        description: '可访问系统中全部部门的数据。'
                    },
                    {
                        value: 4,
                        label: '自定数据权限',
                        description: '选择指定部门，限制该角色可访问的数据。'
                    }
                ],
            }
        },
        computed: {
            selectedDataScope () {
                if (this.roleObj.data_range === null || this.roleObj.data_range === undefined) return null
                return this.dataScopeOptions.find(item => item.value === Number(this.roleObj.data_range)) || null
            }
        },
        created() {
            this.pageRequest()
        },
        methods:{
            // 获取角色
            pageRequest (query) {
                return apiSystemRoleAll(query).then((res) => {

                    // res.map((value, index) => {
                    //     value.node_id = index
                    // })
                    this.data = res.data
                    this.data.map((value, index) => {
                            if(value.dept.length>0){
                                let tempdept = []
                                value.dept.forEach(item=>{
                                    tempdept.push(item.id)
                                })
                                value.dept = tempdept
                            }else{
                                value.dept = []
                            }
                            if(value.menu.length>0){
                                let tempmenu = []
                                value.menu.forEach(item=>{
                                    tempmenu.push(item.id)
                                })
                                value.menu = tempmenu
                            }else{
                                value.menu = []
                            }
                            if(value.permission.length>0){
                                let temppermission = []
                                value.permission.forEach(item=>{
                                    temppermission.push(item.id)
                                })
                                value.permission = temppermission
                            }else{
                                value.permission = []
                            }
                            value.node_id = index
                        })
                    this.$nextTick().then(() => {
                        this.initNode()
                    })
                })
            },
            initNode () {
                if (history.state.id && this.$refs.tree) {
                    this.data.map((value) => {
                        if (history.state.id === value.id) {
                            this.node_id = value.node_id
                        }
                    })
                    const node = this.$refs.tree.getNode(this.node_id)
                    this.$refs.tree.setCurrentKey(node)
                    this.nodeClick(node.data, node)
                }
            },
            // 部门数据
            getapiSystemDept(){
                apiSystemDept().then(res=>{
                    if(res.code ==2000) {
                        this.optionsDataAll = res.data.length > 0 ? res.data : []
                        let childrenList = res.data.filter(item=> item.parent_id)
                        let parentList = res.data.filter(item=> !item.parent_id)
                        if(parentList.length >0) {
                            parentList.forEach(item=>{
                                let children = childrenList.filter(itema=>itema.parent_id == item.id)
                                item.children=[...children]
                            })
                        }
                        this.optionsData = parentList
                    } else {
                        this.$message.warning(res.msg)
                    }
                })
            },

            // 提交修改
            submitPermisson() {
                if (!this.roleObj.name) return
                this.roleObj.menu = this.getMenuAllCheckedKeys() // 获取选中的菜单
                this.roleObj.dept = this.getDeptAllCheckedKeys() // 获取选中的部门
                const menuData = XEUtils.toTreeArray(this.menuOptions)
                const permissionData = []
                menuData.forEach((x) => {
                    const checkedPermission = x.menuButtons.filter((f) => {
                        return f.checked
                    })

                    if (checkedPermission.length > 0) {
                        for (const item of checkedPermission) {
                            permissionData.push(item.id)
                        }
                    }
                })
                this.roleObj.permission = permissionData

                this.updateRequest(this.roleObj)
            },
            updateRequest (row) {
                apiPermissionSave(row).then(res=>{
                    if(res.code ==2000) {
                        this.$message.success(res.msg)
                        this.pageRequest()
                    } else {
                        this.$message.warning(res.msg)
                    }
                })
            },
            // 获取菜单数据
            getMenuData (data) {
                const roleId = data.id
                const requestId = ++this.menuRequestId
                return apiSystemRoleIdToMenuid(roleId).then((res) => {
                    // 忽略快速切换角色时较晚返回的旧权限数据
                    if (requestId !== this.menuRequestId || this.roleObj.id !== roleId) return
                    res.data.forEach((x) => {
                        // 根据当前角色的permission,对menuButtons进行勾选处理
                        x.menuButtons.forEach((a) => {
                            if (data.permission.indexOf(a.id) > -1) {
                                // this.$set(a, 'checked', true)
                                // a.checked = true
                                a.checked = true
                            } else {
                                // this.$set(a, 'checked', false)
                                a.checked = false
                            }
                        })
                    })
                    // 将菜单列表转换为树形列表
                    this.menuOptions = XEUtils.toArrayTree(res.data, { parentKey: 'parent_id' })
                })
            },

            // 所有勾选菜单节点数据
            getMenuAllCheckedKeys () {
                if (!this.$refs.menuTree) return this.menuCheckedKeys || []
                // 目前被选中的菜单节点
                const checkedKeys = this.$refs.menuTree.getCheckedKeys()
                // 半选中的菜单节点
                const halfCheckedKeys = this.$refs.menuTree.getHalfCheckedKeys()
                checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys)
                return checkedKeys
            },
            // 所有自定义权限时,勾选的部门节点数据
            getDeptAllCheckedKeys () {
                if (!this.$refs.dept) return this.deptCheckedKeys || []
                // 目前被选中的部门节点
                const checkedKeys = this.$refs.dept.getCheckedKeys()
                // 半选中的部门节点
                const halfCheckedKeys = this.$refs.dept.getHalfCheckedKeys()
                checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys)
                return checkedKeys
            },
            filterNode (value, data) {
                if (!value) return true
                return String(data.name || '').toLowerCase().includes(String(value).toLowerCase())
            },
            // 获取部门数据
            getDeptData () {
                apiSystemDept({page:1,limit:9999}).then((res) => {
                     res.data.forEach(item=>{
                         item.disabled=false
                     })
                    // 将列表数据转换为树形数据
                    this.deptOptions = XEUtils.toArrayTree(res.data, { parentKey: 'parent_id', strict: false })
                })
            },
            // 角色树被点击
            nodeClick (data, node, self) {
                this.menuOptions = []
                this.roleObj = data
                this.getDeptData()
                this.getMenuData(data)
                this.menuCheckedKeys = data.menu // 加载已勾选的菜单
                this.deptCheckedKeys = data.dept
            },
            /** 选择角色权限范围触发 */
            dataScopeSelectChange (value) {
                if (value !== 4) {
                    // this.$refs.dept.setCheckedKeys([]);
                }
            },
            /**
             * 菜单树点击,全选权限部分数据
             * @param data
             */
            handleCheckClick(data, checked) {
                const setBranchChecked = (menu) => {
                    ;(menu.menuButtons || []).forEach((item) => {
                        item.checked = checked
                    })

                    ;(menu.children || []).forEach((child) => {
                        if (this.$refs.menuTree) {
                            this.$refs.menuTree.setChecked(child.id, checked)
                        }
                        setBranchChecked(child)
                    })
                }

                setBranchChecked(data)
            },

        },
        watch: {
            filterText (val) {
                this.$refs.tree.filter(val)
            }
        },
    }
</script>

<style lang="scss" scoped>
    .lyauthcontainer {
        /* 骨架账本 120 = 顶栏 60 + 标签条净高 40 + 内容区 padding 20（旧值 115 按原 33px 标签条校准） */
        height: calc(100vh - 120px);
        min-height: 480px;
        overflow: hidden;
        color: var(--ly-text-1);
    }

    .authority-workspace {
        display: grid;
        height: 100%;
        min-height: 0;
        grid-template-columns: 260px minmax(0, 1fr);
        gap: 14px;
    }

    .role-panel,
    .permission-panel {
        min-width: 0;
        min-height: 0;
        border: 1px solid var(--ly-glass-border);
        border-radius: 14px;
        background: var(--ly-glass-bg-strong);
        box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card);
        backdrop-filter: var(--ly-glass-blur);
        -webkit-backdrop-filter: var(--ly-glass-blur);
    }

    .role-panel {
        display: flex;
        flex-direction: column;
        padding: 17px 14px 12px;
    }

    .role-panel-heading {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 10px;
        padding: 0 3px;
    }

    .role-panel-heading h2,
    .permission-header h1,
    .tab-intro h2,
    .department-scope-heading h2 {
        margin: 0;
        color: var(--ly-text-1);
        font-weight: 600;
    }

    .role-panel-heading h2 {
        font-size: 15px;
    }

    .role-panel-heading p,
    .permission-header p,
    .tab-intro p {
        margin: 4px 0 0;
        color: var(--ly-text-3);
        font-size: 12px;
        line-height: 1.5;
    }

    :deep(.role-search) {
        margin-top: 15px;
    }

    :deep(.role-search .el-input__wrapper) {
        min-height: 36px;
        border-radius: 8px;
        background: var(--ly-glass-bg-soft);
        box-shadow: 0 0 0 1px var(--ly-line-soft) inset;
    }

    .role-list-scroll {
        flex: 1;
        min-height: 0;
        margin-top: 10px;
    }

    :deep(.role-tree.el-tree) {
        background: transparent;
        color: var(--ly-text-2);
    }

    :deep(.role-tree .el-tree-node__content) {
        height: 40px;
        margin: 2px 0;
        padding: 0 9px;
        border-radius: 8px;
        transition: color .16s ease, background-color .16s ease;
    }

    :deep(.role-tree .el-tree-node__content:hover) {
        color: var(--el-color-primary);
        background: rgba(58, 123, 255, .07);
    }

    :deep(.role-tree .is-current > .el-tree-node__content) {
        color: var(--el-color-primary);
        background: rgba(58, 123, 255, .10);
        font-weight: 600;
    }

    .role-panel-hint {
        display: flex;
        align-items: flex-start;
        gap: 7px;
        margin-top: 10px;
        padding: 10px 8px 0;
        border-top: 1px solid var(--ly-line-soft);
        color: var(--ly-text-3);
        font-size: 11px;
        line-height: 1.5;
    }

    .role-panel-hint .el-icon {
        flex: 0 0 auto;
        margin-top: 1px;
        color: var(--el-color-primary);
    }

    .permission-panel {
        display: flex;
        flex-direction: column;
        overflow: hidden;
        padding: 22px 24px 16px;
    }

    .permission-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 20px;
        padding-bottom: 17px;
        border-bottom: 1px solid var(--ly-line-soft);
    }

    .permission-kicker {
        display: block;
        margin-bottom: 4px;
        color: var(--el-color-primary);
        font-size: 10px;
        font-weight: 700;
        letter-spacing: .12em;
    }

    .permission-header h1 {
        font-size: 20px;
        line-height: 1.35;
    }

    .permission-header-actions {
        display: flex;
        align-items: center;
        gap: 10px;
    }

    :deep(.permission-header-actions .el-button) {
        min-height: 36px;
        border-radius: 8px;
        background: var(--ly-gradient-primary);
        box-shadow: 0 4px 10px rgba(58, 123, 255, .17);
    }

    .permission-empty-state {
        display: flex;
        flex: 1;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        min-height: 0;
        padding: 30px;
        text-align: center;
    }

    .empty-icon {
        display: grid;
        width: 58px;
        height: 58px;
        place-items: center;
        border: 1px solid rgba(58, 123, 255, .12);
        border-radius: 17px;
        color: var(--el-color-primary);
        background: rgba(58, 123, 255, .08);
        font-size: 24px;
    }

    .permission-empty-state h2 {
        margin: 17px 0 0;
        color: var(--ly-text-1);
        font-size: 16px;
        font-weight: 600;
    }

    .permission-empty-state p {
        max-width: 350px;
        margin: 7px 0 0;
        color: var(--ly-text-3);
        font-size: 12px;
        line-height: 1.6;
    }

    :deep(.permission-tabs.el-tabs) {
        display: flex;
        flex: 1;
        flex-direction: column;
        min-height: 0;
        margin-top: 13px;
    }

    :deep(.permission-tabs .el-tabs__header) {
        flex: 0 0 auto;
        margin-bottom: 13px;
    }

    :deep(.permission-tabs .el-tabs__content) {
        flex: 1;
        min-height: 0;
        overflow: hidden;
        padding: 0;
    }

    :deep(.permission-tabs .el-tab-pane) {
        display: flex;
        height: 100%;
        flex-direction: column;
        min-height: 0;
    }

    .permission-tab-label {
        display: inline-flex;
        align-items: center;
        gap: 7px;
    }

    .permission-tree-scroll,
    .department-tree-scroll {
        flex: 1;
        min-height: 0;
    }

    :deep(.permission-tree.el-tree),
    :deep(.department-tree.el-tree) {
        width: 100%;
        background: transparent;
        color: var(--ly-text-2);
    }

    :deep(.permission-tree .el-tree-node__content),
    :deep(.department-tree .el-tree-node__content) {
        min-height: 38px;
        height: auto;
        padding-top: 3px;
        padding-bottom: 3px;
        border-bottom: 1px solid var(--ly-line-soft);
    }

    .menu-data {
        display: flex;
        width: 100%;
        min-width: 0;
        align-items: center;
        gap: 16px;
        padding: 3px 8px 3px 0;
    }

    .menu-name {
        flex: 0 0 170px;
        overflow: hidden;
        color: var(--ly-text-1);
        font-size: 12px;
        font-weight: 500;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .menu-buttons {
        display: flex;
        min-width: 0;
        flex: 1;
        flex-wrap: wrap;
        gap: 0 14px;
    }

    :deep(.menu-buttons .el-checkbox) {
        margin-right: 0;
        color: var(--ly-text-2);
        font-size: 12px;
    }

    .data-scope-layout {
        display: grid;
        flex: 1;
        min-height: 0;
        grid-template-columns: minmax(260px, 340px) minmax(0, 1fr);
        gap: 16px;
        padding-top: 5px;
    }

    .scope-card,
    .department-scope-card,
    .scope-preview-card {
        min-width: 0;
        min-height: 0;
        padding: 18px;
        border: 1px solid var(--ly-line-soft);
        border-radius: 11px;
        background: var(--ly-glass-bg-soft);
    }

    .scope-card .tab-intro {
        display: flex;
        align-items: flex-start;
        justify-content: space-between;
        gap: 12px;
    }

    .tab-intro h2,
    .department-scope-heading h2 {
        font-size: 14px;
        line-height: 1.5;
    }

    .tab-intro p {
        font-size: 11px;
    }

    .scope-help {
        margin-top: 3px;
        color: var(--ly-text-3);
        cursor: help;
    }

    :deep(.data-scope-select) {
        width: 100%;
        margin-top: 20px;
    }

    .scope-note {
        display: flex;
        align-items: center;
        gap: 7px;
        margin-top: 14px;
        padding: 10px;
        border-radius: 8px;
        color: var(--ly-text-3);
        background: var(--ly-glass-bg-strong);
        font-size: 11px;
    }

    .scope-note .el-icon {
        color: var(--el-color-primary);
    }

    .department-scope-card {
        display: flex;
        flex-direction: column;
    }

    .department-scope-heading {
        display: flex;
        align-items: baseline;
        gap: 10px;
        padding-bottom: 12px;
        border-bottom: 1px solid var(--ly-line-soft);
    }

    .department-scope-heading span {
        color: var(--ly-text-3);
        font-size: 11px;
    }

    .department-tree-scroll {
        margin-top: 10px;
    }

    .scope-preview-card {
        display: flex;
        align-items: center;
        gap: 14px;
        padding: 22px;
    }

    .scope-preview-icon {
        display: grid;
        width: 44px;
        height: 44px;
        flex: 0 0 auto;
        place-items: center;
        border-radius: 12px;
        color: var(--el-color-primary);
        background: rgba(58, 123, 255, .09);
        font-size: 19px;
    }

    .scope-preview-card span {
        color: var(--ly-text-3);
        font-size: 11px;
    }

    .scope-preview-card h2 {
        margin: 4px 0;
        color: var(--ly-text-1);
        font-size: 15px;
        font-weight: 600;
    }

    .scope-preview-card p {
        margin: 0;
        color: var(--ly-text-2);
        font-size: 12px;
        line-height: 1.5;
    }

    @media (max-width: 900px) {
        .authority-workspace {
            grid-template-columns: 220px minmax(0, 1fr);
            gap: 10px;
        }

        .permission-panel {
            padding: 18px 18px 14px;
        }

        .menu-name {
            flex-basis: 135px;
        }
    }

    @media (max-width: 680px) {
        .lyauthcontainer {
            height: auto;
            min-height: 0;
            overflow: visible;
        }

        .authority-workspace {
            height: auto;
            grid-template-columns: minmax(0, 1fr);
        }

        .role-panel {
            height: 230px;
        }

        .permission-panel {
            min-height: 540px;
        }

        .data-scope-layout {
            grid-template-columns: minmax(0, 1fr);
        }

        .menu-data {
            align-items: flex-start;
            flex-direction: column;
            gap: 4px;
        }

        .menu-name {
            flex-basis: auto;
        }
    }
</style>
