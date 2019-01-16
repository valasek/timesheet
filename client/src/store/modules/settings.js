import api from '../../api/axiosSettings'

// initial state
const state = {
    version: 'dev' // application version
}

const getters = {}

const actions = {

    getSettings ({ commit, dispatch }) {
        api.apiClient.get(`/api/settings`, { port: 3000, crossDomain: true })
            .then(response => {
                commit('SET_SETTINGS', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read application settings from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    }

}

const mutations = {

    SET_SETTINGS (state, payload) {
        state.version = payload.version
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
