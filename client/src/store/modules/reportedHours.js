import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: []
}

const getters = {}

const actions = {
    getReportedHours ({ commit }) {
        timesheet.getReportedHours(reportedHours => {
            commit('setReportedHours', reportedHours)
        })
    }
}

const mutations = {
    setReportedHours ( state, reportedHours ) {
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