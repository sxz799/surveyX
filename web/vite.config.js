import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// 第9行的注释不能删除，docker镜像打包时会用到 github actions 会用到
export default defineConfig(({mode, command}) => {
    return {
        plugins: [vue()],
        //base:'/dist/',
        server: {
            proxy: {
                '/api': {
                    target: 'http://localhost:3000/api',
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/api/, '')
                }
            }
        },
        resolve: {
            // https://cn.vitejs.dev/config/#resolve-alias
            alias: {
                // 设置路径
                '~': path.resolve(__dirname, './'),
                // 设置别名
                '@': path.resolve(__dirname, './src')
            },
        },
    }
})
