import request from "@/utils/request.js";


export function changPwd(data) {
    return request({
        url: '/api/user/changePwd',
        method: 'put',
        data: data
    })
}

export function resetPwd(data) {
    return request({
        url: '/api/user/resetPwd',
        method: 'put',
        data: data
    })
}

export function register(data) {
    return request({
        url: '/api/user/register',
        method: 'post',
        data: data
    })
}