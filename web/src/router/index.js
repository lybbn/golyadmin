import { createRouter, createWebHistory ,createWebHashHistory } from 'vue-router'
import {useMutitabsStore} from "@/store/mutitabs";
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import '../assets/css/nprogress.scss'
NProgress.inc(0.4)
NProgress.configure({ easing: 'ease', speed: 500, showSpinner: true })
import {setStorage,getStorage} from '@/utils/util'

const routes = [
  {
    path: '/',
    name: 'root',
    component: () => import('../views/login.vue'),
    hidden: true,
    meta: {
      requireAuth: false,
      index: '/',
    }
  },
    {
    path: '/login',
    name: 'login',
    component: () => import('../views/login.vue'),
    hidden: true,
    meta: {
      requireAuth: false,
      index: '/login',
    }
  },
    {
    path: '/404',
    name: '404',
    component: () => import('../views/error/404.vue'),
    hidden: true,
    meta: {
      requireAuth: false,
      index: '/404',
    }
  },
  {
    path: '/index',
    name: 'index',
    component: () => import('../views/index.vue'),
    iconCls: 'el-icon-tickets',
    meta: {
      requireAuth: false,
      index: '/index',
    },
    children: [
      {
        path: '/adminManage',
        name: 'adminManage',
        component: () => import('../views/adminManage/adminManage.vue'),
        meta: {
          requireAuth: true,
          index: '/adminManage',
        }
      },
        {
        path: '/userManageCrud',
        name: 'userManageCrud',
        component: () => import('../views/userManage/userManageCrud.vue'),
        meta: {
          requireAuth: true,
          index: '/userManageCrud',
        }
      },
        {
            path: '/carouselSettingsimg',
            name: 'carouselSettingsimg',
            component: () => import('../views/platformSettings/carouselSettingsimg.vue'),
            meta: {
                requireAuth: true,
                index: '/carouselSettingsimg',
            }
        },
        {
            path: '/platformSettingsother',
            name: 'platformSettingsother',
            component: () => import('../views/platformSettings/platformSettingsother.vue'),
            meta: {
                requireAuth: true,
                index: '/platformSettingsother',
            }
        },
      {
        path: '/departmentManage',
        name: 'departmentManage',
        component: () => import('../views/systemManage/departmentManage/departmentManage.vue'),
        meta: {
          requireAuth: true,
          index: '/departmentManage',
        }
      },
      {
        path: '/menuManage',
        name: 'menuManage',
        component: () => import('../views/systemManage/menuManage/menuManage.vue'),
        meta: {
          requireAuth: true,
          index: '/menuManage',
        }
      },

      {
        path: '/roleManage',
        name: 'roleManage',
        component: () => import('../views/systemManage/roleManage/roleManage.vue'),
        meta: {
          requireAuth: true,
          index: '/roleManage',
        }
      },


      {
        path: '/authorityManage',
        name: 'authorityManage',
        component: () => import('../views/systemManage/authorityManage/authorityManage.vue'),
        meta: {
          requireAuth: true,
          index: '/authorityManage',
        }
      },
      {
        path: 'buttonConfig',
        name: 'buttonConfig',
        component: () => import('../views/systemManage/buttonConfig/buttonConfig.vue'),
        meta: {
          requireAuth: true,
          index: '/buttonConfig',
        }
      },
        {
        path: '/buttonManage',
        name: 'buttonManage',
        component: () => import('../views/systemManage/button/buttonManage.vue'),
        meta: {
          requireAuth: true,
          index: '/buttonManage',
        }
      },
      {
        path: '/messagNotice',
        name: 'messagNotice',
        component: () => import('../views/messageCenter/messagNotice.vue'),
        meta: {
          requireAuth: true,
          index: '/messagNotice',
        }
      },
      {
        path: '/personalCenter',
        name: 'personalCenter',
        component: () => import('../views/personalCenter/personalCenter.vue'),
        meta: {
          requireAuth: true,
          index: '/personalCenter',
        }
      },
      {
        path: '/journalManage',
        name: 'journalManage',
        component: () => import('../views/journalManage/journalManage.vue'),
        meta: {
          requireAuth: true,
          index: '/journalManage',
        }
      },

    ]
  }
]

const viewModules = import.meta.glob('../views/**/*.vue')

