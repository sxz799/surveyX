
import request from '../utils/request'


export function add(data) {
    return request({
        url: '/answer',
        method: 'post',
        data
    })
}

