import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: "0.0.0.0",
    port: 5178,
    proxy: {
      "/business": {
        target: 'https://test.backend.xiuxian.xialoukeji.com',
        changeOrigin: true,// 允许跨域
        rewrite: (path) => path.replace(/^\/business/, ""),
      },
      "/platform": {
        target: 'https://test.xialou.xialoukeji.com',
        changeOrigin: true,// 允许跨域
        rewrite: (path) => path.replace(/^\/platform/, ""),
      },
    },
  },
});