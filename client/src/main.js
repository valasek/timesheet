import 'babel-polyfill'
import Vue from 'vue'
import App from './components/App.vue'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueRouter from 'vue-router'
import './plugins/vuetify'

import ReportTable from './components/ReportTable'
import ShowReportedTime from './components/ShowReportedTime'
import Help from './components/Help'

Vue.use(VueRouter)
Vue.use(VueAxios, axios)

Vue.config.productionTip = false

let router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'report',
      component: ReportTable
    },
    {
      path: '/reported',
      name: 'reported',
      component: ShowReportedTime
    },
    {
      path: '/help',
      name: 'help',
      component: Help
    },
    { path: '*', redirect: '/' }
  ]
})

new Vue({ // eslint-disable-line no-new
  el: '#app',
  store,
  router,
  render: h => h(App)
})
