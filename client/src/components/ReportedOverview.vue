<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>{{ thisYear }}</v-toolbar-title>
    </v-toolbar>
    <v-container grid-list-md text-xs-center>
      <v-layout row wrap>
        <v-flex xs4>
          <v-card>
            <v-data-table :headers="headersV" :items="vacations" hide-actions class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value/dailyWorkingHours }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
        <v-flex xs4>
          <v-card>
            <v-data-table :headers="headersP" :items="personalDays" hide-actions class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value/dailyWorkingHours }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
        <v-flex xs4>
          <v-card>
            <v-data-table :headers="headersS" :items="sickDays" hide-actions class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value/dailyWorkingHours }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-container grid-list-md text-xs-center>
      <v-layout row wrap>
        <v-flex xs6>
          <v-card>
            <v-toolbar flat>
              <v-toolbar-title>{{ thisWeek }} </v-toolbar-title>
            </v-toolbar>
          </v-card>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headers" :items="weeklyOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value/dailyWorkingHours }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
        <v-flex xs6>
          <v-card>
            <v-toolbar flat>
              <v-toolbar-title>{{ thisMonth }}</v-toolbar-title>
            </v-toolbar>
          </v-card>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headers" :items="monthlyOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value/dailyWorkingHours }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import moment from 'moment-timezone'

  export default {
    name: 'ReportedOverview',

    data () {
      return {
        headers: [
          { text: '', align: 'left', value: 'reported', sortable: false },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'dayes', sortable: false, class: 'body-1' }
        ],
        headersV: [
          { text: 'Vacations', align: 'left', value: 'vacations', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'dayes', sortable: false, class: 'body-1' }
        ],
        headersP: [
          { text: 'Personal Days', align: 'left', value: 'personalDays', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'dayes', sortable: false, class: 'body-1' }
        ],
        headersS: [
          { text: 'Sick Days', align: 'left', value: 'sickDays', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'dayes', sortable: false, class: 'body-1' }
        ]
      }
    },

    computed: {
      selectedReportedHours () {
        return this.reportedHours.filter(report => {
          let d = new Date(report.date)
          return (d >= this.dateFrom && d <= this.dateTo && report.consultant === this.selectedConsultants)
        })
      },
      monthlyConsultantReportedHours () {
        return this.reportedHours.filter(report => {
          return (report.consultant === this.selectedConsultants)
        })
      },
      ...mapState({
        reportedHours: state => state.reportedHours.all,
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
        selectedConsultants: state => state.consultants.selected,
        dailyWorkingHours: state => state.context.dailyWorkingHours,
        yearlyVacationDays: state => state.context.yearlyVacationDays,
        yearlyPersonalDays: state => state.context.yearlyPersonalDays,
        yearlySickDays: state => state.context.yearlySickDays
      }),
      thisYear () { return moment(this.dateFrom).year() },
      thisWeek () {
        return moment.tz(this.dateFrom, 'Europe/Prague').format('MMMM D') + ' - ' + moment(this.dateTo).format('D')
      },
      thisMonth () { return moment(this.dateFrom).format('MMMM') },
      vacations () {
        return [
          {
            text: 'Total',
            value: this.yearlyVacationDays * this.dailyWorkingHours
          },
          {
            text: 'Remaining',
            value: this.yearlyVacationDays * this.dailyWorkingHours - this.getTotals(this.selectedReportedHours, 'Vacation')
          },
          {
            text: 'Reported',
            value: this.getTotals(this.selectedReportedHours, 'Vacation')
          }
        ]
      },
      personalDays () {
        return [
          {
            text: 'Total',
            value: this.yearlyPersonalDays * this.dailyWorkingHours
          },
          {
            text: 'Remaining',
            value: 15
          },
          {
            text: 'Reported',
            value: this.getTotals(this.selectedReportedHours, 'Personal Day')
          }
        ]
      },
      sickDays () {
        return [
          {
            text: 'Total',
            value: this.yearlySickDays * this.dailyWorkingHours
          },
          {
            text: 'Remaining',
            value: 10
          },
          {
            text: 'Reported',
            value: this.getTotals(this.selectedReportedHours, 'Sick Day')
          }
        ]
      },
      weeklyOverview () {
        return [
          {
            text: 'Available working hours',
            value: 18 * this.dailyWorkingHours
          },
          {
            text: 'Reported working hours',
            value: this.getTotals(this.selectedReportedHours, 'Off-site') + this.getTotals(this.selectedReportedHours, 'On-site')
          },
          {
            text: 'Reported working hours during holiday',
            value: this.getTotals(this.selectedReportedHours, 'Off-site Holiday')
          },
          {
            text: 'Personal days',
            value: this.getTotals(this.selectedReportedHours, 'Personal Day')
          },
          {
            text: 'Sick days',
            value: this.getTotals(this.selectedReportedHours, 'Sick Day')
          },
          {
            text: 'Vacations',
            value: this.getTotals(this.selectedReportedHours, 'Vacation')
          },
          {
            text: 'Unpaid leave',
            value: this.getTotals(this.selectedReportedHours, 'Unpaid Leave')
          },
          {
            text: 'Sick',
            value: this.getTotals(this.selectedReportedHours, 'Sick')
          },
          {
            text: 'Weekend',
            value: this.getTotals(this.selectedReportedHours, 'Off-site Weekend')
          },
          {
            text: 'Holiday',
            value: this.getTotals(this.selectedReportedHours, 'Public Holiday')
          },
          {
            text: 'Overtime',
            value: 999
          }
        ]
      },
      monthlyOverview () {
        return [
          {
            text: 'Available working hours',
            value: this.getWeekdaysInMonth('', '') * this.dailyWorkingHours
          },
          {
            text: 'Reported working hours',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site') + this.getTotals(this.monthlyConsultantReportedHours, 'On-site')
          },
          {
            text: 'Reported working hours during holiday',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site Holiday')
          },
          {
            text: 'Personal days',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Personal Day')
          },
          {
            text: 'Sick days',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Sick Day')
          },
          {
            text: 'Vacations',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Vacation')
          },
          {
            text: 'Unpaid leave',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Unpaid Leave')
          },
          {
            text: 'Sick',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Sick')
          },
          {
            text: 'Weekend',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site Weekend')
          },
          {
            text: 'Holiday',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Public Holiday')
          },
          {
            text: 'Overtime',
            value: 999
          }
        ]
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Overview of reported hours')
    },

    methods: {
      getDaysInMonth (month, year) {
        let date = new Date(year, month, 0).getDate()
        return date
      },
      isWeekday (year, month, day) {
        let date = new Date(year, month, day).getDay()
        return date !== 0 && date !== 6
      },
      getWeekdaysInMonth (month, year) {
        month = '12'
        year = '2018'
        let days = this.getDaysInMonth(month, year)
        let weekdays = 0
        for (let i = 0; i < days; i++) {
          if (this.isWeekday(year, month, i + 1)) weekdays++
        }
        return weekdays
      },
      getTotals (hours, rate) {
        let h = hours.reduce(
          function (total, current) {
            let h = 0
            if (current.rate === rate) { h = current.hours }
            return total + h
          }, 0)
        return h || ''
      }
    }
  }
</script>

<style scoped>
</style>
