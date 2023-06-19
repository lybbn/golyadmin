const { defineConfig } = require('@vue/cli-service')
const path = require('path');
// 导入compression-webpack-plugin
const CompressionWebpackPlugin = require('compression-webpack-plugin');
// 定义压缩文件类型
const productionGzipExtensions = ['js', 'css'];

const Timestamp = new Date().getTime();

const webpack = require('webpack');

const appConfig = require("./src/config/index.js")


function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = defineConfig({
	transpileDependencies: true,
	assetsDir:'static',//表示打包后，静态资源生成到static文件夹中
	publicPath:'./',
	runtimeCompiler: true, /* 开启vue运行时模板编译功能！！ */
	productionSourceMap: false, //打包后是否生成map文件
	lintOnSave:false, // 是否开启eslint保存检测
	devServer: {//开发环境
		port:8090,
		host:'127.0.0.1',
		open: true, //配置自动启动浏览器
		hot: true, // 热更新
		client: {
        	// 当有错误的时候在客户端进行覆盖显示
            overlay: false,
        },
		historyApiFallback: true,
		allowedHosts: "all",//允许任意ip
		proxy: {
			'/api': {
				target: appConfig.API_BASEURL,
				ws: false, //解决控制台反复输出WebSocket connection to 'ws://xxxxxx:8080/ws' failed:
				pathRewrite: {
					'^/api': ''
				}
			}
		}
	},
	//gzip配置
	configureWebpack(config){
		//性能提示
		config["performance"] = {
			hints: false,
		}
		/* 引入 jquery 插件 */
		config.plugins.push(new webpack.ProvidePlugin({
			jQuery: 'jquery',
			$: 'jquery',
			'windows.jQuery': 'jquery'
		}));
		/* gzip压缩 */
		config.plugins.push(new CompressionWebpackPlugin({
			algorithm: 'gzip', // 使用 gzip 压缩
			test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
			filename: '[path][base].gz', // 压缩后的文件名
			threshold: 10240,// 对超过 10k 的数据压缩
			minRatio: 0.8,// 压缩率小于 1 才会压缩
			deleteOriginalAssets: false // 是否删除未压缩的文件(源文件)，不建议开启
		}))
		config.devtool = 'source-map'
    	config.output.libraryExport = 'default'  /* 解决import UMD打包文件时, 组件install方法执行报错的问题！！ */
		// 打包编译后修改 js 文件名
		// config.output.filename = `assets/js/[name].${Timestamp}.js`
		// config.output.chunkFilename = `assets/js/[name].${Timestamp}.js`

		if (process.env.NODE_ENV === 'production') {
			/* 生产环境 */
			config.mode = 'production';
			// 配置如何展示性能提示
			config['performance'] = {
				hints: 'warning', // 打开/关闭提示
				"maxAssetSize": 250000, // 根据单个资源体积(单位: bytes)，控制 webpack 何时生成性能提示
				"maxEntrypointSize": 250000, // 根据入口起点的最大体积(单位: bytes)，控制 webpack 何时生成性能提示
				// 只给出 .js 文件的性能提示
				assetFilter: function (assetFilename) {
					return assetFilename.endsWith('.js');
				}
			}
			config['optimization'] = {
				splitChunks:{
					chunks: "all",
					automaticNameDelimiter: '~',
					name: "lyadminChunks",
					cacheGroups: {
						//第三方库抽离
						vendor: {
							name: "modules",
							test: /[\\/]node_modules[\\/]/,
							priority: -10
						},
						elicons: {
							name: "elicons",
							test: /[\\/]node_modules[\\/]@element-plus[\\/]icons-vue[\\/]/
						},
						tinymce: {
							name: "tinymce",
							test: /[\\/]node_modules[\\/]tinymce[\\/]/
						},
						echarts: {
							name: "echarts",
							test: /[\\/]node_modules[\\/]echarts[\\/]/
						},
					}
				}
			}
		}else if (process.env.NODE_ENV === 'test') {
			/* 测试环境 */
			config.mode = 'none';
		} else{
			/* 开发环境 */
			config.mode = 'development'
		};

	},
	//解决富文本编辑器报错imports失败
    chainWebpack(config){
		 // 移除 prefetch 插件
		 config.plugins.delete('preload');
		 config.plugins.delete('prefetch');
		 config.resolve.alias.set('vue-i18n', 'vue-i18n/dist/vue-i18n.cjs.js');
		 //设置title
		 config.plugin("html").tap((args) => {
			args[0].title = appConfig.APP_TITLE;
			return args;
		})
		 //处理svg-icon
		 config.module
			 .rule('svg')
			 .exclude.add(resolve('src/icons'))
			 .end()
		 config.module
			.rule('icons')
			.test(/\.svg$/)
			.include.add(resolve('src/icons'))
			.end()
			.use('svg-sprite-loader')
			.loader('svg-sprite-loader')
			.options({
				symbolId: 'icon-[name]'
			})
			.end()
    },
})