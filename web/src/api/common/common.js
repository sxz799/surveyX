import request from '../../utils/request.js'

export function login(data) {
    return request({
        url: '/common/login',
        method: 'post',
        data
    })
}

export function getGithubLoginUrl() {
    return request({
        url: '/common/oauth/loginUrl',
        method: 'GET'
    })
}

export function loginByGithub(code) {
    return request({
        url: '/common/oauth/github?code=' + code,
        method: 'post'
    })
}




export function logout() {
    return request({
        url: '/common/logout',
        method: 'post'
    })
}
