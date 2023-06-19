/**
* 左边菜单
*/
<template>
<!--  background-color="#304156"-->
<!--    active-background-color="#304156"-->
  <div class="lyadmin-wrapper-side">
    <!-- 经典布局 -->
    <el-scrollbar>
      <el-menu
        :default-active="route.meta.index"
        :collapse="collapsed"
        collapse-transition
        router
        unique-opened
        class="el-menu-vertical-demo"
        background-color="var(--l-header-bg)"
        text-color="#ffffff"
        active-text-color="var(--el-color-primary)">
          <lyLeftMenu :navMenus="allmenu"></lyLeftMenu>
      </el-menu>
    </el-scrollbar>
    <!-- 经典布局美化 -->
    <div v-if="lyLayout=='msimple'" class="lyadmin-side-bottom" @click="sideToggle">
      <svg-icon icon-class="expand" v-if="collapsed"></svg-icon>
      <svg-icon icon-class="fold" v-else></svg-icon>
    </div>
  </div>

</template>
<script setup>
    import {ref, onMounted, watch, getCurrentInstance, computed} from 'vue'
    import {useMutitabsStore} from "@/store/mutitabs";
    import {useSiteThemeStore} from "@/store/siteTheme";
    import { useRouter,useRoute,onBeforeRouteUpdate } from 'vue-router';
    import {setStorage, getStorage, deepClone} from '@/utils/util'
    import lyLeftMenu from '@/components/layout/lyLeftMenu'

    const route = useRoute()


    let bus = getCurrentInstance().appContext.config.globalProperties.$Bus; // 声明$Bus
    const mutitabsStore = useMutitabsStore()
    const siteThemeStore = useSiteThemeStore()
    let collapsed = ref(!siteThemeStore.collapsed)
    let allmenu = ref([])
    let menuTitle = ref("")

    function handleOpen2(obj){
        mutitabsStore.editableTabs(obj)
    }

    function getMenu() {
        menuTitle.value=''
        allmenu.value=[]

        //动态加载菜单
        allmenu.value = JSON.parse(getStorage('allmenu'))
    }

    const lyLayout = computed(() => {
      return siteThemeStore.programLayout
    })

    // 切换显示
    function sideToggle() {
        collapsed.value = !collapsed.value
        siteThemeStore.setCollapsed(collapsed.value)
        bus.emit('toggle', !collapsed.value)
    }

    onMounted(()=>{
        let userId = mutitabsStore.userId
        let params = {
            userId: userId
        };
        getMenu()
        bus.on("toggle", value => {
          collapsed.value = !value;
        });
        bus.on("routeReload", value => {
            getMenu()
        });
    })
    onBeforeRouteUpdate(to=>{
        // if(mutitabsStore.logintoken) {
        //     getMenu()
        // }
        mutitabsStore.switchtabNoRoute(to.name)
    })
</script>
<style lang="scss" scoped>
  .lyadmin-wrapper{
    display: flex;
    flex-flow: column;
    flex: 1;
    width: 100%;
  }
  .lyadmin-wrapper-side{
    display: flex;
    overflow: auto;
    overflow-x: hidden;
    flex-flow: column;
    flex: 1;
  }
  .lyadmin-wrapper-side::v-deep(.el-scrollbar){
    width: 100%;
  }
  .lyadmin-side-bottom{
    /*border-top: 1px solid var(--el-color-info-light-3);*/
    height:50px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #ffffff;
  }
  .lyadmin-side-bottom i {
    font-size: 24px;
  }
  .lyadmin-side-bottom:hover {
    color: var(--el-color-primary);
  }

  .el-menu.el-menu--horizontal{
    border-bottom: 0;
  }
  .menu-nav-title{
    height: 56px;
    line-height: 56px;
    padding-left: 20px;
    font-size: 20px;
    font-weight: bold;
    background: #eff6ff;
    color: var(--el-color-primary);
    border-bottom: 1px solid #c9e0ff;
  }
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 100%;
    min-height: 400px;
  }
  .el-menu-vertical-demo{
    i{
      margin-right: 5px;
      font-size: 18px;
    }
  }

  .el-menu-vertical-demo:not(.el-menu--collapse) {
    border: none;
    text-align: left;
  }

  ::v-deep(.el-menu-item-group__title) {
    padding: 0px !important;
  }
  .el-menu-item.is-active {
    position: relative;
    /*background-color: rgb(48, 54, 62) !important;*/
    background-color: var(--l-main-sidebar-menu-active-bg) !important;
    &:before{
      width: 2px;
      height: 100%;
      position: absolute;
      left: 0;
      top: 0;
      background: var(--l-main-sidebar-menu-hover-bg);
      display: block;
    }
  }
  .el-menu-bg {
    background-color: #1f2d3d !important;
  }

  .el-menu {
    border: none;
    overflow: hidden;

  }

  .router-link-active{
    color: #ffd04b;
  }
  .aside span{
    display: none;
  }
  .el-menu--collapse{
      ::v-deep(.el-sub-menu__icon-arrow){
        display: none !important;
      }
  }
  .el-sub-menu{
    /*width: 180px;*/
    width:100%
  }
</style>
