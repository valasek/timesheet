import timesheet from '../../api/timesheet'

// initial state
const state = {
    all: [ '' ] // _id, name
}

const getters = {}

const actions = {

    getProjects ({ commit }) {
        timesheet.getProjects(projects => {
            commit('SET_PROJECTS', projects)
        })
    }

}

const mutations = {

    SET_PROJECTS (state, projects) {
        state.all = projects
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
