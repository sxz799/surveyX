import request from '../../utils/request.js'

export function list(params) {
    return request({
        url: '/api/admin/survey/list',
        method: 'get',
        params: params
    })
}

export function get(id) {
    return request({
        url: '/api/admin/survey/' + id,
        method: 'get',
    })
}

export function add(data) {
    return request({
        url: '/api/admin/survey',
        method: 'post',
        data
    })
}

export function update(data) {
    return request({
        url: '/api/admin/survey',
        method: 'put',
        data
    })
}

export function del(id) {
    return request({
        url: '/api/admin/survey/' + id,
        method: 'delete',
    })
}

export function analysis(id) {
    return request({
        url: '/api/admin/survey/analysis/' + id,
        method: 'get',
    })
}

export function upload(formData) {
    console.log(formData.get("file"))
    return request({
        url: '/api/admin/survey/import',
        method: 'post',
        data:formData,
    })
}