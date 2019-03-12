<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <v-toolbar>
      <v-layout align-center justify-center row fill-height>
        <v-flex xs4>
          <change-week />
        </v-flex>
        <v-flex xs3>
          <select-consultant />
        </v-flex>
        <v-spacer />
      </v-layout>
    </v-toolbar>
    <v-container grid-list-md text-xs-center>
      <v-toolbar dense>
        <v-toolbar-title>Year - {{ thisYear }}</v-toolbar-title>
      </v-toolbar>
      <v-layout row wrap>
        <v-flex xs4>
          <v-card>
            <v-data-table :headers="headersV" :items="vacations" hide-actions class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value / dailyWorkingHours }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
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
                  {{ props.item.value / dailyWorkingHours }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
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
                  {{ props.item.value / dailyWorkingHours }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
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
          <v-toolbar dense>
            <v-toolbar-title>Week - {{ thisWeek }} </v-toolbar-title>
          </v-toolbar>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headersNumbers" :items="weeklyWorkingTimeOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value / dailyWorkingHours }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headersProjects" :items="weeklyProjectsOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value / dailyWorkingHours }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
        <v-flex xs6>
          <v-toolbar dense>
            <v-toolbar-title>Month - {{ selectedMonth | formatMonth }}</v-toolbar-title>
          </v-toolbar>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headersNumbers" :items="monthlyWorkingTimeOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value / dailyWorkingHours }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :headers="headersProjects" :items="monthlyProjectsOverview" hide-actions class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value / dailyWorkingHours }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
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
  import { format, lightFormat, getYear, getDaysInMonth, eachWeekendOfMonth } from 'date-fns'
  import selectConsultant from '../components/SelectConsultant'
  import changeWeek from '../components/ChangeWeek'

  export default {

    components: {
      'select-consultant': selectConsultant,
      'change-week': changeWeek
    },

    filters: {
      formatMonth: function (date) {
        if (!date) return ''
        return format(date, 'MMMM')
      }
    },

    data () {
      return {
        headersNumbers: [
          { text: 'Reported time', align: 'left', value: 'reported', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' }
        ],
        headersProjects: [
          { text: 'Reported projects', align: 'left', value: 'reported', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' }
        ],
        headersV: [
          { text: 'Vacations', align: 'left', value: 'vacations', sortable: false, class: 'subheading' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' }
        ],
        headersP: [
          { text: 'Personal Days', align: 'left', value: 'personalDays', sortable: false, class: 'subheading' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' }
        ],
        headersS: [
          { text: 'Sick Days', align: 'left', value: 'sickDays', sortable: false, class: 'subheading' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' }
        ]
      }
    },

    computed: {
      selectedReportedHoursWeekly () {
        let from = this.dateFrom
        let to = this.dateTo
        return this.reportedHours.filter(function (report) {
          let d = new Date(report.date)
          return (d >= from && d <= to)
        })
      },
      ...mapState({
        reportedHours: state => state.reportedHours.consultantMonthly,
        reportedHoursSummary: state => state.reportedHours.summary,
        rates: state => state.rates.all,
        projects: state => state.projects.all,
        dateFrom: state => state.settings.dateFrom,
        dateTo: state => state.settings.dateTo,
        selectedMonth: state => state.settings.selectedMonth,
        selectedConsultant: state => state.consultants.selected,
        dailyWorkingHours: state => state.settings.dailyWorkingHours,
        yearlyVacationDays: state => state.settings.yearlyVacationDays,
        vacation: state => state.settings.vacation,
        yearlyPersonalDays: state => state.settings.yearlyPersonalDays,
        vacationPersonal: state => state.settings.vacationPersonal,
        yearlySickDays: state => state.settings.yearlySickDays,
        vacationSick: state => state.settings.vacationSick,
        isWorking: state => state.settings.isWorking,
        isNonWorking: state => state.settings.isNonWorking
      }),
      thisYear () { return getYear(this.selectedMonth) },
      thisWeek () {
        return format(this.dateFrom, 'MMM d') + ' - ' + format(this.dateTo, 'MMM d')
      },
      // FIXME - subtract state holidays
      businessMonthly () {
        return getDaysInMonth(this.selectedMonth) - (eachWeekendOfMonth(this.selectedMonth).length)
      },
      // FIXME - subtract state holidays
      businessWeekly () {
        return 5
      },
      vacations () {
        return [
          {
            text: 'Total',
            value: this.yearlyVacationDays * this.dailyWorkingHours
          },
          {
            text: 'Remaining',
            value: this.yearlyVacationDays * this.dailyWorkingHours - this.getTotalsForRate(this.reportedHoursSummary, this.vacation)
          },
          {
            text: 'Reported',
            value: this.getTotalsForRate(this.reportedHoursSummary, this.vacation)
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
            value: this.yearlyPersonalDays * this.dailyWorkingHours - this.getTotalsForRate(this.reportedHoursSummary, this.vacationPersonal)
          },
          {
            text: 'Reported',
            value: this.getTotalsForRate(this.reportedHoursSummary, this.vacationPersonal)
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
            value: this.yearlySickDays * this.dailyWorkingHours - this.getTotalsForRate(this.reportedHoursSummary, this.vacationSick)
          },
          {
            text: 'Reported',
            value: this.getTotalsForRate(this.reportedHoursSummary, this.vacationSick)
          }
        ]
      },
      weeklyWorkingTimeOverview () {
        return this.workingTimeOverview('week')
      },
      weeklyProjectsOverview () {
        return this.projectsOverview('week')
      },
      monthlyWorkingTimeOverview () {
        return this.workingTimeOverview('month')
      },
      monthlyProjectsOverview () {
        return this.projectsOverview('month')
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Overview of reported hours')
      this.$store.dispatch('reportedHours/getYearlySummary', this.selectedMonth)
    },

    methods: {
      projectsOverview (period) {
        var summary = {}
        switch (period) {
          case 'week':
            summary = this.getProjectTotalsInWeek(this.selectedReportedHoursWeekly)
            break
          case 'month':
            summary = this.getProjectTotalsInMonth(this.reportedHoursSummary)
            break
          default:
            console.log('workingTimeOverview unknown period:', period) /* eslint-disable-line no-console */
        }
        return summary
      },
      workingTimeOverview (period) {
        var data = []
        var summary = {}
        switch (period) {
          case 'week':
            summary = this.getTotalsInWeek(this.selectedReportedHoursWeekly)
            data.push({
              text: 'Available working time',
              value: this.businessWeekly * this.dailyWorkingHours
            })
            break
          case 'month':
            data.push({
              text: 'Available working time',
              value: this.businessMonthly * this.dailyWorkingHours
            })
            summary = this.getTotalInMonth(this.reportedHoursSummary)
            break
          default:
            console.log('workingTimeOverview unknown period:', period) /* eslint-disable-line no-console */
        }
        data.push(
          {
            text: 'Reported total',
            value: summary.working + summary.nonWorking
          },
          {
            text: 'Reported working time',
            value: summary.working
          },
          {
            text: 'Reported non-working time',
            value: summary.nonWorking
          }
          // {
          //   remaining working time
          // }
        )
        return data
      },
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
      getTotalsForRate (hours, rate) {
        const consultant = this.selectedConsultant
        let h = hours.reduce(
          function (total, current) {
            let h = 0
            if (current.rate === rate && current.consultant === consultant) { h = current.hours }
            return total + h
          }, 0)
        return h || 0
      },
      getTotalInMonth (hours) {
        let working = 0
        let nonWorking = 0
        const consultant = this.selectedConsultant
        const month = lightFormat(this.selectedMonth, 'M')
        hours.forEach(function (element) {
          if (element.month === month && element.consultant === consultant) {
            var el = this.rates.find(o => o.name === element.rate)
            if (el !== undefined && el.type === this.isWorking) { working += element.hours }
            if (el !== undefined && el.type === this.isNonWorking) { nonWorking += element.hours }
          }
        }, this)
        return { working: working, nonWorking: nonWorking }
      },
      getTotalsInWeek (hours) {
        let working = 0
        let nonWorking = 0
        hours.forEach(function (element) {
          var el = this.rates.find(o => o.name === element.rate)
          if (el !== undefined && el.type === this.isWorking) { working += element.hours }
          if (el !== undefined && el.type === this.isNonWorking) { nonWorking += element.hours }
        }, this)
        return { working: working, nonWorking: nonWorking }
      },
      getProjectTotalsInWeek (hours) {
        const projects = this.projects.map(function (project) {
          return { text: project.name,
                   value: hours.reduce(
                     function (total, current) {
                       let h = 0
                       current.project === project.name ? h = current.hours : h = 0
                       return total + h
                     }, 0) }
        }).filter(item => item.value > 0)
        return projects
      },
      getProjectTotalsInMonth (hours) {
        const consultant = this.selectedConsultant
        const month = lightFormat(this.selectedMonth, 'M')
        const projects = this.projects.map(function (project) {
          return { text: project.name,
                   value: hours.filter(record => (record.consultant === consultant && record.month === month)).reduce(
                     function (total, current) {
                       let h = 0
                       current.project === project.name ? h = current.hours : h = 0
                       return total + h
                     }, 0)
          }
        }).filter(item => item.value > 0)
        return projects
      }
    }
  }
</script>

<style scoped>
</style>
