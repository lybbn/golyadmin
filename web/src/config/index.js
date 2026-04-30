import { version } from '../../package.json'

const isDev = import.meta.env.DEV

const API_DOMAIN = isDev ? "127.0.0.1:9000" : "golyadmin.lybbn.cn"
const API_BASEURL = isDev ? "http://"+ API_DOMAIN +"/api/" : "https://"+ API_DOMAIN +"/api/"
const VITE_APP_PROXY = false

export default {
    API_DOMAIN,
    API_BASEURL,
    VITE_APP_PROXY,
    API_URL: isDev && VITE_APP_PROXY ? "/api/" : API_BASEURL,
    APP_TITLE: "golyadmin后台管理系统",
    APP_VER: version,
    APP_NAME: "golyadmin后台管理系统",
    PROGRAM_LAYOUT: 'msimple',
    ISMULTITABS: true,
    LANG: 'zh-cn',
    ELEMENT_SIZE: 'default',
    ELEMENT_ZINDEX: 3000,
    ELEMENT_BUTTON: false,
    MENU_IS_COLLAPSE: false,
    MENU_WIDTH: 200,
    MENU_HEADER_COLOR: '#272E39',
    COLOR: '#1966ff',
    THEME: 'light',
    PAGING_LAYOUT: 'white',
    STORAGE_METHOD: 'localStorage',
    TIMEOUT: 350000,
}
