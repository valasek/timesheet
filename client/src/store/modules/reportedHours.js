import api from '../../api/axiosSettings'

// initial state
const state = {
    all: [], // _id, date, hours, project, description, rate, consultant
    loading: true
}

const getters = {}

const actions = {
    getReportedHours ({ commit, dispatch }, month) {
        commit('SET_LOADING', true)
        state.loading = true
        let monthNumber = (new Date(month).getMonth() + 1).toString()
        let options = { year: 'numeric', month: 'long' }
        let monthText = new Intl.DateTimeFormat('en-US', options).format(new Date(month))
        api.apiClient.get('/api/reported/month/' + monthNumber)
            .then(response => {
                commit('SET_REPORTED_HOURS', response.data)
                dispatch('context/setNotification', { text: monthText + ' data retrieved', type: '' }, { root: true })
                commit('SET_LOADING', false)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read reported records from server. \n' + e.toString(), type: 'error' }, { root: true })
                commit('SET_LOADING', false)
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    removeRecord ({ commit, dispatch }, id) {
        const index = state.all.findIndex(records => records._id === id)
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
        console.log(payload) /* eslint-disable-line no-console */
        // payload.date = '2018–09–22T12:42:31+07:00'
        // console.log(payload) /* eslint-disable-line no-console */
        api.apiClient.post('/api/reported', payload)
                .then(response => {
                    commit('ADD_RECORD', payload)
                })
                .catch(e => {
                    // rollback !!!
                    dispatch('context/setNotification', { text: 'Couldn\'t duplicate selected record on server. \n' + e.toString(), type: 'error' }, { root: true })
                    console.log(e) /* eslint-disable-line no-console */
                })
        },
    updateAttributeValue ({ commit, dispatch }, payload) {
        api.apiClient.put('/api/reported/' + payload.id, payload)
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
        state.all.map(value => { value.date = value.date.substr(0, 10) })
    },
    ADD_RECORD (state, payload) {
        state.all.push(payload)
    },
    REMOVE_RECORD (state, index) {
        state.all.splice(index, 1)
    },
    UPDATE_ATTRIBUTE_VALUE (state, payload) {
        let index = state.all.findIndex(obj => obj._id === payload.id)
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
