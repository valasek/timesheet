// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import axios from 'axios'

const port = process.env.PORT || '3000'

const apiClient = axios.create({
    baseURL: 'http://localhost:3000',
    // baseURL: 'https://timesheet-cloud.herokuapp.com:' + port,
    withCredentials: false, // This is the default
    crossDomain: true,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    timeout: 10000
})

console.log('PORT:', port) /* eslint-disable-line no-console */

export default {
    apiClient
}
