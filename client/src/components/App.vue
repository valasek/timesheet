<template>
  <v-app>
    <v-snackbar v-model="notification" :color="notificationType" :top="false" :right="true" :timeout="4000">
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
          <v-switch v-model="previousWeeksUnLock" :label="previousWeeksUnLockText" color="error" hide-details class="body-1" />
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar color="blue darken-3" dark fixed clipped-left app>
      <v-toolbar-side-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title class="headline">
        Timesheet /
      </v-toolbar-title>
      <v-toolbar-title>
        <span>{{ page }}</span>
        <!-- <span class="font-weight-light">  management</span> -->
      </v-toolbar-title>
      <v-spacer />
      <v-btn flat to="help">
        <span class="mr-2">
          Help
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
    <v-footer color="blue darken-0" dark fixed app>
      <v-flex text-xs-center>
        2018 - {{ (new Date()).getFullYear() }} &copy; <strong>Stanislav Valasek</strong> Version {{ version }}
      </v-flex>
    </v-footer>
  </v-app>
</template>

<script>
  import ReportTable from './ReportTable'
  import { mapState } from 'vuex'

  export default {
    name: 'App',
    components: {
      ReportTable
    },

    data () {
      return {
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
      previousWeeksUnLock: {
        get () {
          return this.$store.state.context.previousWeeksUnLock
        },
        set (newValue) {
          this.$store.dispatch('context/TogglePreviousWeeksUnLock')
        }
      },
      previousWeeksUnLockText () {
        if (this.$store.state.context.previousWeeksUnLock) {
          return 'Previous weeks unlocked'
        } else {
          return 'Previous weeks locked'
        }
      },
      ...mapState({
        notificationText: state => state.context.notificationText,
        notificationType: state => state.context.notificationType,
        selectedMonth: state => state.context.selectedMonth,
        page: state => state.context.page,
        version: state => state.settings.version
      })
    },

    created () {
      this.$store.dispatch('context/resetNotification')
      this.$store.dispatch('consultants/getConsultants')
      this.$store.dispatch('reportedHours/getReportedHours', this.selectedMonth)
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

</style>
