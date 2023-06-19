<!--
 * @Descripttion: 表格选择器组件
 * @version: 1.0
 * @Program: django-vue-lyadmin
 * @Author: lybbn
 * @Date: 2023/02/11
-->

<template>
	<el-select ref="select" v-model="defaultValue" :size="size" :clearable="clearable" :multiple="multiple" :collapse-tags="collapseTags" :collapse-tags-tooltip="collapseTagsTooltip" :filterable="filterable" :placeholder="placeholder" :disabled="disabled" :filter-method="filterMethod" @remove-tag="removeTag" @visible-change="visibleChange" @clear="clear" style="width: 100%">
		<template #empty>
			<div class="ly-table-select__table" :style="{width: tableWidth+'px'}" v-loading="loading">
				<div class="ly-table-select__header">
					<slot name="header" :form="formData" :submit="formSubmit"></slot>
				</div>
				<el-table ref="table" border :data="tableData" :height="350" :highlight-current-row="!multiple" @row-click="click" @select="select" @select-all="selectAll">
					<el-table-column v-if="multiple" type="selection" width="45"></el-table-column>
					<el-table-column v-else type="index" width="45">
						<template #default="scope">
							<span v-text="getIndex(scope.$index)"></span>
						</template>
					</el-table-column>
					<slot></slot>
				</el-table>
				<div class="ly-table-select__page">
					<Pagination :small="true" v-bind:child-msg="pageparm"  @callFather="callFather"></Pagination>
				</div>
			</div>
		</template>
	</el-select>
</template>

