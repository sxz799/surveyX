
import request from '../../utils/request.js'

export function list(data) {
  return request({
    url: '/admin/survey/list',
    method: 'get',
    params: data
  })
}

export function get(id) {
    return request({
        url: '/admin/survey/'+id,
        method: 'get',
    })
}

export function add(data) {
    return request({
        url: '/admin/survey',
        method: 'post',
        data
    })
}

export function update(data) {
    return request({
        url: '/admin/survey',
        method: 'put',
        data
    })
}

export function del(id) {
    return request({
        url: '/admin/survey/'+id,
        method: 'delete',
    })
}
