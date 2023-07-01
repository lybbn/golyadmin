<template>
    <div>
        <div ref="tableSelect">
            <el-breadcrumb separator-class="el-icon-arrow-right" style="margin-bottom: 16px;margin-top: 10px;">
                <el-breadcrumb-item :to="{ path: '/menuManage' }">菜单管理</el-breadcrumb-item>
                <el-breadcrumb-item >按钮管理</el-breadcrumb-item>
            </el-breadcrumb>
        </div>

        <el-table
            :height="tableHeight"
            border
            row-key="id"
            :data="tableData"
            ref="tableref"
            v-loading="loadingPage"
            style="width: 100%">
            <el-table-column type="index" width="55" align="center" label="序号"></el-table-column>
            <el-table-column min-width="150" prop="name" label="名称"></el-table-column>
            <el-table-column min-width="150" prop="value" label="key值"></el-table-column>
            <!-- <el-table-column min-width="150" prop="update_at" label="更新时间"></el-table-column> -->
            <el-table-column min-width="150" prop="created_at" label="创建时间">
                <template #default="scope">{{ formatDateTime(scope.row.created_at) }}</template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="140">
                <template #header>
                    <el-button type="primary" size="default" @click="handleEdit(null,'新增')" >新增</el-button>
                </template>
                <template #default="scope">
                    <span class="table-operate-btn" @click="handleEdit(scope.row,'编辑')" >编辑</span>
                    <span class="table-operate-btn" @click="handleEdit(scope.row,'delete')" >删除</span>
                </template>
            </el-table-column>
        </el-table>
        <add-button ref="addButtonFlag" @refreshData="getData"></add-button>
    </div>
</template>

<script setup>
    
    import {ref,reactive,onMounted,onBeforeUnmount,nextTick} from 'vue'
    import { ElMessage,ElMessageBox } from 'element-plus';
    import {dateFormats,formatDateTime, hasPermission} from "@/utils/util";
    import addButton from "./components/addButton";
    import {systemButton,systemButtonDelete} from '@/api/api'
    import useTableHight from '@/mixins/useTableHight';
    import { useRouter,useRoute } from 'vue-router'
    const route = useRoute()
    
    let isFull = ref(false)
    let orderStatic = ref(null)
    let tableSelect = ref(null)
    let tableHeight = useTableHight(orderStatic,tableSelect,isFull.value,true,-30)
    let loadingPage = ref(false)
    let formInline = ref({
        page: 1,
        limit: 999,
    })
    let pageparm = ref({
        page: 1,
        limit: 10,
        total: 0
    })
    let tableData = ref([])
    function setFull(){
        isFull.value=!isFull.value
        window.dispatchEvent(new Event('resize'))
    }
    // 表格序列号
    function getIndex($index) {
        // (当前页 - 1) * 当前显示数据条数 + 当前行数据的索引 + 1
        return (pageparm.value.page-1)*pageparm.value.limit + $index +1
    }

    let addButtonFlag = ref(null)

    function handleEdit(row, flag, menu) {
        if (flag == '编辑' || flag=='新增') {
            addButtonFlag.value.addButtonFn(row, flag)
        }
        if (flag == 'delete') {
            let vm = this
            ElMessageBox.confirm('您确定要删除选中的按钮？', {
                closeOnClickModal: false
            }).then(() => {
                systemButtonDelete({id: row.id}).then(res => {
                    if (res.code == 2000) {
                        ElMessage.success(res.msg)
                        getData()
                    } else {
                        ElMessage.warning(res.msg)
                    }
                })

            }).catch(() => {

            })
        }
    }
    function getData() {
        loadingPage.value = true
        systemButton(formInline.value).then(res => {
            loadingPage.value = false
            if (res.code == 2000) {
                tableData.value = res.data
            } else {
                ElMessage.warning(res.msg)
            }
        })
    }

    onMounted(()=>{
        getData()
    })

</script>

<style scoped>

</style>
