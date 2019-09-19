// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import Vue from 'vue'
import Vuex from 'vuex'
import settings from './modules/settings'
import reportedHours from './modules/reportedHours'
import consultants from './modules/consultants'
import projects from './modules/projects'
import rates from './modules/rates'
import holidays from './modules/holidays'
import context from './modules/context'
import createLogger from 'vuex/dist/logger'

Vue.use(Vuex)

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation
 */

const debug = !process.env.NODE_ENV.startsWith('production')

export default function (/* { ssrContext } */) {
  const Store = new Vuex.Store({
    modules: {
      settings,
      reportedHours,
      consultants,
      context,
      projects,
      rates,
      holidays
    },

    // enable strict mode (adds overhead!)
    // for dev mode only
    strict: debug,
    plugins: debug ? [createLogger()] : []
  })

  return Store
}
