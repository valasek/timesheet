<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-toolbar class="q-pa-md bg-grey-3">
      <change-week />
      <div class="q-gutter-x-md">
        <select-consultant class="q-gutter-x-md"/>
        </div>
    </q-toolbar>
    <q-toolbar class="my-toolbar q-pa-md bg-grey-3">
      <q-toolbar-title>Year - {{ thisYear }}</q-toolbar-title>
    </q-toolbar>
    <q-card flat class="row justify-around items-baseline">
      <q-card-section>
        <q-table :columns="headersV" :data="vacations" hide-bottom>
          <template slot="items" slot-scope="props">
            <td class="text-xs-left">
              {{ props.item.text }}
            </td>
            <td class="text-xs-left">
              {{ props.item.value  }}
            </td>
            <td class="text-xs-left">
              {{ props.item.value }}
            </td>
          </template>
        </q-table>
      </q-card-section>
      <q-card-section>
        <q-table :columns="headersP" :data="personalDays" hide-bottom>
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
        </q-table>
      </q-card-section>
      <q-card-section>
        <q-table :columns="headersS" :data="sickDays" hide-bottom>
          <template slot="items" slot-scope="props">
            <td class="text-xs-left">
              {{ props.item.text }}
            </td>
            <td class="text-xs-left">
              {{ props.item.value }}
            </td>
            <td class="text-xs-left">
              {{ props.item.value }}
            </td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>
    <div class="row">
      <q-toolbar class="q-pa-md bg-grey-3">
        <q-toolbar-title>Available minus reported hours in {{ selectedMonth | formatMonth }}: {{ balancePeriod }}</q-toolbar-title>
      </q-toolbar>
      <q-card flat>
        <q-card-section>
            Balance: <span :class="textColor(getBalance())">
              {{ getBalance() | pluralizeHour }}</span>
        </q-card-section>
      </q-card>
    </div>
    <div class="row justify-around items-baseline">
      <div class="column">
        <q-toolbar class="q-pa-md bg-grey-3">
          <q-toolbar-title>Week - {{ thisWeek }} </q-toolbar-title>
        </q-toolbar>
        <div class="row">
        <q-card flat>
          <q-card-section>
            <q-table :columns="headersNumbers" :data="weeklyWorkingTimeOverview" :pagination.sync="myPagination"
              hide-bottom
            >
              <template slot="items" slot-scope="props">
                <td v-if="props.item.value !== ''" class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td v-if="props.item.value !== ''" class="text-xs-left">
                  {{ props.item.value }}
                </td>
                <td v-if="props.item.value !== ''" class="text-xs-left">
                  {{ props.item.value }}
                </td>
              </template>
            </q-table>
          </q-card-section>
        </q-card>
        </div>
        <div class="row">
          <q-card flat>
            <q-card-section>
              <q-table :columns="headersProjects" :data="weeklyProjectsOverview" :pagination.sync="myPagination"
                hide-bottom
              >
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </q-table>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="column">
        <q-toolbar class="q-pa-md bg-grey-3">
          <q-toolbar-title>Month - {{ selectedMonth | formatMonth }}</q-toolbar-title>
        </q-toolbar>
        <div class="row">
          <q-card flat>
            <q-card-section>
              <q-table :columns="headersNumbers" :data="monthlyWorkingTimeOverview" :pagination.sync="myPagination"
                hide-bottom
              >
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </q-table>
            </q-card-section>
          </q-card>
        </div>
        <div class="row">
          <q-card flat>
            <q-card-section>
              <q-table :columns="headersProjects" :data="monthlyProjectsOverview" :pagination.sync="myPagination"
                hide-bottom
              >
                <template slot="items" slot-scope="props">
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                  <td v-if="props.item.value !== ''" class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </q-table>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import { format, lightFormat, getYear, differenceInDays, eachWeekendOfInterval, addDays, startOfMonth, endOfMonth } from 'date-fns'
import { workHoursMixin } from '../mixins/workHoursMixin'
// import selectConsultant from '../components/SelectConsultant'
// import changeWeek from '../components/ChangeWeek'

