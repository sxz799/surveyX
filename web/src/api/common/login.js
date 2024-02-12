import request from '../../utils/request.js'

export function login(data) {
    return request({
        url: '/login',
        method: 'post',
        data
    })
}

export function logout() {
    return request({
        url: '/logout',
        method: 'post'
    })
}