<script>
	import Pagination from "@/components/Pagination";
	export default {
		components: {Pagination},
		props: {
			modelValue: null,
			successCode:{ type: Number, default: 2000 },//网络请求完成代码
			apiObj: { type: Function, default: null},//网络请求的实例
			params: { type: Object, default: () => {} },//网络请求的额外参数
			placeholder: { type: String, default: "请选择" },
			size: { type: String, default: "default" },
			clearable: { type: Boolean, default: false },
			multiple: { type: Boolean, default: false },
			filterable: { type: Boolean, default: false },
			collapseTags: { type: Boolean, default: false },
			collapseTagsTooltip: { type: Boolean, default: false },
			disabled: { type: Boolean, default: false },
			tableWidth: {type: Number, default: 500},
			mode: { type: String, default: "popover" },
			props: { type: Object, default: () => {} }
		},
		data() {
			return {
				loading: false,
				defaultValue: [],
				tableData: [],
				defaultProps: {
					label: 'label',
					value: 'value',
					page: 'page',
					pageSize:'limit',
				},
				formInline:{
                    page: 1,
                    limit: 10,
                },
                pageparm: {
                    page: 1,
                    limit: 10,
                    total: 0
                },
				formData: {}
			}
		},
		computed: {

		},
		watch: {
			modelValue:{
				handler(){
					this.defaultValue = this.modelValue
					this.autoCurrentLabel()
				},
				deep: true
			}
		},
		mounted() {
			this.defaultProps = Object.assign(this.defaultProps, this.props);
			this.defaultValue =  this.modelValue
			this.autoCurrentLabel()
		},
		methods: {
			// 表格序列号
            getIndex($index) {
                // (当前页 - 1) * 当前显示数据条数 + 当前行数据的索引 + 1
                return (this.pageparm.page-1)*this.pageparm.limit + $index +1
            },
			callFather(parm) {
                this.formInline.page = parm.page
                this.formInline.limit = parm.limit
                this.getData()
            },
			checkDefaultValue(){
            	if(this.multiple && this.defaultValue == undefined){
            		this.defaultValue = []
				}
			},
			//表格显示隐藏回调
			visibleChange(visible){
				if(visible){
					this.formInline.page = 1
					this.formData = {}
					this.getData()
				}else{
					this.autoCurrentLabel()
				}
			},
			//获取表格数据
			async getData(){
				this.loading = true;
				var reqData = {
					[this.defaultProps.page]: this.formInline.page,
					[this.defaultProps.pageSize]: this.formInline.limit,
				}
				Object.assign(reqData, this.params, this.formData)
				this.apiObj()(reqData).then(res=>{
					this.loading = false;
					if(res.code == this.successCode) {
						this.tableData = res.data.data
						 this.pageparm.page = res.data.page;
						 this.pageparm.limit = res.data.limit;
						 this.pageparm.total = res.data.total;
					} else {
						this.$message.warning(res.msg)
					}
					//表格默认赋值
					this.$nextTick(() => {
						if(this.multiple){
							this.checkDefaultValue()
							this.defaultValue.forEach(row => {
								var setrow = this.tableData.filter(item => item[this.defaultProps.value]===row[this.defaultProps.value] )
								if(setrow.length > 0){
									this.$refs.table.toggleRowSelection(setrow[0], true);
								}
							})
						}else{
							var setrow = this.tableData.filter(item => item[this.defaultProps.value]===this.defaultValue[this.defaultProps.value] )
							this.$refs.table.setCurrentRow(setrow[0]);
						}
						this.$refs.table.setScrollTop(0)
					})
				});
			},
			//插糟表单提交
			formSubmit(){
				this.formInline.page = 1
				this.getData()
			},
			//自动模拟options赋值
			autoCurrentLabel(){
				this.$nextTick(() => {
					if(this.multiple){
						this.$refs.select.selected.forEach(item => {
							item.currentLabel = item.value[this.defaultProps.label]
						})
					}else{
						this.$refs.select.selectedLabel = this.defaultValue[this.defaultProps.label]
					}
				})
			},
			//表格勾选事件
			select(rows, row){
            	this.checkDefaultValue()
				var isSelect = rows.length && rows.indexOf(row) !== -1
				if(isSelect){
					this.defaultValue.push(row)
				}else{
					this.defaultValue.splice(this.defaultValue.findIndex(item => item[this.defaultProps.value] == row[this.defaultProps.value]), 1)
				}
				this.autoCurrentLabel()
				this.$emit('update:modelValue', this.defaultValue);
				this.$emit('change', this.defaultValue);
			},
			//表格全选事件
			selectAll(rows){
            	this.checkDefaultValue()
				var isAllSelect = rows.length > 0
				if(isAllSelect){
					rows.forEach(row => {
						var isHas = this.defaultValue.find(item => item[this.defaultProps.value] == row[this.defaultProps.value])
						if(!isHas){
							this.defaultValue.push(row)
						}
					})
				}else{
					this.tableData.forEach(row => {
						var isHas = this.defaultValue.find(item => item[this.defaultProps.value] == row[this.defaultProps.value])
						if(isHas){
							this.defaultValue.splice(this.defaultValue.findIndex(item => item[this.defaultProps.value] == row[this.defaultProps.value]), 1)
						}
					})
				}
				this.autoCurrentLabel()
				this.$emit('update:modelValue', this.defaultValue);
				this.$emit('change', this.defaultValue);
			},
			click(row){
				if(this.multiple){
					//处理多选点击行
				}else{
					this.defaultValue = row
					this.$refs.select.blur()
					this.autoCurrentLabel()
					this.$emit('update:modelValue', this.defaultValue);
					this.$emit('change', this.defaultValue);
				}
			},
			//tags删除后回调
			removeTag(tag){
				var row = this.findRowByKey(tag[this.defaultProps.value])
				this.$refs.table.toggleRowSelection(row, false);
				this.$emit('update:modelValue', this.defaultValue);
			},
			//清空后的回调
			clear(){
				this.$emit('update:modelValue', this.defaultValue);
			},
			// 关键值查询表格数据行
			findRowByKey (value) {
				return this.tableData.find(item => item[this.defaultProps.value] === value)
			},
			filterMethod(){
				this.getData()
			},
			// 触发select隐藏
			blur(){
				this.$refs.select.blur();
			},
			// 触发select显示
			focus(){
				this.$refs.select.focus();
			}
		}
	}
</script>

<style scoped>
	.ly-table-select__table {padding:12px;}
	.ly-table-select__page {padding-top: 2px;}
</style>