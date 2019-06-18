<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar class="glossy" >
          <q-btn
            flat
            dense
            round
            icon="menu"
            @click="drawer = !drawer"
          >
          </q-btn>
          <q-breadcrumbs class="q-pa-sm text-grey" active-color="white" style="font-size: 16px">
            <q-breadcrumbs-el label="Timesheet" icon="home" to="/" />
            <q-breadcrumbs-el v-show="subPage" :label="page" :icon="pageIcon" :to="page" />
          </q-breadcrumbs>
        <q-toolbar-title/>
        <q-btn flat to="/help">Help</q-btn>
      </q-toolbar>
    </q-header>
    <q-drawer v-model="drawer" side="left" bordered>
      <q-list>
        <q-item v-for="item in items" :key="item.title" :to="item.route" exact>
          <q-item-section avatar>
            <q-icon :name="item.icon" />
          </q-item-section>
          <q-item-section>
            <q-item-label>{{ item.title }}</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'MyLayout',

  data () {
    return {
      drawer: true,
      items: [
        { title: 'Timesheet', icon: 'home', route: '/' },
        { title: 'Report', icon: 'work_outline', route: 'report' },
        { title: 'Reported Overview', icon: 'show_chart', route: 'overview' },
        { title: 'State Holidays', icon: 'event', route: 'holidays' },
        { title: 'Administration', icon: 'settings', route: 'administration' },
        { title: 'Documentation', icon: 'help_outline', route: 'documentation' }
      ],
      right: null
    }
  },

  computed: {
    subPage () {
      return this.page !== 'Timesheet'
    },
    ...mapState({
      selectedMonth: state => state.settings.selectedMonth,
      page: state => state.context.page,
      pageIcon: state => state.context.pageIcon
    })
  },

  created () {
    this.$store.dispatch('consultants/getConsultants')
    this.$store.dispatch('projects/getProjects')
    this.$store.dispatch('rates/getRates')
    this.$store.dispatch('holidays/getHolidays')
    this.$store.dispatch('settings/getSettings')
  },
  methods: {
  }
}
</script>

<style>
</style>
