<template>
	<div class="ly-aurora" aria-hidden="true">
		<span class="ly-aurora-blob b1"></span>
		<span class="ly-aurora-blob b2"></span>
		<span class="ly-aurora-blob b3"></span>
	</div>
</template>

<script setup>
// 极光光斑背景（纯 CSS 装饰层，可整组件移除，不影响任何功能）
// 强度受 --ly-aurora-opacity 控制（暗色模式自动减弱），prefers-reduced-motion 时静止
</script>

<style lang="scss" scoped>
.ly-aurora {
	position: fixed;
	inset: 0;
	/* -1：垫在所有内容之下（body 提供页面底色，#app 无背景不遮挡），避免光斑蒙在表格/卡片上 */
	z-index: -1;
	overflow: hidden;
	pointer-events: none;
}
.ly-aurora-blob {
	position: absolute;
	border-radius: 50%;
	filter: blur(90px);
	opacity: var(--ly-aurora-opacity);
	will-change: transform;

	&.b1 {
		/* 尺寸随视口伸缩：设计稿小画布上光斑铺满全屏，真实大屏 clamp 保证覆盖比例一致 */
		width: clamp(560px, 50vw, 900px);
		height: clamp(560px, 50vw, 900px);
		left: -160px;
		top: -200px;
		background: radial-gradient(circle, var(--ly-color-primary, #3A7BFF), transparent 65%);
		animation: ly-drift1 52s ease-in-out infinite alternate;
	}
	&.b2 {
		width: clamp(480px, 44vw, 800px);
		height: clamp(480px, 44vw, 800px);
		right: -120px;
		top: 12%;
		background: radial-gradient(circle, #6C9BFF, transparent 65%);
		animation: ly-drift2 64s ease-in-out infinite alternate;
	}
	&.b3 {
		/* 青色光斑收敛向蓝调 + 单独降透明度，避免大屏中央高饱和色块 */
		width: clamp(440px, 36vw, 620px);
		height: clamp(440px, 36vw, 620px);
		left: 32%;
		bottom: -220px;
		background: radial-gradient(circle, #58B8E8, transparent 65%);
		opacity: calc(var(--ly-aurora-opacity) * 0.55);
		animation: ly-drift3 58s ease-in-out infinite alternate;
	}
}
@keyframes ly-drift1 {
	to { transform: translate(140px, 90px) scale(1.12); }
}
@keyframes ly-drift2 {
	to { transform: translate(-110px, 70px) scale(0.92); }
}
@keyframes ly-drift3 {
	to { transform: translate(80px, -90px) scale(1.08); }
}
@media (prefers-reduced-motion: reduce) {
	.ly-aurora-blob { animation: none; }
}
</style>
