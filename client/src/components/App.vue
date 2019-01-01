<template>
  <v-app>
    <v-snackbar v-model="notification" :color="notificationType" :top="false" :right="true" :timeout="4000">
      {{ notificationText }}
    </v-snackbar>

    <v-navigation-drawer app>
      <v-toolbar flat>
        <v-list>
          <v-list-tile>
            <v-list-tile-title class="title">
              Timesheet
            </v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-toolbar>

      <v-divider />

      <v-list dense class="pt-0">
        <v-list-tile v-for="item in items" :key="item.title" :to="item.route">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>

        <v-list-group prepend-icon="settings" value="true">
          <v-list-tile slot="activator">
            <v-list-tile-title>Settings</v-list-tile-title>
          </v-list-tile>

          <v-list-tile>
            <div class="header ">
              Consultant
            </div>
          </v-list-tile>
          <v-list-tile>
            <v-select v-model="selectedConsultant" :items="consultants.all" item-text="name" item-value="name" />
          </v-list-tile>

          <v-list-tile>
            <div class="header ">
              Week
            </div>
          </v-list-tile>
          <v-list-tile>
            <v-container fluid grid-list-xs>
              <v-layout align-center justify-space-between row fill-height>
                <v-flex xs1>
                  <v-icon @click="previousWeek">
                    skip_previous
                  </v-icon>
                </v-flex>
                <v-flex xs8>
                  <div class="text-xs-center">
                    {{ dateFrom.toLocaleDateString("en-US") }} - {{ dateTo.toLocaleDateString("en-US") }}
                  </div>
                </v-flex>
                <v-flex xs1>
                  <v-icon justify-end @click="nextWeek">
                    skip_next
                  </v-icon>
                </v-flex>
              </v-layout>
            </v-container>
          </v-list-tile>

          <v-list-tile>
            <div class="header ">
              Month
            </div>
          </v-list-tile>
          <v-list-tile>
            <v-menu v-model="monthMenu" :close-on-content-click="true" full-width max-width="290">
              <v-text-field slot="activator" :value="dateMonth" readonly />
              <v-date-picker v-model="dateMonth" :landscape="false" type="month" @change="monthMenu = false" />
            </v-menu>
          </v-list-tile>
          <v-list-tile>
            <v-switch v-model="previousWeeksUnLock" :label="previousWeeksUnLockText" color="error" hide-details />
          </v-list-tile>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar app>
      <v-toolbar-title class="headline">
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
    <v-footer app>
      <v-flex text-xs-center>
        2018 &copy; <strong>Stanislav Valasek</strong>
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
        monthMenu: false,
        items: [
          { title: 'Report my work', icon: 'dashboard', route: 'report' },
          { title: 'Show reported work', icon: 'question_answer', route: 'reported' },
          { title: 'State holidays', icon: 'card_travel', route: 'holidays' }
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
      dateMonth: {
        get () {
          return this.$store.state.context.dateMonth
        },
        set (newValue) {
          this.$store.dispatch('context/setMonth', newValue)
          this.$store.dispatch('reportedHours/getReportedHours', newValue)
        }
      },
      selectedConsultant: {
        set (newValue) {
          this.$store.dispatch('consultants/setSelected', newValue)
        },
        get () {
          return this.consultants.selected
        }
      },
      ...mapState({
        notificationText: state => state.context.notificationText,
        notificationType: state => state.context.notificationType,
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
        consultants: state => state.consultants,
        page: state => state.context.page
      })
    },

    created () {
      this.$store.dispatch('context/resetNotification')
      this.$store.dispatch('consultants/getConsultants')
      this.$store.dispatch('reportedHours/getReportedHours', this.dateMonth)
      this.$store.dispatch('projects/getProjects')
      this.$store.dispatch('rates/getRates')
      this.$store.dispatch('holidays/getHolidays')
    },

    methods: {

      previousWeek () {
        this.$store.dispatch('context/changeWeek', 'previous')
      },

      nextWeek () {
        this.$store.dispatch('context/changeWeek', 'next')
      }

    }
  }
</script>

<style>
/* workaround to remove spacing before and above week selector */
div .container.fluid.grid-list-xs {
  padding-left: 0px !important;
  padding-top: 0px !important;
}
</style>
