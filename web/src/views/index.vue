<template>
  <el-container class="index-con">
    <el-header class="index-header lyadmin-header">
      <navcon></navcon>
    </el-header>
    <el-container class="main-con">
      <div :class="showclass" class="lyadmin-side">
        <leftnav ></leftnav>
      </div>
      <el-container>
        <div class="index-main">
          <Mutitabs></Mutitabs>
        </div>
      </el-container>
    </el-container>
  </el-container>
</template>
<script setup>
    import {ref, onMounted,getCurrentInstance,computed} from 'vue'
    import navcon from '@/components/navcon.vue'
    import leftnav from '@/components/leftnav.vue'
    import Mutitabs from "@/components/mutitabs";
    import {useMutitabsStore} from "@/store/mutitabs";
    import {useSiteThemeStore} from "@/store/siteTheme";

    let bus = getCurrentInstance().appContext.config.globalProperties.$Bus; // 声明$Bus
    let showclass = ref("asideshow")
    let showtype = ref(false)
    const mutitabsStore =  useMutitabsStore()
    let isMultiTabs = mutitabsStore.isMultiTabs

    const siteThemeStore = useSiteThemeStore()

    const asideshowWidth = computed(()=>{
        return siteThemeStore.menuWidth +'px'
    })

    onMounted(()=>{
        bus.on('toggle', value => {
            if (value) {
                showclass.value = 'asideshow'
            } else {
                showclass.value = 'aside'
              //   setTimeout(() => {
              //       showclass.value = 'aside'
              // }, 200)
            }
        })
    })
</script>
<style scoped>
  .main-con{
    width:100%;
    height: 100%;
    display: flex;
    flex: 1;
    overflow: auto;
    /*overflow-y: auto;*/
  }
  .hg100{
    height: 100vh !important;
    overflow-y: hidden !important;
  }
  .index-con {
    height: 100%;
    width: 100%;
    box-sizing: border-box;
    overflow-y: hidden;
  }

  .aside {
    display: flex;
    flex-flow: column;
    flex-shrink: 0;
    transition: width .3s;
    width: 64px !important;
    background-color: var(--l-header-bg);
    margin: 0px;
    box-shadow: 0 0 3px #cccccc;
    /*height: calc(100vh - 60px);*/
  }

  .asideshow {
    display: flex;
    flex-flow: column;
    flex-shrink: 0;
    transition: width .3s;
    width: v-bind(asideshowWidth);
    /*height: calc(100vh - 60px);*/
    background-color: var(--l-header-bg);
    margin: 0px;
    box-shadow: 0 0 1px #cccccc;
  }
  .index-main {
    /*display: block;*/
    /*-webkit-box-flex: 1;*/
    /*-ms-flex: 1;*/
    flex: 1;
    /*-ms-flex-preferred-size: auto;*/
    /*flex-basis: auto;*/
    overflow: auto;
    /*padding: 8px 10px 0 10px;*/
    /*width: 100%;*/
    padding: 0;
    height: 100%;
  }
  .index-header{
    padding: 0px;
    width: 100%;
    /*box-shadow: 0px 4px 12px 0px rgba(0, 0, 0, 0.08);*/
  }
  .el-main.noPadding{
    padding: 0px !important;
    border-left: 2px solid #333;
  }

</style>