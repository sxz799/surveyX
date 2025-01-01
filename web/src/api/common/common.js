import request from '../../utils/request.js'

export function login(data) {
    return request({
        url: '/api/common/login',
        method: 'post',
        data
    })
}

export function getGithubLoginUrl() {
    return request({
        url: '/api/common/oauth/loginUrl',
        method: 'GET'
    })
}

export function loginByGithub(code) {
    return request({
        url: '/api/common/oauth/github?code=' + code,
        method: 'post'
    })
}


export function logout() {
    return request({
        url: '/api/common/logout',
        method: 'post'
    })
}
