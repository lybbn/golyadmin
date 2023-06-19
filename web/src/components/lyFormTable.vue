<template>
	<div class="ly-form-table" ref="lyFormTable">
		<el-table :data="data" ref="table" border stripe :row-key="rowKey">
			<el-table-column type="index" width="50" fixed="left">
				<template #header>
					<el-button type="primary" icon="plus" size="small" circle @click="rowAdd"></el-button>
				</template>
				<template #default="scope">
					<div class="ly-form-table-handle">
						<span>{{scope.$index + 1}}</span>
						<el-button type="danger" icon="delete" size="small" plain circle @click="rowDel(scope.row, scope.$index)"></el-button>
					</div>
				</template>
			</el-table-column>
			<el-table-column label="" width="50" v-if="dragSort">
				<template #default>
					<div class="move" style="cursor: move;"><el-icon><DCaret /></el-icon></div>
				</template>
			</el-table-column>
			<slot></slot>
			<template #empty>
				{{placeholder}}
			</template>
		</el-table>
	</div>
</template>

<script>
	import Sortable from 'sortablejs'

	export default {
		props: {
			modelValue: { type: Array, default: () => [] },
			addTemplate: { type: Object, default: () => {} },
			placeholder: { type: String, default: "暂无数据" },
			dragSort: { type: Boolean, default: false },
			rowKey: { type: String, default: "id" },
		},
		data(){
			return {
				data: []
			}
		},
		mounted(){
			this.data = this.modelValue
			if(this.dragSort){
				this.rowDrop()
			}
		},
		watch:{
			modelValue(){
				this.data = this.modelValue
			},
			data: {
				handler(){
					this.$emit('update:modelValue', this.data);
				},
				deep: true
			}
		},
		methods: {
			rowDrop(){
				const _this = this
				const tbody = this.$refs.table.$el.querySelector('.el-table__body-wrapper tbody')
				Sortable.create(tbody, {
					handle: ".move",
					animation: 300,
					ghostClass: "ghost",
					onEnd({ newIndex, oldIndex }) {
						_this.data.splice(newIndex, 0, _this.data.splice(oldIndex, 1)[0])
						const newArray = _this.data.slice(0)
						const tmpHeight = _this.$refs.lyFormTable.offsetHeight
						_this.$refs.lyFormTable.style.setProperty('height', tmpHeight + 'px')
						_this.data = []
						_this.$nextTick(() => {
							_this.data = newArray
							_this.$nextTick(() => {
								_this.$refs.lyFormTable.style.removeProperty('height')
							})

						})
					}
				})
			},
			rowAdd(){
				const temp = JSON.parse(JSON.stringify(this.addTemplate))
				this.data.push(temp)
			},
			rowDel(row, index){
				this.data.splice(index, 1)
			}
		}
	}
</script>

<style scoped>
	.ly-form-table {width: 100%;}
	.ly-form-table .ly-form-table-handle {text-align: center;}
	.ly-form-table .ly-form-table-handle span {display: inline-block;}
	.ly-form-table .ly-form-table-handle button {display: none;}
	.ly-form-table .hover-row .ly-form-table-handle span {display: none;}
	.ly-form-table .hover-row .ly-form-table-handle button {display: inline-block;}
	.ly-form-table .move {text-align: center;font-size: 14px;margin-top: 3px;}
</style>