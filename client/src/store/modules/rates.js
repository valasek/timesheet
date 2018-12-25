import api from '../../api/axiosSettings'

// initial state
const state = {
    all: [ '' ], // all available rates
    allowed: [ '' ] // for selected project
}

const getters = {}

const actions = {

    getRates ({ commit, dispatch }) {
        api.apiClient.get(`/api/rates/list`, { crossDomain: true })
            .then(response => {
                commit('SET_RATES', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read projects from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
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
