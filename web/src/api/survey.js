
import request from '../utils/request'

export function list(data) {
  return request({
    url: '/survey/list',
    method: 'get',
    params: data
  })
}

export function get(id) {
    return request({
        url: '/survey/'+id,
        method: 'get',
    })
}

export function add(data) {
    return request({
        url: '/survey',
        method: 'post',
        data
    })
}

export function update(data) {
    return request({
        url: '/survey',
        method: 'put',
        data
    })
}

export function del(id) {
    return request({
        url: '/survey/'+id,
        method: 'delete',
    })
}
