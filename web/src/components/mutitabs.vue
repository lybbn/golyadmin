<template>
    <div class="myeltas1 lyadmin-body">
        <el-tabs
            v-if="isMultiTabs"
            class="myeltas2"
            v-model="editableTabsValue"
            type="card"
            closable
            @tab-remove="removeTab"
            @tab-click="tabClick($event)"
            @contextmenu.prevent.native="openContextMenu($event)">
            <el-tab-pane
                :key="item.name"
                v-for="item in editableTabs"
                :label="item.title"
                :name="item.name">
            </el-tab-pane>
        </el-tabs>
        <transition name="el-zoom-in-top">
            <ul v-show="contextMenuVisible" :style="{left:left+'px',top:top+'px'}" class="contextmenu" id="lycontextmenu">
                <li @click="reloadPage"><el-icon><Refresh /></el-icon><span class="contextmenu-text">刷新</span></li>
                <li @click="closeAllTabs"><el-icon><CircleCloseFilled /></el-icon><span class="contextmenu-text">关闭所有</span></li>
                <li @click="closeOtherTabs('left')"><el-icon><Back /></el-icon><span class="contextmenu-text">关闭左边</span></li>
                <li @click="closeOtherTabs('right')"><el-icon><Right /></el-icon><span class="contextmenu-text">关闭右边</span></li>
                <li @click="closeOtherTabs('other')"><el-icon><Delete /></el-icon><span class="contextmenu-text">关闭其他</span></li>
                <li @click="maximize"><el-icon><FullScreen /></el-icon><span class="contextmenu-text">最大化</span></li>
                <li @click="openWindow"><el-icon><CopyDocument /></el-icon><span class="contextmenu-text">新窗口打开</span></li>
                <li @click="closeContextMenu()"><el-icon><Close /></el-icon><span class="contextmenu-text">取消操作</span></li>
            </ul>
        </transition>
        <el-main class="lyadmin-main-content">
            <router-view v-slot="{Component,route}">
                <!--              <component :is="Component" v-if="!route.meta.isActive" :key="route.path"></component>-->
    <!--            <transition name="lyfadein" mode="out-in">-->
                <keep-alive :include="keepAliveRoutes" :exclude="excludes">
                    <component :is="Component" :key="route.name" v-if="!isRourterAlive" ref="lyComponent"></component>
                </keep-alive>
    <!--            </transition>-->
            </router-view>
        </el-main>
        <div class="lymain-maximize-exit" @click="exitMaximize"><el-icon><close /></el-icon></div>
    </div>
</template>

