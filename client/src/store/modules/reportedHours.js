import { addDays, subDays } from 'date-fns'
import axios from 'axios'

// initial state
const state = {
    all: [], // _id, date, urs, project, description, rate, consultant
    dateFrom: getMonday(new Date()),
    dateTo: getSunday(new Date())
}

const getters = {
}

const actions = {
    getReportedHours ({ commit }, month) {
        axios.get('http://localhost:3000/api/reported/month/' + month, { crossDomain: true })
            .then(response => {
                commit('SET_REPORTED_HOURS', response.data)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    removeRecord ({ commit }, id) {
        const index = state.all.findIndex(records => records._id === id)
        axios.delete('http://localhost:3000/api/reported/' + id, { crossDomain: true })
            .then(response => {
                commit('REMOVE_RECORD', index)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    changeWeek ({ commit }, direction) {
        commit('SET_WEEK', direction)
    }
}

const mutations = {
    SET_REPORTED_HOURS (state, reportedHours) {
        state.all = reportedHours
    },
    REMOVE_RECORD (state, index) {
        state.all.splice(index, 1)
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
