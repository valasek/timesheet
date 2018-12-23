import axios from 'axios'

const apiClient = axios.create({
    baseURL: `http://localhost:3000`,
    withCredentials: false, // This is the default
    crossDomain: true,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    timeout: 10000
})

// initial state
const state = {
    all: [] // _id, date, urs, project, description, rate, consultant
}

const getters = {}

const actions = {
    getReportedHours ({ commit, dispatch }, month) {
        let monthNumber = (new Date(month).getMonth() + 1).toString()
        let options = { year: 'numeric', month: 'long' }
        let monthText = new Intl.DateTimeFormat('en-US', options).format(new Date(month))
        apiClient.get('/api/reported/month/' + monthNumber)
            .then(response => {
                commit('SET_REPORTED_HOURS', response.data)
                dispatch('context/setNotification', monthText + ' data retrieved', { root: true })
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    removeRecord ({ commit }, id) {
        const index = state.all.findIndex(records => records._id === id)
        apiClient.delete('/api/reported/' + id)
            .then(response => {
                commit('REMOVE_RECORD', index)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    // here should be payload not id and record
    updateRecord ({ commit }, id, record) {
        apiClient.put('/api/reported/' + id, record)
            .then(response => {
                commit('UPDATE_RECORD', record)
            })
            .catch(e => {
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    updateProject ({ commit }, payload) {
        console.log('updateProject', payload) /* eslint-disable-line no-console */
        commit('UPDATE_PROJECT', payload)
    }
}

const mutations = {
    SET_REPORTED_HOURS (state, reportedHours) {
        state.all = reportedHours
    },
    REMOVE_RECORD (state, index) {
        state.all.splice(index, 1)
    },
    UPDATE_PROJECT (state, payload) {
        console.log(payload._id, payload.project) /* eslint-disable-line no-console */
        let index = state.all.findIndex(obj => obj._id === payload._id)
        state.all[index].project = payload.project
        console.log(index, payload.project) /* eslint-disable-line no-console */
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
