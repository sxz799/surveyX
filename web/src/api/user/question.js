
import request from '../../utils/request.js'

export function list(data) {
  return request({
    url: '/user/question/list',
    method: 'get',
    params: data
  })
}


