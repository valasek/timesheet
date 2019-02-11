import 'babel-polyfill'
import Vue from 'vue'
import App from './components/App.vue'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueRouter from 'vue-router'
import './plugins/vuetify'

import ReportTable from './components/ReportTable'
import ReportedOverview from './components/ReportedOverview'
import Holidays from './components/Holidays'
import Administration from './components/Administration'
import Help from './components/Help'

Vue.use(VueRouter)
Vue.use(VueAxios, axios)

Vue.config.productionTip = false

let router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/report',
      name: 'report',
      component: ReportTable
    },
    {
      path: '/overview',
      name: 'overview',
      component: ReportedOverview
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
    { path: '*', redirect: '/report' }
  ]
})

new Vue({ // eslint-disable-line no-new
  el: '#app',
  store,
  router,
  render: h => h(App)
})
