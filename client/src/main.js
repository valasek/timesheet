// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

import 'babel-polyfill'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueRouter from 'vue-router'
import './plugins/vuetify'

import Report from './views/Report'
import Overview from './views/Overview'
import Holidays from './views/Holidays'
import Administration from './views/Administration'
import Help from './views/Help'
import Documentation from './views/Documentation'

Vue.use(VueRouter)
Vue.use(VueAxios, axios)

// Vue.config.productionTip = false
// Vue.config.performance = true

let router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/report',
      name: 'report',
      component: Report
    },
    {
      path: '/overview',
      name: 'overview',
      component: Overview
    },
    {
      path: '/holidays',
      name: 'holidays',
      component: Holidays
    },
    {
      path: '/administration',
      name: 'administration',
      component: Administration
    },
    {
      path: '/help',
      name: 'help',
      component: Help
    },
    {
      path: '/documentation',
      name: 'documentation',
      component: Documentation
    },
    { path: '*', redirect: '/report' }
  ]
})

new Vue({ // eslint-disable-line no-new
  el: '#app',
  store,
  router,
  render: h => h(App)
})
