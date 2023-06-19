<template>
    <div class="lycontainer">
        <el-scrollbar>
            <div>
                <ly-growcard :loading="showloading" :rows="2" v-model="growData"></ly-growcard>
            </div>
            <div class="echarts-inner">
                <ly-echartcard :loading="showloading" :rows="3" v-model="growData"></ly-echartcard>
            </div>
            <el-row :gutter="15" class="echartsMaps">
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option1"></lyEcharts>
                    </el-card>
                </el-col>
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option2"></lyEcharts>
                    </el-card>
                </el-col>
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option3"></lyEcharts>
                    </el-card>
                </el-col>
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option4"></lyEcharts>
                    </el-card>
                </el-col>
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option5"></lyEcharts>
                    </el-card>
                </el-col>
                <el-col :lg="8">
                    <el-card shadow="hover">
                        <lyEcharts height="300px" :option="option6"></lyEcharts>
                    </el-card>
                </el-col>
            </el-row>
        </el-scrollbar>
    </div>
</template>

<script>
    import LyGrowcard from "@/components/analysis/growCard";
    import LyEchartcard from "@/components/analysis/echartCard";
    import lyEcharts from '@/components/analysis/lySmallEcharts/index';
    export default {
        name: "analysis",
        components: {LyEchartcard, LyGrowcard,lyEcharts},
        data(){
            return{
                showloading:true,
                growData:[
                    {id:1,title:"访问数",nums:650309,totalnums:896556,icon:{
                            type:"View",
                            background:"#67c23a",
                        },
                        time:{
                            name:"日",
                            type:"success"
                        }},
                    {id:2,title:"订单数",nums:250108,totalnums:365899,icon:{
                            type:"GoodsFilled",
                            background:"#e6a23c",
                        },
                        time:{
                            name:"月",
                            type:"warning"
                        }},
                    {id:3,title:"下载数",nums:356897,totalnums:568952,icon:{
                            type:"Download",
                            background:"#409eff",
                        },
                        time:{
                            name:"周",
                            type:""
                        }},
                    {id:4,title:"成交数",nums:156889,totalnums:956889,icon:{
                            type:"WalletFilled",
                            background:"#f56c6c",
                        },
                        time:{
                            name:"年",
                            type:"danger"
                        }},
                ],
                echartsData:[

                ],
                option1: {
					title: {
						text: 'Bar Demo',
						subtext: '基础柱状图',
					},
					grid: {
						top: '80px'
					},
					tooltip: {
						trigger: 'axis'
					},
					xAxis: {
						type: 'category',
						data: ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期日']
					},
					yAxis: {
						type: 'value'
					},
					series: [{
						data: [120, 200, 150, 80, 70, 110, 150],
						type: 'bar',
						barWidth: '15px',
					},
					{
						data: [110, 180, 120, 120, 60, 90, 110],
						type: 'bar',
						barWidth: '15px',
					}]
				},
                option2: {
					title: {
						text: 'Line Demo',
						subtext: '基础折线图',
					},
					grid: {
						top: '80px'
					},
					tooltip: {
						trigger: 'axis'
					},
					xAxis: {
						type: 'category',
						data: ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期日']
					},
					yAxis: {
						type: 'value'
					},
					series: [{
						data: [120, 200, 150, 80, 70, 110, 130],
						type: 'line',
					},
					{
						data: [110, 180, 120, 120, 60, 90, 110],
						type: 'line',
					}]
				},
				option3: {
					title: {
						text: 'Pie Demo',
						subtext: '基础饼图',
					},
					tooltip: {
						trigger: 'item'
					},
					series: [
						{
							name: '访问来源',
							type: 'pie',
							radius: ['40%', '70%'],
							center: ['50%', '60%'],
							label: false,
							data: [
								{value: 1600, name: '搜索引擎'},
								{value: 680, name: '直接访问'},
								{value: 490, name: '邮件营销'},
								{value: 660, name: '联盟广告'},
								{value: 350, name: '视频广告'}
							]
						}
					]
				},
				option4: {
					title: {
						text: 'Radar Demo',
						subtext: '基础雷达(属性)图',
					},
					tooltip: {
						trigger: 'item'
					},
					radar: {
						radius: 100,
						center: ['50%', '55%'],
						indicator: [
							{ name: '销售', max: 100},
							{ name: '管理', max: 100},
							{ name: '信息技术', max: 100},
							{ name: '客服', max: 100},
							{ name: '研发', max: 100},
							{ name: '市场', max: 100}
						]
					},
					series: [{
						name: "lyadmin",
						type: 'radar',
						areaStyle: {},
						data: [
							{
								value: [64, 80, 90, 55, 80, 40],
							}
						]
					}]
				},
				option5: {
					title: {
						text: 'Kline Demo',
						subtext: '基础K线图',
					},
					grid: {
						top: '80px'
					},
					tooltip: {
						trigger: 'axis',
						axisPointer: {
							type: 'cross'
						}
					},
					xAxis: {
						data: ['2022-11-23', '2022-11-24', '2022-11-25', '2022-11-26', '2022-11-27', '2022-11-28', '2022-11-29', '2022-11-30'],
					},
					yAxis: {
						scale: true,
					},
					series: [{
						type: 'k',
						data: [
							[2213.10,2199.13,2191.45,2224.22],
							[2203.33,2177.55,2173.77,2210.58],
							[2170.22,2174.12,2161.14,2179.65],
							[2179.05,2205.5,2179.05,2222.81],
							[2212.12,2231.17,2212.5,2236.07],
							[2227.22,2235.45,2219.44,2240.21],
							[2242.12,2246.3,2235.42,2255.21],
							[2246.47,2232.45,2221.12,2247.86]
						]
					}]
				},
				option6: {
					title: {
						text: 'Gauge Demo',
						subtext: '基础仪表盘',
					},
					series: [{
						center: ['50%', '60%'],
						type: 'gauge',
						anchor: {
							show: true,
							showAbove: true,
							size: 20,
							itemStyle: {
								borderWidth: 5
							}
						},
						progress: {
							show: true
						},
						data: [{
							value: 80
						}]
					}]
				}
            }
        },
        methods:{
            setFull(){
                window.dispatchEvent(new Event('resize'))
            },
        },
        created() {
            setTimeout(() => {
                this.showloading = false
            }, 600)
        },
    }
</script>
<style lang="scss" scoped>
    .lycontainer{
        /*width: 100%;*/
        height: calc(100vh - 113px); //动态计算长度值
        /*overflow-x: hidden;*/
        /*overflow-y:auto;*/
    }
    .echarts-inner{
        margin-top: 1px;
    }
    ::v-deep(.el-scrollbar__bar.is-horizontal) {
        display: none;
    }
    .echartsMaps{
        margin-top: 10px;
        /*margin-bottom: 40px;*/
        :deep(.is-hover-shadow){
            margin-bottom: 10px;
        }
        :deep(.is-never-shadow){
            margin-bottom: 10px;
        }
    }
</style>