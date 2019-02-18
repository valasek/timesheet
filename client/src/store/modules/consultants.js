// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'

// initial state
const state = {
    all: [],
    selected: ''
}

const getters = {}

const actions = {

    getConsultants ({ commit, dispatch }) {
        api.apiClient.get('/api/consultants', { crossDomain: true })
        .then(response => {
            commit('SET_CONSULTANTS', response.data)
        })
        .catch(e => {
            dispatch('context/setNotification', { text: 'Couldn\'t read consultants from server. \n' + e.toString(), type: 'error' }, { root: true })
            console.log(e) /* eslint-disable-line no-console */
        })
    },

    setSelected ({ commit }, selectedConsultant) {
        commit('SET_SELECTED', selectedConsultant)
    }

}

const mutations = {

    SET_CONSULTANTS (state, consultants) {
        state.all = consultants
        let c = localStorage.getItem('selectedConsultant')
        if (c && c.length > 0) {
            state.selected = c
        } else {
            state.selected = consultants[0].name
        }
    },

    SET_SELECTED (state, selectedConsultant) {
        state.selected = selectedConsultant
        localStorage.setItem('selectedConsultant', selectedConsultant)
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
