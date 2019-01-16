import Vue from 'vue'
import Vuex from 'vuex'
import settings from './modules/settings'
import reportedHours from './modules/reportedHours'
import consultants from './modules/consultants'
import projects from './modules/projects'
import rates from './modules/rates'
import holidays from './modules/holidays'
import context from './modules/context'
import createLogger from './modules/logger'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        settings,
        reportedHours,
        consultants,
        context,
        projects,
        rates,
        holidays
    },
    strict: debug,
    plugins: debug ? [createLogger()] : []
})
