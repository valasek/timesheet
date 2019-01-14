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
    dailyWorkingHours: 8, // Used for weekly and monthly expected working hours
    yearlyVacationDays: 20, // Used for  weekly and monthly expected working hours
    yearlyPersonalDays: 3, // Used for  weekly and monthly expected working hours
    yearlySickDays: 2, // Used for  weekly and monthly expected working hours
    previousWeeksUnLock: false
}

const getters = {}

const actions = {

    // setMonth ({ commit, dispatch }, month) {
    //     let monday = moment.tz({}, timeZone).startOf('isoWeek')
    //     dispatch('jumpToWeek', monday)
    //     commit('SET_MONTH', month)
    // },
    setNotification ({ commit }, payload) {
        commit('SET_NOTIFICATION', payload)
    },
    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
    TogglePreviousWeeksUnLock ({ commit }) {
        commit('TOGGLE_PREVIOUS_WEEKS_UNLOCK')
    },
    changeWeek ({ dispatch, commit }, direction) {
        let oldDateFrom = state.dateFrom
        let oldDateTo = state.dateTo
        commit('SET_WEEK', direction)
        // read changed month if required
        if (state.dateFrom.isAfter(oldDateFrom, 'month')) {
            dispatch('reportedHours/getReportedHours', state.dateFrom.format('YYYY-MM'), { root: true })
        }
        if (state.dateTo.isBefore(oldDateTo, 'month')) {
            dispatch('reportedHours/getReportedHours', state.dateTo.format('YYYY-MM'), { root: true })
        }
    },
    // monday
    jumpToWeek ({ commit }, payload) {
        commit('JUMP_TO_WEEK', payload)
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
    },
    // payload is any day within a week
    JUMP_TO_WEEK (state, payload) {
        state.dateFrom = moment(payload).startOf('isoWeek')
        state.dateTo = moment(payload).endOf('isoWeek')
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

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
