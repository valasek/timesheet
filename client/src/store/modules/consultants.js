import timesheet from '../../api/timesheet'

var allConsultants = { _id: 0, name: 'All' }

// initial state
const state = {
    all: [ allConsultants ],
    selected: [ allConsultants ]
}

const getters = {}

const actions = {

    getConsultants ({ commit }) {
        timesheet.getConsultants(consultants => {
            consultants.unshift(allConsultants)
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
