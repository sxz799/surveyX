import request from '../../utils/request.js'

export function login(data) {
    return request({
        url: '/login',
        method: 'post',
        data
    })
}

export function getGithubLoginUrl() {
    return request({
        url: '/oauth/loginUrl',
        method: 'GET'
    })
}

export function loginByGithub(code) {
    return request({
        url: '/oauth/github?code=' + code,
        method: 'post'
    })
}

export function logout() {
    return request({
        url: '/logout',
        method: 'post'
    })
}
