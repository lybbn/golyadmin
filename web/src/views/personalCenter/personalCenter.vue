<template>
  <div class="personal-center">
    <header class="page-heading">
      <h1>个人中心</h1>
      <p>管理个人资料与账户安全</p>
    </header>

    <section class="settings-card">
      <div class="account-overview">
        <el-avatar class="profile-avatar" :size="60" :src="userInfo.avatar">
          {{ profileInitial }}
        </el-avatar>
        <div class="overview-copy">
          <div class="overview-title-row">
            <h2>{{ profileName }}</h2>
            <el-tag class="account-tag" effect="light" round>{{ identityLabel }}</el-tag>
          </div>
        </div>
      </div>

        <el-tabs v-model="activeName" class="settings-tabs">
          <el-tab-pane name="userInfo">
            <template #label>
              <span class="tab-label"><el-icon><User /></el-icon>个人资料</span>
            </template>
            <div class="section-heading">
              <p>更新您的联系信息，方便账户识别与日常沟通。</p>
            </div>

            <el-form
              ref="userInfoForm"
              :model="userInfo"
              :disabled="!hasPermission(this.$route.name, 'Update')"
              :rules="userInforules"
              label-position="top"
              class="account-form profile-form"
            >
              <div class="field-grid">
                <el-form-item prop="name" required label="姓名">
                  <el-input v-model="userInfo.name" clearable placeholder="请输入姓名" />
                </el-form-item>
                <el-form-item prop="nickname" label="昵称">
                  <el-input v-model="userInfo.nickname" clearable placeholder="请输入昵称" />
                </el-form-item>
                <el-form-item prop="mobile" label="电话号码">
                  <el-input v-model="userInfo.mobile" clearable placeholder="请输入电话号码" />
                </el-form-item>
                <el-form-item prop="email" label="邮箱">
                  <el-input v-model="userInfo.email" clearable placeholder="请输入邮箱" />
                </el-form-item>
                <el-form-item prop="gender" label="性别" class="gender-item">
                  <el-radio-group v-model="userInfo.gender" class="gender-options">
                    <el-radio value="男">男</el-radio>
                    <el-radio value="女">女</el-radio>
                    <el-radio value="未知">未知</el-radio>
                  </el-radio-group>
                </el-form-item>
              </div>
              <div class="form-actions" v-show="hasPermission(this.$route.name, 'Update')">
                <el-button @click="resetForm('info')" plain>恢复原值</el-button>
                <el-button @click="updateInfo" type="primary" icon="Check">保存资料</el-button>
              </div>
            </el-form>
          </el-tab-pane>

          <el-tab-pane name="password">
            <template #label>
              <span class="tab-label"><el-icon><Lock /></el-icon>密码与安全</span>
            </template>
            <div class="section-heading">
              <p>定期更新密码有助于保护您的账户安全。</p>
            </div>

            <div class="security-note">
              <span class="note-icon"><el-icon><Lock /></el-icon></span>
              <p>修改成功后，请使用新密码重新登录。</p>
            </div>

            <el-form ref="userPasswordForm" :model="userPasswordInfo" :rules="rules" label-position="top" class="account-form password-form">
              <el-form-item label="当前密码" prop="oldPassword">
                <el-input v-model="userPasswordInfo.oldPassword" type="password" show-password placeholder="请输入当前密码" />
                <div class="field-hint">验证当前密码后，才能设置新密码。</div>
              </el-form-item>
              <el-form-item label="新密码" prop="newPassword">
                <el-input v-model="userPasswordInfo.newPassword" type="password" show-password placeholder="请输入新密码" />
                <lyPasswordStrength v-model="userPasswordInfo.newPassword" />
                <div class="field-hint">密码需包含英文字母和数字，长度为 8–30 位。</div>
              </el-form-item>
              <el-form-item label="确认新密码" prop="newPassword2">
                <el-input v-model="userPasswordInfo.newPassword2" type="password" show-password placeholder="请再次输入新密码" />
              </el-form-item>
              <div class="form-actions" v-show="hasPermission(this.$route.name, 'Changepassword')">
                <el-button @click="resetPasswordForm" plain>清空</el-button>
                <el-button type="primary" icon="Check" @click="settingPassword">更新密码</el-button>
              </div>
            </el-form>
          </el-tab-pane>
        </el-tabs>
    </section>
  </div>
</template>