<script setup>
    import {ref, onMounted ,reactive,provide,watch,computed,nextTick} from 'vue'
    import {useKeepAliveStore} from "@/store/keepAlive";
    import {useMutitabsStore} from "@/store/mutitabs";
    import {useSiteThemeStore} from "@/store/siteTheme";
    import NProgress from 'nprogress'
    import 'nprogress/nprogress.css'
    import {setStorage,getStorage} from '@/utils/util'
    import { useRouter,useRoute,onBeforeRouteUpdate } from 'vue-router';
    import { ElMessage,ElMessageBox } from 'element-plus'

    const route = useRoute()
    const router = useRouter()

    const keepAliveStore = useKeepAliveStore()
    const mutitabsstore = useMutitabsStore()
    const siteThemeStore = useSiteThemeStore()

    let isMultiTabs = computed(()=>{
        return mutitabsstore.isMultiTabs
    })

    //右键自定义菜单
    let contextMenuVisible = ref(false)
    let left = ref(0)
    let top = ref(0)
    //刷新当前页
    let isRourterAlive = ref(false)
    let excludes = ref("")

    const editableTabs = computed({
        get() {
            return mutitabsstore.tabsPage;
        },
        set(val) {
            mutitabsstore.tabsPage = val;
        },
    })

    const editableTabsValue = computed({
        get() {
            return mutitabsstore.TabsValue;
        },
        set(val) {
            mutitabsstore.TabsValue = val;
        },
    })

    const keepAliveRoutes = computed({
        get() {
            return keepAliveStore.keepAliveRoute
        },
    })


    let lyComponent = ref(null)
    //高亮tag
    function isActive(routes) {
        return routes.name === route.name
    }
    //退出最大化
    function exitMaximize(){
        document.getElementById('app').classList.remove('lymain-maximize')
        if(lyComponent.value.setFull){
            lyComponent.value.setFull()
        }
    }
     //标签页最大化
    function maximize(){
        var TabsValue = mutitabsstore.TabsValue
        contextMenuVisible.value = false
        //判断是否当前路由，否的话跳转
        if(route.name != TabsValue){
            route.push({
                name: TabsValue,
            })
        }
        document.getElementById('app').classList.add('lymain-maximize')

        if(lyComponent.value.setFull){
            lyComponent.value.setFull()
        }
    }
    //新窗口打开
    function openWindow(){
        let currentPath = mutitabsstore.TabsValue
        let routeData = router.resolve({path:currentPath})
        window.open(routeData.href,'_blank')
        contextMenuVisible.value = false
    }
    function relogin(){
        mutitabsstore.logout('false')
        siteThemeStore.setSiteTheme('light')
        sessionStorage.clear()
        localStorage.clear()
        ElMessage.warning('请重新登录!')
        router.push({path: '/login'})
    }
    function removeTab(targetName) {
        let tabs = editableTabs.value;
        let activeName = editableTabsValue.value;
        //只有一个标签不允许关闭
        if(tabs.length === 1){
            return
        }
        if (activeName === targetName) {
            tabs.forEach((tab, index) => {
                if (tab.name === targetName) {
                    let nextTab = tabs[index + 1] || tabs[index - 1];
                    // console.log(nextTab);
                    if (nextTab) {
                        activeName = nextTab.name;
                    }
                }
        });
        }
        editableTabsValue.value = activeName;
        editableTabs.value = tabs.filter((tab) => tab.name !== targetName);
        mutitabsstore.tabsPage = editableTabs.value;
        setStorage("tabsPage",JSON.stringify(editableTabs.value));
        //解决刷新消失
        setStorage("TabsValue", activeName);
        var thetabsPage = getStorage("tabsPage")
        // 删除时跳转不在停留被删除页
        if (thetabsPage === "[]"||thetabsPage==""||thetabsPage==null) {
            relogin()
        } else {
            router.push({ name: activeName });
        }
    }
    function tabClick(event) {
        //关闭自定义菜单
        closeContextMenu()
        //写一个点击tabs跳转
        mutitabsstore.switchtab(event.props.name)
    }
    //自定义菜单
    function openContextMenu(e) {
        var obj =  e.srcElement ? e.srcElement : e.target;
        if (obj.id) {
            let currentContextTabId = obj.id.split("-")[1];
            contextMenuVisible.value = true;
            mutitabsstore.saveCurContextTabId(currentContextTabId);
            left.value = e.clientX + 1;
            top.value = e.clientY + 1;
            //右键菜单边缘化位置处理（防止在最边缘右键右健菜单消失部分）
            nextTick(() => {
                let ct = document.getElementById("lycontextmenu");
                if(document.body.offsetWidth - e.clientX < ct.offsetWidth){
                    left.value = document.body.offsetWidth - ct.offsetWidth - 1;
                    top.value = e.clientY + 1;
                }
            })
        }
    }
    function reloadPage(){
        NProgress.start()
        contextMenuVisible.value = false
        const currentRoute = router.currentRoute.value;
        currentRoute.matched.forEach((r)=> {
            if (r.path === currentRoute.fullPath) {
                //获取到当前页面的name
                const comName = r.components.default.name!= undefined?r.components.default.name:r.components.default.__name;
                if (comName != undefined) {
                    excludes.value = comName
                    isRourterAlive.value = true
                }
            }
        })
        nextTick(() => {
            isRourterAlive.value = false
            excludes.value = ""
            NProgress.done()
        })
    }
    // 关闭所有标签页
    function closeAllTabs() {
        mutitabsstore.closeAllTabs()
        contextMenuVisible.value = false;
    }
    // 关闭其它标签页
    function closeOtherTabs(par) {
        mutitabsstore.closeOtherTabs(par);
        contextMenuVisible.value = false;
    }
    // 关闭contextMenu
    function closeContextMenu() {
        contextMenuVisible.value = false;
    }

    onMounted(()=>{
        document.addEventListener("click", (e) => {
            if (e.target.className !="myeltas2") {
                contextMenuVisible.value = false; //点击其他区域关闭右键菜单
            }
        });
        //刷新加载localStorage存着地址
        let lytabsPage = getStorage("tabsPage")
        if (lytabsPage) {
            var TabsValue = getStorage("TabsValue");
            if (lytabsPage === "[]"||lytabsPage==""||lytabsPage==null || TabsValue === 'login') {
                relogin()//重新登录
            }
            mutitabsstore.tabsPage = JSON.parse(lytabsPage);
            const currentRouteName = route.name
            if(currentRouteName == 'login' || currentRouteName == 'root'){
                mutitabsstore.TabsValue = TabsValue;
                router.push({ name: TabsValue })
            }else{
                mutitabsstore.switchtabNoRoute(currentRouteName)
            }

        }else{
            relogin()//重新登录
        }
    })

