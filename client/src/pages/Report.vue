<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-toolbar class="q-pa-md bg-grey-3">
      <change-week />
      <div class="q-gutter-x-md">
        <select-consultant class="q-gutter-x-md"/>
      </div>
      <q-toolbar-title>
        <q-input v-model="filter" dense clearable label="Search" single-line>
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </q-toolbar-title>
      <q-toggle v-model="weekUnlocked" :disable="isCurrentWeek===true">
        <q-tooltip>
          Edit this week
        </q-tooltip>
      </q-toggle>
      <q-btn color="primary" :disabled="!weekUnlocked" rounded label="new" icon="add" @click="addItem" />
    </q-toolbar>
    <q-toolbar class="q-pa-md bg-grey-2">
      <span class="text-subtitle1 ">
        Weekly: <span :class="textColorWeek(reportedThisWeek)">
          {{ reportedThisWeek }} hrs
        </span>
      </span>
      <q-space />
      <div class="text-subtitle1">
        Mon: <span :class="textColor(reportedOnMonday)">
          {{ reportedOnMonday }} hrs
        </span>
      </div>
      <div class="text-subtitle1 my-hours">
        Tue: <span :class="textColor(reportedOnTuesday)">
          {{ reportedOnTuesday }} hrs
        </span>
      </div>
      <span class="text-subtitle1 my-hours">
        Wed: <span :class="textColor(reportedOnWednesday)">
          {{ reportedOnWednesday }} hrs
        </span>
      </span>
      <span class="text-subtitle1 my-hours">
        Thu: <span :class="textColor(reportedOnThursday)">
          {{ reportedOnThursday }} hrs
        </span>
      </span>
      <span class="text-subtitle1 my-hours">
        Fri: <span :class="textColor(reportedOnFriday)">
          {{ reportedOnFriday }} hrs
        </span>
      </span>
      <span class="text-subtitle1 my-hours">
        Weekend: {{ reportedOnWeekend }} hrs
      </span>
    </q-toolbar>
    <q-table :columns="headers" row-key="name" :data="selectedReportedHours" :filter="filter" :loading="loading"
      no-data-label="No hours reported this week" :pagination.sync="myPagination" :rows-per-page-options="[30,50,0]"
      binary-state-sort dense bordered>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td key="date" :props="props">
            <span  v-if="!weekUnlocked">
              {{ props.row.date | formatDate }}
            </span>
            <span v-else>
              <q-input :value="props.row.date | formatDate" dense
                @input="val => onUpdateDate({id: props.row.id, date: val.name})"
              >
                <template v-slot:prepend>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy ref="qDateProxy" transition-show="scale" transition-hide="scale">
                      <q-date :value="props.row.date" @input="(val) => onDate(props.row.id, val)"
                        mask="YYYY-MM-DD" :rules="['date']" first-day-of-week="1"
                      />
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
            </span>
          </q-td>
          <q-td key="hours" :props="props">
            <span  v-if="!weekUnlocked">
              {{props.row.hours}}
            </span>
            <span v-else>
              <q-input :value="props.row.hours" type="number" step="0.5" dense
                @change="val => onUpdateHours({id: props.row.id, hours: val.target.value})"
              />
            </span>
          </q-td>
          <q-td key="project" :props="props">
            <span  v-if="!weekUnlocked">
              {{ props.row.project }}
            </span>
            <span v-else>
              <q-select :value="props.row.project" :options="filteredProjects" option-name="name" option-label="name"
                @filter="filterProject" dense options-dense use-input hide-selected fill-input input-debounce="0"
                @input="val => onUpdateProject({id: props.row.id, project: val.name})"
              >
                <template v-slot:no-option>
                  <q-item>
                    <q-item-section class="text-grey">
                      No results
                    </q-item-section>
                  </q-item>
                </template>
              </q-select>
            </span>
          </q-td>
          <q-td key="description" :props="props">
            <span  v-if="!weekUnlocked">
              {{ props.row.description }}
            </span>
            <span v-else>
              <q-input :value="props.row.description" dense
                @change="val => onUpdateDescription({id: props.row.id, description: val.target.value})"
                />
            </span>
          </q-td>
          <q-td key="rate" :props="props">
            <span  v-if="!weekUnlocked">
              {{ props.row.rate }}
            </span>
            <span v-else>
              <q-select :value="props.row.rate" :options="rates" option-label="name" option-name="id"
                dense options-dense
                @input="val => onUpdateRate({id: props.row.id, rate: val.name})"
              />
            </span>
          </q-td>
          <q-td key="actions" :props="props">
            <span  v-if="!weekUnlocked">
            </span>
            <span v-else>
              <q-icon name="insert_drive_file" small color="light-blue-4" size="1.5em" @click="duplicateItem(props.row, 'same')">
                <q-tooltip>Duplicate on the same day</q-tooltip>
              </q-icon>
              <q-icon name="file_copy" small color="light-blue-9" size="1.5em" @click="duplicateItem(props.row, 'next')">
                <q-tooltip>Duplicate to the next day</q-tooltip>
              </q-icon>
              <q-icon name="delete" small color="red" size="1.5em" @click="deleteItem(props.row)">
                <q-tooltip>Delete</q-tooltip>
              </q-icon>
            </span>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <!-- Dialog to confirm delete and edit date in not current -->
    <confirm ref="confirm" />
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import { format, isWithinInterval, getISODay, parseISO, addDays } from 'date-fns'
import { workHoursMixin } from '../mixins/workHoursMixin'

