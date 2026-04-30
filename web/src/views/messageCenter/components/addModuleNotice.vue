<template>
    <div>
        <ly-dialog v-model="dialogVisible" :title="loadingTitle" width="50%" :before-close="handleClose">
            <el-form :inline="false" :model="formData" :rules="rules" ref="rulesForm" label-position="right" label-width="auto">
                <el-form-item label="公告标题：" prop="msg_title">
                    <el-input type="text" v-model.trim="formData.msg_title"></el-input>
                </el-form-item>
                <el-form-item label="跳转路径：" prop="to_path">
                    <el-input type="text" v-model.trim="formData.to_path"></el-input>
                </el-form-item>
                <el-form-item label="目标类型：" prop="target_type">
                    <el-radio-group v-model="formData.target_type">
                        <el-radio :label="1"  border>平台公告</el-radio>
                        <el-radio :label="2"  border>按用户</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="发送对象：" prop="target_user" v-if="formData.target_type == 2" class="is-required">
                    <ly-table-select v-model="formData.target_user" :apiObj="getUserList" :table-width="800" multiple clearable collapse-tags collapse-tags-tooltip :props="tableSelectProps" @change="selectChange">
                        <template #header="{form, submit}">
                            <el-form :inline="true" :model="form">
                                <el-form-item>
                                    <el-input type="text" v-model="form.username" placeholder="请输入用户名"></el-input>
                                </el-form-item>
                                <el-form-item>
                                    <el-button type="primary" @click="submit">查询</el-button>
                                </el-form-item>
                            </el-form>
                        </template>
                        <el-table-column prop="id" label="ID" width="100" show-overflow-tooltip></el-table-column>
                        <el-table-column prop="username" label="用户名" width="100"></el-table-column>
                        <el-table-column prop="nickname" label="昵称" width="100"></el-table-column>
                        <el-table-column prop="mobile" label="手机号" width="150"></el-table-column>
                        <el-table-column prop="create_datetime" label="注册时间"></el-table-column>
                    </ly-table-select>
                </el-form-item>
                <el-form-item label="公告内容：" prop="msg_content">
                    <TEditor v-model="formData.msg_content" ></TEditor>
                </el-form-item>
<!--                <el-form-item label="是否发布：" prop="status">-->
<!--                    <el-switch-->
<!--                            v-model="formData.status"-->
<!--                            active-color="#13ce66"-->
<!--                            inactive-color="#ff4949">-->
<!--                    </el-switch>-->
<!--                </el-form-item>-->
            </el-form>
            <template #footer>
                <el-button @click="handleClose" :loading="loadingSave">取消</el-button>
                <el-button type="primary" @click="submitData" :loading="loadingSave">确定</el-button>
            </template>
        </ly-dialog>
    </div>
</template>

<script>
    import {messagesMessagenoticeAdd,messagesMessagenoticeEdit,UsersUsers} from "@/api/api";
    import TEditor from '@/components/TEditor.vue'
    import LyDialog from "@/components/dialog/dialog.vue";
    import {deepClone} from "@/utils/util";
    import LyTableSelect from "@/components/lyTableSelect.vue";
    export default {
        components: {LyDialog, TEditor,LyTableSelect},
        emits: ['refreshData'],
        name: "addModuleNotice",
        data() {
            return {
                dialogVisible:false,
                loadingSave:false,
                loadingTitle:'',
                formData:{
                    msg_title:'',
                    to_path:'',
                    msg_content:'',
                    target_type:1,
                    target_user:[],
                    status:true
                },
                rules:{
                    msg_title: [
                        {required: true, message: '请填写公告标题',trigger: 'blur'}
                    ],
                    msg_content: [
                        {required: true, message: '请填写公告内容',trigger: 'blur'}
                    ]
                },
                //指定表格选择器的回显映射
                tableSelectProps: {
					label: 'username',
					value: 'id',
				}
            }
        },
        mounted() {
            window.addEventListener("focusin", this.onFocusIn,true);
        },
        unmounted() {
            window.removeEventListener("focusin", this.onFocusIn);
        },
        methods:{
            onFocusIn(e){
                e.stopImmediatePropagation()//阻止当前和后面的一系列事件
            },
            getUserList(){
                return UsersUsers
            },
            //值变化
			selectChange(val){
				//change事件，返回详情查看控制台
				console.log(val)
			},
            handleClose() {
                this.dialogVisible=false
                this.loadingSave=false
                this.formData = {
                    msg_title:'',
                    to_path:'',
                    msg_content:'',
                    target_type:1,
                    target_user:[],
                    status:true
                }
                this.$emit('refreshData')
            },
            addModuleFn(item,flag) {
                this.loadingTitle=flag
                this.dialogVisible=true
                if(item){
                    this.formData=deepClone(item)
                }
            },
            submitData() {
                this.$refs['rulesForm'].validate(obj=>{
                    if(obj) {
                        this.loadingSave=true
                        let param = {
                            ...this.formData
                        }
                        //处理表单组件返回的选中数据(去除杂数据)为id数组['1','2']
                        if(param.target_user.length>0){
                            param.target_user = param.target_user.map(item =>{return item.id} )
                        }
                        if(this.formData.id){
                            messagesMessagenoticeEdit(param).then(res=>{
                                this.loadingSave=false
                                if(res.code ==2000) {
                                    this.$message.success(res.msg)
                                    this.handleClose()
                                } else {
                                    this.$message.warning(res.msg)
                                }
                            })
                        }else{
                            messagesMessagenoticeAdd(param).then(res=>{
                                this.loadingSave=false
                                if(res.code ==2000) {
                                    this.$message.success(res.msg)
                                    this.handleClose()
                                } else {
                                    this.$message.warning(res.msg)
                                }
                            })
                        }

                    }
                })
            },

        }
    }
</script>

<style>
    .set-specs .el-form-item__content{
        background: #e6e6e6 !important;
    }
</style>

