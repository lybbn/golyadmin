<template>
  <el-card shadow="hover" header="个人中心">
    <div class="lyadmin-body">
      <el-tabs v-model="activeName" @tab-click="handleClick" tab-position="left">
        <el-tab-pane label="用户设置" name="userInfo" >
            <el-alert title="修改用户信息，无特殊需求可以不修改。" type="info" show-icon style="margin-bottom: 15px;"/>
            <el-form ref="userInfoForm" :model="userInfo" :disabled="!isShowBtn('personalCenter','个人中心','Update')"  required-asterisk :rules="userInforules" :label-position="position" center label-width="120px" style="margin-top: 20px;">
              <el-form-item prop="name" required label="姓名:">
                <el-input v-model="userInfo.name" clearable></el-input>
              </el-form-item>
              <el-form-item label="电话号码:" prop="mobile">
                <el-input v-model="userInfo.mobile" clearable></el-input>
              </el-form-item>
              <el-form-item label="邮箱:" prop="email">
                <el-input v-model="userInfo.email" clearable></el-input>
              </el-form-item>
              <el-form-item label="性別:" prop="gender">
                <el-radio-group v-model="userInfo.gender">
                  <el-radio :label="1">男</el-radio>
                  <el-radio :label="0">女</el-radio>
                  <el-radio :label="-1">未知</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item>
                <el-button @click="updateInfo" icon="Check" type="primary" v-show="isShowBtn('personalCenter','个人中心','Update')">
                  更新
                </el-button>
                <el-button icon="Refresh" @click="resetForm('info')" type="info" v-show="isShowBtn('personalCenter','个人中心','Update')">
                  重置
                </el-button>
              </el-form-item>
            </el-form>
        </el-tab-pane>
        <el-tab-pane label="密码设置" name="passwrod" >
            <el-alert title="密码更新成功后，您需要重新使用新密码登录页面。" type="info" show-icon style="margin-bottom: 15px;"/>
            <el-form ref="userPasswordForm" :model="userPasswordInfo" :rules="rules" label-width="120px" style="margin-top:20px;">
              <el-form-item label="当前密码" prop="oldPassword">
                <el-input v-model="userPasswordInfo.oldPassword" type="password" show-password placeholder="请输入当前密码"></el-input>
                <div class="el-form-item-msg">必须提供当前登录用户密码才能进行更改</div>
              </el-form-item>
              <el-form-item label="新密码" prop="newPassword">
                <el-input v-model="userPasswordInfo.newPassword" type="password" show-password placeholder="请输入新密码"></el-input>
                <lyPasswordStrength v-model="userPasswordInfo.newPassword"></lyPasswordStrength>
                <div class="el-form-item-msg">请输入包含英文、数字的8位以上密码</div>
              </el-form-item>
              <el-form-item label="确认新密码" prop="newPassword2">
                <el-input v-model="userPasswordInfo.newPassword2" type="password" show-password placeholder="请再次输入新密码"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button  type="primary" icon="Check" @click="settingPassword" v-show="isShowBtn('personalCenter','个人中心','Changepassword')">
                  提交
                </el-button>
                <el-button icon="Refresh" @click="resetForm('passwordForm')" type="info" v-show="isShowBtn('personalCenter','个人中心','Changepassword')">
                  重置
                </el-button>
              </el-form-item>
            </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </el-card>
</template>

