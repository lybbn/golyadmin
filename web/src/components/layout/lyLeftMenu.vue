<template>
	<div v-if="navMenus.length<=0" style="padding:20px;">
		<el-alert title="无菜单" center type="info" :closable="false"></el-alert>
	</div>
	<template v-for="menu in navMenus" :key="menu.id">
		<el-menu-item v-if="!menu.hasChildren" :index="menu.attributes.url?'/'+menu.attributes.url:menu.id" :key="menu.id">
			<a v-if="isLink(menu.attributes.url)" :href="menu.attributes.url" target="_blank" @click.stop='()=>{}'></a>
			<svg-icon :icon-class="menu.attributes.icon?menu.attributes.icon:'Menu'"></svg-icon>
			<template #title>{{menu.text}}</template>
		</el-menu-item>
		<el-sub-menu v-else :index="menu.attributes.url?'/'+menu.attributes.url:menu.id" :key="menu.id">
			<template #title >
                <svg-icon :icon-class="menu.attributes.icon?menu.attributes.icon:'Menu'"></svg-icon>
                <span>{{menu.text}}</span>
            </template>
			<lyLeftMenu :navMenus="menu.children"></lyLeftMenu>
		</el-sub-menu>
	</template>
</template>

<script setup>
    const props = defineProps({
        navMenus :{
            type:Array,
            default:[]
        },
    })
    function isLink(str) {
        let reg = /(https|http):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/
        if (!reg.test(str)) {
            return false
        }
        return true
    }
</script>

<style lang="scss" scoped>
	.el-menu.el-menu--horizontal{
    	border-bottom: 0;
  	}
	.menu-nav-title{
		height: 56px;
		line-height: 56px;
		padding-left: 20px;
		font-size: 20px;
		font-weight: bold;
		background: #eff6ff;
		color: var(--el-color-primary);
		border-bottom: 1px solid #c9e0ff;
	}
	.el-menu-vertical-demo:not(.el-menu--collapse) {
		width: 100%;
		min-height: 400px;
	}
	.el-menu-vertical-demo{
		i{
			margin-right: 5px;
			font-size: 18px;
		}
	}

	.el-menu-vertical-demo:not(.el-menu--collapse) {
		border: none;
		text-align: left;
	}

	::v-deep(.el-menu-item-group__title) {
		padding: 0px !important;
	}
	.el-menu-item.is-active {
		position: relative;
		/*background-color: rgb(48, 54, 62) !important;*/
		background-color: var(--l-main-sidebar-menu-active-bg) !important;
		&:before{
			width: 2px;
			height: 100%;
			position: absolute;
			left: 0;
			top: 0;
			background: var(--l-main-sidebar-menu-hover-bg);
			display: block;
		}
	}

</style>