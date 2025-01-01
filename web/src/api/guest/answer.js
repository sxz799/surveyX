import request from '../../utils/request.js'


export function add(data) {
    return request({
        url: '/api/guest/answer',
        method: 'post',
        data
    })
}

