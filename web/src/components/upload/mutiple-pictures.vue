<template>
	<div class="ly-upload-multiple">
		<el-upload ref="lyuploader" list-type="picture-card"
			:auto-upload="autoUpload"
			:disabled="disabled"
			:action="action"
			:name="name"
			:data="data"
			:http-request="request"
			v-model:file-list="defaultFileList"
			:show-file-list="showFileList"
			:accept="accept"
			:multiple="multiple"
			:limit="limit"
			:before-upload="before"
			:on-success="success"
			:on-error="error"
			:on-preview="handlePreview"
			:on-exceed="handleExceed">
			<slot>
				<el-icon><Plus/></el-icon>
			</slot>
			<template #tip>
				<div v-if="tip" class="el-upload__tip">{{tip}}</div>
			</template>
			<template #file="{ file }">
				<div class="ly-upload-list-item">
					<el-image class="el-upload-list__item-thumbnail" :src="file.url" fit="cover" :preview-src-list="preview" :initial-index="preview.findIndex(n=>n==file.url)" hide-on-click-modal append-to-body :z-index="9999">
						<template #placeholder>
							<div class="ly-upload-multiple-image-slot">
								Loading...
							</div>
						</template>
					</el-image>
					<div v-if="!disabled && file.status=='success'" class="ly-upload__item-actions">
						<span class="del" @click="handleRemove(file)"><el-icon><Delete /></el-icon></span>
					</div>
					<div v-if="file.status=='ready' || file.status=='uploading'" class="ly-upload__item-progress">
						<el-progress :percentage="file.percentage" :text-inside="true" :stroke-width="16"/>
					</div>
				</div>
			</template>
		</el-upload>
		<span style="display:none!important"><el-input v-model="value"></el-input></span>
	</div>
</template>

