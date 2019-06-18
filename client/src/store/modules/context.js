// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>
import { isWithinInterval } from 'date-fns'

// initial state
const state = {
  page: 'Timesheet',
  pageIcon: 'Home',
  isCurrentWeek: true,
  weekUnlocked: true
}

const getters = {}

const actions = {

  setWeekUnlocked ({ commit }, payload) {
    commit('SET_WEEK_UNLOCKED', payload)
  },
  setIsCurrentWeek ({ commit, rootState }) {
    let today = new Date()
    if (isWithinInterval(today, { start: rootState.settings.dateFrom, end: rootState.settings.dateTo })) {
      commit('SET_IS_CURRENT_WEEK', true)
      commit('SET_WEEK_UNLOCKED', true)
    } else {
      commit('SET_IS_CURRENT_WEEK', false)
      commit('SET_WEEK_UNLOCKED', false)
    }
  }
}

const mutations = {

  SET_PAGE (state, page) {
    state.page = page
  },

  SET_PAGE_ICON (state, pageIcon) {
    state.pageIcon = pageIcon
  },

  SET_IS_CURRENT_WEEK (state, payload) {
    state.isCurrentWeek = payload
  },

  SET_WEEK_UNLOCKED (state, payload) {
    state.weekUnlocked = payload
  }

}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
