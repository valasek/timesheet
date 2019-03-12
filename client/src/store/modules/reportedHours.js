// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import api from '../../api/axiosSettings'
import { format, getYear, parseISO } from 'date-fns'

// initial state
const state = {
    consultantMonthly: [], // {id, date, hours, project, description, rate, consultant}
    summary: [], // {consultant, year, month, project, rate, hours }
    loading: false
}

const getters = {}

const actions = {
    // get monthly data
    getMonthlyData ({ commit, dispatch }, payload) {
        commit('SET_LOADING', true)
        api.apiClient.get('/api/reported/year/' + getYear(payload.date) + '/month/' + format(payload.date, 'MM') + '/consultant/' + payload.consultant)
            .then(response => {
                commit('SET_REPORTED_HOURS', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read reported records from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    // get yearly summary
    getYearlySummary ({ commit, dispatch }, date) {
        api.apiClient.get('/api/reported/summary/' + getYear(date))
            .then(response => {
                commit('SET_REPORTED_HOURS_SUMMARY', response.data)
            })
            .catch(e => {
                dispatch('context/setNotification', { text: 'Couldn\'t read reported records summary from server. \n' + e.toString(), type: 'error' }, { root: true })
                console.log(e) /* eslint-disable-line no-console */
            })
    },
    removeRecord ({ commit, dispatch }, id) {
        const index = state.consultantMonthly.findIndex(records => records.id === id)
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
                    payload.date = format(parseISO(payload.date), 'yyyy-MM-dd')
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
        state.consultantMonthly = reportedHours
        // convert dates to ISO strings for Vuetify
        state.consultantMonthly.map(value => { value.date = value.date.substr(0, 10) })
        state.loading = false
    },
    SET_REPORTED_HOURS_SUMMARY (state, payload) {
        state.summary = payload
    },
    ADD_RECORD (state, payload) {
        state.consultantMonthly.push(payload)
    },
    REMOVE_RECORD (state, index) {
        state.consultantMonthly.splice(index, 1)
    },
    UPDATE_ATTRIBUTE_VALUE (state, payload) {
        let index = state.consultantMonthly.findIndex(obj => obj.id === payload.id)
        switch (payload.type) {
        case 'description':
            state.consultantMonthly[index].description = payload.value
            break
        case 'hours':
            state.consultantMonthly[index].hours = payload.value
            break
        case 'rate':
            state.consultantMonthly[index].rate = payload.value
            break
        case 'project':
            state.consultantMonthly[index].project = payload.value
            break
        case 'date':
            state.consultantMonthly[index].date = payload.value
            break
        default:
            console.log('UPDATE_ATTRIBUTE_VALUE unknown attribute type', payload.type) /* eslint-disable-line no-console */
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
