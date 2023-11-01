import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
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
        }
    }
})
