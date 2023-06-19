<!--
/**
 * description：icon选择器
 * author: lybbn
 * program django-vue-lyadmin
 * email: 1042594286@qq.com
 * website: http://www.lybbn.cn
 * date: 2022.11.17
 * @version: 1.0
 * remark: 如果要分发django-vue-lyadmin源码或其中组件等，需在本文件顶部保留此文件头信息！！！
 */
-->
<template>
    <div>
        <el-button  @click="handleClick">
            <template #icon v-if="!!iconText">
                <svg-icon :icon-class="iconText" style="font-size: 18px;"></svg-icon>
            </template>
            {{iconText?iconText:"请选择图标"}}
        </el-button>
        <ly-dialog  v-model="dialogVisible" :title="dialogTitle" width="50%" :before-close="handleClose">
            <div style="display: flex">
                <el-input  size="large" v-model="searchIconText" clearable class="searchinput" placeholder="搜索如：avatar" prefix-icon="Search"></el-input>
                <el-button size="large" icon="Delete" @click="deleteAll" type="danger">清除</el-button>
            </div>
            <el-tabs>
                <el-tab-pane  v-for="item in iconDataList" :key="item.name" lazy>
                    <template #label>
                        {{item.name}} <el-tag size="small" type="info">{{item.icons.length}}</el-tag>
                    </template>
                    <el-scrollbar height="600px">
                        <div class="icons-container">
                            <div v-for="(v, i) in item.icons" :key="i" class="lyicon" @click="chooseIcon(v)">
                                <div class="lyicon2">
                                    <svg-icon :icon-class="v" style="font-size: 20px;margin-top:10px;"></svg-icon>
                                    <span class="icon-text"> {{ v }}</span>
                                </div>
                            </div>
                        </div>
                    </el-scrollbar>
                </el-tab-pane>
            </el-tabs>
        </ly-dialog>
    </div>
</template>

<script>

    import * as Icons from '@element-plus/icons-vue'
    import LyDialog from "@/components/dialog/dialog";
    import {getIconList} from "@/icons/icon-list"
    import {deepClone} from "@/utils/util";

    export default {
        name: "LYChooseIcons",
        components: {LyDialog},
        props: {
            modelValue: {
                type: String,
                default: ""
            },
        },
        data(){
            return{
                dialogTitle:"Icon图标选择器",
                dialogVisible:false,
                oldIcons:[],
                iconText:"",
                searchIconText:"",
                iconDataList:[],
            }
        },
        created() {
            this.iconText = this.modelValue
            let extendIcons = getIconList()
            let eleIcons = Object.keys(Icons)
            this.iconDataList = [
                {
                    name:"默认",
                    icons:eleIcons,
                },
                {
                    name:"扩展",
                    icons:extendIcons,
                },
            ]
            this.oldIcons = deepClone(this.iconDataList)
        },
        watch:{
            modelValue: function(nval){
                this.iconText = nval; // modelValue改变是同步子组件visible的值
            },
            iconText: function(nval) {
                this.$emit('update:modelValue', nval); // visible改变时同步父组件modelValue的值
            },
            searchIconText(val){
				this.handleChange(val)
			}
        },
        methods:{
            handleClick(){
                this.dialogVisible = true
            },
            handleClose(){
                this.dialogVisible = false
            },
            chooseIcon(v){
                this.iconText = v
                this.handleClose()
            },
            handleChange(val){
                if(val){
                    let filterData = this.iconDataList
					filterData.forEach(t => {
						t.icons = t.icons.filter(item => item.toLowerCase().indexOf(val.toLowerCase()) === 0)
					})
                    this.iconDataList = filterData
                }else{
                    this.iconDataList = deepClone(this.oldIcons)
                }

            },
            deleteAll(){
                this.iconText=''
                this.searchIconText = ""
                this.iconDataList = deepClone(this.oldIcons)
            },

        },

    }
</script>

<style lang="scss" scoped>
    .searchinput{
       padding-bottom: 20px;
    }
    .icons-container {
      display: flex;
      flex-wrap: wrap;
      .lyicon {
        /*border-right:  1px solid var(--el-border-color);*/
        /*border-bottom:  1px solid var(--el-border-color);*/
        width: 20%;
        height: 70px;
        display: flex;
        flex-direction: column;
        align-items: center;
        box-sizing: border-box;
        cursor: pointer;
        .lyicon2{
            display: flex;
            align-items: center;
            flex-direction: column;
            height: 60px;
            width: 85%;
            border:  1px solid var(--el-border-color);
        }
        .inicon{
          margin-top:5px;
        }
        .icon-text {
          flex: 1;
          font-size: 12px;
        }
        &:hover {
          color: #409efa;
        }
      }
      svg {
        width: 1.5em;
        height: 1.5em;
      }
    }
</style>