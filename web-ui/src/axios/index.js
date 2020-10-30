import axios from 'axios';
axios.defaults.baseURL = 'apiURL'; // 请求的默认域名
// 添加一个请求拦截器
axios.interceptors.request.use(config => {
        config.headers.languagetype = 'CN'; // 举例，加上一个公共头部
        return config;
    },
    err => {
        return Promise.reject(err);
    });
//添加一个响应拦截器
axios.interceptors.response.use(res => {
    //在这里对返回的数据进行处理
    console.log(res.data, '网络正常');
    return res.data;
}, err => {
    console.log('网络开了小差！请重试...');
    return Promise.reject(err);
})

export default axios

