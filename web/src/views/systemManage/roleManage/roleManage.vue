<template>
    <div :class="{'ly-is-full':isFull}">
        <div class="tableSelect" ref="tableSelect">
            <el-form :inline="true" :model="formInline" label-position="left">
                <el-form-item label="关键词：">
                    <el-input v-model.trim="formInline.search" maxlength="60" placeholder="关键词" clearable @change="search" style="width:200px"></el-input>
                </el-form-item>
                <el-form-item label="状态：">
                    <el-select v-model="formInline.status" placeholder="请选择" clearable @change="search">
                        <el-option
                                v-for="item in statusList"
                                :key="item.id"
                                :label="item.name"
                                :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-button  @click="search" type="primary" icon="Search" v-show="hasPermission(route.name,'Search')">查询</el-button>
                <el-button  @click="handleEdit('','reset')" icon="Refresh">重置</el-button>
                <el-button type="primary" @click="handelAddRole(null,'新增')" icon="Plus" v-show="hasPermission(route.name,'Create')">新增</el-button>
            </el-form>
        </div>

        <div class="table">
            <el-table  :height="'calc('+(tableHeight)+'px)'" border :data="tableData" ref="tableref" v-loading="loadingPage" style="width: 100%">
                <!-- <el-table-column width="80" type="index" align="center" label="序号">
                    <template #default="scope">
                        <span v-text="getIndex(scope.$index)"></span>
                    </template>
                </el-table-column> -->
                <el-table-column type="" width="70" prop="id" label="ID"></el-table-column>
                <el-table-column min-width="120" prop="name" label="角色名称"></el-table-column>
                <el-table-column min-width="120" prop="key" label="权限字符"></el-table-column>
<!--                <el-table-column min-width="120" label="是否管理员">-->
<!--                    <template #default="scope">-->
<!--                        <el-tag v-if="scope.row.admin==1" type="">是</el-tag>-->
<!--                        <el-tag v-else type="danger">否</el-tag>-->
<!--                    </template>-->
<!--                </el-table-column>-->
                <el-table-column min-width="120" label="状态">
                    <template #default="scope">
                        <el-tag v-if="scope.row.status==1" type="">启用</el-tag>
                        <el-tag v-else>禁用</el-tag>
                    </template>
                </el-table-column>
                <el-table-column min-width="120" prop="sort" label="排序"></el-table-column>
                <el-table-column min-width="180" prop="created_at" label="创建时间">
                    <template #default="scope">{{ formatDateTime(scope.row.created_at) }}</template>
                </el-table-column>
                <el-table-column label="操作" fixed="right" width="280">
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
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'detail')" v-show="hasPermission(route.name,'Detail')">详情</span>
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'edit')" v-show="hasPermission(route.name,'Update')">编辑</span>
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'delete')" v-show="hasPermission(route.name,'Delete')">删除</span>
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'authority')" v-show="hasPermission(route.name,'Detail')">权限管理</span>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <Pagination v-bind:child-msg="pageparm" @callFather="callFather"></Pagination>
        <add-role ref="addRoleFlag"  @refreshData="getData"></add-role>
    </div>
</template>
<script setup>
    import {ref,reactive,onMounted,onBeforeUnmount,nextTick} from 'vue'
    import { ElMessage,ElMessageBox } from 'element-plus';
    import addRole from "./components/addRole";
    import Pagination from "@/components/Pagination";
    import {formatDateTime,hasPermission} from "@/utils/util";
    import {apiSystemRole,apiSystemRoleDelete} from '@/api/api'
    import useTableHight from '@/mixins/useTableHight';
    import { useRouter,useRoute } from 'vue-router'
    const route = useRoute()

    let isFull = ref(false)
    let orderStatic = ref(null)
    let tableSelect = ref(null)
    let tableHeight = useTableHight(orderStatic,tableSelect,isFull.value)
    let loadingPage = ref(false)
    let formInline = ref({
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
    function setFull(){
        isFull.value=!isFull.value
        window.dispatchEvent(new Event('resize'))
    }
    // 表格序列号
    function getIndex($index) {
        // (当前页 - 1) * 当前显示数据条数 + 当前行数据的索引 + 1
        return (pageparm.value.page-1)*pageparm.value.limit + $index +1
    }
    let addRoleFlag = ref(null)
    function handelAddRole(){
        addRoleFlag.value.addRoleFn(null,'新增')
    }

    function handleEdit(row,flag) {
        if(flag=='edit') {
            addRoleFlag.value.addRoleFn(row,'编辑')
        }
        else if(flag == 'detail') {
            addRoleFlag.value.addRoleFn(row,'详情')
        }
        else if(flag == 'authority') {
            // this.$router.push({name:'authorityManage',params:{id:row.id}})//已失效
            this.$router.push({name:'authorityManage',state:{id:row.id}})
            // this.$router.push({name:'authorityManage',query:{id:row.id}})
        }
        else if(flag=='delete') {
            ElMessageBox.confirm('您确定要删除选中的角色？',{
                closeOnClickModal:false
            }).then(()=>{
                apiSystemRoleDelete({id:row.id}).then(res=>{
                    if(res.code == 2000) {
                        ElMessage.success(res.msg)
                        search()
                    } else {
                        ElMessage.warning(res.msg)
                    }
                })
            }).catch(()=>{

            })
        }
        else if(flag=="reset"){
            formInline.value = {
                page:1,
                limit: 10
            }
            pageparm.value={
                page: 1,
                limit: 10,
                total: 0
            }
            getData()
        }
    }

    function callFather(parm) {
        formInline.value.page = parm.page
        formInline.value.limit = parm.limit
        getData()
    }
    function search() {
        formInline.value.page = 1
        formInline.value.limit = 10
        getData()
    }
    //获取列表
    function getData(){
        loadingPage.value = true
        apiSystemRole(formInline.value).then(res => {
            loadingPage.value = false
            if(res.code ==2000) {
                tableData.value = res.data.data;
                pageparm.value.page = res.data.page;
                pageparm.value.limit = res.data.limit;
                pageparm.value.total = res.data.total;
            }
        })
    }

    onMounted(()=>{
        getData()
    })
</script>
<style  scoped>
    .tableNav{
        overflow: hidden;
        margin-bottom: 20px;
    }
</style>



