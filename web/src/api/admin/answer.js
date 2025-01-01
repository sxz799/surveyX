import request from '../../utils/request.js'


export function list(data) {
    return request({
        url: '/api/admin/answer/list',
        method: 'get',
        params: data
    })
}



