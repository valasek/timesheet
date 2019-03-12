// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>
import { isWithinInterval } from 'date-fns'

const defaultNotificationType = 'info'

// initial state
const state = {
    page: 'Timesheet',
    notification: false,
    notificationText: '',
    notificationType: defaultNotificationType, // success, info, error - snackbar types https://vuetifyjs.com/en/components/snackbars#introduction
    isCurrentWeek: true,
    weekUnlocked: true
}

const getters = {}

const actions = {

    setNotification ({ commit }, payload) {
        commit('SET_NOTIFICATION', payload)
    },
    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
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

    SET_NOTIFICATION (state, payload) {
        state.notificationText = payload.text
        if (payload.type.length > 1) {
            state.notificationType = payload.type
        }
        state.notification = true
    },

    RESET_NOTIFICATION (state) {
        state.notification = false
        state.notificationText = ''
        state.notificationType = defaultNotificationType
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
