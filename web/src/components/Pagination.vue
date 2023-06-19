/**
* 分页组件
*/
<template>
    <div :class="isBackgroud?'lyPagination-page-bk':'lyPagination-page'">
        <el-pagination class="page-box" @size-change="handleSizeChange" @current-change="handleCurrentChange" background :small="small" :current-page="childMsg.page" :page-sizes="pageSizes" :page-size="childMsg.limit" :layout="layout" :total="childMsg.total"></el-pagination>
    </div>
</template>
<script setup>
    import {ref, onMounted, reactive, computed} from 'vue'
    import {useSiteThemeStore} from "@/store/siteTheme";

    const emit = defineEmits(["callFather"])

    const siteThemeStore = useSiteThemeStore()

    const pagingLayout = computed(() => {
        return siteThemeStore.pagingLayout
    })

    const isBackgroud = computed(()=>{
        return pagingLayout.value == 'backgroud'?true:false
    })

    const props = defineProps({
        childMsg: { type: Object, default: () => {} },
        pageSizes: { type: Array, default: [10,20,30,40,50,100] },
        layout: { type: String, default: "total, sizes, prev, pager, next, jumper" },
        small: {type:Boolean, default:false}
    })

    let pageparm = ref({
        page: props.childMsg.page || 1,
        limit: props.childMsg.limit || 20,
    })

    function handleSizeChange(val) {
        pageparm.value.limit = val
        pageparm.value.page = 1
        emit('callFather', pageparm.value)
    }

    function handleCurrentChange(val) {
        pageparm.value.page = val
        emit('callFather', pageparm.value)
    }

</script>

<style lang="scss" scoped>
    .lyPagination-page{
        display: flex;
        align-items: center;
        background: var(--el-fill-color-blank);
        border-bottom: 1px solid var(--el-border-color-lighter);
        border-left: 1px solid var(--el-border-color-lighter);
        border-right: 1px solid var(--el-border-color-lighter);
    }
    .lyPagination-page-bk{
        display: flex;
        align-items: center;
    }
    .page-box {
        margin: 10px auto;
        text-align: center;
        .el-pagination__editor.el-input{
            width: 70px !important;
            .el-input__inner{
                text-indent: 0 !important;
            }
        }
    }
</style>