export default {

  components: {
    /* webpackChunkName: "core" */
    'confirm': () => import('components/Confirm'),
    /* webpackChunkName: "core" */
    'select-consultant': () => import('components/SelectConsultant'),
    /* webpackChunkName: "core" */
    'change-week': () => import('components/ChangeWeek')
  },

  filters: {
    formatDate: function (date) {
      if (!date) return ''
      return format(parseISO(date), 'iii, M/d')
    }
  },

  mixins: [ workHoursMixin ],

  data () {
    return {
      filter: '',
      myPagination: { 'rowsPerPage': 30, 'sortBy': 'date', 'descending': false },
      headers: [
        { name: 'date', label: 'Date', align: 'left', sortable: true, field: 'date', style: 'width: 20%' },
        { name: 'hours', label: 'Hours', align: 'left', sortable: true, field: 'hours', style: 'width: 5%' },
        { name: 'project', label: 'Project', align: 'left', sortable: true, field: 'project', style: 'width: 15%' },
        { name: 'description', label: 'Description', align: 'left', sortable: false, field: 'description', style: 'width: 50%' },
        { name: 'rate', label: 'Rate', align: 'left', sortable: false, field: 'rate', width: '15%' },
        { name: 'actions', label: 'Actions', align: 'center', sortable: false, value: '', style: 'width: 5%' }
      ],
      reported: [],
      filteredProjects: this.assignedProjects
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
    weeklyHolidays () {
      return this.getHolidays(this.dateFrom, this.dateTo) * this.dailyWorkingHours
    },
    selectedReportedHours () {
      const from = this.dateFrom
      const to = this.dateTo
      return this.reportedHours.filter(function (report) {
        let d = new Date(report.date)
        return (d >= from && d <= to)
      })
    },
    reportedThisWeek () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        rep = rep + this.selectedReportedHours[i].hours
      }
      return rep
    },
    ...mapState({
      loading: state => state.reportedHours.loading,
      isCurrentWeek: state => state.context.isCurrentWeek,
      dateFrom: state => state.settings.dateFrom,
      dateTo: state => state.settings.dateTo,
      selectedMonth: state => state.settings.selectedMonth,
      reportedHours: state => state.reportedHours.consultantMonthly,
      assignedProjects: state => state.projects.all,
      rates: state => state.rates.all,
      selectedConsultant: state => state.consultants.selected,
      selectedAllocation: state => state.consultants.selectedAllocation,
      dailyWorkingHours: state => state.settings.dailyWorkingHours,
      dailyWorkingHoursMax: state => state.settings.dailyWorkingHoursMax,
      dailyWorkingHoursMin: state => state.settings.dailyWorkingHoursMin,
      holidays: state => state.holidays.all
    }),
    reportedOnMonday () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 1) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    },
    reportedOnTuesday () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 2) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    },
    reportedOnWednesday () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 3) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    },
    reportedOnThursday () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 4) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    },
    reportedOnFriday () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 5) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    },
    reportedOnWeekend () {
      let rep = 0.0
      for (let i = 0; i < this.selectedReportedHours.length; i++) {
        if (getISODay(parseISO(this.selectedReportedHours[i].date)) === 6 || getISODay(parseISO(this.selectedReportedHours[i].date)) === 7) {
          rep = rep + this.selectedReportedHours[i].hours
        }
      }
      return rep
    }
  },

  watch: {
    selectedConsultant (newValue, oldValue) {
      this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: this.selectedConsultant })
    }
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Reported hours')
    this.$store.commit('context/SET_PAGE_ICON', 'work_outline')
    if (this.reportedHours.length === 0) {
      this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: this.selectedConsultant })
    }
  },

  methods: {
    // FIXME return green of this day is holiday
    textColor (item) {
      let colorClass = ''
      if (item < this.dailyWorkingHoursMin * this.selectedAllocation) { colorClass = 'text-red-6' }
      if (item >= this.dailyWorkingHoursMax) { colorClass = 'text-orange-6' }
      if (item >= 24) { colorClass = 'text-red-6' }
      return colorClass
    },
    textColorWeek (item) {
      let colorClass = ''
      if (item < (this.dailyWorkingHours * this.selectedAllocation * 5 - this.weeklyHolidays)) { colorClass = 'text-red-6' }
      return colorClass
    },
    filterProject (val, update, abort) {
      update(() => {
        const needle = val.toLowerCase()
        this.filteredProjects = this.assignedProjects.filter(v => v.name.toLowerCase().indexOf(needle) > -1)
      })
    },
    onUpdateProject (newValue) {
      let payload = {
        id: newValue.id,
        type: 'project',
        value: newValue.project
      }
      this.$store.dispatch('reportedHours/updateAttributeValue', payload)
      // set rate to default rate
      const newRate = this.assignedProjects.find(function (element) {
        if (element.name === newValue.project) { return element.rate }
      })
      let payloadRate = {
        id: newValue.id,
        type: 'rate',
        value: newRate.rate
      }
      this.$store.dispatch('reportedHours/updateAttributeValue', payloadRate)
    },
    async onUpdateDate (id, date) {
      this.$refs.qDateProxy.hide()
      let payload = {
        id: id,
        type: 'date',
        value: date
      }
      if (isWithinInterval(parseISO(date), { start: this.dateFrom, end: this.dateTo })) {
        this.$store.dispatch('reportedHours/updateAttributeValue', payload)
      } else {
        if (await this.$refs.confirm.open('Please confirm', 'You selected ' + format(parseISO(date), 'iiii, MMM do') + '. The record will be moved to another week. Continue?', { color: 'bg-warning' })) {
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
          this.$store.dispatch('settings/jumpToWeek', parseISO(date))
        }
      }
    },
    onUpdateHours (newValue) {
      const hrs = parseFloat(newValue.hours)
      let payload = {
        id: newValue.id,
        type: 'hours',
        value: parseFloat(newValue.hours)
      }
      if (isNaN(hrs) || hrs < 0 || hrs > 24) {
        payload.value = 0
      }
      this.$store.dispatch('reportedHours/updateAttributeValue', payload)
    },
    onUpdateDescription (newValue) {
      let payload = {
        id: newValue.id,
        type: 'description',
        value: newValue.description
      }
      this.$store.dispatch('reportedHours/updateAttributeValue', payload)
    },
    onUpdateRate (newValue) {
      let payload = {
        id: newValue.id,
        type: 'rate',
        value: newValue.rate
      }
      this.$store.dispatch('reportedHours/updateAttributeValue', payload)
    },
    remainingHoursDaily (date, hours) {
      let totalDailyHours = this.reportedHours.filter(x => x.date === date).reduce(
        function (total, current) {
          return total + current.hours
        }, 0)
      if (typeof (totalDailyHours) === 'object') {
        totalDailyHours = totalDailyHours.hours
      }
      const totalDailyHoursNew = totalDailyHours + hours
      if (totalDailyHoursNew <= 24) {
        return hours
      }
      if (totalDailyHoursNew < 48) {
        this.$q.notify({
          message: 'Over 24 hours reported on ' + format(parseISO(date), 'EEEE'),
          icon: 'warning'
        })
        return hours
      }
      if (totalDailyHours >= 48) {
        this.$q.notify({
          message: totalDailyHours + ' hours reported on ' + format(parseISO(date), 'EEEE') + ' and you want to add additional ' + hours + ' hours. Record was not created.',
          icon: 'report_problem'
        })
        return -1
      }
      if (totalDailyHoursNew > 48) {
        this.$q.notify({
          message: 'Only ' + (48 - totalDailyHours).toString() + ' hours added on ' + format(parseISO(date), 'EEEE') + '. You wanted to add ' + hours + ' to already reported ' + totalDailyHours + ' hours.',
          icon: 'warning'
        })
        return 48 - totalDailyHours
      }
      return hours
    },
    addItem (item) {
      let d = new Date()
      if (!isWithinInterval(d, { start: this.dateFrom, end: this.dateTo })) {
        d = this.dateFrom
      }
      const newRecord = {
        id: null,
        consultant: this.selectedConsultant,
        date: format(d, 'yyyy-MM-dd'),
        hours: 8,
        rate: '',
        description: '',
        project: ''
      }
      const newHrs = this.remainingHoursDaily(newRecord.date, newRecord.hours)
      if (newHrs > 0 && newHrs <= newRecord.hours) {
        // newRecord.date = format(d, "yyyy-MM-dd'T'HH:mm:ssXXX")
        newRecord.date = format(d, "yyyy-MM-dd'T'00:00:00XXX")
        newRecord.hours = newHrs
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      }
    },
    duplicateItem (item, day) {
      let nextDay = ''
      if (day === 'same') {
        nextDay = format(parseISO(item.date), 'yyyy-MM-dd')
      } else {
        nextDay = format(addDays(parseISO(item.date), 1), 'yyyy-MM-dd')
      }
      const newHrs = this.remainingHoursDaily(nextDay, item.hours)
      if (newHrs > 0 && newHrs <= item.hours) {
        const newRecord = {
          id: null,
          date: nextDay + 'T00:00:00Z',
          hours: newHrs,
          project: item.project,
          description: item.description,
          rate: item.rate,
          consultant: item.consultant
        }
        console.log('creating', newRecord)
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      } else {
        console.log('skipping record', newHrs, item.hours)
      }
    },
    async deleteItem (item) {
      if (await this.$refs.confirm.open('Please confirm', 'Are you sure you want to delete the record?', { color: 'bg-warning' })) {
        this.$store.dispatch('reportedHours/removeRecord', parseInt(item.id, 10))
        this.$q.notify({
          message: this.$options.filters.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted',
          color: 'teal'
        })
      } else {
        console.log('canceled record delete') /* eslint-disable-line no-console */
      }
    }
  }
}
</script>

<style lang="stylus">
/* add space between days on reported hours weekly row */
.my-hours {
  margin-left: 1em !important;
}
</style>
