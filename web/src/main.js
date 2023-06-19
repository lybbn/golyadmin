import { createApp } from 'vue'
//引入ElementPlus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import i18n from './locales'
// 统一导入el-icon图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
//引入ElementPlus 结束

//多标签bus通信
import $bus from '@/utils/bus'

//导入自定义css
import './assets/css/common.scss'
//elementplus暗黑主题从（element-plus/theme-chalk/dark/css-vars.css）拷贝
import './assets/css/elementplus-theme-dark-css-vars.css'

import App from './App.vue'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'

import router from './router'

const app = createApp(App)

app.config.globalProperties.axios = axios

// vue3.x的全局实例，要挂载在config.globalProperties上
app.config.globalProperties.$Bus = $bus

//引入状态管理
import {isShowBtn,hasPermission,formatUnitSize} from '@/utils/util'
app.config.globalProperties.isShowBtn = isShowBtn
app.config.globalProperties.hasPermission = hasPermission
app.config.globalProperties.formatUnitSize = formatUnitSize
// 过滤器
import * as custom from './utils/util'
Object.keys(custom).forEach(key => {
  app.component(key, custom[key])
})

// //进入自定义指令
import directivePlugin from '@/utils/directive.js'

// 注册全局elementplus icon组件
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
import registerSvgIcon from "@/icons"
registerSvgIcon(app)

//注册自定义组件
import lyComponents from '@/components/index' //加载公共组件

app.use(ElementPlus)
app.use(i18n)
app.use(lyComponents)
app.use(store)
app.use(router)
app.use(VueAxios,axios)
app.use(directivePlugin)
app.mount('#app')


