import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: [ '' ]
}

const getters = {}

const actions = {

    getRates ({ commit }) {
        timesheet.getRates(rates => {
            commit('SET_RATES', rates)
        })
    }

}

const mutations = {

    SET_RATES (state, rates) {
        state.all = rates
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
