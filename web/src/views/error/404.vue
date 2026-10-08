<template>
  <main class="not-found-page">
    <section class="not-found-state" aria-labelledby="not-found-title">
      <div class="error-code">404</div>
      <div class="error-label">PAGE NOT FOUND</div>
      <h1 id="not-found-title">页面不存在</h1>
      <p>当前地址无法访问，可能已被移动、删除，或输入有误。</p>

      <div class="not-found-actions">
        <el-button type="primary" @click="backhome">返回首页</el-button>
        <el-button @click="goback">返回上一页</el-button>
      </div>
      <el-button class="relogin-button" link @click="exit">重新登录</el-button>
    </section>
  </main>
</template>

<script>
    import {useMutitabsStore} from "@/store/mutitabs";
    import {getStorage} from '@/utils/util'

    export default {
        name: "404",
        setup(){
            const mutitabsstore = useMutitabsStore()
            return { mutitabsstore}
        },
        methods:{
            backhome(){
                let allmenu = getStorage('allmenu')
                if(allmenu){
                    allmenu = JSON.parse(allmenu)
                    if(allmenu.length>0){
                        let tabsPage = allmenu[0].attributes.url
                        this.mutitabsstore.switchtab(tabsPage)
                    }
                }
            },
            goback(){
                this.$router.go(-1);
            },
            // 退出登录
            exit() {
                this.$confirm('退出登录, 是否继续?', '提示', {
                  confirmButtonText: '确定',
                  cancelButtonText: '取消',
                  type: 'warning'
                }).then(() => {
                    this.mutitabsstore.logout('false')
                    this.$router.push({path: '/login'})
                    sessionStorage.clear()
                    localStorage.clear()
                    this.$message.success('已退出登录!')
                }).catch(() => {})
            },
        },
    }
</script>

<style lang="scss" scoped>
.not-found-page {
  display: grid;
  width: 100%;
  height: 100%;
  min-height: 280px;
  place-items: center;
  box-sizing: border-box;
  padding: 24px;
  color: var(--ly-text-1);
  background: transparent;
}

.not-found-state {
  width: min(100%, 460px);
  padding: 24px;
  box-sizing: border-box;
  text-align: center;
}

.error-code {
  color: var(--el-color-primary);
  font-size: clamp(64px, 9vw, 88px);
  font-weight: 650;
  letter-spacing: -.065em;
  line-height: 1;
  font-variant-numeric: lining-nums;
}

.error-label {
  margin-top: 10px;
  color: var(--ly-text-3);
  font-size: 10px;
  font-weight: 600;
  letter-spacing: .16em;
}

.not-found-state h1 {
  margin: 18px 0 8px;
  color: var(--ly-text-1);
  font-size: 21px;
  font-weight: 600;
  line-height: 1.45;
}

.not-found-state p {
  margin: 0;
  color: var(--ly-text-2);
  font-size: 13px;
  line-height: 1.8;
}

.not-found-actions {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 22px;
}

:deep(.not-found-actions .el-button) {
  min-width: 112px;
  height: 38px;
  margin: 0;
  border-radius: 8px;
}

.relogin-button {
  margin-top: 10px;
  color: var(--ly-text-3);
}

.relogin-button:hover {
  color: var(--el-color-primary);
}

@media (max-width: 480px) {
  .not-found-page {
    min-height: 240px;
    padding: 16px;
  }

  .not-found-state {
    padding: 16px 8px;
  }

  .not-found-state h1 {
    font-size: 19px;
  }
}
</style>