function getAutoRouterList() {
    const routerList = [];
    const viewKeys = Object.keys(viewModules)
    viewKeys.forEach((name) => {
        if(name.indexOf("/components/")==-1 && !name.includes('index.vue') && !name.includes('login.vue') && !name.includes('lyterminal.vue')){
            let isSame = false
            const componentName = name.split('/').pop()?.split('.')[0]
            for(let i=0;i<routes.length;i++){
                if(routes[i].path=="/"||routes[i].path=="/login" ||routes[i].path=="/lyterminal"){
                    continue
                }
                if(routes[i].name === componentName){
                    isSame = true
                    break
                }
                if(routes[i].path === '/index' && routes[i].children.length>0){
                    for(let s=0;s<routes[i].children.length;s++){
                          if(routes[i].children[s].name === componentName){
                              isSame = true
                              break
                          }
                    }
                }
            }
            if(!isSame){
                const path = "/"+componentName
                routerList.push({
                    path: path,
                    name: componentName,
                    component: viewModules[name],
                    meta: {
                        requireAuth: true,
                        index: path,
                    }
                });
            }
        }

    });
    for(let t=0;t<routes.length;t++){
        if(routes[t].path == '/index'){
            routerList.forEach(drouter=>{
                routes[t].children.push(drouter)
            })
            break
        }
    }
    return routerList;
}

getAutoRouterList()

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes
})

function isRouterNameExist(name){
    let isExist = false
    if(name){
        if(routes.some(item=> name == item.name)){
            return true
        }
        for(let t=0;t<routes.length;t++){
            if(routes[t].children && routes[t].children.length>0){
                if(routes[t].children.some(item=> name == item.name)){
                    isExist = true
                    break
                }
            }
        }
    }else{
        return false
    }
    return isExist
}
function isRouterPathExist(path){
    let isExist = false
    if(path){
        if(routes.some(item=>path == item.path)){
            return true
        }
        for(let t=0;t<routes.length;t++){
            if(routes[t].children && routes[t].children.length>0){
                if(routes[t].children.some(item=> path == item.path)){
                    isExist = true
                    break
                }
            }
        }
    }else{
        return false
    }
    return isExist
}

router.beforeEach((to, from, next) => {
    const store = useMutitabsStore()
    const whiteList = ['buttonConfig', 'menuManage', 'lyterminal', 'buttonManage','lyFilePreview']
    NProgress.start()
    let userId = store.userId ? store.userId : ''
    if (to.meta.requireAuth) {
        if (userId) {
            let menuList = JSON.parse(getStorage('menuList'))
            if(menuList && (menuList.filter(item=>item.url == to.name).length > 0 || (whiteList.indexOf(to.name) !== -1))) {
                if(to.path){
                    next()
                }else if(isRouterPathExist(to.path)){
                    next()
                }else{
                    next({
                        path: '/404'
                    })
                }
            } else {
                next({
                    path: '/404'
                })
            }
        } else {
            store.logout('false')
            next({
              path: '/login'
            })
        }
    } else {
        if(to.path=="/login" ||to.path=="/"){
            if(userId){
                let tabsValue = getStorage("TabsValue")
                if(tabsValue){
                    if(isRouterNameExist(tabsValue)){
                        if(tabsValue === 'login'){
                            next({
                                path: '/404'
                            })
                        }else{
                            store.switchtab(tabsValue)
                        }
                    }else{
                        next({
                            path: '/404'
                        })
                    }
                }else{
                    let tabsPage = JSON.parse(getStorage("tabsPage"))
                    if (tabsPage) {
                        if(isRouterNameExist(tabsPage[0].name)){
                            store.switchtab(tabsPage[0].name)
                        }else{
                          next({
                              path: '/404'
                          })
                        }

                    }else{
                        next({
                            path:'/404'
                        })
                    }
                }
            }else{
                if(to.name){
                    next()
                }else if(isRouterPathExist(to.path)){
                    next()
                }
                else{
                    next({
                        path: '/404'
                    })
                }
            }
        }else{
            if(to.name){
                next()
            }else if(isRouterPathExist(to.path)){
                next()
            }else{
                next({
                    path: '/404'
                })
            }
        }
    }
})
router.afterEach(() => {
  NProgress.done()
})
export default router
