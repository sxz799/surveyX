import request from '../../utils/request.js'

export function list(data) {
    return request({
        url: '/api/user/guest/list',
        method: 'get',
        params: data
    })
}


