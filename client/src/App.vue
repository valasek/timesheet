<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <v-app :dark="goDark">
    <v-snackbar v-model="notification" :color="notificationType" :left="true" :timeout="5000">
      {{ notificationText }}
    </v-snackbar>

    <v-navigation-drawer v-model="drawer" clipped fixed app>
      <v-list dense class="pt-0">
        <v-list-tile v-for="item in items" :key="item.title" :to="item.route">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>

        <v-divider class="menuSettings" />

        <v-list-tile class="menuSettings">
          <v-switch v-model="weekUnlocked" :disabled="isCurrentWeek===true" label="Enable editing of this week" color="info" hide-details class="body-1" />
        </v-list-tile>
        <v-list-tile class="menuSettings">
          <v-switch v-model="goDark" label="Enable dark mode" color="info" hide-details class="body-1" />
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar color="blue darken-3" dense dark fixed clipped-left app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title class="headline">
        <router-link to="/" class="my_title">
          Timesheet
        </router-link> /
      </v-toolbar-title>
      <v-toolbar-title>
        <span>{{ page }}</span>
      </v-toolbar-title>
      <v-spacer />
      <v-btn flat to="about">
        <span class="mr-2">
          About
        </span>
      </v-btn>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <router-view>
          <report-table />
        </router-view>
      </v-container>
    </v-content>
    <!-- <v-footer color="blue darken-0" dark fixed app>
      <v-flex text-xs-center>
        2018 - {{ (new Date()).getFullYear() }} &copy; <strong>Stanislav Valasek</strong> Version {{ version }}
      </v-flex>
    </v-footer> -->
  </v-app>
</template>

<script>
  import ReportTable from './views/ReportTable'
  import { mapState } from 'vuex'

  export default {
    components: {
      ReportTable
    },

    data () {
      return {
        goDark: false,
        drawer: true,
        items: [
          { title: 'Report my work', icon: 'work_outline', route: 'report' },
          { title: 'Reported overview', icon: 'show_chart', route: 'overview' },
          { title: 'State holidays', icon: 'event', route: 'holidays' },
          { title: 'Administration', icon: 'settings', route: 'administration' }
        ],
        right: null
      }
    },

    computed: {
      notification: {
        get () {
          return this.$store.state.context.notification
        },
        set (newValue) {
          this.$store.dispatch('context/resetNotification')
        }
      },
      weekUnlocked: {
        get () {
          return this.$store.state.context.weekUnlocked
        },
        set (newValue) {
          this.$store.dispatch('context/setWeekUnlocked', newValue)
        }
      },
      ...mapState({
        notificationText: state => state.context.notificationText,
        notificationType: state => state.context.notificationType,
        isCurrentWeek: state => state.context.isCurrentWeek,
        selectedMonth: state => state.settings.selectedMonth,
        page: state => state.context.page
      })
    },

    created () {
      this.$store.dispatch('context/resetNotification')
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
/* workaround to remove spacing before and above week selector */
.menuSettings {
  padding-top: 1px !important;
  padding-bottom: 1px !important;
}

.my_title {
  text-decoration: none !important;
  color: white;
}
</style>
