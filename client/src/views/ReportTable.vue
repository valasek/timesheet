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
        <v-flex xs2>
          <v-toolbar-title>
            <v-text-field v-model="search" clearable append-icon="search" label="Search" single-line />
          </v-toolbar-title>
        </v-flex>
        <v-spacer />
        <v-flex xs2>
          <v-btn color="primary" :disabled="!weekUnlocked" class="mb-2" @click="addItem">
            new record
          </v-btn>
        </v-flex>
      </v-layout>
    </v-toolbar>
    <v-toolbar dense>
      <v-toolbar-title>
        <v-label>Weekly: {{ reportedThisWeek }} hrs</v-label>
      </v-toolbar-title>
      <v-spacer />
      <v-toolbar-title>
        <v-label>
          Mon: <span :class="textColor(reportedOnMonday)">
            {{ reportedOnMonday }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Tue: <span :class="textColor(reportedOnTuesday)">
            {{ reportedOnTuesday }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Wed: <span :class="textColor(reportedOnWednesday)">
            {{ reportedOnWednesday }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Thu: <span :class="textColor(reportedOnThursday)">
            {{ reportedOnThursday }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Fri: <span :class="textColor(reportedOnFriday)">
            {{ reportedOnFriday }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>Weekend: {{ reportedOnWeekend }} hrs</v-label>
      </v-toolbar-title>
    </v-toolbar>
    <v-data-table :headers="headers" :items="selectedReportedHours" :search="search" :loading="loading" :disable-initial-sort="false" class="elevation-1 fixed-header" :rows-per-page-items="rowsPerPage">
      <template slot="items" slot-scope="props">
        <td>
          <template v-if="!weekUnlocked">
            <label class="body-1">{{ props.item.date | formatDate }}</label>
          </template>
          <template v-else>
            <v-menu :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px" @keyup.esc="model = false">
              <v-text-field slot="activator" :value="props.item.date | formatDate" readonly class="body-1" />
              <v-date-picker first-day-of-week="1" :value="props.item.date" @input="onUpdateDate({id: props.item.id, date: $event})" />
            </v-menu>
          </template>
        </td>
        <td class="text-xs-left">
          <template v-if="!weekUnlocked">
            <label class="body-1">{{ props.item.hours }}</label>
          </template>
          <template v-else>
            <v-text-field :value="props.item.hours" :rules="hoursRules"
                          type="number" min="0" max="24" step="0.5" maxlength="2"
                          class="body-1" single-line @change="onUpdateHours({id: props.item.id, hours: $event})"
            />
          </template>
        </td>
        <td class="text-xs-left">
          <template v-if="!weekUnlocked">
            <label class="body-1">{{ props.item.project }}</label>
          </template>
          <template v-else>
            <v-autocomplete :value="props.item.project" item-text="name" item-value="name"
                            :items="assignedProjects" :dense="true" :hide-selected="false" class="body-1"
                            @change="onUpdateProject({id: props.item.id, project: $event})"
            />
          </template>
        </td>
        <td class="text-xs-left">
          <template v-if="!weekUnlocked">
            <label class="body-1">{{ props.item.description }}</label>
          </template>
          <template v-else>
            <v-text-field slot="input" :value="props.item.description" single-line class="body-1"
                          maxlength="200" @change="onUpdateDescription({id: props.item.id, description: $event})"
            />
          </template>
        </td>
        <td class="text-xs-left">
          <template v-if="!weekUnlocked">
            <label class="body-1">{{ props.item.rate }}</label>
          </template>
          <template v-else>
            <v-autocomplete slot="input" :value="props.item.rate" item-text="name" item-value="name"
                            :items="rates" :dense="true" :hide-selected="false" class="body-1"
                            @change="onUpdateRate({id: props.item.id, rate: $event})"
            />
          </template>
        </td>
        <td v-if="weekUnlocked" class="text-xs-center px-0">
          <v-tooltip bottom>
            <v-icon slot="activator" small class="mr-2" @click="duplicateItem(props.item)">
              file_copy
            </v-icon>
            <span>Duplicate record</span>
          </v-tooltip>
          <v-tooltip bottom>
            <v-icon slot="activator" small @click="deleteItem(props.item)">
              delete
            </v-icon>
            <span>Remove record</span>
          </v-tooltip>
        </td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" color="info" icon="info">
          No hours reported this week ...
        </v-alert>
      </template>
    </v-data-table>
    <!-- Dialog to confirm delete and edit date in not current -->
    <confirm ref="confirm" />
    <inform ref="inform" />
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import { format, isWithinInterval, getISODay, parseISO } from 'date-fns'
  import confirm from '../components/Confirm'
  import inform from '../components/Inform'
  import selectConsultant from '../components/SelectConsultant'
  import changeWeek from '../components/ChangeWeek'

  export default {

    components: {
      'confirm': confirm,
      'inform': inform,
      'select-consultant': selectConsultant,
      'change-week': changeWeek
    },

    filters: {
      formatDate: function (date) {
        if (!date) return ''
        return format(parseISO(date), 'iii, MMM do')
      }
    },

    data () {
      return {
        search: '',
        hoursRules: [
          (v) => !!v || 'Working hours empty',
          (v) => (parseFloat(v) <= 24.0) || 'Working hours should be between 0 and 24',
          (v) => (parseFloat(v) >= 0) || 'Working hours should be between 0 and 24'
        ],
        rowsPerPage: [ 30, 50, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
        headers: [
          { text: 'Date', align: 'left', sortable: true, value: 'date', width: '12%', class: 'body-1' },
          { text: 'Hours', align: 'left', sortable: true, value: 'hours', width: '5%', class: 'body-1' },
          { text: 'Project', align: 'left', sortable: true, value: 'project', width: '15%', class: 'body-1' },
          { text: 'Description', align: 'left', value: 'description', width: '48%', class: 'body-1' },
          { text: 'Rate', align: 'left', sortable: true, value: 'rate', width: '15%', class: 'body-1' },
          { text: 'Actions', align: 'center', sortable: false, value: 'actions', width: '5%', class: 'body-1' }
        ],
        reported: []
      }
    },

    computed: {
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
        weekUnlocked: state => state.context.weekUnlocked,
        dateFrom: state => state.settings.dateFrom,
        dateTo: state => state.settings.dateTo,
        selectedMonth: state => state.settings.selectedMonth,
        reportedHours: state => state.reportedHours.consultantMonthly,
        assignedProjects: state => state.projects.all,
        rates: state => state.rates.all,
        selectedConsultant: state => state.consultants.selected,
        dailyWorkingHoursMax: state => state.settings.dailyWorkingHoursMax,
        dailyWorkingHoursMin: state => state.settings.dailyWorkingHoursMin
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
    },

    methods: {
      // FIXME return green of this day is holiday
      textColor (item) {
        let colorClass = ''
        if (item < this.dailyWorkingHoursMin) { colorClass = 'red--text lighten-2' }
        if (item >= this.dailyWorkingHoursMax) { colorClass = 'orange--text lighten-2' }
        return colorClass
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
      async onUpdateDate (newValue) {
        let payload = {
          id: newValue.id,
          type: 'date',
          value: newValue.date
        }
        if (isWithinInterval(parseISO(newValue.date), { start: this.dateFrom, end: this.dateTo })) {
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
        } else {
          if (await this.$refs.confirm.open('Please confirm', 'You selected ' + format(parseISO(newValue.date), 'iiii, MMM do') + '. The record will be moved to another week. Continue?', { color: 'orange lighten-2' })) {
            this.$store.dispatch('reportedHours/updateAttributeValue', payload)
            this.$store.dispatch('settings/jumpToWeek', parseISO(newValue.date))
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
      addItem (item) {
        let d = new Date()
        if (!isWithinInterval(d, { start: this.dateFrom, end: this.dateTo })) {
          d = this.dateFrom
        }
        const newRecord = {
          id: null,
          consultant: this.selectedConsultant,
          date: format(d, "yyyy-MM-dd'T'HH:mm:ssXXX"),
          hours: 8,
          rate: '',
          description: '',
          project: ''
        }
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      },
      duplicateItem (item) {
        let totalDailyHours = this.reportedHours.filter(x => x.date === item.date).reduce(
          function (total, current) {
            return total + current.hours
          }, 0)
        if (typeof (totalDailyHours) === 'object') {
          totalDailyHours = totalDailyHours.hours
        }
        totalDailyHours = totalDailyHours + item.hours
        if (totalDailyHours > 33) {
          this.$refs.inform.open('Too many hours per day', 'Attempt to report ' + totalDailyHours + ' hours on ' + format(parseISO(item.date), 'ddd, MMM do') + '. Record was not duplicated.', { color: 'orange lighten-2' })
        } else {
          let newRecord = Object.assign({}, item)
          newRecord.id = null
          newRecord.date = format(parseISO(item.date), 'yyyy-MM-dd') + 'T00:00:00Z'
          this.$store.dispatch('reportedHours/addRecord', newRecord)
        }
      },
      async deleteItem (item) {
        if (await this.$refs.confirm.open('Please confirm', 'Are you sure you want to delete the record?', { color: 'orange lighten-2' })) {
          this.$store.dispatch('reportedHours/removeRecord', parseInt(item.id, 10))
          this.$store.dispatch('context/setNotification', { text: this.$options.filters.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted', type: 'success' })
        } else {
          console.log('canceled record delete') /* eslint-disable-line no-console */
        }
      }
    }
  }
</script>

<style lang="stylus">
/* workaround to decrease spacing of date field using elements v-menu amnd v-text-field */
table .v-messages.theme--light {
  display: none;
}
table .v-input--is-readonly.theme--light {
  padding-top: 0px !important;
  margin-top: 0px !important;
}
/* remove spacing above table text fields */
html body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-text-field--single-line {
  padding-top: 0px !important;
}
/* remove spacing above table select field */
html body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-select {
  padding-top: 0px !important;
}
/* .menuable__content__active {
    transform: scale(0.875);
    transform-origin: left;
} */
@import '~vuetify-stylus-fixed-table-header';
.fixed-header {
  fixed-table-header(65vh)
}
</style>
