import Vue from 'vue'
import Vuex from 'vuex'
import reportedHours from './modules/reportedHours'
import consultants from './modules/consultants'
import context from './modules/context'
import createLogger from './modules/logger'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        reportedHours,
        consultants,
        context
    },
    strict: debug,
    plugins: debug ? [createLogger()] : []
})
