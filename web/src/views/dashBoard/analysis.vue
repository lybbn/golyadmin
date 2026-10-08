<template>
    <div class="lycontainer">
        <el-scrollbar>
            <div>
                <ly-growcard :loading="showloading" :rows="2" v-model="growData"></ly-growcard>
            </div>
            <div class="echarts-inner">
                <ly-echartcard :loading="showloading" :rows="3" v-model="growData"></ly-echartcard>
            </div>
            <div class="dashboard-lower">
                <div class="dashboard-lower-top">
                    <section v-if="visibleShortcuts.length" class="dashboard-panel">
                        <div class="panel-heading">
                            <div>
                                <h2>常用入口</h2>
                                <p>快速进入常用管理功能</p>
                            </div>
                        </div>
                        <div class="shortcut-grid">
                            <button
                                v-for="item in visibleShortcuts"
                                :key="item.name"
                                class="shortcut-item"
                                type="button"
                                @click="openShortcut(item.name)"
                            >
                                <span class="shortcut-icon" :style="{'--shortcut-color': item.color}">
                                    <el-icon><component :is="item.icon" /></el-icon>
                                </span>
                                <span class="shortcut-copy">
                                    <strong>{{ item.title }}</strong>
                                    <small>{{ item.description }}</small>
                                </span>
                                <el-icon class="shortcut-arrow"><ArrowRight /></el-icon>
                            </button>
                        </div>
                    </section>

                    <section v-if="canViewServer" class="dashboard-panel system-panel">
                        <div class="panel-heading">
                            <div>
                                <h2>系统运行状态</h2>
                                <p>{{ systemInfo?.system || '服务器资源概况' }}</p>
                            </div>
                            <el-tag v-if="systemInfo" :type="systemHealth.type" effect="light" round>
                                <span class="status-dot" :class="'is-' + systemHealth.type"></span>{{ systemHealth.text }}
                            </el-tag>
                            <el-tag v-else-if="systemInfoError" type="info" effect="light" round>暂不可用</el-tag>
                            <el-tag v-else effect="light" round>获取中</el-tag>
                        </div>
                        <div v-if="systemInfo" class="system-metrics">
                            <div v-for="metric in systemMetrics" :key="metric.label" class="system-metric">
                                <div class="metric-label">
                                    <span>{{ metric.label }}</span>
                                    <strong :class="'is-' + metric.level">{{ metric.value.toFixed(1) }}%</strong>
                                </div>
                                <div class="metric-track">
                                    <span :class="'is-' + metric.level" :style="{width: metric.value + '%'}"></span>
                                </div>
                            </div>
                            <div class="system-meta">
                                <span><el-icon><Timer /></el-icon> 已运行 {{ systemInfo.time || '—' }}</span>
                                <span><el-icon><Monitor /></el-icon> {{ systemInfo.is_windows ? 'Windows' : 'Linux / Unix' }}</span>
                            </div>
                        </div>
                        <div v-else class="panel-empty">
                            {{ systemInfoError ? '无法读取服务器状态，请稍后重试。' : '正在读取服务器状态…' }}
                        </div>
                    </section>
                </div>

                <section v-if="canViewOperations" class="dashboard-panel operations-panel">
                    <div class="panel-heading">
                        <div>
                            <h2>最近操作</h2>
                            <p>系统记录的最新管理操作</p>
                        </div>
                        <el-button text type="primary" @click="openShortcut('journalManage')">
                            查看全部<el-icon class="el-icon--right"><ArrowRight /></el-icon>
                        </el-button>
                    </div>
                    <div v-if="recentOperations.length" class="operation-list">
                        <div v-for="item in recentOperations" :key="item.id" class="operation-row">
                            <span class="method-badge" :class="methodClass(item.method)">{{ item.method || 'API' }}</span>
                            <div class="operation-copy">
                                <strong>{{ item.path || '—' }}</strong>
                                <span>{{ item.user?.username || '系统' }} · {{ formatDateTime(item.created_at) }}</span>
                            </div>
                            <el-tag :type="Number(item.code) < 400 ? 'success' : 'warning'" effect="light" size="small">
                                {{ item.code || '—' }}
                            </el-tag>
                        </div>
                    </div>
                    <div v-else class="panel-empty">
                        {{ operationsLoading ? '正在读取操作记录…' : '暂无操作记录' }}
                    </div>
                </section>
            </div>
        </el-scrollbar>
    </div>
