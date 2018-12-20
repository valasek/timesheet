// initial state
const state = {
    page: 'Time Sheet',
    notification: false,
    notificationText: ''
}

const getters = {}

const actions = {
    // getPage ({ commit, page }) {
    //     commit('SET_PAGE', page)
    // },

    setNotification ({ commit }, text) {
        commit('SET_NOTIFICATION', text)
    },

    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    }

}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    },

    SET_NOTIFICATION (state, text) {
        state.notificationText = text
        state.notification = true
    },

    RESET_NOTIFICATION (state, text) {
        state.notification = false
        state.notificationText = ''
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
