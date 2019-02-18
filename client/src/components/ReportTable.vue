<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <v-toolbar>
      <v-btn :disabled="isCurrentWeek" class="mb-2" @click="currentWeek">
        today
      </v-btn>
      <v-btn fab small @click="previousWeek">
        <v-icon>
          skip_previous
        </v-icon>
      </v-btn>
      <v-btn fab small @click="nextWeek">
        <v-icon>
          skip_next
        </v-icon>
      </v-btn>
      <v-flex xs2>
        <v-label class="text-xs-center">
          {{ formatWeek(dateFrom) }} - {{ formatWeek(dateTo) }}
        </v-label>
      </v-flex>
      <v-autocomplete v-model="selectedConsultant" prepend-icon="person_outline" :dense="true" :items="consultants.all" item-text="name" item-value="name" class="body-1" />
      <v-spacer />
      <v-toolbar-title>
        <v-text-field v-model="search" clearable append-icon="search" label="Search" single-line />
      </v-toolbar-title>
      <v-spacer />
      <v-btn color="primary" :disabled="btnNewRecordDisabled" class="mb-2" @click="addItem">
        new record
      </v-btn>
    </v-toolbar>
    <v-toolbar dense>
      <v-toolbar-title>
        <v-label>Weekly: {{ reportedThisWeek }} hrs</v-label>
      </v-toolbar-title>
      <v-spacer />
      <v-toolbar-title>
        <v-label>
          Mon: <span :class="textColor(reportedThisDay(1))">
            {{ reportedThisDay(1) }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Tue: <span :class="textColor(reportedThisDay(2))">
            {{ reportedThisDay(2) }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Wed: <span :class="textColor(reportedThisDay(3))">
            {{ reportedThisDay(3) }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Thu: <span :class="textColor(reportedThisDay(4))">
            {{ reportedThisDay(4) }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>
          Fri: <span :class="textColor(reportedThisDay(5))">
            {{ reportedThisDay(5) }} hrs
          </span>
        </v-label>
      </v-toolbar-title>
      <v-toolbar-title>
        <v-label>Weekend: {{ reportedThisDay(6) + reportedThisDay(7) }} hrs</v-label>
      </v-toolbar-title>
    </v-toolbar>
    <v-data-table :headers="headers" :items="selectedReportedHours" :search="search" :loading="loading" class="elevation-1" :rows-per-page-items="rowsPerPage">
      <template slot="items" slot-scope="props">
        <td>
          <v-menu :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px" @keyup.esc="model = false">
            <v-text-field slot="activator" :value="props.item.date | formatDate" prepend-icon="today" readonly class="body-1" />
            <v-date-picker first-day-of-week="1" :value="props.item.date" @input="onUpdateDate({id: props.item.id, date: $event})" />
          </v-menu>
        </td>
        <td class="text-xs-left">
          <v-text-field :value="props.item.hours" :rules="[ruleFloat]"
                        type="number" min="0" max="24" step="0.5" maxlength="2"
                        class="body-1" single-line @change="onUpdateHours({id: props.item.id, hours: $event})"
          />
        </td>
        <td class="text-xs-left">
          <v-autocomplete :value="props.item.project" item-text="name" item-value="name"
                          :items="assignedProjects" :dense="true" :hide-selected="false" class="body-1"
                          @change="onUpdateProject({id: props.item.id, project: $event})"
          />
        </td>
        <td class="text-xs-left">
          <v-text-field slot="input" :value="props.item.description" single-line class="body-1"
                        :rules="[ruleMaxChars]" @change="onUpdateDescription({id: props.item.id, description: $event})"
          />
        </td>
        <td class="text-xs-left">
          <v-autocomplete slot="input" :value="props.item.rate" item-text="name" item-value="name"
                          :items="rates" :dense="true" :hide-selected="false" class="body-1"
                          @change="onUpdateRate({id: props.item.id, rate: $event})"
          />
        </td>
        <td class="text-xs-center px-0">
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
    <!-- Dialog to confirm delete and unlock -->
    <confirm ref="confirm" />
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import moment from 'moment-timezone'
  import confirm from './Confirm'

  export default {

    components: {
      'confirm': confirm
    },

    filters: {
      formatDate: function (date) {
        if (!date) return ''
        return moment(date).format('ddd, MMM Do')
      }
    },

    data () {
      return {
        search: '',
        ruleMaxChars: v => v.length <= 80 || v.length + ' / 80',
        ruleFloat: v => !isNaN(parseFloat(v)) || 'Input should be a muber with one decimal',
        repDate: '',
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
      selectedConsultant: {
        set (newValue) {
          this.$store.dispatch('consultants/setSelected', newValue)
        },
        get () {
          return this.$store.state.consultants.selected
        }
      },
      isCurrentWeek () {
        let today = moment.tz({}, this.timeZone)
        if (today.isBetween(this.dateFrom, this.dateTo, null, '[]')) {
          return true
        }
        return false
      },
      btnNewRecordDisabled () {
        if (this.previousWeeksUnLock) {
          return false
        }
        if (this.isCurrentWeek) { return false }
        return true
      },
      formTitle () {
        return this.editedIndex === -1 ? 'New Record' : 'Edit Record'
      },
      selectedReportedHours () {
        let from = this.dateFrom
        let to = this.dateTo
        let consultant = this.selectedConsultant
        return this.reportedHours.filter(function (report) {
          let d = new Date(report.date)
          return (d >= from && d <= to && report.consultant === consultant)
        })
      },
      reportedThisWeek () {
        let rep = 0.0
        for (let i = 0; i < this.selectedReportedHours.length; i++) {
          rep = rep + this.selectedReportedHours[i].hours
        }
        return rep
      },
      previousWeeksUnLock: {
        set () {
          this.$store.dispatch('context/TogglePreviousWeeksUnLock')
        },
        get () {
          return this.$store.state.context.previousWeeksUnLock
        }
      },
      ...mapState({
        loading: state => state.reportedHours.loading,
        dateFrom: state => state.settings.dateFrom,
        dateTo: state => state.settings.dateTo,
        reportedHours: state => state.reportedHours.all,
        assignedProjects: state => state.projects.all,
        rates: state => state.rates.all,
        consultants: state => state.consultants,
        timeZone: state => state.settings.timeZone,
        dailyWorkingHoursMax: state => state.settings.dailyWorkingHoursMax,
        dailyWorkingHoursMin: state => state.settings.dailyWorkingHoursMin
      })
    },

    // watch: {
    //   dialog (val) {
    //     val || this.close()
    //   }
    // },

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
      reportedThisDay (weekDay) {
        let rep = 0.0
        for (let i = 0; i < this.selectedReportedHours.length; i++) {
          if (moment.tz(this.selectedReportedHours[i].date, this.timeZone).weekday() === weekDay) {
            rep = rep + this.selectedReportedHours[i].hours
          }
        }
        return rep
      },
      currentWeek () {
        this.$store.dispatch('settings/jumpToWeek', moment.tz({}, this.timeZone))
      },
      editPreviousWeeks (itemID) {
        const editedDay = moment(this.selectedReportedHours.filter(item => item.id === itemID)[0].date)
        if (this.previousWeeksUnLock) {
          return true
        } else {
          if (editedDay.isBetween(moment.tz({}, this.timeZone).startOf('isoWeek'), moment.tz({}, this.timeZone).endOf('isoWeek'), null, '[]')) {
            return true
          }
        }
        this.$refs.confirm.open('Unlock previous weeks?', 'Data might be already reported to the clients. Do you want to unlock editing of previous weeks and repeat the edit?', { color: 'orange lighten-2' })
          .then(confirm => {
            if (confirm) { this.previousWeeksUnLock = true }
          })
          .catch(e => {
            console.log('No') /* eslint-disable-line no-console */
          })
        return false
      },
      onUpdateProject (newValue) {
        if (this.editPreviousWeeks(newValue.id)) {
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
        }
      },
      onUpdateDate (newValue) {
        if (this.editPreviousWeeks(newValue.id)) {
          let payload = {
            id: newValue.id,
            type: 'date',
            value: newValue.date
          }
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
        }
      },
      onUpdateHours (newValue) {
        if (this.editPreviousWeeks(newValue.id)) {
          const hrs = parseFloat(newValue.hours)
          if (!isNaN(hrs)) {
            let payload = {
              id: newValue.id,
              type: 'hours',
              value: parseFloat(newValue.hours)
            }
            this.$store.dispatch('reportedHours/updateAttributeValue', payload)
          }
        }
      },
      onUpdateDescription (newValue) {
        if (this.editPreviousWeeks(newValue.id)) {
          let payload = {
            id: newValue.id,
            type: 'description',
            value: newValue.description
          }
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
        }
      },
      onUpdateRate (newValue) {
        if (this.editPreviousWeeks(newValue.id)) {
          let payload = {
            id: newValue.id,
            type: 'rate',
            value: newValue.rate
          }
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
        }
      },
      addItem (item) {
        let d = moment.tz({}, this.timeZone)
        if (!d.isBetween(this.dateFrom, this.dateTo, '[]')) {
          d = this.dateFrom
        }
        const newRecord = {
          id: null,
          consultant: this.selectedConsultant,
          date: d.format('YYYY-MM-DDTHH:mm:ssZ'),
          hours: 8,
          rate: '',
          description: '',
          project: ''
        }
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      },
      duplicateItem (item) {
        if (this.editPreviousWeeks(item.id)) {
          let newRecord = Object.assign({}, item)
          newRecord.id = null
          newRecord.date = moment.tz(item.date, this.timeZone).format('YYYY-MM-DDTHH:mm:ssZ')
          this.$store.dispatch('reportedHours/addRecord', newRecord)
        }
      },
      async deleteItem (item) {
        if (this.editPreviousWeeks(item.id)) {
          if (await this.$refs.confirm.open('Please confirm', 'Are you sure you want to delete the record?', { color: 'orange lighten-2' })) {
            this.$store.dispatch('reportedHours/removeRecord', parseInt(item.id, 10))
            this.$store.dispatch('context/setNotification', { text: this.$options.filters.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted', type: 'success' })
          } else {
            console.log('canceled record delete') /* eslint-disable-line no-console */
          }
        }
      },
      previousWeek () {
        this.$store.dispatch('settings/changeWeek', 'previous')
      },
      nextWeek () {
        this.$store.dispatch('settings/changeWeek', 'next')
      },
      // FIXME move to filter
      formatWeek (date) {
        let a = moment.tz(date, 'YYYY-MM-DD', this.timeZone).format('MMM Do')
        return a
      }
    }
  }
</script>

<style>
/* workaround to decrease spacing of date field using elements v-menu amnd v-text-field */
table .v-messages.theme--light {
  display: none;
}
table .v-input--is-readonly.theme--light {
  padding-top: 0px !important;
  margin-top: 0px !important;
}
/* remove spacing above table text fields */
/* html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-text-field--single-line.v-input--is-label-active.v-input--is-dirty.theme--light {
html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-text-field--single-line.theme--light */
html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-text-field--single-line {
  padding-top: 0px !important;
}
/* remove spacing above table select field */
/* html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-select.v-input--is-label-active.v-input--is-dirty.theme--light {
html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-select.v-autocomplete.theme--light */
html.gr__localhost body div#app.application.theme--light div.application--wrap main.v-content div.v-content__wrap div.container.fluid div div.elevation-1 div.v-table__overflow table.v-datatable.v-table.theme--light tbody tr td.text-xs-left div.v-input.body-1.v-text-field.v-select {
  padding-top: 0px !important;
}
/* .menuable__content__active {
    transform: scale(0.875);
    transform-origin: left;
} */
</style>
