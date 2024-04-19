
import request from '../../utils/request.js'


export function get(id) {
    return request({
        url: '/user/survey/'+id,
        method: 'get',
    })
}
