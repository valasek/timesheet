<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
          <q-btn
            flat
            dense
            round
            icon="menu"
            @click="drawer = !drawer"
          >
          </q-btn>
        <q-toolbar-title>
          <router-link to="/" class="my_title">
            Timesheet
          </router-link> / {{ page }}
        </q-toolbar-title>
        <q-btn flat to="/help">Help</q-btn>
      </q-toolbar>
    </q-header>
    <q-drawer v-model="drawer" side="left" bordered>
      <q-list>
        <q-item v-for="item in items" :key="item.title" :to="item.route">
          <q-item-section avatar>
            <q-icon :name="item.icon" />
          </q-item-section>
          <q-item-section>
            <q-item-label>{{ item.title }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />
        <q-item>
          <q-item-section>
            <q-toggle v-model="weekUnlocked" :disable="isCurrentWeek===true" label="Edit this week"/>
          </q-item-section>
        </q-item>
      </q-list>
      <!-- <q-list>
        <q-item to="/report">
          <q-item-section avatar>
            <q-icon name="work_outline" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Report</q-item-label>
            <q-item-label caption>Report weekly work</q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable to="/overview">
          <q-item-section avatar>
            <q-icon name="show_chart" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Overview</q-item-label>
            <q-item-label caption>Projects and rates overview</q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable to="/holidays">
          <q-item-section avatar>
            <q-icon name="event" />
          </q-item-section>
          <q-item-section>
            <q-item-label>State Holiday</q-item-label>
            <q-item-label caption></q-item-label>
          </q-item-section>
        </q-item>
        <q-item clickable to="/administration">
          <q-item-section avatar>
            <q-icon name="settings" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Settings</q-item-label>
            <q-item-label caption></q-item-label>
          </q-item-section>
        </q-item>
      </q-list> -->
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
        { title: 'Report my work', icon: 'work_outline', route: 'report' },
        { title: 'Reported Overview', icon: 'show_chart', route: 'overview' },
        { title: 'State Holidays', icon: 'event', route: 'holidays' },
        { title: 'Administration', icon: 'settings', route: 'administration' }
      ],
      right: null
    }
  },

  computed: {
    weekUnlocked: {
      get () {
        return this.$store.state.context.weekUnlocked
      },
      set (newValue) {
        this.$store.dispatch('context/setWeekUnlocked', newValue)
      }
    },
    ...mapState({
      isCurrentWeek: state => state.context.isCurrentWeek,
      selectedMonth: state => state.settings.selectedMonth,
      page: state => state.context.page
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
.my_title {
  text-decoration: none !important;
  color: white;
}
</style>