</template>

<script>
    import LyGrowcard from "@/components/analysis/growCard.vue";
    import LyEchartcard from "@/components/analysis/echartCard.vue";
    import {monitorGetSystemInfo, systemOperationlog} from "@/api/api";
    import {formatDateTime, hasPermission} from "@/utils/util";
    import {useMutitabsStore} from "@/store/mutitabs";
    export default {
        name: "analysis",
        components: {LyEchartcard, LyGrowcard},
        setup() {
            return {mutitabsStore: useMutitabsStore()}
        },
        data(){
            return{
                showloading:true,
                systemInfo:null,
                systemInfoError:false,
                recentOperations:[],
                operationsLoading:false,
                shortcutItems:[
                    {name:'adminManage',title:'管理员管理',description:'维护后台账号',icon:'User',color:'#3a7bff'},
                    {name:'roleManage',title:'角色管理',description:'配置角色权限',icon:'Key',color:'#7967e8'},
                    {name:'menuManage',title:'菜单管理',description:'维护导航菜单',icon:'Menu',color:'#29a69a'},
                    {name:'server',title:'服务监控',description:'查看资源状态',icon:'Monitor',color:'#e6a23c'},
                ],
                growData:[
                    {id:1,title:"访问数",nums:650309,totalnums:896556,icon:{
                            type:"View",
                            background:"#3a7bff",
                        },
                        time:{
                            name:"日",
                            type:"primary"
                        }},
                    {id:2,title:"订单数",nums:250108,totalnums:365899,icon:{
                            type:"GoodsFilled",
                            background:"#5272e7",
                        },
                        time:{
                            name:"月",
                            type:"primary"
                        }},
                    {id:3,title:"下载数",nums:356897,totalnums:568952,icon:{
                            type:"Download",
                            background:"#398ecf",
                        },
                        time:{
                            name:"周",
                            type:"primary"
                        }},
                    {id:4,title:"成交数",nums:156889,totalnums:956889,icon:{
                            type:"WalletFilled",
                            background:"#345fc9",
                        },
                        time:{
                            name:"年",
                            type:"primary"
                        }},
                ],
                echartsData:[

                ],
            }
        },
        computed:{
            canViewServer(){
                return hasPermission('server','Search')
            },
            canViewOperations(){
                return hasPermission('journalManage','Search')
            },
            visibleShortcuts(){
                return this.shortcutItems.filter(item=>hasPermission(item.name,'Search'))
            },
            systemMetrics(){
                if(!this.systemInfo) return []
                // 三档语义色：>=85% 资源紧张(红) / >=60% 略高(黄) / 正常(绿)，与监控告警惯例一致
                const levelOf = v => v >= 85 ? 'danger' : v >= 60 ? 'warning' : 'success'
                const cpu = this.metricPercent(this.systemInfo.cpu?.[0])
                const memory = this.metricPercent(this.systemInfo.mem?.percent)
                const disk = this.metricPercent(this.systemInfo.disk?.[0]?.size?.[3])
                return [
                    {label:'CPU',value:cpu,level:levelOf(cpu)},
                    {label:'内存',value:memory,level:levelOf(memory)},
                    {label:'磁盘',value:disk,level:levelOf(disk)},
                ]
            },
            // 整体健康状态：取 CPU/内存/磁盘最差一档，联动顶部标签
            systemHealth(){
                const levels = this.systemMetrics.map(m=>m.level)
                if(levels.includes('danger')) return {type:'danger',text:'资源紧张'}
                if(levels.includes('warning')) return {type:'warning',text:'部分偏高'}
                return {type:'success',text:'运行正常'}
            }
        },
        methods:{
            setFull(){
                window.dispatchEvent(new Event('resize'))
            },
            metricPercent(value){
                const parsed = Number(value)
                return Number.isFinite(parsed) ? Math.max(0,Math.min(100,parsed)) : 0
            },
            loadDashboardDetails(){
                if(this.canViewServer){
                    monitorGetSystemInfo({}).then(res=>{
                        if(res.code === 2000) this.systemInfo = res.data
                        else this.systemInfoError = true
                    }).catch(()=>{
                        this.systemInfoError = true
                    })
                }
                if(this.canViewOperations){
                    this.operationsLoading = true
                    systemOperationlog({page:1,limit:5}).then(res=>{
                        if(res.code === 2000) this.recentOperations = res.data?.data || []
                    }).catch(()=>{
                        this.recentOperations = []
                    }).finally(()=>{
                        this.operationsLoading = false
                    })
                }
            },
            openShortcut(name){
                this.mutitabsStore.switchtab(name)
            },
            formatDateTime,
            methodClass(method){
                return `method-${String(method || '').toLowerCase()}`
            },
        },
        created() {
            this.loadDashboardDetails()
            setTimeout(() => {
                this.showloading = false
            }, 600)
        },
    }
