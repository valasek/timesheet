// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { Notify } from 'quasar'

// initial state
const state = {
  all: [], // id, name, allocation, disabled
  selected: '',
  selectedAllocation: 1
}

const getters = {}

const actions = {

  getConsultants ({ commit }) {
    api.apiClient.get('/api/consultants', { crossDomain: true })
      .then(response => {
        commit('SET_CONSULTANTS', response.data)
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t read consultants from server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  createConsultant ({ commit }, payload) {
    api.apiClient.post(`/api/consultants`, payload, { crossDomain: true })
      .then(response => {
        commit('CREATE_CONSULTANT', response.data)
        Notify.create({
          message: payload.name + ' was created with ' + payload.allocation * 100 + '% allocation',
          color: 'teal',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t create ' + payload.name + ' on server. ' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  updateConsultant ({ dispatch }, payload) {
    api.apiClient.put(`/api/consultants`, payload, { crossDomain: true })
      .then(response => {
        dispatch('getConsultants', { root: true })
        Notify.create({
          message: payload.name + ' allocation was updated to ' + payload.allocation * 100 + '%',
          color: 'teal',
          icon: ''
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t update allocation of ' + payload.name + ' on the server. ' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  removeConsultant ({ commit }, payload) {
    api.apiClient.delete(`/api/consultants/` + payload.id, { crossDomain: true })
      .then(response => {
        commit('REMOVE_CONSULTANT', payload.id)
        Notify.create({
          message: payload.name + ' and ' + response.data + ' linked reported records were deleted',
          color: 'teal',
          icon: 'report_problem'
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t delete ' + payload.name + ' on the server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  toggleConsultant ({ commit }, payload) {
    api.apiClient.post(`/api/consultants/toggle/` + payload, { crossDomain: true })
      .then(response => {
        commit('TOGGLE_CONSULTANT', response.data.id)
        Notify.create({
          message: 'Consultant visibility switched',
          color: 'teal',
          icon: ''
        })
      })
      .catch(e => {
        Notify.create({
          message: 'Couldn\'t change consultant visibitity on server. \n' + e.toString(),
          color: 'negative',
          icon: 'report_problem'
        })
        console.log(e) /* eslint-disable-line no-console */
      })
  },

  setSelected ({ commit }, selectedConsultant) {
    commit('SET_SELECTED', selectedConsultant)
  }

}

const mutations = {

  SET_CONSULTANTS (state, consultants) {
    state.all = consultants.sort(function (a, b) {
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
    let c = localStorage.getItem('selectedConsultant')
    if (c && c.length > 0) {
      state.selected = c
      state.selectedAllocation = state.all.find(cc => cc.name === c).allocation
    } else {
      state.selected = consultants[0].name
      state.selectedAllocation = consultants[0].allocation
    }
  },

  CREATE_CONSULTANT (state, payload) {
    state.all.push(payload)
  },

  REMOVE_CONSULTANT (state, id) {
    for (let i = 0; i < state.all.length; ++i) {
      if (state.all[i].id === id) {
        state.all.splice(i, 1)
        break
      }
    }
  },

  TOGGLE_CONSULTANT (state, id) {
    for (let i = 0; i < state.all.length; ++i) {
      if (state.all[i].id === id) {
        state.all[i].disabled = !state.all[i].disabled
      }
    }
  },

  SET_SELECTED (state, selectedConsultant) {
    state.selected = selectedConsultant
    state.selectedAllocation = state.all.find(cc => cc.name === selectedConsultant).allocation
    localStorage.setItem('selectedConsultant', selectedConsultant)
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
