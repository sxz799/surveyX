
import request from '../../utils/request.js'


export function add(data) {
    return request({
        url: '/user/answer',
        method: 'post',
        data
    })
}

