import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import i18n from './locales'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import $bus from '@/utils/bus'

import './assets/css/common.scss'
import './assets/css/elementplus-theme-dark-css-vars.css'

import 'virtual:svg-icons-register'

import App from './App.vue'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'

import router from './router'

const app = createApp(App)

app.config.globalProperties.axios = axios

app.config.globalProperties.$Bus = $bus

import {isShowBtn,hasPermission,formatUnitSize} from '@/utils/util'
app.config.globalProperties.isShowBtn = isShowBtn
app.config.globalProperties.hasPermission = hasPermission
app.config.globalProperties.formatUnitSize = formatUnitSize
import * as custom from './utils/util'
Object.keys(custom).forEach(key => {
  app.component(key, custom[key])
})

import directivePlugin from '@/utils/directive.js'

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
import registerSvgIcon from "@/icons"
registerSvgIcon(app)

import lyComponents from '@/components/index'

app.use(ElementPlus)
app.use(i18n)
app.use(lyComponents)
app.use(store)
app.use(router)
app.use(VueAxios,axios)
app.use(directivePlugin)
app.mount('#app')
