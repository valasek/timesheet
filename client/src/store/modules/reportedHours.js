import timesheet from '../../api/timesheet'
import { addDays, subDays } from 'date-fns'

// initial state
const state = {
    all: [],
    dateFrom: getMonday(new Date()),
    dateTo: getSunday(new Date())
}

const getters = {
}

const actions = {
    getReportedHours ({ commit }, month) {
        console.log('month ' + month) /* eslint-disable-line no-console */
        timesheet.getReportedHours(reportedHours => {
            commit('SET_REPORTED_HOURS', reportedHours)
            // commit('SET_WEEKLY_HOURS', state.all)
        })
    },
    changeWeek ({ commit }, direction) {
        commit('SET_WEEK', direction)
        // commit('SET_WEEKLY_HOURS', state.all)
    }
}

const mutations = {
    SET_REPORTED_HOURS (state, reportedHours) {
        state.all = reportedHours
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
