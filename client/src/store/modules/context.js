import moment from 'moment-timezone'

const defaultNotificationType = 'info'
const timeZone = 'Europe/Prague'

// initial state
const state = {
    page: 'Timesheet',
    notification: false,
    notificationText: '',
    notificationType: defaultNotificationType, // success, info, error - snackbar types https://vuetifyjs.com/en/components/snackbars#introduction
    dateMonth: new Date().toISOString().substr(0, 7),
    dateFrom: moment.tz({}, timeZone).startOf('isoWeek'),
    dateTo: moment.tz({}, timeZone).endOf('isoWeek'),
    dailyWorkingHours: 8,
    previousWeeksUnLock: false
}

const getters = {}

const actions = {

    setMonth ({ commit, dispatch }, month) {
        let monday = moment.tz({}, timeZone).startOf('isoWeek')
        dispatch('jumpToWeek', monday)
        commit('SET_MONTH', month)
    },
    setNotification ({ commit }, payload) {
        commit('SET_NOTIFICATION', payload)
    },
    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
    TogglePreviousWeeksUnLock ({ commit }) {
        commit('TOGGLE_PREVIOUS_WEEKS_UNLOCK')
    },
    changeWeek ({ commit }, direction) {
        commit('SET_WEEK', direction)
    }
    // jumpToWeek ({ commit }, monday) {
    //     commit('JUMP_TO_WEEK', monday)
    // }
}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    },

    SET_MONTH (state, month) {
        state.dateMonth = month
    },

    SET_WEEK (state, direction) {
        console.log(direction, state.dateFrom.format(), state.dateTo.format()) /* eslint-disable-line no-console */
        switch (direction) {
            case 'previous':
                state.dateFrom = moment(state.dateFrom).subtract(7, 'days')
                state.dateTo = moment(state.dateTo).subtract(7, 'days')
                break
            case 'next':
                state.dateFrom = moment(state.dateFrom).add(7, 'days')
                state.dateTo = moment(state.dateTo).add(7, 'days')
                break
        }
        console.log(direction, state.dateFrom.format(), state.dateTo.format()) /* eslint-disable-line no-console */
    },

    JUMP_TO_WEEK (state, month) {
        state.dateFrom = month.startOf('isoWeek')
        state.dateTo = month.endOf('isoWeek')
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

    TOGGLE_PREVIOUS_WEEKS_UNLOCK (state) {
        state.previousWeeksUnLock = !state.previousWeeksUnLock
    }

}

// function getMonday (date) {
//     let d = date.getDate()
//     let day = date.getDay() || 7
//     let m = date.getMonth()
//     let y = date.getFullYear()
//     let newDate = new Date(y, m, d)
//     console.log('newDate:', newDate) /* eslint-disable-line no-console */
//     if (day !== 1) { newDate.setHours(-24 * (day - 1)) }
//     console.log('newDate:', newDate) /* eslint-disable-line no-console */
//     return newDate
// }

// function getSunday (date) {
//     let d = date.getDate()
//     let day = date.getDay() || 7
//     let m = date.getMonth()
//     let y = date.getFullYear()
//     let newDate = new Date(y, m, d)
//     if (day !== 7) { newDate.setHours(24 * (7 - day)) }
//     return newDate
// }

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
