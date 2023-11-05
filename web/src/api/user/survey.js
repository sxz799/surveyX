
import request from '../../utils/request.js'

export function list(data) {
  return request({
    url: '/user/survey/list',
    method: 'get',
    params: data
  })
}

export function get(id) {
    return request({
        url: '/user/survey/'+id,
        method: 'get',
    })
}