</script>
<style>
    .myeltas1{
        /*padding: 0 8px 0 8px;*/
        height: 100%;
        min-height: 0;
        display: flex;
        flex-direction: column;
        overflow: hidden;
    }
    .myeltas1 > .lyadmin-main-content {
        flex: 1 1 0%;
        flex-basis: 0;
        min-height: 0;
        min-width: 0;
        overflow: auto;
    }
    .myeltas2{
        /*height: 48px;*/
        /* 顶部呼吸空隙（悬浮式标签条）：必须用 margin 上8/下-8 抵消实现，净占位与原版自然高度一致。
           禁止用固定 height+padding 占位（曾改 height:41px+padding-top:8 多占 7px）——
           useTableHight.js 的表格高度公式（innerHeight-搜索区-178）按原骨架常量校准，
           标签条多占几像素 .lyadmin-main-content 就会冒出几像素空滚动条（实测恰好 scrolling=7） */
        margin-top: 8px;
        margin-bottom: -8px;
    }
    .myeltas2 .el-tabs__header{
        margin: 0;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item.is-active {
        color: #fff;
        border-color: transparent;
        background: var(--ly-gradient-primary) !important;
        box-shadow: var(--ly-sidebar-active-shadow);
    }
    .myeltas2 .el-tabs__nav .el-tabs__item{
        transition:none;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item.is-closable .is-icon-close {
        width: 14px;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item.is-closable:hover{
        padding-left: 15px;
        padding-right: 15px;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item.is-active.is-closable{
        padding-left: 15px;
        padding-right: 15px;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item.is-active:before {
        content: "";
        background: #fff;
        display: inline-block;
        width: 8px;
        height: 8px;
        border-radius: 50%;
        position: relative;
        margin-right: 5px;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item:last-child {
        padding-right: 15px !important;
    }
    .myeltas2 .el-tabs__nav .el-tabs__item:nth-child(2) {
        padding-left: 15px !important;
    }
    .myeltas2 .el-tabs__nav-wrap{
        background: var(--l-changetab-bg);
        border: none;
    }
    /*去除顶部线*/
    .myeltas2 .el-tabs__header {
        /*border: none;*/
        /*margin: 0 0 2px;*/
        border-bottom :none;
    }
    .myeltas2 .el-tabs__nav{
        /*background-color: lightgrey;*/
        background-color: var(--l-changetab-bg);
        /*border: none !important;*/
        border-radius:0 !important;
    }
    /*字体大小*/
    .myeltas2  .el-tabs__item{
        font-size: 13px;
        color: var(--ly-text-2);
        height: 30px;
        line-height: 30px;
        padding: 0 15px;
    }
    /*标签左右箭头*/
    .myeltas2 .el-tabs__nav-next, .el-tabs__nav-prev {
        line-height: 35px;
        font-size: 17px;
    }
    /*自定义右键菜单（玻璃规格与 app.scss 全局规则一致，此处为兜底来源）*/
    .contextmenu {
        width: 130px;
        margin: 0;
        border: 1px solid var(--ly-glass-border);
        background: var(--l-changetab-right-menu);
        z-index: 3000;
        position: absolute;
        list-style-type: none;
        padding: 5px 0;
        border-radius: var(--ly-radius-md);
        font-size: 14px;
        box-shadow: var(--ly-shadow-card);
    }
    .contextmenu li {
        margin: 0;
        padding: 9px 16px;
        display: flex;
        align-items: center;
    }
    .contextmenu li:hover {
        background: var(--l-changetab-right-menu-hover-bg);
        cursor: pointer;
        color: var(--el-color-primary);
    }
    .contextmenu-text{
        margin-left: 5px;
    }
        .lyfadein-enter-from {
        transform: translateX(-0.5%);
    }
    .lyadmin-main-content{
        padding: 10px !important;
    }
    .lyfadein-leave-from {
        transform: translateX(0);
    }
    .lyfadein-leave-to {
        transform: translateX(0.5%);
        opacity: 0.1;
    }

    .lyfadein-enter-active{
        transition: all 0.4s;
    }
    .lyfadein-leave-active{
        transition: all 0.3s;
    }

    .lyfadein-enter-to {
        transform: translateX(0);
    }

</style>
