import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import viteCompression from 'vite-plugin-compression'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

export default defineConfig(({mode}) => {
    const env = loadEnv(mode, process.cwd())
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
        base: env.VITE_BASE_PATH,
        server: {
            port: 5173,
            proxy: {
                '/dev-api': {
                    target: 'http://localhost:65534',
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/dev-api/, '')
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
                    chunkFileNames: 'js/[hash].js',
                    entryFileNames: 'js/[hash].js',
                    assetFileNames: '[ext]/[hash].[ext]',
                }
            }
        },
    }
})