<script>
	import Sortable from 'sortablejs'
    import {platformsettingsUploadPlatformImg} from "@/api/api";

	export default {
		emits: ['onSuccess','update:modelValue'],
		props: {
			modelValue: { type: [String, Array], default: "" },
            successCode:{ type: Number, default: 2000 },//请求完成代码
			tip: { type: String, default: "" },
			action: { type: String, default: "" },
			apiObj: { type: Function, default: platformsettingsUploadPlatformImg },//上传请求API对象
			name: { type: String, default: 'file' },//form请求时文件的key
			data: { type: Object, default: () => {} },
			accept: { type: String, default: "image/gif, image/jpeg, image/png" },
			maxSize: { type: Number, default: 10 },//最大文件大小 默认10MB
			limit: { type: Number, default: 0 },
			autoUpload: { type: Boolean, default: true },
			showFileList: { type: Boolean, default: true },
			multiple: { type: Boolean, default: false },
			disabled: { type: Boolean, default: false },
			draggable: { type: Boolean, default: false },
			onSuccess: { type: Function, default: () => { return true } }
		},
		data(){
			return {
				value: "",
				defaultFileList: []
			}
		},
		watch:{
			modelValue(val){
				if(Array.isArray(val)){
					if (JSON.stringify(val) != JSON.stringify(this.formatArr(this.defaultFileList))) {
						this.defaultFileList = val
						this.value = val
					}
				}else{
					if (val != this.toStr(this.defaultFileList)) {
						this.defaultFileList = this.toArr(val)
						this.value = val
					}
				}
			},
			defaultFileList: {
				handler(val){
					this.$emit('update:modelValue', Array.isArray(this.modelValue) ? this.formatArr(val) : this.toStr(val))
					this.value = this.toStr(val)
				},
				deep: true
			}
		},
		computed: {
			preview(){
				return this.defaultFileList.map(v => v.url)
			}
		},
		mounted() {
			this.defaultFileList = Array.isArray(this.modelValue) ? this.modelValue : this.toArr(this.modelValue)
			this.value = this.modelValue
			if(!this.disabled && this.draggable){
				this.$nextTick(() => {
					this.rowDrop();
				})
			}
		},
		methods: {
			//默认值转换为数组
			toArr(str){
				var _arr = [];
				if(!str){
				    return _arr;
                }
				var arr = str.split(",");
				arr.forEach(item => {
					if(item){
						var urlArr = item.split('/');
						var fileName = urlArr[urlArr.length - 1]
						_arr.push({
							name: fileName,
							url: item
						})
					}
				})
				return _arr;
			},
			//数组转换为原始值
			toStr(arr){
				return arr.map(v => v.url).join(",")
			},
			//格式化数组值
			formatArr(arr){
				var _arr = []
				arr.forEach(item => {
					if(item){
						_arr.push({
							name: item.name,
							url: item.url
						})
					}
				})
				return _arr
			},
			//拖拽
			rowDrop(){
				const _this = this
				const itemBox = this.$refs.lyuploader.$el.querySelector('.el-upload-list')
				if(itemBox){
					Sortable.create(itemBox, {
						handle: ".el-upload-list__item",
						animation: 200,
						ghostClass: "ghost",
						onEnd({ newIndex, oldIndex }) {
							const tableData = _this.defaultFileList
							const currRow = tableData.splice(oldIndex, 1)[0]
							tableData.splice(newIndex, 0, currRow)
						}
					})
				}
			},
			before(file){
				if(!['image/jpeg','image/png','image/gif'].includes(file.type)){
					this.$message.warning(`选择的文件类型 ${file.type} 非图像类文件`);
					return false;
				}
				const maxSize = file.size / 1024 / 1024 < this.maxSize;
				if (!maxSize) {
					this.$message.warning(`上传图片大小不能超过 ${this.maxSize}MB!`);
					return false;
				}
			},
			success(res, file){
			    if(res){
			        var os = this.onSuccess(res, file)
                    if(os!=undefined && os==false){
                        return false
                    }
                    let src=''
                    if (res.data.data[0].indexOf("://")>=0){
                        src = res.data.data[0]

                    }else{
                        src = url.split('/api')[0]+res.data.data[0]
                    }
                    file.url = src
                }
			    this.$emit('onSuccess',this.value)
			},
			error(err){
				this.$notify.error({
					title: '上传文件未成功',
					message: err
				})
			},
			beforeRemove(uploadFile){
				return this.$confirm(`是否移除 ${uploadFile.name} ?`, '提示', {
					type: 'warning',
				}).then(() => {
					return true
				}).catch(() => {
					return false
				})
			},
			handleRemove(file){
				this.$refs.lyuploader.handleRemove(file)
				//this.defaultFileList.splice(this.defaultFileList.findIndex(item => item.uid===file.uid), 1)
			},
			handleExceed(){
				this.$message.warning(`当前设置最多上传 ${this.limit} 个文件，请移除后上传!`)
			},
			handlePreview(uploadFile){
				window.open(uploadFile.url)
			},
			async request(param){
			    var vm = this
				var apiObj = vm.apiObj;
                let obj= await apiObj(param)
                if(obj.code == vm.successCode) {
                    param.onSuccess(obj)
                } else {
                    param.onError(obj.msg || "未知错误")
                }
			}
		}
	}
</script>

<style scoped>
	.el-form-item.is-error .ly-upload-multiple:deep(.el-upload--picture-card) {border-color: var(--el-color-danger);}
	:deep(.el-upload-list__item) {transition:none;border-radius: 0;}
	.ly-upload-multiple:deep(.el-upload-list__item.el-list-leave-active) {position: static!important;}
	.ly-upload-multiple:deep(.el-upload--picture-card) {border-radius: 0;}
	.ly-upload-list-item {width: 100%;height: 100%;position: relative;}
	.ly-upload-multiple .el-image {display: block;}
	.ly-upload-multiple .el-image:deep(img) {-webkit-user-drag: none;}
	.ly-upload-multiple-image-slot {display: flex;justify-content: center;align-items: center;width: 100%;height: 100%;font-size: 12px;}
	.ly-upload-multiple .el-upload-list__item:hover .ly-upload__item-actions{display: block;}
	.ly-upload__item-actions {position: absolute;top:0;right: 0;display: none;}
	.ly-upload__item-actions span {display: flex;justify-content: center;align-items: center;;width: 25px;height:25px;cursor: pointer;color: #fff;}
	.ly-upload__item-actions span i {font-size: 12px;}
	.ly-upload__item-actions .del {background: #F56C6C;}
	.ly-upload__item-progress {position: absolute;width: 100%;height: 100%;top: 0;left: 0;background-color: var(--el-overlay-color-lighter);}
</style>