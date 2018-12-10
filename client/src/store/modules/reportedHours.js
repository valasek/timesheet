import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: []
}

const getters = {}

const actions = {
    getReportedHours ({ commit }) {
        timesheet.getReportedHours(reportedHours => {
            commit('SET_REPORTED_HOURS', reportedHours)
        })
    }
}

const mutations = {
    SET_REPORTED_HOURS ( state, reportedHours ) {
        state.all = reportedHours
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}