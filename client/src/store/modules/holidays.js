// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  all: [ '' ] // date, description
}

const getters = {}

const actions = {

  getHolidays ({ commit, dispatch }) {
    api.apiClient.get(`/api/holidays`, { crossDomain: true })
      .then(response => {
        commit('SET_HOLIDAYS', response.data)
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t read holidays from server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  }

}

const mutations = {

  SET_HOLIDAYS (state, payload) {
    state.all = payload
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
