<template>
	<div class="ly-upload-file">
		<el-upload
			:disabled="disabled"
			:auto-upload="autoUpload"
			:action="action"
			:name="name"
			:data="data"
			:http-request="request"
			v-model:file-list="defaultFileList"
			:show-file-list="showFileList"
			:drag="drag"
			:accept="accept"
			:multiple="multiple"
			:limit="limit"
			:before-upload="before"
			:on-success="success"
			:on-error="error"
			:on-preview="handlePreview"
			:on-exceed="handleExceed">
			<slot>
				<el-button :type="buttonType" :disabled="disabled" :icon="icon">{{title}}</el-button>
			</slot>
			<template #tip>
				<div v-if="tip" class="el-upload__tip">{{tip}}</div>
			</template>
		</el-upload>
		<span style="display:none!important"><el-input v-model="value"></el-input></span>
	</div>
</template>

<script>
    import {platformsettingsUploadPlatformImg} from "@/api/api";
	export default {
		emits: ['onSuccess','update:modelValue'],
		props: {
			modelValue: { type: [String, Array], default: "" },
            successCode:{ type: Number, default: 2000 },//请求完成代码
			buttonType:{ type: String, default: "primary" },
            title: { type: String, default: "点击上传" },
            icon: { type: String, default: "" },
			tip: { type: String, default: "" },
			action: { type: String, default: "" },
			apiObj: { type: Function, default: platformsettingsUploadPlatformImg },
			name: { type: String, default: 'file' },
			data: { type: Object, default: () => {} },
			accept: { type: String, default: "" },
			maxSize: { type: Number, default: 10 },
			limit: { type: Number, default: 1 },
			autoUpload: { type: Boolean, default: true },
			showFileList: { type: Boolean, default: true },
			drag: { type: Boolean, default: false },
			multiple: { type: Boolean, default: true },
			disabled: { type: Boolean, default: false },
			onSuccess: { type: Function, default: () => { return true } }
		},
		data() {
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
		mounted() {
			this.defaultFileList = Array.isArray(this.modelValue) ? this.modelValue : this.toArr(this.modelValue)
			this.value = this.modelValue
		},
		methods: {
			//默认值转换为数组
			toArr(str){
				var _arr = []
                if(!str){
                    return _arr
                }
				var arr = str.split(",")
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
				return _arr
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
			before(file){
				const maxSize = file.size / 1024 / 1024 < this.maxSize;
				if (!maxSize) {
					this.$message.warning(`上传文件大小不能超过 ${this.maxSize}MB!`);
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
	.el-form-item.is-error .ly-upload-file:deep(.el-upload-dragger) {border-color: var(--el-color-danger);}
	.ly-upload-file {width: 100%;}
	.ly-upload-file:deep(.el-upload-list__item) {transition: none !important;}
</style>