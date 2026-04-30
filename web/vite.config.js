import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import viteCompression from 'vite-plugin-compression'
import path from 'path'

export default defineConfig(({ mode }) => {
  const isDev = mode === 'development'
  const API_DOMAIN = isDev ? '127.0.0.1:9000' : 'golyadmin.lybbn.cn'
  const API_BASEURL = isDev
    ? `http://${API_DOMAIN}/api/`
    : `https://${API_DOMAIN}/api/`

  return {
    base: './',
    plugins: [
      vue(),
      vueJsx(),
      createSvgIconsPlugin({
        iconDirs: [path.resolve(process.cwd(), 'src/icons/svg')],
        symbolId: 'icon-[name]',
      }),
      viteCompression({
        algorithm: 'gzip',
        ext: '.gz',
        threshold: 10240,
        deleteOriginFile: false,
      }),
    ],
    resolve: {
      extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue'],
      alias: {
        '@': path.resolve(__dirname, 'src'),
        'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
      },
    },
    css: {
      preprocessorOptions: {
        scss: {
          api: 'modern-compiler',
        },
      },
    },
    server: {
      port: 8090,
      host: '127.0.0.1',
      open: true,
      hmr: true,
      proxy: {
        '/api': {
          target: API_BASEURL,
          ws: false,
          changeOrigin: true,
          rewrite: (p) => p.replace(/^\/api/, ''),
        },
      },
    },
    build: {
      outDir: 'dist',
      assetsDir: 'static',
      sourcemap: false,
      chunkSizeWarningLimit: 250000,
      rollupOptions: {
        output: {
          manualChunks: {
            modules: ['vue', 'vue-router', 'pinia', 'axios'],
            elicons: ['@element-plus/icons-vue'],
            tinymce: ['tinymce'],
            echarts: ['echarts'],
            elementplus: ['element-plus'],
          },
        },
      },
    },
  }
})
