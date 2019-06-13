// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  all: [ '' ] // id, name, rate
}

const getters = {}

const actions = {

  getProjects ({ commit, dispatch }) {
    api.apiClient.get(`/api/projects`, { crossDomain: true })
      .then(response => {
        commit('SET_PROJECTS', response.data)
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t read projects from server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  }

}

const mutations = {

  SET_PROJECTS (state, projects) {
    state.all = projects
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
