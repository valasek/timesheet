import api from '../../api/axiosSettings'

// initial state
const state = {
    all: [ '' ] // _id, name
}

const getters = {}

const actions = {

    getProjects ({ commit, dispatch }) {
        api.apiClient.get(`/api/projects`, { crossDomain: true })
            .then(response => {
                commit('SET_PROJECTS', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read projects from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
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
