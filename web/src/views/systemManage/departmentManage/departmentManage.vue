<template>
    <div :class="{'ly-is-full':isFull}">
        <div class="tableSelect" ref="tableSelect">
            <el-form :inline="true" :model="formInline" label-position="left">
                <el-form-item label="关键词：">
                    <el-input size="default" v-model.trim="formInline.search" maxlength="60" placeholder="关键词" clearable @change="getData" style="width:200px"></el-input>
                </el-form-item>
                <el-form-item label="状态：">
                    <el-select v-model="formInline.status" placeholder="请选择" clearable @change="getData" size="default" style="width:100px">
                        <el-option
                            v-for="item in statusList"
                            :key="item.id"
                            :label="item.name"
                            :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label=""><el-button  @click="getData" type="primary" icon="Search" v-show="hasPermission(route.name,'Search')">查询</el-button></el-form-item>
                <el-form-item label=""><el-button  @click="handleEdit('','reset')" icon="Refresh">重置</el-button></el-form-item>
                <el-form-item label="">
                    <el-button icon="Plus" type="primary" @click="addDepart" v-show="hasPermission(route.name,'Create')">新增</el-button>
                </el-form-item>
            </el-form>
        </div>

        <div class="table">
            <el-table
                :max-height="tableHeight"
                border
                row-key="id"
                :data="tableData"
                ref="tableref"
                v-loading="loadingPage"
                style="width: 100%"
                lazy
                :load="loadChild"
                :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
                <!-- <el-table-column type="index" width="70" label="序号">
                    <template #default="scope">
                        <span v-text="getIndex(scope.$index)"></span>
                    </template>
                </el-table-column> -->
                <el-table-column type="" width="70" prop="id" label="ID"></el-table-column>
                <el-table-column min-width="180" prop="name" label="部门名称"></el-table-column>
                <el-table-column min-width="100" prop="owner" label="负责人"></el-table-column>
                <el-table-column min-width="120" prop="phone" label="联系电话"></el-table-column>
                <el-table-column min-width="120" prop="email" label="邮箱"></el-table-column>
                <el-table-column min-width="90"  label="状态">
                    <template #default="scope">
                        <el-tag v-if="scope.row.status==1" type="">启用</el-tag>
                        <el-tag v-else type="danger">禁用</el-tag>
                    </template>
                </el-table-column>
                <el-table-column min-width="70" prop="sort" label="排序"></el-table-column>
                <el-table-column min-width="160" prop="created_at" label="创建时间">
                    <template #default="scope">{{ formatDateTime(scope.row.created_at) }}</template>
                </el-table-column>
                <el-table-column label="操作" fixed="right" width="180">
                    <template #header>
                        <div style="display: flex;justify-content: space-between;align-items: center;">
                            <div>操作</div>
                            <div @click="setFull">
                                <el-tooltip content="全屏" placement="bottom">
                                    <el-icon ><full-screen /></el-icon>
                                </el-tooltip>
                            </div>
                        </div>
                    </template>
                    <template #default="scope">
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'edit')" v-show="hasPermission(route.name,'Update')">编辑</span>
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'detail')" v-show="hasPermission(route.name,'Detail')">详情</span>
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'delete')" v-show="hasPermission(route.name,'Delete')">删除</span>
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <add-department ref="addDepartmentFlag"  @refreshData="editAddrefreshData"></add-department>
    </div>
