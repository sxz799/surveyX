
import request from '../../utils/request.js'

export function list(data) {
  return request({
    url: '/admin/question/list',
    method: 'get',
    params: data
  })
}

export function get(id) {
    return request({
        url: '/admin/question/'+id,
        method: 'get',
    })
}

export function add(data) {
    return request({
        url: '/admin/question',
        method: 'post',
        data
    })
}

export function update(data) {
    return request({
        url: '/admin/question',
        method: 'put',
        data
    })
}


export function del(id) {
    return request({
        url: '/admin/question/'+id,
        method: 'delete',
    })
}

export function analysis(id) {
    return request({
        url: '/admin/question/analysis/'+id,
        method: 'get',
    })
}
