// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import axios from 'axios'

const debug = process.env.NODE_ENV !== 'production'
// baseURL: 'https://timesheet-cloud.herokuapp.com:' + port,
let baseURL = ''

if (debug) {
  baseURL = 'http://localhost:3000'
}

const apiClient = axios.create({
    baseURL: baseURL,
    withCredentials: false, // This is the default
    crossDomain: true,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
      // 'Accept-Encoding': 'gzip, deflate, br'
    },
    timeout: 10000
})

// apiClient.interceptors.response.use(function (response) {
//   var ctype = response.headers['content-type']
//   response.data = ctype.includes('charset=GB2312') ? iconv.decode(response.data, 'gb2312') : iconv.decode(response.data, 'utf-8')
//   return response
// })

export default {
    apiClient
}
