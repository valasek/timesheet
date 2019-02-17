import moment from 'moment-timezone'
import api from '../../api/axiosSettings'

// initial state, updated from configuration file
const state = {
    version: 'dev', // application version
    timeZone: 'Europe/Prague', // time zone for all dates
    dailyWorkingHours: 8, // Used for weekly and monthly expected working hours
    dailyWorkingHoursMin: 8, // Used to highlight if reported less
    dailyWorkingHoursMax: 12, // Used to highlight if reported more
    yearlyVacationDays: 20, // Used for  weekly and monthly expected working hours
    vacation: 'Vacation', // Rate for vacations
    yearlyPersonalDays: 3, // Used for  weekly and monthly expected working hours
    vacationPersonal: 'Vacation Personal', // Rate for additonal vacations
    yearlySickDays: 2, // Used for  weekly and monthly expected working hours
    vacationSick: 'Vacation Sick', // Rate for additonal vacations as sick days
    isWorking: 'work', // all rates are categorized ising one of this two rates used on Overview page
    isNonWorking: 'not-work', // all rates are categorized ising one of this two rates used on Overview page
    selectedMonth: moment.tz({}, 'Europe/Prague'),
    dateFrom: moment.tz({}, 'Europe/Prague').startOf('isoWeek'),
    dateTo: moment.tz({}, 'Europe/Prague').endOf('isoWeek')
}

const getters = {}

const actions = {

    getSettings ({ commit, dispatch, state }) {
        api.apiClient.get(`/api/settings`, { port: 3000, crossDomain: true })
            .then(response => {
                commit('SET_SETTINGS', response.data)
                // update selectedMonth, dateFrom, dateTo usingrespecting retrieved timezone
                const today = moment.tz({}, state.timeZone)
                commit('SET_MONTH', today)
                commit('JUMP_TO_WEEK', today)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read application settings from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    },

    setHours ({ commit }, payload) {
        commit('SET_HOURS', payload)
    },

    setDays ({ commit }, payload) {
        commit('SET_DAYS', payload)
    },

    setRate ({ commit }, payload) {
        commit('SET_RATE', payload)
    },

    setRateType ({ commit }, payload) {
        commit('SET_RATE_TYPE', payload)
    },

    setTimeZone ({ commit }, payload) {
        commit('SET_TIME_ZONE', payload)
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

    SET_SETTINGS (state, payload) {
        state.version = payload.version
        state.dailyWorkingHoursMax = payload.dailyWorkingHoursMax
        state.dailyWorkingHoursMin = payload.dailyWorkingHoursMin
        state.timeZone = payload.timeZone
        state.dailyWorkingHours = payload.dailyWorkingHours
        state.yearlyVacationDays = payload.yearlyVacationDays
        state.vacation = payload.vacation
        state.yearlyPersonalDays = payload.yearlyPersonalDays
        state.vacationPersonal = payload.vacationPersonal
        state.yearlySickDays = payload.yearlySickDays
        state.vacationSick = payload.vacationSick
        state.isWorking = payload.isWorking
        state.isNonWorking = payload.isNonWorking
    },
    SET_HOURS (state, hours) {
        switch (hours.hourType) {
            case 'dailyWorkingHours':
            state.dailyWorkingHours = hours.hourValue
            break
            case 'dailyWorkingHoursMin':
            state.dailyWorkingHoursMin = hours.hourValue
            break
            case 'dailyWorkingHoursMax':
            state.dailyWorkingHoursMax = hours.hourValue
            break
            default:
            console.log('SET_HOURS unknown value', hours) /* eslint-disable-line no-console */
        }
    },
    SET_DAYS (state, days) {
        switch (days.dayType) {
            case 'yearlyVacationDays':
            state.yearlyVacationDays = days.dayValue
            break
            case 'yearlyPersonalDays':
            state.yearlyPersonalDays = days.dayValue
            break
            case 'yearlySickDays':
            state.yearlySickDays = days.dayValue
            break
            default:
            console.log('SET_DAYS unknown value', days) /* eslint-disable-line no-console */
        }
    },
    SET_RATE (state, rate) {
        switch (rate.rateType) {
            case 'vacation':
            state.vacation = rate.rateValue
            break
            case 'vacationPersonal':
            state.vacationPersonal = rate.rateValue
            break
            case 'vacationSick':
            state.vacationSick = rate.rateValue
            break
            default:
            console.log('SET_RATE unknown value', rate) /* eslint-disable-line no-console */
        }
    },
    SET_RATE_TYPE (state, rateType) {
        switch (rateType.rateTypeType) {
            case 'isWorking':
            state.isWorking = rateType.rateTypeValue
            break
            case 'isNonWorking':
            state.isNonWorking = rateType.rateTypeValue
            break
            default:
            console.log('SET_RATE_TYPE unknown value', rateType) /* eslint-disable-line no-console */
        }
    },
    SET_MONTH (state, month) {
        state.selectedMonth = month
    },
    SET_TIME_ZONE (state, timeZone) {
        state.timeZone = timeZone
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
    JUMP_TO_WEEK (state, date) {
        state.dateFrom = moment(date).startOf('isoWeek')
        state.dateTo = moment(date).endOf('isoWeek')
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
