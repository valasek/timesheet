// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  all: [] // id, name, rate, disabled
}

const getters = {}

const actions = {

  getProjects ({ commit }) {
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
  },

  createProject ({ commit }, payload) {
    api.apiClient.post(`/api/projects`, payload, { crossDomain: true })
      .then(response => {
        commit('CREATE_PROJECT', response.data)
        Notify.create({
          message: 'Project ' + payload.name + ' created and a default rate set to ' + payload.rate,
          color: 'teal',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t create project on server. ' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  removeProject ({ commit }, payload) {
    api.apiClient.delete(`/api/projects/` + payload.id, { crossDomain: true })
      .then(response => {
        commit('REMOVE_PROJECT', payload.id)
        Notify.create({
          message: 'Project ' + payload.name + ' and ' + response.data + ' corresponding reported records were deleted',
          color: 'teal',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t delete project ' + payload.name + ' on server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  toggleProject ({ commit }, payload) {
    api.apiClient.post(`/api/projects/toggle/` + payload, { crossDomain: true })
      .then(response => {
        commit('TOGGLE_PROJECT', response.data.id)
        Notify.create({
          message: 'Project visibility switched',
          color: 'teal',
          icon: ''
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t change project visibitity on server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  }

}

const mutations = {

  SET_PROJECTS (state, projects) {
    state.all = projects.sort(function (a, b) {
      var nameA = a.name.toUpperCase()
      var nameB = b.name.toUpperCase()
      if (nameA < nameB) {
        return -1
      }
      if (nameA > nameB) {
        return 1
      }
      return 0
    })
  },

  CREATE_PROJECT (state, payload) {
    state.all.push(payload)
  },

  REMOVE_PROJECT (state, id) {
    for (let i = 0; i < state.all.length; ++i) {
      if (state.all[i].id === id) {
        state.all.splice(i, 1)
        break
      }
    }
  },

  TOGGLE_PROJECT (state, id) {
    for (let i = 0; i < state.all.length; ++i) {
      if (state.all[i].id === id) {
        state.all[i].disabled = !state.all[i].disabled
      }
    }
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
