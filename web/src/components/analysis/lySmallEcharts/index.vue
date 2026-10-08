<template>
	<div ref="lyEcharts" :style="{height:height, width:width}"></div>
</template>

<script>
	import * as echarts from 'echarts';
	import { registerLyChartThemes, lyChartTheme } from '@/components/analysis/echartsTheme.js';
	// 全量 echarts 与 echarts/core 主题注册表相互独立，此处需单独注册（v4 亮/暗主题）
	registerLyChartThemes(echarts);
	const unwarp = (obj) => obj && (obj.__v_raw || obj.valueOf() || obj);

	export default {
		...echarts,
		name: "lyEcharts",
		props: {
			height: { type: String, default: "100%" },
			width: { type: String, default: "100%" },
			nodata: {type: Boolean, default: false },
			option: { type: Object, default: () => {} }
		},
		data() {
			return {
				isActivat: false,
				myChart: null
			}
		},
		watch: {
			option: {
				deep:true,
				handler (v) {
					if(this.myChart){
						unwarp(this.myChart).setOption(v);
					}
				}
			}
		},
		computed: {
			myOptions: function() {
				return this.option || {};
			}
		},
		activated(){
			if(!this.isActivat){
				this.$nextTick(() => {
					this.myChart.resize()
				})
			}
		},
		deactivated(){
			this.isActivat = false;
		},
		mounted(){
		    let that = this
			that.isActivat = true;
			setTimeout(() => {
                that.$nextTick(() => {
                    that.draw();
                })
            }, 200)

		},
		methods: {
			draw(){
				var myChart = echarts.init(this.$refs.lyEcharts, lyChartTheme());
				myChart.setOption(this.myOptions);
				this.myChart = myChart;
				window.addEventListener('resize', () => myChart.resize());
			}
		}
	}
</script>