import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import viteCompression from 'vite-plugin-compression'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// 第9行的注释不能删除，docker镜像打包时会用到 github actions 会用到
export default defineConfig(({ mode, command }) => {
    return {
        plugins: [
            vue(),
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
            viteCompression({
                threshold: 10240 // 对大于 1mb 的文件进行压缩
            })
        ],
        //sed_tag/base: '/dist/',
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
        build: {
            chunkSizeWarningLimit: 1500,
            rollupOptions: {
                output: {
                    manualChunks(id) {
                        if (id.includes("node_modules")) {
                            return id.toString().split("node_modules/")[1].split("/")[0].toString()
                        }
                    },
                    // 用于从入口点创建的块的打包输出格式[name]表示文件名,[hash]表示该文件内容hash值
                    entryFileNames: 'js/[name].[hash].js',
                    // 用于命名代码拆分时创建的共享块的输出命名
                    // 　　chunkFileNames: 'js/[name].[hash].js',
                    // 用于输出静态资源的命名，[ext]表示文件扩展名
                    assetFileNames: '[ext]/[name].[hash].[ext]',
                    // 拆分js到模块文件夹
                    chunkFileNames: (chunkInfo) => {
                        const facadeModuleId = chunkInfo.facadeModuleId ? chunkInfo.facadeModuleId.split('/') : [];
                        const fileName = facadeModuleId[facadeModuleId.length - 2] || '[name]';
                        return `js/${fileName}/[name].[hash].js`;
                    },
                }
            }
        },
    }
})