</template>
<script setup>
    import {ref,reactive,onMounted,onBeforeUnmount,nextTick} from 'vue'
    import { ElMessage,ElMessageBox } from 'element-plus';
    import AddDepartment from "./components/addDepartment.vue";
    import {apiSystemDept, apiSystemDeptDelete} from '@/api/api'
    import {hasPermission,formatDateTime} from "@/utils/util";
    import XEUtils from 'xe-utils'
    import useTableHight from '@/mixins/useTableHight';
    import { useRouter,useRoute } from 'vue-router'
    const route = useRoute()

    let isFull = ref(false)
    let orderStatic = ref(null)
    let tableSelect = ref(null)
    let tableHeight = useTableHight(orderStatic,tableSelect,isFull.value,true,-50)
    let loadingPage = ref(false)
    let formInline = ref({
        lazy:true,
        page: 1,
        limit: 10,
    })
    let pageparm = ref({
        page: 1,
        limit: 10,
        total: 0
    })
    let tableData = ref([])
    let statusList = ([
        {id:1,name:'启用'},
        {id:0,name:'禁用'},
    ])
    let loadMap = ref(new Map())
    function setFull(){
        isFull.value=!isFull.value
        window.dispatchEvent(new Event('resize'))
    }
    // 表格序列号
    function getIndex($index) {
        // (当前页 - 1) * 当前显示数据条数 + 当前行数据的索引 + 1
        return (pageparm.value.page-1)*pageparm.value.limit + $index +1
    }
    // 选项框选中数组
    let ids = ref([])
    // 选项框非单个禁用
    let single = ref(true)
    // 非多个禁用
    let multiple = ref(true)
    //多选项框被选中数据
    function handleSelectionChange(selection) {
        ids.value = selection.map(item => item.id);
        single.value = selection.length !== 1;
        multiple.value = !selection.length;
    }
    function loadChild(tree, treeNode, resolve){
        //将获取到的子节点存储到loadMap中
        loadMap.value.set(tree.id,{tree,treeNode,resolve})
        var params = {
            lazy:true,
            parent_id:tree.id,
            page:1,
            limit:999
        }
        apiSystemDept(params).then(async res => {
            if(res.code == 2000) {
                resolve(res.data)
            }else {
                ElMessage.warning(res.msg)
            }
        })
    }
    function refreshTree(parentid){
        if(!!parentid){
            const {tree,treeNode,resolve} = this.loadMap.get(parentid)
            this.$refs.tableref.store.states.lazyTreeNodeMap.value[parentid] = []//清空节点数据
            this.loadChild(tree,treeNode,resolve)
        }
    }
    function editAddrefreshData(parentid){
        getData()
        refreshTree(parentid)
    }
    let addDepartmentFlag = ref(null)
    function addDepart() {
        addDepartmentFlag.value.addDepartmentFn(null,'新增')
    }
    function handleEdit(row,flag) {
        if(flag=='edit') {
            addDepartmentFlag.value.addDepartmentFn(row,'编辑')
        }
        else if(flag == 'detail') {
            addDepartmentFlag.value.addDepartmentFn(row,'详情')
        }
        else if(flag=='delete') {
            ElMessageBox.confirm('您确定要删除选中的部门？',{
                closeOnClickModal:false
            }).then(()=>{
                apiSystemDeptDelete({id:row.id}).then(res=>{
                    if(res.code == 2000) {
                        ElMessage.message.success(res.msg)
                        getData()
                        refreshTree(row.parent_id)
                    } else {
                        ElMessage.message.warning(res.msg)
                    }
                })
            }).catch(()=>{

            })
        }
        else if(flag=="reset"){
            formInline.value = {
                lazy:true,
                page:1,
                limit: 999
            }
            pageparm.value={
                page: 1,
                limit: 10,
                total: 0
            }
            getData()
        }
    }
    //获取列表
    function getData(){
        loadingPage.value = true
        apiSystemDept(formInline.value).then(res => {
                loadingPage.value = false
                if(res.code == 2000) {
                    tableData.value = XEUtils.toArrayTree(res.data, { parentKey: 'parent_id', strict: false })
                } else {
                    ElMessage.warning(res.msg)
                }
            })
    }
    onMounted(()=>{
        getData()
    })
</script>
<style lang="scss" scoped>
    .tableNav{
        overflow: hidden;
        margin-bottom: 20px;
    }
    /*::v-deep(.el-table__placeholder){*/
    /*    display: unset;*/
    /*}*/
</style>