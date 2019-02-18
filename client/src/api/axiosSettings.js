// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import axios from 'axios'

const apiClient = axios.create({
    baseURL: 'http://localhost:3000',
    withCredentials: false, // This is the default
    crossDomain: true,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    timeout: 10000
})

export default {
    apiClient
}