</script>
<style lang="scss" scoped>
    .lycontainer{
        /*width: 100%;*/
        /* 骨架账本 120 = 顶栏 60 + 标签条净高 40 + 内容区上下 padding 20，必须随骨架变化同步
           （旧值 113 按原版标签条 33px 校准，v4 标签条净高 40px 后差 7px 会冒空滚动条） */
        height: calc(100vh - 120px); //动态计算长度值
        /*overflow-x: hidden;*/
        /*overflow-y:auto;*/
    }
    .echarts-inner{
        margin-top: 8px;
    }
    .dashboard-lower{
        display: grid;
        gap: 12px;
        margin-top: 12px;
        padding-bottom: 12px;
    }
    .dashboard-lower-top{
        display: grid;
        grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
        gap: 12px;
    }
    .dashboard-panel{
        min-width: 0;
        padding: 18px 20px;
        border: 1px solid var(--ly-glass-border);
        border-radius: var(--ly-radius-md);
        background: var(--ly-glass-bg);
        box-shadow: var(--ly-glass-highlight), var(--ly-shadow-card);
        backdrop-filter: var(--ly-glass-blur);
        -webkit-backdrop-filter: var(--ly-glass-blur);
    }
    .panel-heading{
        display: flex;
        align-items: flex-start;
        justify-content: space-between;
        gap: 16px;
        margin-bottom: 16px;
    }
    .panel-heading h2{
        margin: 0;
        color: var(--ly-text-1);
        font-size: 15px;
        font-weight: 600;
        line-height: 1.45;
    }
    .panel-heading p{
        margin: 4px 0 0;
        color: var(--ly-text-3);
        font-size: 12px;
        line-height: 1.4;
    }
    .shortcut-grid{
        display: grid;
        grid-template-columns: repeat(2, minmax(0, 1fr));
        gap: 10px;
    }
    .shortcut-item{
        display: flex;
        min-width: 0;
        align-items: center;
        gap: 11px;
        padding: 12px;
        border: 1px solid var(--ly-line-soft);
        border-radius: 10px;
        color: var(--ly-text-1);
        background: var(--ly-glass-bg-soft);
        text-align: left;
        cursor: pointer;
        transition: border-color .18s ease, transform .18s ease, background-color .18s ease;
    }
    .shortcut-item:hover{
        transform: translateY(-1px);
        border-color: color-mix(in srgb, var(--el-color-primary) 30%, var(--ly-line-soft));
        background: var(--ly-glass-bg-strong);
    }
    .shortcut-icon{
        display: grid;
        width: 36px;
        height: 36px;
        flex: 0 0 auto;
        place-items: center;
        border-radius: 10px;
        color: var(--shortcut-color);
        background: color-mix(in srgb, var(--shortcut-color) 11%, transparent);
        font-size: 17px;
    }
    .shortcut-copy{
        display: grid;
        min-width: 0;
        gap: 3px;
    }
    .shortcut-copy strong{
        overflow: hidden;
        color: var(--ly-text-1);
        font-size: 12px;
        font-weight: 600;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .shortcut-copy small{
        overflow: hidden;
        color: var(--ly-text-3);
        font-size: 11px;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .shortcut-arrow{
        flex: 0 0 auto;
        margin-left: auto;
        color: var(--ly-text-3);
        font-size: 13px;
    }
    .system-panel .panel-heading{
        align-items: center;
    }
    .status-dot{
        display: inline-block;
        width: 6px;
        height: 6px;
        margin-right: 5px;
        border-radius: 50%;
        background: var(--el-color-success);
        vertical-align: 1px;
    }
    .status-dot.is-warning{ background: var(--el-color-warning); }
    .status-dot.is-danger{ background: var(--el-color-danger); }
    .system-metrics{
        display: grid;
        gap: 14px;
    }
    .system-metric{
        display: grid;
        gap: 7px;
    }
    .metric-label{
        display: flex;
        align-items: center;
        justify-content: space-between;
        color: var(--ly-text-2);
        font-size: 12px;
    }
    .metric-label strong{
        color: var(--ly-text-1);
        font-size: 12px;
        font-weight: 600;
        font-variant-numeric: tabular-nums;
    }
    /* 百分比数值跟随档位着色（正常=默认墨色，偏高/紧张用语义色加重提示） */
    .metric-label strong.is-warning{ color: var(--el-color-warning); }
    .metric-label strong.is-danger{ color: var(--el-color-danger); }
    .metric-track{
        height: 6px;
        overflow: hidden;
        border-radius: 10px;
        background: var(--ly-glass-bg-soft);
    }
    .metric-track span{
        display: block;
        height: 100%;
        border-radius: inherit;
        background: var(--el-color-success);
        transition: width .25s ease, background-color .25s ease;
    }
    /* 三档语义色（EP 语义变量，暗色自动跟随） */
    .metric-track span.is-success{ background: var(--el-color-success); }
    .metric-track span.is-warning{ background: var(--el-color-warning); }
    .metric-track span.is-danger{ background: var(--el-color-danger); }
    .system-meta{
        display: flex;
        flex-wrap: wrap;
        gap: 8px 18px;
        padding-top: 12px;
        border-top: 1px solid var(--ly-line-soft);
        color: var(--ly-text-3);
        font-size: 11px;
    }
    .system-meta span{
        display: inline-flex;
        align-items: center;
        gap: 5px;
    }
    .panel-empty{
        display: grid;
        min-height: 100px;
        place-items: center;
        color: var(--ly-text-3);
        font-size: 12px;
    }
    .operations-panel .panel-heading{
        align-items: center;
    }
    .operation-list{
        display: grid;
    }
    .operation-row{
        display: flex;
        min-width: 0;
        align-items: center;
        gap: 12px;
        padding: 11px 0;
        border-top: 1px solid var(--ly-line-soft);
    }
    .method-badge{
        display: inline-flex;
        min-width: 54px;
        height: 24px;
        flex: 0 0 auto;
        align-items: center;
        justify-content: center;
        border-radius: 6px;
        color: var(--el-color-primary);
        background: rgba(58, 123, 255, .08);
        font-size: 10px;
        font-weight: 600;
    }
    .method-post{ color: #268a74; background: rgba(38, 138, 116, .09); }
    .method-put,.method-patch{ color: #bd8124; background: rgba(189, 129, 36, .10); }
    .method-delete{ color: #d35b65; background: rgba(211, 91, 101, .09); }
    .operation-copy{
        display: grid;
        min-width: 0;
        flex: 1;
        gap: 4px;
    }
    .operation-copy strong{
        overflow: hidden;
        color: var(--ly-text-1);
        font-size: 12px;
        font-weight: 500;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .operation-copy span{
        overflow: hidden;
        color: var(--ly-text-3);
        font-size: 11px;
        text-overflow: ellipsis;
        white-space: nowrap;
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
    @media (max-width: 980px){
        .dashboard-lower-top{
            grid-template-columns: minmax(0, 1fr);
        }
    }
    @media (max-width: 560px){
        .dashboard-panel{
            padding: 15px;
        }
        .shortcut-grid{
            grid-template-columns: minmax(0, 1fr);
        }
        .operation-row{
            gap: 8px;
        }
    }
</style>
