// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  all: [ '' ] // id, name, rate
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
    api.apiClient.post(`/api/projects/` + payload, { crossDomain: true })
      .then(response => {
        commit('CREATE_PROJECT', response.data)
        Notify.create({
          message: 'Project ' + payload + ' created. Set the default rate.',
          color: 'warning',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t create project on server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  deleteProject ({ commit }, payload) {
    api.apiClient.delete(`/api/projects/` + payload, { crossDomain: true })
      .then(response => {
        commit('DELETE_PROJECT', response.data)
        Notify.create({
          message: 'Deleted project ' + payload + ' and corresponding reported records.',
          color: 'positive',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t delete project on server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  toggleProject ({ commit }, payload) {
    api.apiClient.delete(`/api/projects/toggle/` + payload, { crossDomain: true })
      .then(response => {
        commit('TOGGLE_PROJECT', response.data)
        var toggleText = ''
        payload === true ? toggleText = ' enabled' : toggleText = ' disabled'
        Notify.create({
          message: 'Project ' + payload + toggleText,
          color: 'positive',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        var toggleText = ''
        payload === true ? toggleText = ' enable' : toggleText = ' disable'
        Notify.create({
          message: 'Couldn\'t ' + toggleText + ' project on server. \n' + e.toString(),
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
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
