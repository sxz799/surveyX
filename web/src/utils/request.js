import axios from 'axios'



axios.defaults.headers['Content-Type'] = 'application/json;charset=utf-8'
// 创建axios实例
const service = axios.create({
  // axios中请求配置有baseURL选项，表示请求URL公共部分
  baseURL: import.meta.env.VITE_APP_BASE_API,
  // 超时
  timeout: 10000
})

// request拦截器
service.interceptors.request.use(
    (config) => {
      // 在发送请求之前可以进行一些操作，比如添加请求头
      // config.headers.Authorization = `Bearer ${token}`;
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
);

// 响应拦截器
service.interceptors.response.use(
    (response) => {
        if(response.data.message==="Token Expired"){
            window.location.href="/login"
        }
      // 在这里可以对响应数据进行处理
      return response.data;
    },
    (error) => {
      return Promise.reject(error);
    }
);


export default service
