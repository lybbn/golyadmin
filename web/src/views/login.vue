<template>
    <div class="login-page">
      <div class="login-tools">
        <el-button
          :icon="siteThemeStore.siteTheme == 'dark' ? 'sunny' : 'moon'"
          circle
          plain
          class="tool-button"
          :aria-label="$t('login.toggleTheme')"
          @click="setSiteTheme"
        ></el-button>
        <el-dropdown trigger="click" placement="bottom-end" @command="changeLang">
          <el-button circle plain class="tool-button" :aria-label="$t('login.language')">
            <svg viewBox="0 0 24 24" width="18" height="18" aria-hidden="true">
              <path fill="currentColor" d="m18.5 10 4.4 11h-2.155l-1.201-3h-4.09l-1.199 3h-2.154L16.5 10h2zM10 2v2h6v2h-1.968a18.222 18.222 0 0 1-3.62 6.301 14.864 14.864 0 0 0 2.336 1.707l-.751 1.878A17.015 17.015 0 0 1 9 13.725a16.676 16.676 0 0 1-6.201 3.548l-.536-1.929a14.7 14.7 0 0 0 5.327-3.042A18.078 18.078 0 0 1 4.767 8h2.24A16.032 16.032 0 0 0 9 10.877a16.165 16.165 0 0 0 2.91-4.876L2 6V4h6V2h2zm7.5 10.885L16.253 16h2.492L17.5 12.885z"></path>
            </svg>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-for="item in lang" :key="item.value" :command="item" :class="{'lydpselected': language == item.value}">{{ item.name }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>

      <main class="login-main">
        <div class="login-shell">
          <section class="brand-panel" :aria-label="APPName">
            <div class="brand-orbit brand-orbit-large"></div>
            <div class="brand-orbit brand-orbit-small"></div>
            <div class="brand-grid"></div>

            <div class="brand-lockup">
              <div class="brand-logo">
                <img src="../assets/logo.png" :alt="APPName">
              </div>
              <span class="brand-name">{{ APPName }}</span>
            </div>

            <div class="brand-message">
              <span class="brand-kicker">GOLYADMIN</span>
              <h1>{{ $t('login.workspaceTitle') }}</h1>
              <p>{{ $t('login.workspaceSubtitle') }}</p>
            </div>

            <div class="brand-footer">
              <span>01</span>
              <span class="brand-footer-line"></span>
              <span>ADMIN CONSOLE</span>
            </div>
          </section>

          <section class="form-panel">
            <el-form
              label-position="top"
              :model="ruleForm"
              :rules="rules"
              ref="ruleForm"
              label-width="0px"
              class="login-form"
            >
              <div class="login-heading">
                <span class="login-kicker">{{ APPName }}</span>
                <h2>{{ $t('login.loginInTitle') }}</h2>
                <p>{{ $t('login.loginPrompt') }}</p>
              </div>

              <el-form-item prop="username">
                <label class="field-label" for="login-username">{{ $t('login.accountLabel') }}</label>
                <el-input
                  id="login-username"
                  v-model.trim="ruleForm.username"
                  type="text"
                  size="large"
                  autocomplete="username"
                  :placeholder="$t('login.loginAccount')"
                  maxlength="60"
                >
                  <template #prefix><el-icon><User /></el-icon></template>
                </el-input>
              </el-form-item>

              <el-form-item prop="password">
                <label class="field-label" for="login-password">{{ $t('login.passwordLabel') }}</label>
                <el-input
                  id="login-password"
                  v-model.trim="ruleForm.password"
                  type="password"
                  size="large"
                  autocomplete="current-password"
                  :placeholder="$t('login.loginPWD')"
                  maxlength="60"
                  @keyup.enter="submitForm('ruleForm')"
                >
                  <template #prefix><el-icon><Lock /></el-icon></template>
                </el-input>
              </el-form-item>

              <el-form-item prop="captcha">
                <label class="field-label" for="login-captcha">{{ $t('login.codeLabel') }}</label>
                <el-input
                  id="login-captcha"
                  v-model.trim="ruleForm.captcha"
                  type="text"
                  size="large"
                  class="captcha-input"
                  autocomplete="off"
                  :placeholder="$t('login.code')"
                  @keyup.enter="submitForm('ruleForm')"
                >
                  <template #prefix><el-icon><CircleCheck /></el-icon></template>
                  <template #append>
                    <button type="button" class="captcha-refresh" :aria-label="$t('login.refreshCaptcha')" @click="getCaptchas">
                      <img class="login-code" :src="image_base" :alt="$t('login.refreshCaptcha')">
                    </button>
                  </template>
                </el-input>
              </el-form-item>

              <el-checkbox class="remember" v-model="rememberpassword">{{ $t('login.rememberMe') }}</el-checkbox>
              <el-form-item class="submit-item">
                <el-button type="primary" size="large" :loading="loadingLg" class="login-submit" @click="submitForm('ruleForm')">
                  {{ $t('login.login') }}
                  <el-icon class="submit-icon"><Right /></el-icon>
                </el-button>
              </el-form-item>
            </el-form>
          </section>
        </div>
      </main>

      <footer class="login-copyright">
        Copyright © 2022 golyadmin All rights reserved.
      </footer>
    </div>
</template>
<script >
  import {login,apiSystemWebRouter,getCaptcha} from '@/api/api'
  import {delCookie, getCookie, setCookie, transArrayMenuToTree} from '@/utils/util'
  import {useMutitabsStore} from "@/store/mutitabs";
  import {useSiteThemeStore} from "@/store/siteTheme";
  import {setStorage} from '@/utils/util'
  import i18n from '@/locales'
  import config from "@/config"

  export default {
    name: 'login',
    setup(){
        const mutitabsstore = useMutitabsStore()
        const siteThemeStore = useSiteThemeStore()
        const { t } = i18n.global
        return { mutitabsstore,siteThemeStore,t}

    },
    data() {
      return {
        loadingLg:false,
        logining: false,
        rememberpassword: false,
        APPName:config.APP_NAME,
        ruleForm: {
            username: '',
            password: '',
            captcha:'',
            captchaKey: null,
        },
        loginFlag:false,
        rules: {
            username: [{required: true, message: this.t('login.AccountError'), trigger: 'blur'}],
            password: [{required: true, message: this.t('login.PWError'), trigger: 'blur'}],
            captcha: [{required: true, message: this.t('login.codeError'), trigger: 'blur'}],
        },
        image_base: null,
        allmenu:[],
        language:this.siteThemeStore.language,
        lang: [
            {
                name: '简体中文',
                value: 'zh-cn',
            },
            {
                name: 'English',
                value: 'en',
            }
        ],
      }
    },
      created() {
        //动态添加该页面meta viewport 手机适配
        if(document.querySelector("meta[name='viewport']")){
            document.querySelector("meta[name='viewport']")["content"] = "width=device-width,initial-scale=1.0,maximum-scale=1.0,minimum-scale=1.0,user-scalable=no"
        }
        //请求数据
        this.getuserpassword()
        this.getCaptchas()
      },
      beforeRouteLeave(to, form, next){
          //离开页面去除动态添加该页面meta viewport 手机适配
          document.querySelector("meta[name='viewport']")["content"] = this.getCurrentWith()
          next()
      },
      methods: {
        //设置主题
        setSiteTheme(){
            if(this.siteThemeStore.siteTheme=='light'){
                this.siteThemeStore.setSiteTheme('dark')
            }else{
                this.siteThemeStore.setSiteTheme('light')
            }
        },
        //设置语言
        changeLang(command){
            this.language = command.value
            this.siteThemeStore.setLanguage(command.value)
        },
        getCurrentWith(){
            var designWidth = 375;
            var deviceWidth = parseInt(window.screen.width) || parseInt(document.documentElement.clientWidth);  //获取当前设备的屏幕宽度
            // var deviceScale = deviceWidth/designWidth;
            var deviceScale = 0.6;
            var ua = navigator.userAgent;
            //获取当前设备类型（安卓或苹果）

            if (ua && /Android (\d+.\d+)/.test(ua)) {
                // +",user-scalable=no"
                return "width=680,initial-scale="+deviceScale+",minimum-scale="+deviceScale+",maximum-scale="+deviceScale;
            }
            else if (ua && /iPhone|ipad|ipod|ios/.test(ua)){
                return "width=680,initial-scale="+deviceScale+",minimum-scale="+deviceScale+",maximum-scale="+deviceScale;
            }
            else {
                return '';
            }
        },
      // 获取用户名密码
      // 获取菜单
      getMenu() {
        this.menuTitle=''
        this.allmenu=[]
        this.loadingLg=true
        apiSystemWebRouter().then(res=>{
          if(res.code == 2000) {
            let menuTree = []
            if(res.data.length > 0) {
              menuTree = transArrayMenuToTree(res.data)
              // 操作权限管控
              let menuList=[]
              res.data.forEach(item=>{
                menuList.push({
                  url:item.web_path,
                  moduleName:item.name,
                  menuPermission:item.menuPermission
                })
              })
              setStorage('menuList', JSON.stringify(menuList))
            }
            this.allmenu =  menuTree
            if(this.allmenu.length >0) {
              this.$nextTick(()=>{
                this.$router.replace({path: `/${this.allmenu[0].attributes.url}`})
              })
            } else {
               this.mutitabsstore.logout('false')
               this.$router.push({path: '/login'})
               sessionStorage.clear()
               localStorage.clear()
               this.loadingLg=false
               this.$message.warning('暂无授权任何菜单权限~')
            }

            setStorage('allmenu', JSON.stringify(this.allmenu))
            //优化首次登录第一个标签显示问题
            let tabsPage = ""
            let TabsValue = ""
            if(menuTree[0].hasChildren){
                tabsPage = [{"title":menuTree[0].children[0].text,"name":menuTree[0].children[0].attributes.url}]
                TabsValue = menuTree[0].children[0].attributes.url
            }else{
                tabsPage = [{"title":menuTree[0].text,"name":menuTree[0].attributes.url}]
                TabsValue = menuTree[0].attributes.url
            }
            this.mutitabsstore.firstTabs([tabsPage,TabsValue])
            this.$forceUpdate()
          } else {
            this.$message.warning(res.msg)
          }

          this.loadingLg=false
        })

        this.$forceUpdate()
      },


      getuserpassword() {
        if (getCookie('username') != '' && getCookie('password') != '') {
          this.ruleForm.username = getCookie('username')
          this.ruleForm.password = getCookie('password')
          this.rememberpassword = true
        }
      },

      /**
       * 获取验证码
       */
      getCaptchas () {
        getCaptcha().then((res) => {
          this.ruleForm.captcha = null
          this.ruleForm.captchaKey = res.data.captchaKey
          this.image_base = res.data.captcha
        })
      },

      //获取info列表
      submitForm(formName) {

        this.$refs[formName].validate(valid => {
          if (valid) {
            this.loadingLg = true
            login(this.ruleForm).then(async res => {
              this.loadingLg = false
              if (res.code === 2000) {
                if (this.rememberpassword) {
                  //保存帐号到cookie，有效期7天
                  await setCookie('username', this.ruleForm.username, 7)
                  //保存密码到cookie，有效期7天
                  await setCookie('password', this.ruleForm.password, 7)
                } else {
                  await delCookie('username')
                  await delCookie('password')
                }
                this.mutitabsstore.setLogintoken(res.data.access)
                this.mutitabsstore.setUserName(res.data.user.name)
                this.mutitabsstore.setUserId(res.data.user.id)
                this.mutitabsstore.setIdentity(res.data.user.identity)
                this.getMenu()
              } else {
                this.getCaptchas()
                this.$message.error(res.msg)
                return false
              }
            })
          } else {
            this.$message.error('请输入用户名密码/验证码！')
            return false
          }
        })

      }
    }
  }
</script>

<style lang="scss" scoped>
.login-page {
  position: relative;
  isolation: isolate;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  min-height: 100svh;
  padding: 88px 48px 78px;
  overflow: hidden;
  color: var(--ly-text-1);
  background-color: var(--ly-bg-page);
  background-image: var(--ly-page-gradient, none);
  box-sizing: border-box;
}

.login-page::before {
  position: absolute;
  z-index: -1;
  inset: 0;
  background:
    radial-gradient(ellipse at 16% 82%, rgba(58, 123, 255, 0.13), transparent 34%),
    radial-gradient(ellipse at 88% 18%, rgba(108, 155, 255, 0.16), transparent 30%),
    radial-gradient(ellipse at 50% 48%, rgba(255, 255, 255, 0.24), transparent 58%);
  content: '';
  pointer-events: none;
}

.login-page::after {
  position: absolute;
  z-index: -1;
  inset: 0;
  background-image:
    linear-gradient(rgba(58, 123, 255, 0.035) 1px, transparent 1px),
    linear-gradient(90deg, rgba(58, 123, 255, 0.035) 1px, transparent 1px);
  background-size: 64px 64px;
  content: '';
  mask-image: radial-gradient(ellipse at center, transparent 18%, #000 100%);
  pointer-events: none;
}

.login-tools {
  position: absolute;
  z-index: 4;
  top: 24px;
  right: 32px;
  display: flex;
  align-items: center;
  gap: 10px;
}

:deep(.tool-button.el-button) {
  width: 38px;
  height: 38px;
  margin: 0;
  border-color: var(--ly-line-soft);
  color: var(--ly-text-2);
  background: var(--ly-glass-bg-strong);
  box-shadow: var(--ly-glass-highlight);
}

:deep(.tool-button.el-button:hover) {
  border-color: color-mix(in srgb, var(--el-color-primary) 34%, transparent);
  color: var(--el-color-primary);
  background: var(--ly-glass-bg-strong);
}

:deep(.lydpselected) {
  color: var(--el-dropdown-menuItem-hover-color);
  background-color: var(--el-dropdown-menuItem-hover-fill);
}

.login-main {
  display: flex;
  width: 100%;
  justify-content: center;
}

.login-shell {
  display: grid;
  grid-template-columns: minmax(0, 0.94fr) minmax(0, 1.06fr);
  width: min(1050px, 100%);
  min-height: 580px;
  overflow: hidden;
  border: 1px solid var(--ly-glass-border);
  border-radius: 22px;
  background: var(--ly-glass-bg-strong);
  box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card);
  backdrop-filter: var(--ly-glass-blur);
  -webkit-backdrop-filter: var(--ly-glass-blur);
}

.brand-panel {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-width: 0;
  overflow: hidden;
  padding: 46px 48px 38px;
  color: #fff;
  background: linear-gradient(145deg, var(--ly-color-primary-deep), var(--el-color-primary));
}

.brand-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.055) 1px, transparent 1px);
  background-position: center;
  background-size: 38px 38px;
  mask-image: linear-gradient(140deg, transparent 8%, #000 70%);
  opacity: 0.48;
  pointer-events: none;
}

