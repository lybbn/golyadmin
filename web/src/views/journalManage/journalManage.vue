<template>
    <div :class="{'ly-is-full':isFull}">
        <div class="tableSelect" ref="tableSelect">
            <el-form :inline="true" :model="formInline" label-position="left">
                <el-form-item label="关键词：">
                    <el-input size="default" v-model.trim="formInline.search" maxlength="60" style="width:160px;" clearable placeholder="关键词" @change="search"></el-input>
                </el-form-item>
                <el-form-item label="请求地址：">
                    <el-input size="default" v-model.trim="formInline.path" maxlength="60" style="width:150px;" clearable placeholder="请求地址" @change="search"></el-input>
                </el-form-item>
                <el-form-item label="请求方法：" v-if="showOtherSearch">
                    <el-input size="default" v-model.trim="formInline.method" maxlength="30" style="width:100px;" clearable placeholder="请求方法" @change="search"></el-input>
                </el-form-item>
                <el-form-item label="IP地址：" v-if="showOtherSearch">
                    <el-input size="default" v-model.trim="formInline.ip" maxlength="60" style="width:150px;" clearable placeholder="IP地址" @change="search"></el-input>
                </el-form-item>
                <el-form-item label="状态码：" v-if="showOtherSearch">
                    <el-input size="default" v-model.trim="formInline.code" maxlength="60" style="width:100px;" clearable placeholder="状态码" @change="search"></el-input>
                </el-form-item>
                <el-form-item label="创建时间：" v-if="showOtherSearch">
                    <el-date-picker
                            style="width:350px"
                            v-model="timers"
                            type="datetimerange"
                            @change="timeChange"
                            range-separator="至"
                            start-placeholder="开始日期"
                            end-placeholder="结束日期">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label=""><el-button  @click="search" type="primary" icon="Search" v-show="hasPermission(route.name,'Search')">查询</el-button></el-form-item>
                <el-form-item label=""><el-button  @click="handleEdit('','reset')" icon="Refresh">重置</el-button></el-form-item>
                <el-form-item label=""><el-button  @click="deleteAlllogs" type="danger" v-show="hasPermission(route.name,'Delete')">全部清空</el-button></el-form-item>
                <el-form-item label="" @click="clickMore" v-if="!showOtherSearch">
                    <span class="lysearchmore">展开
                        <el-icon><ArrowDown /></el-icon>
                    </span>
                </el-form-item>
                <el-form-item label="" @click="clickMore" v-if="showOtherSearch">
                    <span class="lysearchmore">收起
                        <el-icon><ArrowUp /></el-icon>
                    </span>
                </el-form-item>
            </el-form>
        </div>
        <el-table  :height="tableHeight"  border :data="tableData" ref="tableref" v-loading="loadingPage" style="width: 100%">
            <el-table-column type="index" width="60" align="center" label="序号">
                <template #default="scope">
                    <span v-text="getIndex(scope.$index)"></span>
                </template>
            </el-table-column>
            <el-table-column min-width="150" prop="path" label="请求地址" show-overflow-tooltip></el-table-column>
            <el-table-column width="90" prop="method" label="请求方法" show-overflow-tooltip></el-table-column>
            <el-table-column width="160" prop="ip" label="IP地址" show-overflow-tooltip></el-table-column>
            <el-table-column min-width="160" prop="agent" label="请求浏览器" show-overflow-tooltip></el-table-column>
            <el-table-column width="90" prop="body" label="请求数据">
                <template #default="scope">
                    <div>
                      <el-popover v-if="scope.row.body" placement="left-start" trigger="click" width="450px">
                        <div class="lypopoverbox">
                          <pre>{{ formatBody(scope.row.body) }}</pre>
                        </div>
                        <template #reference>
                          <el-icon style="cursor: pointer;"><warning /></el-icon>
                        </template>
                      </el-popover>
                      <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column width="70" label="状态码">
                <template #default="scope">
                    <el-tag :type="scope.row.code === 200?'success':'warning'">{{ scope.row.code }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column width="90" prop="resp" label="返回信息">
                <template #default="scope">
                    <div>
                      <el-popover v-if="scope.row.resp" placement="left-start" trigger="click" width="450px">
                        <div class="lypopoverbox">
                          <pre>{{ formatBody(scope.row.resp) }}</pre>
                        </div>
                        <template #reference>
                          <el-icon style="cursor: pointer;"><warning /></el-icon>
                        </template>
                      </el-popover>
                      <span v-else>无</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column min-width="160" label="操作人" show-overflow-tooltip>
                <template #default="scope">
                    <div v-if="scope.row.user_id !=0">{{ scope.row.user.username }}({{ scope.row.user.name }})</div>
                </template>
            </el-table-column>
            <el-table-column width="180" prop="created_at" label="创建时间" show-overflow-tooltip>
                <template #default="scope">{{ formatDateTime(scope.row.created_at) }}</template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="120">
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
                    <span class="table-operate-btn" @click="handleEdit(scope.row,'delete')" v-show="hasPermission(route.name,'Delete')">删除</span>
                </template>
            </el-table-column>
        </el-table>
        <Pagination v-bind:child-msg="pageparm" @callFather="callFather"></Pagination>
    </div>
</template>

<script setup>
    import {ref,reactive,onMounted,onBeforeUnmount,nextTick} from 'vue'
    import { ElMessage,ElMessageBox } from 'element-plus';
    import {systemOperationlog,systemOperationlogDelete,systemOperationlogDeletealllogsDelete} from '@/api/api'
    import Pagination from "@/components/Pagination";
    import {dateFormats,hasPermission,formatDateTime} from "@/utils/util";
    import useTableHight from '@/mixins/useTableHight';
    import { useRouter,useRoute } from 'vue-router'
    const route = useRoute()

    let isFull = ref(false)
    let orderStatic = ref(null)
    let tableSelect = ref(null)
    let tableHeight = useTableHight(orderStatic,tableSelect,isFull.value)
    let loadingPage = ref(false)
    let showOtherSearch = ref(false)//隐藏过长的搜索条件
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
    let timers = ref([])
    function setFull(){
        isFull.value=!isFull.value
        window.dispatchEvent(new Event('resize'))
    }
    // 表格序列号
    function getIndex($index) {
        // (当前页 - 1) * 当前显示数据条数 + 当前行数据的索引 + 1
        return (pageparm.value.page-1)*pageparm.value.limit + $index +1
    }
    const formatBody = (value) => {
        try {
            return JSON.parse(value)
        } catch (err) {
            return value
        }
    }

    function clickMore(){
        showOtherSearch.value = !showOtherSearch.value
        window.dispatchEvent(new Event('resize'))
    }
    function  deleteAlllogs(){
        ElMessageBox.confirm('是否确认清空全部日志数据', "警告", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning"
        }).then(function() {
            return systemOperationlogDeletealllogsDelete().then(res=>{
                if(res.code == 2000) {
                    ElMessage.success(res.msg)
                    search()
                } else {
                    ElMessage.warning(res.msg)
                }
            })
        })
    }
    function handleEdit(row,flag) {
        if(flag=='delete') {
            ElMessageBox.confirm('您确定要删除选中的数据吗？',{
                closeOnClickModal:false
            }).then(res=>{
                systemOperationlogDelete({id:row.id}).then(res=>{
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
        formInline.value.page = 1
        formInline.value.limit = 10
        getData()
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
    function getData() {
        loadingPage.value = true
        systemOperationlog(formInline.value).then(res => {
            loadingPage.value = false
            if(res.code ==2000) {
                tableData.value = res.data.data
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

<style scoped>
    .lypopoverbox {
        background: #132738;
        color: #c26a3e;
        height: 560px;
        width: 425px;
        overflow: auto;
    }
    .lypopoverbox::-webkit-scrollbar {
        display: none; /* Chrome Safari */
    }
</style>