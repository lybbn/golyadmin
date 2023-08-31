<template>
    <div :class="{'ly-is-full':isFull}">
        <div class="tableSelect" ref="tableSelect">
            <el-form :inline="true" :model="formInline" label-position="left">
                <el-form-item label="管理员姓名：">
                    <el-input  v-model.trim="formInline.name" maxlength="60"  clearable placeholder="管理员姓名" @change="search" style="width:200px"></el-input>
                </el-form-item>
                <el-form-item label="管理员账号：">
                    <el-input v-model.trim="formInline.username" maxlength="60"  clearable placeholder="管理员账号" @change="search" style="width:200px"></el-input>
                </el-form-item>
<!--                <el-form-item label="权限字符：">-->
<!--                    <el-input size="small" v-model.trim="formInline.name" maxlength="60" placeholder="权限字符" @change="search" style="width:200px"></el-input>-->
<!--                </el-form-item>-->
                <el-form-item label="状态：">
                    <el-select v-model="formInline.is_active" placeholder="请选择" clearable style="width: 120px" @change="search">
                        <el-option
                                v-for="item in statusList"
                                :key="item.id"
                                :label="item.name"
                                :value="item.id">
                        </el-option>
                    </el-select>
                </el-form-item>
<!--                <el-form-item label="创建时间：">-->
<!--                    <el-date-picker-->
<!--                            style="width:350px"-->
<!--                            v-model="timers"-->
<!--                            size="small"-->
<!--                            type="datetimerange"-->
<!--                            @change="timeChange"-->
<!--                            range-separator="至"-->
<!--                            start-placeholder="开始日期"-->
<!--                            end-placeholder="结束日期">-->
<!--                    </el-date-picker>-->
<!--                </el-form-item>-->
                <el-form-item label=""><el-button  @click="search" type="primary" icon="Search" v-show="hasPermission(route.name,'Search')">查询</el-button></el-form-item>
                <el-form-item label=""><el-button  @click="handleEdit('','reset')" icon="Refresh">重置</el-button></el-form-item>
                <el-form-item label=""><el-button  icon="Plus" @click="handelAddAdmin" type="primary" v-show="hasPermission(route.name,'Create')">新增</el-button></el-form-item>
            </el-form>
        </div>

        <div class="table">
            <el-table  :height="'calc('+(tableHeight)+'px)'" border :data="tableData" ref="tableref" v-loading="loadingPage" style="width: 100%">
                <!-- <el-table-column type="index" width="60" align="center" label="序号">
                    <template #default="scope">
                        <span v-text="getIndex(scope.$index)"></span>
                    </template>
                </el-table-column> -->
                <el-table-column width="70" prop="id" label="ID"></el-table-column>
                <el-table-column min-width="150" prop="username" label="管理员账号"></el-table-column>
                <el-table-column width="180" prop="name" label="姓名"></el-table-column>
                <el-table-column width="160" prop="dept" label="部门">
                    <template #default="scope">
                        <el-tag v-if="scope.row.dept">{{scope.row.dept.name}}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column min-width="120" prop="role" label="角色">
                    <template #default="scope">
                        <el-tag v-for="(item,index) in scope.row.role" :key="index" v-if="scope.row.role">{{item.name}}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column width="160" prop="mobile" label="手机号"></el-table-column>
                <el-table-column width="100" prop="gender" label="性别"></el-table-column>
                <el-table-column width="100" label="状态">
                    <template #default="scope">
                        <el-tag v-if="scope.row.is_active" type="success">正常</el-tag>
                        <el-tag v-else type="danger">禁用</el-tag>

                    </template>
                </el-table-column>
                <el-table-column min-width="150" prop="created_at" label="创建时间">
                    <template #default="scope">{{ formatDateTime(scope.row.created_at) }}</template>
                </el-table-column>
                <el-table-column label="操作" fixed="right" width="130">
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
                        <span class="table-operate-btn" @click="handleEdit(scope.row,'delete')" v-show="hasPermission(route.name,'Delete')">删除</span>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <Pagination v-bind:child-msg="pageparm" @callFather="callFather"></Pagination>
        <add-admin ref="addAdminFlag" @refreshData="getData"></add-admin>
    </div>
</template>
<script setup>
    import {ref,reactive,onMounted,onBeforeUnmount,nextTick} from 'vue'
    import { ElMessage,ElMessageBox } from 'element-plus';
    import addAdmin from "./components/addAdmin";
    import Pagination from "@/components/Pagination";
    import {dateFormats,formatDateTime, hasPermission} from "@/utils/util";
    import {apiSystemUser,apiSystemUserDelte} from '@/api/api'
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
    let timers = ref([])
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
    let addAdminFlag = ref(null)
    function handelAddAdmin() {
        addAdminFlag.value.addAdminFn(null,'新增')
    }

    function handleEdit(row,flag) {
        if(flag=='edit') {
            addAdminFlag.value.addAdminFn(row,'编辑')
        }
        else if(flag=='delete') {
            ElMessageBox.confirm('您确定要删除选中的管理员？',{
                closeOnClickModal:false,
                type: 'warning',
            }).then(res=>{
                apiSystemUserDelte({id:row.id}).then(res=>{
                    if(res.code == 2000) {
                        ElMessage.success(res.msg)
                        getData()
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
            timers.value = []
            getData()
        }
    }

    function callFather(parm) {
        formInline.value.page = parm.page
        formInline.value.limit = parm.limit
        getData()
    }
    function search() {
        formInline.page = 1
        formInline.limit = 10
        getData()
    }
    //获取列表
    function getData(){
        loadingPage.value = true
        apiSystemUser(formInline.value).then(res => {
            loadingPage.value = false
            if(res.code ==2000) {
                tableData.value = res.data.data
                pageparm.value.page = res.data.page;
                pageparm.value.limit = res.data.limit;
                pageparm.value.total = res.data.total;
            }
        })
    }

    function timeChange(val){
        if (val) {
            formInline.value.beginAt=dateFormats(val[0],'yyyy-MM-dd hh:mm:ss');
            formInline.value.endAt=dateFormats(val[1],'yyyy-MM-dd hh:mm:ss');
        } else {
            formInline.value.beginAt = null
            formInline.value.endAt = null
        }
        search()
    }
    onMounted(()=>{
        getData()
    })
</script>
