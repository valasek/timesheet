import api from '../../api/axiosSettings'
import moment from 'moment'

// initial state
const state = {
    all: [], // {id, date, hours, project, description, rate, consultant}
    summary: [], // {consultant, year, month, rate, hours }
    loading: true
}

const getters = {}

const actions = {
    getReportedHours ({ commit, dispatch }, yearMonth) {
        commit('SET_LOADING', true)
        state.loading = true
        let [year, month] = yearMonth.split('-')
        // get yearly summary
        api.apiClient.get('/api/reported/summary')
            .then(response => {
                commit('SET_REPORTED_HOURS_SUMMARY', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read reported records summary from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
        // get monthly data
        api.apiClient.get('/api/reported/year/' + year + '/month/' + month)
            .then(response => {
                commit('SET_REPORTED_HOURS', response.data)
                dispatch('context/setNotification', { text: moment.tz(yearMonth, 'Europe/Prague').format('MMMM') + ' ' + year + ' data', type: '' }, { root: true })
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read reported records from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
            commit('SET_LOADING', false)
    },
    removeRecord ({ commit, dispatch }, id) {
        const index = state.all.findIndex(records => records.id === id)
        api.apiClient.delete('/api/reported/' + id)
            .then(response => {
                commit('REMOVE_RECORD', index)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t remove selected record from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    addRecord ({ commit, dispatch }, payload) {
        api.apiClient.post('/api/reported', payload)
                .then(response => {
                    // format date for Vuetify which works with ISO format
                    payload.date = moment(payload.date).format('YYYY-MM-DD')
                    payload.id = response.data.id
                    commit('ADD_RECORD', payload)
                })
                .catch(e => {
                    // rollback !!!
                    dispatch('context/setNotification', { text: 'Couldn\'t duplicate selected record on server. \n' + e.toString(), type: 'error' }, { root: true })
                    console.log(e) /* eslint-disable-line no-console */
                })
        },
    updateAttributeValue ({ commit, dispatch }, payload) {
        let payloadDB = {
            id: payload.id,
            type: payload.type,
            value: String(payload.value)
        }
        api.apiClient.put('/api/reported/' + payload.id, payloadDB)
            .then(response => {
                commit('UPDATE_ATTRIBUTE_VALUE', payload)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t update' + payload.type + 'on server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    }
}

const mutations = {
    SET_REPORTED_HOURS (state, reportedHours) {
        state.all = reportedHours
        // convert dates to ISO strings for Vuetify
        state.all.map(value => { value.date = value.date.substr(0, 10) })
    },
    SET_REPORTED_HOURS_SUMMARY (state, payload) {
        state.summary = payload
        // convert dates to ISO strings for Vuetify
        // state.all.map(value => { value.date = value.date.substr(0, 10) })
    },
    ADD_RECORD (state, payload) {
        state.all.push(payload)
    },
    REMOVE_RECORD (state, index) {
        state.all.splice(index, 1)
    },
    UPDATE_ATTRIBUTE_VALUE (state, payload) {
        let index = state.all.findIndex(obj => obj.id === payload.id)
        switch (payload.type) {
        case 'description':
            state.all[index].description = payload.value
            break
        case 'hours':
            state.all[index].hours = payload.value
            break
        case 'rate':
            state.all[index].rate = payload.value
            break
        case 'project':
            state.all[index].project = payload.value
            break
        case 'date':
            state.all[index].date = payload.value
            break
        default:
            console.log('UPDATE_ATTRIBUTE_VALUE unknown attribute type', payload) /* eslint-disable-line no-console */
        }
    },
    SET_LOADING (state, payload) {
        state.loading = payload
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}
