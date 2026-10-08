<template>
    <el-row :gutter="20">
        <el-col :span="24" >
            <div  class="space-inner">
                <el-tabs type="border-card" class="lycard" v-model="activeName" @tab-change="handleTabChage">
                    <el-skeleton :rows="rows" :animated="animated" :count="count" :loading="loading" style="padding: 20px;width: auto;overflow: hidden;">
                        <template #default>
                            <el-tab-pane label="订单分析" name="tab1">
                                <ly-line-echart ref="lyecharts1" v-if="activeName == 'tab1'"></ly-line-echart>
                            </el-tab-pane>
                            <el-tab-pane label="访问量" name="tab2" >
                                <ly-bar-echart ref="lyecharts2" v-if="activeName == 'tab2'"></ly-bar-echart>
                            </el-tab-pane>
                        </template>
                    </el-skeleton>
                </el-tabs>
            </div>
        </el-col>
    </el-row>
</template>

<script>
    import LyBarEchart from "./barEchart.vue";
    import LyLineEchart from "./lineEchart.vue";
    export default {
        name: "LyEchartcard",
        components: {LyLineEchart, LyBarEchart},
        data(){
            return{
                activeName:"tab1",
                dataList:"",
            }
        },
        created() {
            this.dataList = this.modelValue
        },
        props:{
            loading: {
                type: Boolean,
                default: true
            },
            count:{
                type:Number,
                default:1,
            },
            rows:{
                type:Number,
                default:4,
            },
            animated:{
                type:Boolean,
                default:true,
            },
            modelValue: {
              type: Array,
              default: []
            },
            height:{
                type:Number,
                default:300,
            }
        },
        watch:{
            modelValue: function(nval){
                this.dataList = nval;
            },
            dataList: function(nval) {
                this.$emit('update:modelValue', nval);
            },
        },
        methods:{
            handleTabChage(e){
            }
        },
    }
</script>

<style scoped>
    .space-inner{
    }
    .lycard{
        border-radius: var(--ly-radius-md);
        background: var(--ly-glass-bg);
        backdrop-filter: var(--ly-glass-blur);
        -webkit-backdrop-filter: var(--ly-glass-blur);
        border: 1px solid var(--ly-glass-border);
        box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card);
        overflow: hidden;
    }
    .lycard:hover{
        box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card-hover);
    }
    /* border-card tabs：表头玻璃化内嵌于玻璃卡（双类 (0,4,0) 稳压全局 seg 化规则与 EP border-card 实底） */
    .lycard.el-tabs--border-card > :deep(.el-tabs__header){
        background: transparent !important;
        border-bottom: none !important;
        padding: 10px 12px 0 12px !important;
        width: auto !important;
        border-radius: 0;
    }
    .lycard.el-tabs--border-card > :deep(.el-tabs__content){
        background: transparent;
    }
</style>