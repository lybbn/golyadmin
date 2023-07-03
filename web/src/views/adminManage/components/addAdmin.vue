<template>
    <div>
        <ly-dialog v-model="dialogVisible" :title="loadingTitle" width="640px" :before-close="handleClose">
            <el-form :inline="false" :model="formData" :rules="rules" ref="rulesForm" label-position="right" label-width="auto">
                <!--<el-form-item label="管理员编号：" prop="id">-->
                    <!--<el-input v-model.trim="formData.id" style="width: 300px"></el-input>-->
                <!--</el-form-item>-->
                <el-form-item label="管理员姓名：" prop="name">
                    <el-input v-model="formData.name"></el-input>
                </el-form-item>
                <el-form-item label="登陆账号：" prop="username">
                    <el-input v-model.trim="formData.username"></el-input>
                </el-form-item>
                <el-form-item label="登录密码：" prop="password">
                    <el-input v-model.trim="formData.password" clearable show-password></el-input>
                </el-form-item>
                <!--<el-form-item label="排序：" prop="sort">-->
                    <!--<el-input-number v-model="formData.sort" :min="1" :max="999999"></el-input-number>-->
                <!--</el-form-item>-->
<!--                <el-form-item label="角色：" prop="role">-->
<!--                    <el-checkbox-group v-model="formData.role">-->
<!--                        <el-checkbox :label="item.id" v-for="(item,index) in rolelist" :key="index">{{item.name}}</el-checkbox>-->
<!--                    </el-checkbox-group>-->
<!--                </el-form-item>-->
                <el-form-item label="角色：" prop="roleIds">
                    <el-select v-model="formData.roleIds" multiple filterable clearable placeholder="请选择" style="width:100%">
                        <el-option
                            v-for="item in rolelist"
                            :key="item.id"
                            :label="item.name"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="部门：" prop="dept_id">
                    <el-tree-select v-model="formData.dept_id" node-key="id" :data="options"
                            check-strictly filterable clearable :render-after-expand="false"
                            :props="{label:'name',value: 'id'}"
                            style="width: 100%" placeholder="请选择" />
                </el-form-item>
                <el-form-item label="状态：" prop="is_active">
                    <el-switch
                        v-model="formData.is_active"
                        active-color="#13ce66"
                        inactive-color="#ff4949">
                    </el-switch>
                </el-form-item>
            </el-form>
            <template #footer>
                <el-button @click="handleClose" :loading="loadingSave">取消</el-button>
                <el-button type="primary" @click="submitData" :loading="loadingSave">确定</el-button>
            </template>
        </ly-dialog>
    </div>
</template>

<script>
    import {apiSystemUserAdd,apiSystemUserEdit,apiSystemRole,apiSystemDept} from "@/api/api";
    import LyDialog from "@/components/dialog/dialog";
    import XEUtils from "xe-utils";
    export default {
        components: {LyDialog},
        emits: ['refreshData'],
        name: "addAdmin",
        data() {
            return {
                dialogVisible:false,
                loadingSave:false,
                loadingTitle:'',
                formData:{
                    name:'',
                    username:'',
                    password:'123456',
                    dept_id:'',
                    roleIds:[],
                    role:[],
                    is_active:true
                },
                rules:{
                    name: [
                        {required: true, message: '请输入管理员名称',trigger: 'blur'}
                    ],
                    role: [
                        {required: true, message: '请选择角色',trigger: 'blur'}
                    ],
                    // dept: [
                    //     {required: true, message: '请选择部门',trigger: 'blur'}
                    // ],
                    username: [
                        {required: true, message: '请输入管理员用户名',trigger: 'blur'}
                    ],
                    is_active: [
                        {required: true, message: '请选择是否启用',trigger: 'blur'}
                    ]
                },
                rolelist:[],
                options:[],
            }
        },
        methods:{
            handleClose() {
                this.dialogVisible=false
                this.loadingSave=false
                this.formData = {
                    name:'',
                    username:'',
                    password:'123456',
                    dept_id:'',
                    roleIds:[],
                    role:[],
                    is_active:true
                }
                this.$emit('refreshData')
            },
            getNewArray(arr,){},
            addAdminFn(item,flag) {
                this.getapiSystemRole()
                this.getapiSystemDept()
                this.loadingTitle=flag
                this.dialogVisible=true
                // console.log(item,'item----')
                // if(item && item.dept) {
                //     item.dept = item.dept.split(" ")
                // }
                this.formData=item ? item : {
                    name:'',
                    username:'',
                    password:'123456',
                    dept_id:'',
                    roleIds:[],
                    role:[],
                    is_active:true
                }
            },
            submitData() {
                this.$refs['rulesForm'].validate(obj=>{
                    if(obj) {
                        this.loadingSave=true
                        let param = {
                            ...this.formData
                        }
                        if(param.role){
                            let rolearr = []
                            for(var r in param.role){
                                rolearr.push(param.role[r])
                            }
                            param.role = rolearr
                        }else{
                            param.role = []
                        }
                        if(param.dept){
                            if(typeof  param.dept == 'object') {
                                param.dept=param.dept[param.dept.length-1]
                            }
                        }else{
                            param.dept = ''
                        }

                        if(this.formData.nickname=="" || this.formData.nickname== undefined || this.formData.nickname.length<=0 || this.formData.nickname=='""'){
                            param.nickname = this.formData.name
                        }
                        if(this.formData.id){
                            apiSystemUserEdit(param).then(res=>{
                                this.loadingSave=false
                                if(res.code ==2000) {
                                    this.$message.success(res.msg)
                                    this.handleClose()
                                    this.$emit('refreshData')
                                } else {
                                    this.$message.warning(res.msg)
                                }
                            })
                        }else{
                            apiSystemUserAdd(param).then(res=>{
                                this.loadingSave=false
                                if(res.code ==2000) {
                                    this.$message.success(res.msg)
                                    this.handleClose()
                                    this.$emit('refreshData')
                                } else {
                                    this.$message.warning(res.msg)
                                }
                            })
                        }

                    }
                })
            },
            getapiSystemRole(){
                apiSystemRole({page:1,limit:999}).then(res=>{
                    if(res.code ==2000) {
                        this.rolelist = res.data.data
                    } else {
                        this.$message.warning(res.msg)
                    }
                })
            },
            loadChild(treeNode, resolve){
                if(treeNode.level == 0){
                    return resolve([])
                }
                var params = {
                    lazy:true,
                    parent:treeNode.data.id,
                    page:1,
                    limit:999
                }
                apiSystemDept(params).then(async res => {
                     if(res.code == 2000) {
                         resolve(res.data)
                     }else {
                         this.$message.warning(res.msg)
                     }
                })
            },
            getapiSystemDept(){
                apiSystemDept({page:1,limit:999}).then(res=>{
                    if(res.code ==2000) {
                        this.options = XEUtils.toArrayTree(res.data, { parentKey: 'parent_id' })
                    } else {
                        this.$message.warning(res.msg)
                    }
                })
            },


        }
    }
</script>