<script>
    import {systemUserUserInfoEdit,systemUserUserInfo,systemUserChangePassword} from '@/api/api'
    import {useMutitabsStore} from "@/store/mutitabs";
    import lyPasswordStrength from "@/components/password/lyPasswordStrength.vue";
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
                activeName: 'userInfo',
                userInfo: {
                  name: '',
                  gender: '男',
                  mobile: '',
                  avatar: '',
                  email: ''
                },
                userInforules: {
                  name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
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
        computed: {
            profileName () {
                return this.userInfo.nickname || this.userInfo.name || this.mutitabsstore.getUserName || '用户'
            },
            profileInitial () {
                return Array.from(String(this.profileName).trim())[0] || '用'
            },
            identityLabel () {
                const identity = Number(this.mutitabsstore.getIdentity)
                if (identity === 1) return '超级管理员'
                if (identity === 2) return '后台用户'
                if (identity === 3) return '前台用户'
                return '用户'
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
                        this.userInfo=res.data
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
              if (name === 'info') {
                this.getCurrentUserInfo()
              }
            },
            resetPasswordForm () {
              this.$refs.userPasswordForm?.resetFields()
              this.userPasswordInfo = {
                oldPassword: '',
                newPassword: '',
                newPassword2: ''
              }
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
                            _self.resetPasswordForm()
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
.personal-center {
  width: 100%;
  max-width: none;
  min-height: calc(100vh - 130px);
  padding: 4px 6px 8px;
  box-sizing: border-box;
  color: var(--ly-text-1);
}

.page-heading {
  margin-bottom: 10px;
}

.settings-card {
  width: 100%;
  max-width: none;
  min-height: calc(100vh - 240px);
  margin: 0;
  box-sizing: border-box;
}

.page-heading h1 {
  margin: 0;
  color: var(--ly-text-1);
  font-size: 22px;
  font-weight: 650;
  line-height: 1.4;
}

.page-heading p {
  margin: 4px 0 0;
  color: var(--ly-text-2);
  font-size: 13px;
  line-height: 1.5;
}

.settings-card {
  padding: 0 34px 30px;
  overflow: hidden;
  border: 1px solid var(--ly-glass-border);
  border-radius: 16px;
  background: var(--ly-glass-bg-strong);
  box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card);
  backdrop-filter: var(--ly-glass-blur);
  -webkit-backdrop-filter: var(--ly-glass-blur);
}

.account-overview {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 24px 0 20px;
  border-bottom: 1px solid var(--ly-line-soft);
}

:deep(.profile-avatar.el-avatar) {
  flex: 0 0 auto;
  border: 3px solid rgba(58, 123, 255, 0.14);
  color: #fff;
  background: var(--ly-gradient-primary);
  box-shadow: 0 4px 12px rgba(35, 79, 184, 0.14);
  font-size: 22px;
  font-weight: 600;
}

.overview-copy {
  min-width: 0;
}

.overview-title-row {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 9px;
}

.overview-title-row h2 {
  margin: 0;
  color: var(--ly-text-1);
  font-size: 18px;
  font-weight: 650;
  line-height: 1.45;
  overflow-wrap: anywhere;
}

:deep(.account-tag.el-tag) {
  border-color: rgba(58, 123, 255, 0.16);
  color: var(--el-color-primary);
  background: rgba(58, 123, 255, 0.08);
}

.settings-tabs {
  background: transparent;
}

:deep(.settings-tabs .el-tabs__content) {
  overflow: visible;
  background: transparent !important;
}

:deep(.settings-tabs .el-tabs__header) {
  margin: 12px 0 16px;
}

:deep(.settings-tabs .el-tabs__nav-wrap::after) {
  display: none;
}

:deep(.settings-tabs .el-tabs__nav) {
  display: inline-flex;
  gap: 3px;
  padding: 3px;
  border-radius: 10px;
  background: var(--ly-glass-bg-soft);
}

:deep(.settings-tabs .el-tabs__item) {
  height: 36px;
  padding: 0 14px;
  border-radius: 7px;
  color: var(--ly-text-2);
  font-size: 13px;
  line-height: 36px;
  transition: color 0.18s ease, background-color 0.18s ease;
}

:deep(.settings-tabs .el-tabs__item.is-active) {
  color: var(--el-color-primary);
  font-weight: 600;
  background: var(--ly-glass-bg-strong);
  box-shadow: 0 1px 3px rgba(27, 35, 64, 0.08);
}

:deep(.settings-tabs .el-tabs__active-bar) {
  display: none;
}

.tab-label {
  display: inline-flex;
  align-items: center;
  gap: 7px;
}

.tab-label .el-icon {
  font-size: 14px;
}

.section-heading {
  margin-bottom: 16px;
}

.section-heading p {
  margin: 0;
  color: var(--ly-text-2);
  font-size: 12px;
  line-height: 1.6;
}

.field-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  column-gap: 24px;
}

:deep(.account-form .el-form-item) {
  min-width: 0;
  margin-bottom: 16px;
}

:deep(.account-form .el-form-item__label) {
  padding-bottom: 6px;
  color: var(--ly-text-2);
  font-size: 12px;
  font-weight: 500;
  line-height: 1.4;
}

:deep(.account-form .el-input__wrapper) {
  min-height: 40px;
  padding: 0 11px;
  border-radius: 8px;
  background: var(--ly-glass-bg-strong);
  box-shadow: 0 0 0 1px var(--ly-line-soft) inset;
}

:deep(.account-form .el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px color-mix(in srgb, var(--el-color-primary) 38%, transparent) inset;
}

:deep(.account-form .el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px var(--el-color-primary) inset, var(--ly-input-focus-ring);
}

:deep(.account-form .el-input__inner) {
  color: var(--ly-text-1);
  font-size: 13px;
}

:deep(.account-form .el-input__inner::placeholder) {
  color: var(--ly-text-3);
  font-size: 12px;
}

:deep(.account-form .el-form-item.is-error .el-input__wrapper) {
  box-shadow: 0 0 0 1px var(--el-color-danger) inset;
}

:deep(.gender-item.el-form-item) {
  grid-column: 1 / -1;
  margin-bottom: 1px;
}

:deep(.gender-item .el-form-item__content) {
  min-height: 30px;
  align-items: center;
}

.gender-options {
  display: flex;
  flex-wrap: wrap;
  gap: 18px;
}

:deep(.gender-options .el-radio) {
  margin-right: 0;
  color: var(--ly-text-2);
  font-size: 12px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 9px;
  margin-top: 13px;
  padding-top: 16px;
  border-top: 1px solid var(--ly-line-soft);
}

:deep(.form-actions .el-button) {
  min-height: 36px;
  border-radius: 8px;
  font-size: 12px;
}

:deep(.form-actions .el-button--primary) {
  border: 0;
  background: var(--ly-gradient-primary);
  box-shadow: 0 4px 10px rgba(58, 123, 255, 0.18);
}

.security-note {
  display: flex;
  max-width: 680px;
  align-items: center;
  gap: 10px;
  margin-bottom: 18px;
  padding: 11px 13px;
  border: 1px solid rgba(58, 123, 255, 0.10);
  border-radius: 9px;
  background: rgba(58, 123, 255, 0.045);
}

.note-icon {
  display: grid;
  width: 28px;
  height: 28px;
  flex: 0 0 auto;
  place-items: center;
  border-radius: 8px;
  color: var(--el-color-primary);
  background: var(--ly-glass-bg-strong);
  font-size: 13px;
}

.security-note p {
  margin: 0;
  color: var(--ly-text-2);
  font-size: 12px;
  line-height: 1.5;
}

.password-form {
  max-width: 680px;
}

:deep(.password-form .el-form-item) {
  margin-bottom: 16px;
}

:deep(.password-form .ly-password-strength) {
  width: 100%;
}

.field-hint {
  width: 100%;
  margin-top: 5px;
  color: var(--ly-text-3);
  font-size: 11px;
  line-height: 1.45;
}

.password-form .form-actions {
  margin-top: 18px;
}

@media (max-width: 720px) {
  .personal-center {
    min-height: auto;
    padding: 18px 14px 24px;
  }

  .settings-card {
    min-height: 0;
    padding-right: 22px;
    padding-left: 22px;
  }
}

@media (max-width: 560px) {
  .personal-center {
    padding-right: 8px;
    padding-left: 8px;
  }

  .page-heading h1 {
    font-size: 20px;
  }

  .settings-card {
    padding-right: 16px;
    padding-left: 16px;
  }

  .account-overview {
    gap: 12px;
    padding: 18px 0 16px;
  }

  :deep(.profile-avatar.el-avatar) {
    width: 50px;
    height: 50px;
    font-size: 18px;
  }

  :deep(.settings-tabs .el-tabs__item) {
    padding: 0 12px;
    font-size: 12px;
  }

  .field-grid {
    grid-template-columns: minmax(0, 1fr);
  }

  :deep(.gender-item.el-form-item) {
    grid-column: auto;
  }

  .form-actions {
    margin-top: 8px;
  }
}
</style>
