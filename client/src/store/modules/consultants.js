import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: [ ],
    selected: ''
}

const getters = {}

const actions = {

    getConsultants ({ commit }) {
        timesheet.getConsultants(consultants => {
            commit('SET_CONSULTANTS', consultants)
        })
    },

    setSelected ({ commit }, selectedConsultant) {
        commit('SET_SELECTED', selectedConsultant)
    }

}

const mutations = {

    SET_CONSULTANTS (state, consultants) {
        state.all = consultants
    },

    SET_SELECTED (state, selectedConsultants) {
        state.selected = selectedConsultants
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
