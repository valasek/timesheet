// initial state
const state = {
    page: 'Time Sheet'
}

const getters = {}

const actions = {
    // getPage ({ commit, page }) {
    //     commit('SET_PAGE', page)
    // }
}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