.brand-orbit {
  position: absolute;
  border: 1px solid rgba(255, 255, 255, 0.20);
  border-radius: 50%;
  pointer-events: none;
}

.brand-orbit-large {
  top: 112px;
  right: -195px;
  width: 440px;
  height: 440px;
  box-shadow: 0 0 0 34px rgba(255, 255, 255, 0.035), 0 0 0 76px rgba(255, 255, 255, 0.025);
}

.brand-orbit-small {
  top: 202px;
  right: 52px;
  width: 96px;
  height: 96px;
  border-color: rgba(255, 255, 255, 0.28);
  box-shadow: inset 0 0 0 17px rgba(255, 255, 255, 0.045);
}

.brand-lockup,
.brand-message,
.brand-footer {
  position: relative;
  z-index: 1;
}

.brand-lockup {
  display: flex;
  align-items: center;
  gap: 14px;
}

.brand-logo {
  display: grid;
  width: 48px;
  height: 48px;
  flex: 0 0 auto;
  place-items: center;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.54);
  border-radius: 13px;
  background: rgba(255, 255, 255, 0.94);
  box-shadow: 0 5px 15px rgba(18, 47, 118, 0.16);
}

.brand-logo img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.brand-name {
  display: block;
  min-width: 0;
  max-width: 280px;
  overflow: hidden;
  flex: 0 1 auto;
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.01em;
  line-height: 24px;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.brand-message {
  max-width: 370px;
  margin: auto 0;
  padding: 58px 0 68px;
}

.brand-kicker,
.login-kicker {
  display: inline-block;
  color: var(--el-color-primary);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.16em;
  line-height: 1.4;
}

.brand-kicker {
  color: rgba(255, 255, 255, 0.72);
}

.brand-message h1 {
  max-width: 360px;
  margin: 18px 0 12px;
  font-size: clamp(28px, 3.2vw, 38px);
  font-weight: 600;
  letter-spacing: 0.02em;
  line-height: 1.42;
}

.brand-message p {
  max-width: 330px;
  margin: 0;
  color: rgba(255, 255, 255, 0.74);
  font-size: 14px;
  line-height: 1.9;
}

.brand-footer {
  display: flex;
  align-items: center;
  gap: 12px;
  color: rgba(255, 255, 255, 0.68);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.14em;
}

.brand-footer > span:first-child {
  color: #fff;
  font-variant-numeric: tabular-nums;
}

.brand-footer-line {
  width: 36px;
  height: 1px;
  background: rgba(255, 255, 255, 0.42);
}

.form-panel {
  display: flex;
  align-items: center;
  justify-content: center;
  min-width: 0;
  padding: 52px clamp(38px, 5.2vw, 66px);
  background: var(--ly-glass-bg-strong);
}

.login-form {
  width: 100%;
  max-width: 380px;
  margin: 0 auto;
}

.login-heading {
  margin-bottom: 30px;
}

.login-kicker {
  color: var(--el-color-primary);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.login-heading h2 {
  margin: 12px 0 7px;
  color: var(--ly-text-1);
  font-size: 28px;
  font-weight: 650;
  letter-spacing: 0.01em;
  line-height: 1.35;
}

.login-heading p {
  margin: 0;
  color: var(--ly-text-2);
  font-size: 13px;
  line-height: 1.65;
}

:deep(.login-form .el-form-item) {
  display: block;
  margin-bottom: 18px;
}

:deep(.login-form .el-form-item__content) {
  display: block;
  line-height: normal;
}

.field-label {
  display: block;
  margin-bottom: 8px;
  color: var(--ly-text-1);
  font-size: 13px;
  font-weight: 600;
  line-height: 1.4;
}

:deep(.login-form .el-input__wrapper) {
  min-height: 48px;
  padding: 0 13px;
  border-radius: 10px;
  background: var(--ly-glass-bg-strong);
  box-shadow: 0 0 0 1px var(--ly-line-soft) inset;
  transition: box-shadow var(--ly-duration-fast) ease, background-color var(--ly-duration-fast) ease;
}

:deep(.login-form .el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px color-mix(in srgb, var(--el-color-primary) 38%, transparent) inset;
}

:deep(.login-form .el-input__wrapper.is-focus) {
  background: var(--ly-glass-bg-strong);
  box-shadow: 0 0 0 1px var(--el-color-primary) inset, var(--ly-input-focus-ring);
}

:deep(.login-form .el-form-item.is-error .el-input__wrapper) {
  box-shadow: 0 0 0 1px var(--el-color-danger) inset;
}

:deep(.login-form .el-input__prefix-inner > .el-icon) {
  color: var(--ly-text-3);
  font-size: 16px;
}

:deep(.login-form .el-input__inner) {
  color: var(--ly-text-1);
  font-size: 14px;
}

:deep(.login-form .el-input__inner::placeholder) {
  color: var(--ly-text-3);
  font-size: 13px;
}

:deep(.login-form .el-form-item__error) {
  padding-top: 4px;
  font-size: 12px;
}

:deep(.captcha-input .el-input-group__append) {
  width: 132px;
  min-width: 132px;
  flex: 0 0 132px;
  box-sizing: border-box;
  padding: 2px 6px;
  border: 0;
  background: transparent !important;
  box-shadow: none;
}

:deep(.captcha-input .el-input-group__append::before) {
  display: none;
}

.captcha-refresh {
  display: flex;
  width: 120px;
  min-width: 120px;
  height: 44px;
  align-items: center;
  justify-content: center;
  padding: 0;
  overflow: hidden;
  border: 0;
  border-radius: 7px;
  background: var(--ly-glass-bg-strong);
  cursor: pointer;
}

.login-code {
  display: block;
  width: 120px;
  max-width: 120px;
  height: auto;
  object-fit: contain;
}

.remember {
  margin: 0 0 21px;
}

:deep(.remember .el-checkbox__label) {
  color: var(--ly-text-2);
  font-size: 13px;
}

:deep(.submit-item.el-form-item) {
  margin: 0;
}

:deep(.login-submit.el-button) {
  display: flex;
  width: 100%;
  height: 50px;
  align-items: center;
  border: 0;
  border-radius: 10px;
  background: var(--ly-gradient-primary);
  box-shadow: 0 8px 20px rgba(58, 123, 255, 0.22);
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.04em;
}

:deep(.login-submit.el-button:hover) {
  filter: brightness(1.04);
}

:deep(.submit-icon) {
  margin-left: auto;
  font-size: 16px;
}

.login-copyright {
  position: absolute;
  right: 16px;
  bottom: 22px;
  left: 16px;
  color: var(--ly-text-3);
  font-size: 12px;
  text-align: center;
}

@media (max-width: 900px) {
  .login-page {
    padding-right: 30px;
    padding-left: 30px;
  }

  .brand-panel {
    padding-right: 34px;
    padding-left: 34px;
  }

  .form-panel {
    padding-right: 38px;
    padding-left: 38px;
  }
}

@media (max-width: 740px) {
  .login-page {
    padding: 76px 24px 72px;
  }

  .login-shell {
    grid-template-columns: minmax(0, 1fr);
    width: min(480px, 100%);
    min-height: 0;
    border-radius: 18px;
  }

  .brand-panel {
    min-height: 104px;
    justify-content: center;
    padding: 22px 28px;
  }

  .brand-message,
  .brand-footer,
  .brand-orbit-small {
    display: none;
  }

  .brand-orbit-large {
    top: -170px;
    right: -108px;
    width: 300px;
    height: 300px;
  }

  .brand-logo {
    width: 42px;
    height: 42px;
    border-radius: 11px;
  }

  .brand-name {
    font-size: 14px;
    line-height: 22px;
  }

  .form-panel {
    padding: 34px 34px 38px;
  }

  .login-heading {
    margin-bottom: 25px;
  }
}

@media (max-width: 420px) {
  .login-page {
    padding: 68px 16px 72px;
  }

  .login-tools {
    top: 14px;
    right: 16px;
  }

  .form-panel {
    padding: 28px 23px 30px;
  }

  .login-heading h2 {
    font-size: 25px;
  }
}

@media (max-height: 680px) and (min-width: 741px) {
  .login-page {
    padding-top: 68px;
    padding-bottom: 62px;
  }

  .login-shell {
    min-height: 540px;
  }

  .brand-panel {
    padding-top: 38px;
    padding-bottom: 30px;
  }

  .form-panel {
    padding-top: 38px;
    padding-bottom: 38px;
  }

  .login-heading {
    margin-bottom: 22px;
  }

  :deep(.login-form .el-form-item) {
    margin-bottom: 14px;
  }

  .remember {
    margin-bottom: 16px;
  }
}
</style>
