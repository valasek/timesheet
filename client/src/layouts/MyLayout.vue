<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-primary text-secondary">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          @click="drawer = !drawer"
        />
        <img :src="logo">
        <q-toolbar-title />
        <q-breadcrumbs class="q-pa-sm text-secondary absolute-center" active-color="white" style="font-size: 16px">
          <template v-slot:separator>
            <q-icon
              v-show="subPage"
              size="1.2em"
              name="arrow_forward"
              color="secondary"
            />
          </template>
          <q-breadcrumbs-el label="Timesheet" icon="home" to="/" class="text-secondary" />
          <q-breadcrumbs-el v-show="subPage" :label="page" :icon="pageIcon" :to="page" />
        </q-breadcrumbs>
        <q-icon :name="darkIcon" size="sm" @click="dark = !dark">
          <q-tooltip>
            Toggle dark / light mode
          </q-tooltip>
        </q-icon>

        <q-btn flat to="/help">
          Help
        </q-btn>
      </q-toolbar>
    </q-header>
    <q-drawer v-model="drawer" side="left" bordered :width="220" content-class="text-secondary">
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
      right: null,
      logo: ''
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
    }),
    dark: {
      get () {
        return this.$q.dark.isActive
      },
      set (val) {
        this.darkIcon = 'brightness_high'
        this.$q.dark.set(val)
      }
    },
    darkIcon () {
      if (this.dark === true) {
        return 'brightness_high'
      } else {
        return 'brightness_low'
      }
    }
  },

  created () {
    this.$store.dispatch('consultants/getConsultants')
    this.$store.dispatch('projects/getProjects')
    this.$store.dispatch('rates/getRates')
    this.$store.dispatch('holidays/getHolidays')
    this.$store.dispatch('settings/getSettings')
    this.logo = 'statics/' + process.env.APP_LOGO
  },

  methods: {
  }
}
</script>

<style>
</style>
