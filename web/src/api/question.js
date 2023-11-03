
import request from '../utils/request'

export function list(data) {
  return request({
    url: '/question/list',
    method: 'get',
    params: data
  })
}

export function get(id) {
    return request({
        url: '/question/'+id,
        method: 'get',
    })
}

export function add(data) {
    return request({
        url: '/question',
        method: 'post',
        data
    })
}

export function update(data) {
    return request({
        url: '/question',
        method: 'put',
        data
    })
}


export function del(id) {
    return request({
        url: '/question/'+id,
        method: 'delete',
    })
}
