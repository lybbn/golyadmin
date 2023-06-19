<template>
	<div class="lyeditor">
		<Editor v-model="contentValue" :init="init" :disabled="disabled" :placeholder="placeholder" @onClick="onClick" @blur="onHandleBlur" @change="onHandleChange"/>
	</div>
</template>

<script>
    import {url} from "@/api/url";
    import axios from "axios";
	import Editor from '@tinymce/tinymce-vue'
	import tinymce from 'tinymce/tinymce'

	//import 'tinymce/models/dom'

	// 引入编辑器插件
	import 'tinymce/themes/silver' // 编辑器主题，不引入则报错
	import 'tinymce/icons/default' // 引入编辑器图标icon，不引入则不显示对应图标
	// import 'tinymce/skins/ui/oxide/content.css'
	// import 'tinymce/skins/content/default/content.css'
	// // 引入编辑器插件（基本免费插件都在这儿了）
	import 'tinymce/plugins/advlist' // 高级列表
	import 'tinymce/plugins/anchor' // 锚点
	import 'tinymce/plugins/autolink' // 自动链接
	import 'tinymce/plugins/autoresize' // 编辑器高度自适应,注：plugins里引入此插件时，Init里设置的height将失效
	import 'tinymce/plugins/autosave' // 自动存稿
	import 'tinymce/plugins/charmap' // 特殊字符
	import 'tinymce/plugins/code' // 编辑源码
	import 'tinymce/plugins/codesample' // 代码示例
	import 'tinymce/plugins/directionality' // 文字方向
	import 'tinymce/plugins/emoticons' // 表情

	import 'tinymce/plugins/fullpage' // 文档属性
	import 'tinymce/plugins/fullscreen' // 全屏
	import 'tinymce/plugins/help' // 帮助
	import 'tinymce/plugins/hr' // 水平分割线
	import 'tinymce/plugins/image' // 插入编辑图片
	import 'tinymce/plugins/importcss' // 引入css
	import 'tinymce/plugins/insertdatetime' // 插入日期时间
	import 'tinymce/plugins/link' // 超链接
	import 'tinymce/plugins/lists' // 列表插件
	import 'tinymce/plugins/media' // 插入编辑媒体
	import 'tinymce/plugins/nonbreaking' // 插入不间断空格
	import 'tinymce/plugins/pagebreak' // 插入分页符
	import 'tinymce/plugins/paste' // 粘贴插件
	import 'tinymce/plugins/preview'// 预览
	import 'tinymce/plugins/print'// 打印
	import 'tinymce/plugins/quickbars' // 快速工具栏
	import 'tinymce/plugins/save' // 保存
	import 'tinymce/plugins/searchreplace' // 查找替换
	// import 'tinymce/plugins/spellchecker'  //拼写检查，暂未加入汉化，不建议使用
	import 'tinymce/plugins/tabfocus' // 切入切出，按tab键切出编辑器，切入页面其他输入框中
	import 'tinymce/plugins/table' // 表格
	import 'tinymce/plugins/template' // 内容模板
	import 'tinymce/plugins/textcolor' // 文字颜色
	import 'tinymce/plugins/textpattern' // 快速排版
	import 'tinymce/plugins/toc' // 目录生成器
	import 'tinymce/plugins/visualblocks' // 显示元素范围
	import 'tinymce/plugins/visualchars' // 显示不可见字符
	import 'tinymce/plugins/wordcount' // 字数统计
	import 'tinymce/plugins/emoticons/js/emojis';

	import {getToken} from '@/utils/util'

	export default {
		components: {
			Editor
		},
		props: {
			modelValue: {
				type: String,
				default: ""
			},
			placeholder: {
				type: String,
				default: ""
			},
			height: {
				type: Number,
				default: 300,
			},
			disabled: {
				type: Boolean,
				default: false
			},
			plugins: {
				type: [String, Array],
				default: 'code image media link preview table quickbars template pagebreak lists advlist fullscreen'
			},
			toolbar: {
				type: [String, Array],
				default: 'fullscreen undo redo restoredraft | cut copy paste pastetext | forecolor backcolor bold italic underline strikethrough link anchor \ table image | alignleft aligncenter alignright alignjustify outdent indent | \ styleselect fontselect fontsizeselect | bullist numlist | blockquote subscript superscript removeformat | media charmap emoticons hr pagebreak insertdatetime print preview | code selectall searchreplace visualblocks | indent2em lineheight formatpainter axupimgs'
			},
			templates: {
				type: Array,
				default: () => []
			}
		},
		data() {
			return {
				init: {
					language_url: 'static/tinymce/langs/zh_CN.js',
					language: 'zh_CN',
					skin_url: 'static/tinymce/skins/ui/oxide',
					content_css: "static/tinymce/skins/content/default/content.css",
					emoticons_database_url: 'tinymce/plugins/emoticons/js/emojis.js',
					lineheight_formats: '0.5 0.8 1 1.2 1.5 1.75 2 2.5 3 4 5', // 行高配置，也可配置成"12px 14px 16px 20px"这种形式
					menubar: false,
					statusbar: true,
					plugins: this.plugins,
					toolbar: this.toolbar,
					toolbar_mode: 'sliding',
					fontsize_formats: '12px 14px 16px 18px 20px 22px 24px 28px 32px 36px 48px 56px 72px', // 字体大小
					height: this.height,
					placeholder: this.placeholder,
					branding: false,
					resize: 'both',
					elementpath: true,
					content_style: false,
					templates: this.templates,
					quickbars_selection_toolbar: 'forecolor backcolor bold italic underline strikethrough link fontselect fontsizeselect ',
					quickbars_image_toolbar: 'alignleft aligncenter alignright',
					quickbars_insert_toolbar: false,
					image_caption: true,
					image_advtab: true,
                    images_upload_handler:function(blobInfo, success, failure){
                        const params = new FormData()
                        params.append('file', blobInfo.blob())
                        const config = {
                            headers: {
                                'Content-Type': 'multipart/form-data',
                                'Authorization': 'JWT '+getToken(),  // 可选参数(服务器上传验证需要) 如果需要token验证，假设你的token有存放在sessionStorage
                            }
                        }
                        const uploadurl = url + 'platformsettings/uploadplatformimg/'
                        // 图片上传
                        axios.post(uploadurl, params, config).then(res => {
                            if (res.data.code == 2000) {
                                //这里很重要，你图片上传成功后，img的src需要在这里添加，res.path就是你服务器返回的图片链接。
                                let imgpath=''
                                if (res.data.data.data[0].indexOf("://")>=0){
                                    imgpath = res.data.data.data[0]

                                }else{
                                    imgpath = url.split('/api')[0]+res.data.data.data[0]
                                }
                                success(imgpath) // 上传成功，在成功函数里填入图片路径
                                // console.log('[文件上传]', res.data)
                            } else {
                                failure('上传失败')
                            }
                        }).catch(() => {
                            failure('上传出错，服务器开小差了呢')
                        })
                    },
					setup: function(editor) {
						editor.on('init', function() {
							this.getBody().style.fontSize = '14px';
						})
						editor.on('OpenWindow', function(e) {
							//FIX 编辑器在el-drawer中，编辑器的弹框无法获得焦点
							var D = document.querySelector('.el-drawer.open')
							var E = e.target.editorContainer
							if(D && D.contains(E)){
								var nowDA = document.activeElement
								setTimeout(()=>{
									document.activeElement.blur()
									nowDA.focus()
								},0)
							}
						})
					}
				},
				contentValue: this.modelValue
			}
		},
		watch: {
			modelValue(val) {
				this.contentValue = val
			},
			contentValue(val){
				this.$emit('update:modelValue', val);
			}
		},
		mounted() {
			tinymce.init({})
		},
		methods: {
			onClick(e){
				this.$emit('onClick', e, tinymce)
			},
			onHandleBlur(e){
				this.$emit('blur', this.contentValue)
			},
			onHandleChange(e){
				this.$emit('change', this.contentValue)
			}
		}
	}
</script>

<style lang="scss" scoped>
	.lyeditor{
		:deep(.tox-editor-header) {
			z-index:3000 !important;
		}
	}
</style>