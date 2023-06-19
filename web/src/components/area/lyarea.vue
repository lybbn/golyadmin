<!--
 * @Descripttion: 地区选择组件
 * @version: 1.0
 * @Program: django-vue-lyadmin
 * @Author: lybbn
 * @Date: 2023/02/11
-->
<template>
    <div class="lyarea">
        <el-cascader
            style="width: 100%"
            :size="size"
            :show-all-levels="showAllLevels"
            :options="pcdData"
            v-model="lyselectedOptions"
            :props="{expandTrigger:'hover', checkStrictly: false ,label:'name',value:'name',children:'childlist'}"
            clearable
            :placeholder="placeholder2"
            :disabled="disabled"
            @change="lyhandleChange">
        </el-cascader>
    </div>
</template>

<script setup>
    import {getAllAreasList} from "@/api/api"
    import { onMounted,onUnmounted, ref, watch,nextTick,computed } from 'vue'

    const emit = defineEmits(['update:modelValue'])

    const props = defineProps({
        size:{
            type:String,
            default:'default',
        },
        // 绑定文本值
        modelValue: {
            type: [String,null],
            default: ''
        },
        showAllLevels:{
            type:Boolean,
            default:true
        },
        disabled:{
            type:Boolean,
            default:false
        },
        placeholder:{
            type:String,
            default:"请选择"
        }
    })

    let placeholder2 = computed(()=>{
        return props.modelValue?props.modelValue:props.placeholder
    })

    let lyselectedOptions = ref("")
    let pcdData = ref([])

    function lyhandleChange(value) {
        // console.log(value);
        // console.log(lyselectedOptions.value.join(""))
        emit('update:modelValue', lyselectedOptions.value.join(""))
    }

    function getData(){
        getAllAreasList().then(res=>{
            if(res.code == 2000){
                pcdData.value = res.data.data
            }
        })
    }

    onMounted(()=>{
        lyselectedOptions.value = ""
        getData()
    })


</script>

<style scoped>
    .lyarea{
        width: 100%;
    }
</style>