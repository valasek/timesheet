// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  // rates: [ '' ], // array of rates:string
  types: [ '' ], // array of rateType:string
  all: [ {} ] // id: '', name: '', type: ''
}

const getters = {}

const actions = {

  getRates ({ commit, dispatch }) {
    api.apiClient.get(`/api/rates`, { crossDomain: true })
      .then(response => {
        commit('SET_RATES', response.data)
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t read rates from server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  }

}

const mutations = {

  SET_RATES (state, rates) {
    state.all = rates
    // state.rates = rates.map(function (obj) {
    //     return obj.name
    // })
    state.types = [...new Set(rates.map(function (obj) {
      return obj.type
    }))]
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