<script>
    import {systemUserUserInfoEdit,systemUserUserInfo,systemUserChangePassword} from '@/api/api'
    import {useMutitabsStore} from "@/store/mutitabs";
    import lyPasswordStrength from "@/components/password/lyPasswordStrength";
    export default {
        components:{lyPasswordStrength},
        name: "personalCenter",
        setup(){
            const mutitabsstore = useMutitabsStore()
            return { mutitabsstore}
        },
        data() {
            var validatePass = (rule, value, callback) => {
              const pwdRegex = new RegExp('(?=.*[0-9])(?=.*[a-zA-Z]).{8,30}')
              if (value === '') {
                callback(new Error('请输入密码'))
              } else if (value === this.userPasswordInfo.oldPassword) {
                callback(new Error('原密码与新密码一致'))
              } else if (!pwdRegex.test(value)) {
                callback(new Error('您的密码复杂度太低(密码中必须包含字母、数字)'))
              } else {
                callback()
              }
            }
            return{
                position: 'right',
                activeName: 'userInfo',
                userInfo: {
                  name: '',
                  gender: 1,
                  mobile: '',
                  avatar: '',
                  email: ''
                },
                userInforules: {
                  name: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
                  mobile: [
                    { pattern: /^1[3|4|5|7|8|9|6]\d{9}$/, message: '请输入正确手机号' }
                  ]
                },
                userPasswordInfo: {
                  oldPassword: '',
                  newPassword: '',
                  newPassword2: ''
                },
                rules: {
                  oldPassword: [
                    { required: true, message: '请输入当前密码'}
                  ],
                  newPassword: [
                    { required: true,validator: validatePass}
                  ],
                  newPassword2: [
                    { required: true, message: '请再次输入新密码'},
                    {validator: (rule, value, callback) => {
                      if (value !== this.userPasswordInfo.newPassword) {
                        callback(new Error('两次输入密码不一致'));
                      }else{
                        callback();
                      }
                    }}
                  ]
                }
              }
        },
        mounted () {
        this.getCurrentUserInfo()
        },
        methods:{
            /**
             * 获取当前用户信息
             */
            getCurrentUserInfo () {
                systemUserUserInfo().then(res=>{
                    if(res.code == 2000) {
                        this.userInfo=res.data.data
                    }

                })
            },
            /**
             * 更新用户信息
             */
            updateInfo () {
              const _self = this

              _self.$refs.userInfoForm.validate((valid) => {
                if (valid) {
                    //console.log(_self.userInfo)
                    systemUserUserInfoEdit(_self.userInfo).then(res=>{
                            if(res.code ==2000) {
                                this.$message.success(res.msg)
                                _self.getCurrentUserInfo()
                            } else {
                                this.$message.warning(res.msg)
                            }
                        })
                } else {
                  // 校验失败
                  // 登录表单校验失败
                  this.$message.error('表单校验失败，请检查')
                }
              })
            },
            // 重置
            resetForm (name) {
              const _self = this
              if (name === 'info') {
                _self.getCurrentUserInfo()
              } else {
                _self.userPasswordForm = {}
              }
            },
            // tab切换,默认切换清除原字符
            handleClick (tab, event) {
              const _self = this
              // if (tab.paneName === 'userInfo') {
              //   _self.$refs.userPasswordForm.resetFields()
              // } else {
              //   _self.$refs.userInfoForm.resetFields()
              // }
            },
            /**
             * 重新设置密码
             */
            settingPassword () {
              const _self = this
              _self.$refs.userPasswordForm.validate((valid) => {
                if (valid) {
                  const userId = this.mutitabsstore.getUserId
                  if (userId) {
                    const params = JSON.parse(JSON.stringify(_self.userPasswordInfo))
                      params.id = userId
                    systemUserChangePassword(params).then(res=>{
                        if(res.code ==2000) {
                            _self.activeName = 'userInfo'
                            this.$message.success(res.msg)
                        } else {
                            this.$message.warning(res.msg)
                        }
                    })
                  }
                } else {
                  // 校验失败
                  return false
                }
              })
            }

        }
    }
</script>

<style scoped>
  .el-tabs .el-tabs__content{
    background-color: var(--el-bg-color) !important;
  }
  .el-tabs{
    background-color: var(--el-bg-color) !important;
  }
  .el-tabs{
    padding: 20px;
  }
  /*::v-deep(.el-tabs__header){*/
  /*  margin-top: 20px !important;*/
  /*}*/
  .el-form-item-msg{
    font-size: 12px;
    color: #999;
    clear: both;
    width: 100%;
  }
</style>
