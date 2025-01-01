import request from '../../utils/request.js'


export function get(id) {
    return request({
        url: '/api/guest/survey/' + id,
        method: 'get',
    })
}
