<!--
 * @Descripttion: 弹窗扩展组件
 * @Program: djangolyadmin
 * @version: 1.1
 * @Author: lybbn
 * @Date: 2023年04月11日
-->
<template>
  <div class="ly-dialog">
    <el-dialog
      v-model="visible"
      :close-on-click-modal="closeOnClickModal"
      :title="title"
      :width="width"
      :top="top"
      :fullscreen="screeFull"
      :center="center"
      :before-close="beforeClose"
      :append-to-body="appendToBody"
      :destroy-on-close="true"
      :draggable="draggable"
      :show-close="showClose"
      @closed="closed">
      <template #header="{ close, titleId, titleClass }">
        <div>
          <slot name="header">
            <span :id="titleId" :class="titleClass">{{ title }}</span>
          </slot>
          <div class="ly-dialog__headerbtn">
            <button aria-label="fullscreen" type="button" @click="handleFullScreenClick">
              <el-icon v-if="screeFull" class="el-dialog__close"><Minus /></el-icon>
              <el-icon v-else class="el-dialog__close"><full-screen /></el-icon>
            </button>
            <button aria-label="close" type="button" @click="close">
              <el-icon class="el-dialog__close"><close /></el-icon>
            </button>
          </div>
        </div>
      </template>
      <div v-loading="loading">
        <slot></slot>
      </div>
      <template v-if="$slots.footer" #footer>
        <slot name="footer"></slot>
      </template>
    </el-dialog>
  </div>
</template>

<script>
    import 'element-plus/es/components/dialog/style/css'
    export default {
        name: 'LyDialog',
        data(){
            return{
                visible:false,
                screeFull:false
            }
        },
        props: {
            title: {
              type: String,
              default: ''
            },
            modelValue: {
              type: Boolean,
              default: true
            },
            width: {
              type: String,
              default: '50%'
            },
            center: {
              type: Boolean,
              default: false
            },
            top: {
              type: String,
              default: '10vh'
            },
            draggable: {
              type: Boolean,
              default: true
            },
            appendToBody: {
              type: Boolean,
              default: false
            },
            closeOnClickModal: {
              type: Boolean,
              default: false
            },
            fullscreen: {
              type: Boolean,
              default: false
            },
            showClose: {
              type: Boolean,
              default: false
            },
            loading: {
              type: Boolean,
              default: false
            },
            beforeClose:Function// 关闭回调函数
        },
        watch:{
            modelValue: function(nval){
                this.visible = nval; // modelValue改变是同步子组件visible的值
            },
            fullscreen:function(nval) {
                this.screeFull = nval
            },
        },
        mounted() {
            this.screeFull = this.fullscreen
            this.visible = this.modelValue
        },
        methods:{
            openDialog() {
                this.visible = true
            },
            closeDialog() {
                this.visible = false
            },
            closed() {
                this.$emit('closed')
            },
            handleFullScreenClick(){
                this.screeFull = !this.screeFull
            }
        }
    }
</script>

<style scoped>
  .ly-dialog__headerbtn {
      position: absolute;
      top: var(--el-dialog-padding-primary);
      right: var(--el-dialog-padding-primary);
  }
  .ly-dialog__headerbtn button {
      padding: 0;
      background: transparent;
      border: none;
      outline: none;
      cursor: pointer;
      font-size: var(--el-message-close-size,16px);
      margin-left: 15px;
      color: var(--el-color-info);
  }
  .ly-dialog__headerbtn button:hover .el-dialog__close {
      color: var(--el-color-primary);
  }
  .ly-dialog:deep(.el-dialog).is-fullscreen {
      display: flex;
      flex-direction: column;
      top:0px !important;
      left:0px !important;
  }
  .ly-dialog:deep(.el-dialog).is-fullscreen .el-dialog__header {
      border-bottom:var(--el-border);
      margin-right:0 !important;
  }
  .ly-dialog:deep(.el-dialog).is-fullscreen .el-dialog__body {
      flex:1;
      overflow: auto;
  }
  .ly-dialog:deep(.el-dialog).is-fullscreen .el-dialog__footer {
      padding-bottom: 10px;
      border-top:var(--el-border);
  }
</style>