import moment from 'moment-timezone'

const defaultNotificationType = 'info'
const timeZone = 'Europe/Prague'

// initial state
const state = {
    page: 'Timesheet',
    notification: false,
    notificationText: '',
    notificationType: defaultNotificationType, // success, info, error - snackbar types https://vuetifyjs.com/en/components/snackbars#introduction
    selectedMonth: moment.tz({}, timeZone),
    dateFrom: moment.tz({}, timeZone).startOf('isoWeek'),
    dateTo: moment.tz({}, timeZone).endOf('isoWeek'),
    dailyWorkingHours: 8, // Used for weekly and monthly expected working hours
    dailyWorkingHoursMin: 8, // Used to highlight if reported less
    dailyWorkingHoursMax: 12, // Used to highlight if reported more
    yearlyVacationDays: 20, // Used for  weekly and monthly expected working hours
    yearlyPersonalDays: 3, // Used for  weekly and monthly expected working hours
    yearlySickDays: 2, // Used for  weekly and monthly expected working hours
    isWorking: 'work',
    isNonWorking: 'not-work',
    previousWeeksUnLock: false
}

const getters = {}

const actions = {

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
            dispatch('reportedHours/getMonthlyData', state.dateFrom, { root: true })
            commit('SET_MONTH', state.dateFrom)
        }
        if (state.dateTo.isBefore(oldDateTo, 'month')) {
            dispatch('reportedHours/getMonthlyData', state.dateTo, { root: true })
            commit('SET_MONTH', state.dateTo)
        }
    },
    // monday
    jumpToWeek ({ commit, dispatch }, payload) {
        if (payload.month() !== state.selectedMonth.month()) {
            dispatch('reportedHours/getMonthlyData', payload, { root: true })
            commit('SET_MONTH', payload)
        }
        commit('JUMP_TO_WEEK', payload)
    }
}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    },

    SET_MONTH (state, month) {
        state.selectedMonth = month
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