export default {

  components: {
    /* webpackChunkName: "core" */
    'select-consultant': () => import('components/SelectConsultant'),
    /* webpackChunkName: "core" */
    'change-week': () => import('components/ChangeWeek')
  },

  filters: {
    formatMonth: function (date) {
      if (!date) return ''
      return format(date, 'MMMM')
    },
    formatDay: function (date) {
      if (!date) return ''
      return format(date, 'MM/dd')
    },
    pluralizeHour: function (count) {
      if (count === 0) {
        return '0 hours'
      } else if (count === 1) {
        return '1 hour'
      } else {
        return count + ' hours'
      }
    }
  },

  mixins: [ workHoursMixin ],

  data () {
    return {
      balancePeriod: '',
      filter: '',
      myPagination: { 'rowsPerPage': 50 },
      headersNumbers: [
        { label: 'Reported time', align: 'left', field: 'text', sortable: false },
        { label: 'Days', align: 'left', field: 'value', format: val => `${val}` / this.dailyWorkingHours, sortable: false },
        { label: 'Hours', align: 'left', field: 'value', sortable: false }
      ],
      headersProjects: [
        { label: 'Reported projects', align: 'left', field: 'text', sortable: false },
        { label: 'Days', align: 'left', field: 'value', format: val => `${val}` / this.dailyWorkingHours, sortable: false },
        { label: 'Hours', align: 'left', field: 'value', sortable: false }
      ],
      headersV: [
        { label: 'Vacations', field: 'text', sortable: false },
        { label: 'Days', align: 'left', field: 'value', format: val => `${val}` / this.dailyWorkingHours, sortable: false },
        { label: 'Hours', align: 'left', field: 'value', sortable: false }
      ],
      headersP: [
        { label: 'Personal Days', field: 'text', sortable: false },
        { label: 'Days', align: 'left', field: 'value', format: val => `${val}` / this.dailyWorkingHours, sortable: false },
        { label: 'Hours', align: 'left', field: 'value', sortable: false }
      ],
      headersS: [
        { label: 'Sick Days', field: 'text', sortable: false },
        { label: 'Days', align: 'left', field: 'value', format: val => `${val}` / this.dailyWorkingHours, sortable: false },
        { label: 'Hours', align: 'left', field: 'value', sortable: false }
      ]
    }
  },

  computed: {
    ...mapState({
      reportedHours: state => state.reportedHours.consultantMonthly,
      reportedHoursSummary: state => state.reportedHours.summary,
      rates: state => state.rates.all,
      projects: state => state.projects.all,
      dateFrom: state => state.settings.dateFrom,
      dateTo: state => state.settings.dateTo,
      selectedMonth: state => state.settings.selectedMonth,
      selectedConsultant: state => state.consultants.selected,
      selectedAllocation: state => state.consultants.selectedAllocation,
      dailyWorkingHours: state => state.settings.dailyWorkingHours,
      yearlyVacationDays: state => state.settings.yearlyVacationDays,
      vacation: state => state.settings.vacation,
      yearlyPersonalDays: state => state.settings.yearlyPersonalDays,
      vacationPersonal: state => state.settings.vacationPersonal,
      yearlySickDays: state => state.settings.yearlySickDays,
      vacationSick: state => state.settings.vacationSick,
      isWorking: state => state.settings.isWorking,
      isNonWorking: state => state.settings.isNonWorking,
      holidays: state => state.holidays.all
    }),
    startOfMonth () {
      return startOfMonth(new Date())
    },
    yesterday () {
      return addDays(new Date(), -1)
    },
    thisYear () { return getYear(this.selectedMonth) },
    thisWeek () {
      return format(this.dateFrom, 'MMM d') + ' - ' + format(this.dateTo, 'MMM d')
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

  watch: {
    selectedConsultant (newValue, oldValue) {
      this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: this.selectedConsultant })
    }
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Overview of reported hours')
    this.$store.commit('context/SET_PAGE_ICON', 'show_chart')
    this.$store.dispatch('reportedHours/getYearlySummary', this.selectedMonth)
  },

  methods: {
    textColor (item) {
      let colorClass = ''
      if (item < 0) { colorClass = 'red--text lighten-2' }
      return colorClass
    },
    getBalance () {
      const today = new Date()
      if (this.selectedMonth.getMonth() !== today.getMonth() || format(today, 'd') === '1') {
        this.balancePeriod = this.$options.filters.formatDay(startOfMonth(this.selectedMonth)) + ' - ' + this.$options.filters.formatDay(endOfMonth(this.selectedMonth))
        return this.monthlyWorkingTimeOverview[4].value * -1
      }
      const worked = this.getTotalsInPeriod(this.selectedReportedHoursInPeriod(this.startOfMonth, this.yesterday))
      const available = this.workingDaysInPeriod(this.startOfMonth, this.yesterday)
      this.balancePeriod = this.$options.filters.formatDay(this.startOfMonth) + ' - ' + this.$options.filters.formatDay(this.yesterday)
      return (worked.working + worked.nonWorking - available * this.dailyWorkingHours) * this.selectedAllocation
    },
    workingDaysInPeriod (dFrom, dTo) {
      return (differenceInDays(dFrom, dTo) - 1) * -1 - eachWeekendOfInterval({ start: dFrom, end: dTo }).length - this.getHolidays(dFrom, dTo)
    },
    selectedReportedHoursInPeriod (from, to) {
      return this.reportedHours.filter(function (report) {
        const d = new Date(report.date)
        return (d >= from && d <= to)
      })
    },
    projectsOverview (period) {
      var summary = {}
      switch (period) {
        case 'week':
          summary = this.getProjectTotalsInWeek(this.selectedReportedHoursInPeriod(this.dateFrom, this.dateTo))
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
          summary = this.getTotalsInPeriod(this.selectedReportedHoursInPeriod(this.dateFrom, this.dateTo))
          data.push({
            text: 'Available working time',
            value: this.workingDaysInPeriod(this.dateFrom, this.dateTo) * this.dailyWorkingHours
          })
          break
        case 'month':
          data.push({
            text: 'Available working time',
            value: this.workingDaysInPeriod(startOfMonth(this.selectedMonth), endOfMonth(this.selectedMonth)) * this.dailyWorkingHours
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
          text: ' — Reported working time',
          value: summary.working
        },
        {
          text: ' — Reported non-working time',
          value: summary.nonWorking
        }
      )
      switch (period) {
        case 'week':
          data.push({
            text: 'Remaining total',
            value: this.workingDaysInPeriod(this.dateFrom, this.dateTo) * this.dailyWorkingHours - summary.working + summary.nonWorking
          })
          break
        case 'month':
          data.push({
            text: 'Remaining total',
            value: this.workingDaysInPeriod(startOfMonth(this.selectedMonth), endOfMonth(this.selectedMonth)) * this.dailyWorkingHours - summary.working + summary.nonWorking
          })
          break
        default:
          console.log('workingTimeOverview unknown period:', period) /* eslint-disable-line no-console */
      }
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
    getTotalsInPeriod (hours) {
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
.my-toolbar {
  margin-top: 1em !important;
}
</style>
