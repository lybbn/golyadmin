import {ref,reactive,onMounted,onUnmounted,nextTick} from 'vue'
import {getTableHeight} from "@/utils/util";
export default function(orderStatic,tableSelect,isFull=false,allowPage=true) {

    let tableHeight = ref(500)
    let orderstaticHeight = ref(0)
    // 计算搜索栏的高度
    function  listenResize() {
        nextTick(() => {
            getTheTableHeight()
        })
    }

    function getTheTableHeight(){
        if(orderStatic.value && orderStatic.value !==undefined){
            orderstaticHeight.value = orderStatic.value.offsetHeight
        }
        let tableSelectHeight =  tableSelect.value?tableSelect.value.offsetHeight:0
        let tableSelectTop = tableSelect.value?tableSelect.value.offsetTop:0
        let isTrueFull =  isFull || tableSelectTop == 0?true:false
        tableHeight.value =  getTableHeight(isTrueFull?tableSelectHeight+orderstaticHeight.value - 110:tableSelectHeight+orderstaticHeight.value,allowPage)
    }
    onMounted(()=>{
        // 监听页面宽度变化搜索框的高度
        window.addEventListener('resize', listenResize);
        listenResize()
    })
    onUnmounted(()=>{
        // 页面销毁，去掉监听事件
        window.removeEventListener("resize", listenResize);
    })

    return tableHeight
}