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
              Time sheets
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
      </v-list>

      <v-divider />

      <v-list>
        <v-list-tile>
          <div class="header ">
            Consultant
          </div>
        </v-list-tile>
        <v-list-tile>
          <v-select :items="consultants" item-text="name" item-value="name" @change="changeConsultant" />
        </v-list-tile>
      </v-list>

      <v-list>
        <v-list-tile>
          <div class="header ">
            Week
          </div>
        </v-list-tile>
        <v-list-tile>
          <v-icon @click="previousWeek">
            skip_previous
          </v-icon>
          {{ dateFrom.toLocaleDateString("en-US") }} - {{ dateTo.toLocaleDateString("en-US") }}
          <v-icon @click="nextWeek">
            skip_next
          </v-icon>
        </v-list-tile>
      </v-list>

      <v-list>
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
        2018 &copy; <strong>DataArch</strong>
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
          { title: 'Show reported work', icon: 'question_answer', route: 'reported' }
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
      dateMonth: {
        get () {
          return this.$store.state.context.dateMonth
        },
        set (newValue) {
          this.$store.dispatch('context/setMonth', newValue)
          this.$store.dispatch('reportedHours/getReportedHours', newValue)
        }
      },
      ...mapState({
        notificationText: state => state.context.notificationText,
        notificationType: state => state.context.notificationType,
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
        consultants: state => state.consultants.all,
        selectedConsultants: {
          set (newValue) {
            this.$store.dispatch('consultants/setSelected', newValue)
          },
          get () {
            return this.consultants.selected
          }
        },
        page: state => state.context.page
      })
    },

    created () {
      this.$store.dispatch('context/resetNotification')
      this.$store.dispatch('consultants/getConsultants')
      this.$store.dispatch('projects/getProjects')
      this.$store.dispatch('rates/getRates')
      this.$store.dispatch('reportedHours/getReportedHours', this.dateMonth)
    },

    methods: {
      changeConsultant (ids) {
        this.$store.dispatch('consultants/setSelected', ids)
      },

      previousWeek () {
        this.$store.dispatch('context/changeWeek', 'previous')
      },

      nextWeek () {
        this.$store.dispatch('context/changeWeek', 'next')
      }

      // computedDateFormattedDatefns () {
      //   return this.dateMonth ? format(this.dateMonth, 'dddd, MMMM Do YYYY') : ''
      // }
    }
  }
</script>
