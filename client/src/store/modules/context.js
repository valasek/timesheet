import { addDays, subDays } from 'date-fns'

const defaultNotificationType = 'info'

// initial state
const state = {
    page: 'Timesheet',
    notification: false,
    notificationText: '',
    notificationType: defaultNotificationType, // success, info, error - snackbar types https://vuetifyjs.com/en/components/snackbars#introduction
    dateMonth: new Date().toISOString().substr(0, 7),
    dateFrom: getMonday(new Date()),
    dateTo: getSunday(new Date()),
    dailyWorkingHours: 8
}

const getters = {}

const actions = {

    setMonth ({ commit, dispatch }, month) {
        let monday = getMonday(new Date(month))
        dispatch('jumpToWeek', monday)
        commit('SET_MONTH', month)
    },

    setNotification ({ commit }, payload) {
        commit('SET_NOTIFICATION', payload)
    },

    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
    changeWeek ({ commit }, direction) {
        commit('SET_WEEK', direction)
    },
    jumpToWeek ({ commit }, monday) {
        commit('JUMP_TO_WEEK', monday)
    }
}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    },

    SET_MONTH (state, month) {
        state.dateMonth = month
    },

    SET_WEEK (state, direction) {
        let targetMonday = {}
        switch (direction) {
            case 'previous':
                targetMonday = subDays(state.dateFrom, 7)
                break
            case 'next':
                targetMonday = addDays(state.dateTo, 7)
                break
            default:
                return
        }
        state.dateFrom = targetMonday
        state.dateTo = addDays(targetMonday, 7)
    },

    JUMP_TO_WEEK (state, monday) {
        state.dateFrom = monday
        state.dateTo = addDays(monday, 7)
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
    }

}

function getMonday (date) {
    let day = date.getDay() || 7
    if (day !== 1) { date.setHours(-24 * (day - 1)) }
    return date
}

function getSunday (date) {
    let day = date.getDay() || 7
    if (day !== 7) { date.setHours(24 * (7 - day)) }
    return date
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
