/**
 * v4 玻璃拟态 ECharts 主题（对齐 dvlyadmin-mini design-mockups/dashboard-preview-v4.html）
 *
 * 图表语言规范（设计稿 §ECharts）：
 * - 系列调色板：蓝色系梯度（深→浅），禁止 echarts 默认绿/红/黄杂色与旧主色 #409EFF
 * - 轴文字：亮 #9AA3BC / 暗 #6B7490，fontSize 11
 * - 分割线：亮 rgba(27,35,64,.06) / 暗 rgba(255,255,255,.08)
 * - tooltip：亮 白底 0.94 + 蓝描边 rgba(58,123,255,.2) 深字 #1B2340；暗 深玻璃底 + 亮字
 * - 画布背景必须透明（融入玻璃卡与极光，禁止写死白色实底）
 *
 * 用法：
 *   import { lyChartTheme } from "@/components/analysis/echartsTheme";
 *   echarts.init(el, lyChartTheme())   // 按当前亮/暗模式自动选主题
 *
 * 注意：echarts 主题是 init 时确定的，页面驻留期间切换暗色需重新进入页面才会重绘主题。
 */

// v4 系列调色板：深蓝 → 主蓝 → 青 → 浅蓝（多系列堆叠呈现蓝色梯度，与极光主题融合）
const LY_V4_PALETTE = ['#234FB8', '#2E66E8', '#3A7BFF', '#58B8E8', '#9DBEFF', '#C9DBFF'];

// 亮色主题
const lyV4ThemeLight = {
    color: LY_V4_PALETTE,
    backgroundColor: 'transparent',
    title: {
        textStyle: { color: '#1B2340' },
        subtextStyle: { color: '#9AA3BC' },
    },
    legend: {
        textStyle: { color: '#5A647C', fontSize: 12 },
        inactiveColor: 'rgba(128,128,128,0.4)',
    },
    categoryAxis: {
        axisLine: { show: true, lineStyle: { color: 'rgba(27,35,64,.10)' } },
        axisTick: { show: false },
        axisLabel: { color: '#9AA3BC', fontSize: 11 },
        splitLine: { show: false },
    },
    valueAxis: {
        axisLine: { show: false },
        axisTick: { show: false },
        axisLabel: { color: '#9AA3BC', fontSize: 11 },
        splitLine: { show: true, lineStyle: { color: 'rgba(27,35,64,.06)' } },
    },
    tooltip: {
        backgroundColor: 'rgba(255,255,255,.94)',
        borderColor: 'rgba(58,123,255,.2)',
        borderWidth: 1,
        textStyle: { color: '#1B2340', fontSize: 12 },
    },
};

// 暗色主题（html.dark）
const lyV4ThemeDark = {
    color: LY_V4_PALETTE,
    backgroundColor: 'transparent',
    title: {
        textStyle: { color: '#E8ECF6' },
        subtextStyle: { color: '#6B7490' },
    },
    legend: {
        textStyle: { color: '#A8B3CC', fontSize: 12 },
        inactiveColor: 'rgba(128,128,128,0.4)',
    },
    categoryAxis: {
        axisLine: { show: true, lineStyle: { color: 'rgba(255,255,255,.12)' } },
        axisTick: { show: false },
        axisLabel: { color: '#6B7490', fontSize: 11 },
        splitLine: { show: false },
    },
    valueAxis: {
        axisLine: { show: false },
        axisTick: { show: false },
        axisLabel: { color: '#6B7490', fontSize: 11 },
        splitLine: { show: true, lineStyle: { color: 'rgba(255,255,255,.08)' } },
    },
    tooltip: {
        backgroundColor: 'rgba(30,37,58,.94)',
        borderColor: 'rgba(58,123,255,.35)',
        borderWidth: 1,
        textStyle: { color: '#E8ECF6', fontSize: 12 },
    },
};

// 已注册模块记录（ES 模块命名空间对象冻结不可扩展，禁止往 echarts 模块上挂属性，用 WeakSet 幂等）
const registeredModules = new WeakSet();

/**
 * 在指定 echarts 实例（echarts/core 或全量 echarts）上注册 v4 主题
 * 两个模块的主题注册表相互独立，需各自注册一次（WeakSet 幂等保护）
 * @param {Object} echartsModule echarts/core 或 * as echarts
 */
export function registerLyChartThemes(echartsModule) {
    if (!echartsModule || typeof echartsModule.registerTheme !== 'function') return;
    if (registeredModules.has(echartsModule)) return;
    echartsModule.registerTheme('lyv4', lyV4ThemeLight);
    echartsModule.registerTheme('lyv4-dark', lyV4ThemeDark);
    registeredModules.add(echartsModule);
}

/**
 * 按当前亮/暗模式返回应使用的主题名（init 时调用）
 * @returns {string} 'lyv4' | 'lyv4-dark'
 */
export function lyChartTheme() {
    return document.documentElement.classList.contains('dark') ? 'lyv4-dark' : 'lyv4';
}
