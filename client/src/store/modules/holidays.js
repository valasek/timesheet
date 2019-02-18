// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'

// initial state
const state = {
    all: [ '' ] // date, description
}

const getters = {}

const actions = {

    getHolidays ({ commit, dispatch }) {
        api.apiClient.get(`/api/holidays`, { crossDomain: true })
            .then(response => {
                commit('SET_HOLIDAYS', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read holidays from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    }

}

const mutations = {

    SET_HOLIDAYS (state, payload) {
        state.all = payload
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
