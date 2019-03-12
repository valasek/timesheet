// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { subDays, addDays, startOfWeek, endOfWeek, getMonth, parseISO } from 'date-fns'

// initial state, updated from configuration file
const state = {
    version: 'dev', // application version
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
    selectedMonth: new Date(),
    dateFrom: startOfWeek(new Date(), { weekStartsOn: 1 }),
    dateTo: endOfWeek(new Date(), { weekStartsOn: 1 })
}

const getters = {}

const actions = {

    getSettings ({ commit, dispatch, state }) {
        api.apiClient.get(`/api/settings`, { port: 3000, crossDomain: true })
            .then(response => {
                commit('SET_SETTINGS', response.data)
                // update selectedMonth, dateFrom, dateTo
                const today = new Date()
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

    changeWeek ({ dispatch, commit, rootState }, direction) {
        let oldDateFrom = state.dateFrom
        let oldDateTo = state.dateTo
        commit('SET_WEEK', direction)
        // read changed month if required
        if (getMonth(state.dateFrom) > getMonth(oldDateFrom)) {
            dispatch('reportedHours/getMonthlyData', { date: state.dateFrom, consultant: rootState.consultants.selected }, { root: true })
            commit('SET_MONTH', state.dateFrom)
        }
        if (getMonth(state.dateTo) < getMonth(oldDateTo)) {
            dispatch('reportedHours/getMonthlyData', { date: state.dateTo, consultant: rootState.consultants.selected }, { root: true })
            commit('SET_MONTH', state.dateTo)
        }
        dispatch('context/setIsCurrentWeek', {}, { root: true })
    },

    // monday
    jumpToWeek ({ commit, dispatch, rootState }, day) {
        if (getMonth(parseISO(day)) !== getMonth(state.selectedMonth)) {
            dispatch('reportedHours/getMonthlyData', { date: day, consultant: rootState.consultants.selected }, { root: true })
            commit('SET_MONTH', day)
        }
        commit('JUMP_TO_WEEK', day)
        dispatch('context/setIsCurrentWeek', {}, { root: true })
    }

}

const mutations = {

    SET_SETTINGS (state, payload) {
        state.version = payload.version
        state.dailyWorkingHoursMax = payload.dailyWorkingHoursMax
        state.dailyWorkingHoursMin = payload.dailyWorkingHoursMin
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
    SET_WEEK (state, direction) {
        switch (direction) {
            case 'previous':
                state.dateFrom = subDays(state.dateFrom, 7)
                state.dateTo = subDays(state.dateTo, 7)
                break
            case 'next':
                state.dateFrom = addDays(state.dateFrom, 7)
                state.dateTo = addDays(state.dateTo, 7)
                break
        }
    },
    // payload is any day within a week
    JUMP_TO_WEEK (state, date) {
        state.dateFrom = startOfWeek(date, { weekStartsOn: 1 })
        state.dateTo = endOfWeek(date, { weekStartsOn: 1 })
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
