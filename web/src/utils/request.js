import axios from 'axios'
import router from "@/utils/router.js";

// 创建axios实例

const request = axios.create({
    // axios中请求配置有baseURL选项，表示请求URL公共部分
    baseURL: import.meta.env.VITE_BASE_URL,
    // 超时
    timeout: 10000
})

// request拦截器
request.interceptors.request.use(
    (config) => {
        // 在发送请求之前可以进行一些操作，比如添加请求头
        if (localStorage.getItem("token")) {
            config.headers.Authorization = `Bearer ` + localStorage.getItem("token");
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器
request.interceptors.response.use((response) => {
        if (response.data.code === 401) { // 401 token过去 重定向到登录页
            router.push({path: '/login'})
        }
        return response.data;
    },
    (error) => {
        return Promise.reject(error);
    }
);


export default request
