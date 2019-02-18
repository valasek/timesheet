<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

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
                  {{ props.item.value / dailyWorkingHours }}
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
                  {{ props.item.value / dailyWorkingHours }}
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
                  {{ props.item.value / dailyWorkingHours }}
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
                    {{ props.item.value / dailyWorkingHours }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
        <v-flex xs6>
          <v-card>
            <v-toolbar flat>
              <v-toolbar-title>{{ selectedMonth.format('MMMM') }}</v-toolbar-title>
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

    data () {
      return {
        headers: [
          { text: '', align: 'left', value: 'reported', sortable: false },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' }
        ],
        headersV: [
          { text: 'Vacations', align: 'left', value: 'vacations', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' }
        ],
        headersP: [
          { text: 'Personal Days', align: 'left', value: 'personalDays', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' }
        ],
        headersS: [
          { text: 'Sick Days', align: 'left', value: 'sickDays', sortable: false, class: 'subheading' },
          { text: 'Hours', align: 'left', value: 'hours', sortable: false, class: 'body-1' },
          { text: 'Days', align: 'left', value: 'days', sortable: false, class: 'body-1' }
        ]
      }
    },

    computed: {
      selectedReportedHoursWeekly () {
        let from = this.dateFrom
        let to = this.dateTo
        let consultant = this.selectedConsultant
        return this.reportedHours.filter(function (report) {
          let d = new Date(report.date)
          return (d >= from && d <= to && report.consultant === consultant)
        })
      },
      // summaryRates () {
      //   let allowProjects = ['_Public Holiday', '_Personal Day', '_Sick Day']
      //   return this.reportedHoursSummary.filter(record => {
      //     return (record.consultant === this.selectedConsultant && allowProjects.includes(record.project))
      //   })
      // },
      ...mapState({
        reportedHours: state => state.reportedHours.all,
        reportedHoursSummary: state => state.reportedHours.summary,
        rates: state => state.rates.all,
        dateFrom: state => state.settings.dateFrom,
        dateTo: state => state.settings.dateTo,
        selectedMonth: state => state.settings.selectedMonth,
        timeZone: state => state.settings.timeZone,
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
      thisYear () { return this.selectedMonth.year() },
      thisWeek () {
        return moment.tz(this.dateFrom, this.timeZone).format('MMM D') + ' - ' + moment(this.dateTo).format('MMM D')
      },
      // FIXME - subtract state holidays
      businessMonthly () {
        var param = moment.tz(this.selectedMonth, this.timeZone).startOf('month')
        var param1 = moment.tz(this.selectedMonth, this.timeZone).endOf('month')
        var signal = param.unix() < param1.unix() ? 1 : -1
        var start = moment.min(param, param1).clone()
        var end = moment.max(param, param1).clone()
        var startOffset = start.day() - 7
        var endOffset = end.day()

        var endSunday = end.clone().subtract(endOffset, 'd')
        var startSunday = start.clone().subtract(startOffset, 'd')
        var weeks = endSunday.diff(startSunday, 'days') / 7

        startOffset = Math.abs(startOffset)
        if (startOffset === 7) {
          startOffset = 5
        } else {
          if (startOffset === 1) {
            startOffset = 0
          } else {
            startOffset -= 2
          }
        }
        if (endOffset === 6) {
          endOffset--
        }
        return signal * (weeks * 5 + startOffset + endOffset)
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
      weeklyOverview () {
        var data = []
        data.push({
          text: 'Available working time',
          value: this.businessWeekly * this.dailyWorkingHours
        })
        let summary = this.getTotalsInWeek(this.selectedReportedHoursWeekly)
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
      monthlyOverview () {
        var data = []
        data.push({
          text: 'Available working time',
          value: this.businessMonthly * this.dailyWorkingHours
        })
        let summary = this.getTotalInMonth(this.reportedHoursSummary)
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
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Overview of reported hours')
      this.$store.dispatch('reportedHours/getYearlySummary', this.selectedMonth)
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
        const month = this.selectedMonth.format('M')
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
      }
    }
  }
</script>

<style scoped>
</style>
