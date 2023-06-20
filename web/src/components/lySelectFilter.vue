<!--
 * @Descripttion: 分类筛选器
 * @version: 1.0
 * @program: djangolyadmin
 * @Author: lybbn
 * @Date: 2023.05.01
-->

<template>
	<div class="ly-select-filter">
		<div v-if="data.length<=0" class="ly-select-filter__no-data">
			暂无数据
		</div>
		<div v-for="item in data" :key="item.key" class="ly-select-filter__item">
			<div class="ly-select-filter__item-title" :style="{'width':labelWidth+'px'}"><label>{{item.title}}：</label></div>
			<div class="ly-select-filter__item-options">
				<el-badge v-for="option in item.options" style="margin-right: 22px;" :key="option.value" :type="option.type?option.type:''" :value="option.nums?option.nums:''">
					<el-button :size="size" :round="round" :plain="plain" :circle="circle" :icon="option.icon?option.icon:''" @click="select(option, item)" :type="selected[item.key]&&selected[item.key].includes(option.value)?'primary':''">{{option.label}}</el-button>
				</el-badge>
			</div>
		</div>
	</div>
</template>

<script setup>
	import {ref, onMounted, reactive, nextTick, watch, computed} from 'vue'
	import {deepClone} from "@/utils/util";

	const emit = defineEmits(['onChange'])
	const props = defineProps({
        data: { type: Array, default: () => [] },
		selectedValues: { type: Object, default: () => { return {} } },
		labelWidth: {type: Number, default: 80},
		outputValueTypeToArray: { type: Boolean, default: false },
		size:{ type:String,default:"small" },
		round:{ type: Boolean, default: true },
		plain:{ type: Boolean, default: false },
		circle:{ type: Boolean, default: false },
    })
	let selected =  ref({})

	watch(() => props.data, (val) => {
        val.forEach(item => {
			selected.value[item.key] = props.selectedValues[item.key] || (Array.isArray(item.options) && item.options.length) ? [item.options[0].value] : []
		})
    })

	const selectedString = computed(() => {
        var outputData = deepClone(selected.value)
		for (var key in outputData) {
			outputData[key] = outputData[key].join(",")
		}
		return outputData
    })

	function select(option, item){
		//判断单选多选
		if(item.multiple){
			//如果多选选择的第一个
			if(option.value === item.options[0].value){
				//就赋值第一个的值
				selected.value[item.key] = [option.value]
			}else{
				//如果选择的值已有
				if(selected.value[item.key].includes(option.value)){
					//删除选择的值
					selected.value[item.key].splice(selected.value[item.key].findIndex(s => s === option.value), 1)
					//当全删光时，把第一个选中
					if(selected.value[item.key].length == 0){
						selected.value[item.key] = [item.options[0].value]
					}
				}else{
					//未有值的时候，追加选中值
					selected.value[item.key].push(option.value)
					//当含有第一个的值的时候，把第一个删除
					if(selected.value[item.key].includes(item.options[0].value)){
						selected.value[item.key].splice(selected.value[item.key].findIndex(s => s === item.options[0].value), 1)
					}
				}
			}
		}else{
			//单选时，如果点击了已有值就赋值
			if(!selected.value[item.key].includes(option.value)){
				selected.value[item.key] = [option.value]
			}else{
				return false
			}
		}
		change()
	}
	function change(){
		if(props.outputValueTypeToArray){
			emit('onChange', selected.value)
		}else{
			emit('onChange', selectedString.value)
		}
	}

	onMounted(()=>{
		//默认赋值
		props.data.forEach(item => {
			selected.value[item.key] = props.selectedValues[item.key] || (Array.isArray(item.options) && item.options.length) ? props.selectedValues[item.key] : []
			// selected.value[item.key] = props.selectedValues[item.key] || (Array.isArray(item.options) && item.options.length) ? [item.options[0].value] : []
		})
		change()
	})
</script>

<style scoped>
	.ly-select-filter {width: 100%;font-size: 14px;}
	.ly-select-filter__item {display: flex;align-items: center}
	.ly-select-filter__item-title {width: 80px;}
	.ly-select-filter__item-title label {font-size: 14px;display: inline-block;color: #999;}
	.ly-select-filter__item-options {flex: 1;border-bottom: 1px dashed var(--el-border-color-light);}
	.ly-select-filter__item:last-of-type .ly-select-filter__item-options {border: 0;}
	.ly-select-filter__no-data {color: #999;}
</style>