const defaultNotificationType = 'info'

// initial state
const state = {
    page: 'Timesheet',
    notification: false,
    notificationText: '',
    notificationType: defaultNotificationType, // success, info, error - snackbar types https://vuetifyjs.com/en/components/snackbars#introduction
    previousWeeksUnLock: false
}

const getters = {}

const actions = {

    setNotification ({ commit }, payload) {
        commit('SET_NOTIFICATION', payload)
    },
    resetNotification ({ commit }) {
        commit('RESET_NOTIFICATION')
    },
    TogglePreviousWeeksUnLock ({ commit }) {
        commit('TOGGLE_PREVIOUS_WEEKS_UNLOCK')
    }
}

const mutations = {

    SET_PAGE (state, page) {
        state.page = page
    },

    SET_NOTIFICATION (state, payload) {
        state.notificationText = payload.text
        if (payload.type.length > 1) {
            state.notificationType = payload.type
        }
        state.notification = true
    },

    RESET_NOTIFICATION (state) {
        state.notification = false
        state.notificationText = ''
        state.notificationType = defaultNotificationType
    },

    TOGGLE_PREVIOUS_WEEKS_UNLOCK (state) {
        state.previousWeeksUnLock = !state.previousWeeksUnLock
    }

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
