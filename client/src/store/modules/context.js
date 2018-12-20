import { addDays, subDays } from 'date-fns'

// initial state
const state = {
    page: 'Time Sheet',
    notification: false,
    notificationText: '',
    dateMonth: new Date().toISOString().substr(0, 7),
    dateFrom: getMonday(new Date()),
    dateTo: getSunday(new Date())
}

const getters = {
    // dateMonth: new Date() // new Date().toISOString().substr(0, 7)
}

const actions = {
    // getPage ({ commit, page }) {
    //     commit('SET_PAGE', page)
    // },

    setMonth ({ commit }, month) {
        commit('SET_MONTH', month)
    },

    setNotification ({ commit }, text) {
        commit('SET_NOTIFICATION', text)
    },

    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
    changeWeek ({ commit }, direction) {
        commit('SET_WEEK', direction)
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

    SET_NOTIFICATION (state, text) {
        state.notificationText = text
        state.notification = true
    },

    RESET_NOTIFICATION (state, text) {
        state.notification = false
        state.notificationText = ''
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
