import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: []
}

const getters = {}

const actions = {
    getConsultants ({ commit }) {
        timesheet.getConsultants(consultants => {
            commit('SET_CONSULTANTS', consultants)
        })
    }
}

const mutations = {
    SET_CONSULTANTS (state, consultants) {
        state.all = consultants
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
