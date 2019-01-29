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
      <v-select v-model="selectedConsultant" prepend-icon="person" :dense="true" :items="consultants.all" item-text="name" item-value="name" class="body-1" />
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
          <!-- <v-menu :value="props.item.date" :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px"> -->
          <v-menu :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px" @keyup.esc="model = false">
            <v-text-field slot="activator" :value="formatDate(props.item.date)" prepend-icon="today" readonly class="body-1" />
            <v-date-picker first-day-of-week="1" :value="props.item.date" @input="onUpdateDate({id: props.item.id, date: $event})" />
          </v-menu>
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.hours" lazy>
            {{ props.item.hours }}
            <v-text-field slot="input" suffix="hours" :value="props.item.hours" label="Hours" single-line
                          :rules="[ruleFloat]"
                          type="number" min="0" max="20" step="0.5" maxlength="2"
                          @change="onUpdateHours({id: props.item.id, hours: $event})"
            />
          </v-edit-dialog>
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.project" lazy>
            {{ props.item.project }}
            <v-select slot="input" :value="props.item.project" item-text="name" item-value="name"
                      :items="assignedProjects" label="Project" :dense="true" :hide-selected="false"
                      @change="onUpdateProject({id: props.item.id, project: $event})"
            />
          </v-edit-dialog>
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.description" lazy>
            {{ props.item.description }}
            <v-text-field slot="input" :value="props.item.description" label="Description" single-line counter
                          :rules="[ruleMaxChars]" @change="onUpdateDescription({id: props.item.id, description: $event})"
            />
          </v-edit-dialog>
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.rate" lazy>
            {{ props.item.rate }}
            <v-select slot="input" :value="props.item.rate" item-text="name" item-value="name"
                      :items="rates" label="Rate" :dense="true" :hide-selected="false"
                      @change="onUpdateRate({id: props.item.id, rate: $event})"
            />
          </v-edit-dialog>
        </td>
        <td class="text-xs-center px-0">
          <v-tooltip bottom>
            <v-icon slot="activator" small class="mr-2" @click="duplicateItem(props.item)">
              file_copy
            </v-icon>
            <span>Duplicate this record</span>
          </v-tooltip>
          <v-tooltip bottom>
            <v-icon slot="activator" small @click="deleteItem(props.item)">
              delete
            </v-icon>
            <span>Remove this record</span>
          </v-tooltip>
        </td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" color="info" icon="info">
          No reported hours this week :(
        </v-alert>
      </template>
    </v-data-table>
    <!-- Dialog if use attempts to edit previous weeks without unlocking -->
    <!-- <v-dialog v-model="dialogEditingUnlockedWeek" width="30%">
      <v-card>
        <v-card-title class="headline">
          Editing the previous week
        </v-card-title>

        <v-card-text multiline>
          Data might be already reported to the clients.
          Do you want to unlock editing of previous weeks and repeat the edit?
        </v-card-text>

        <v-card-actions>
          <v-spacer />

          <v-btn color="blue darken-1" flat="flat" @click="unlockPreviousWeeks">
            Unlock
          </v-btn>

          <v-btn color="primary" dark class="mb-2" @click="dialogEditingUnlockedWeek = false">
            Cancel
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog> -->
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

    data () {
      return {
        search: '',
        // dialogEditingUnlockedWeek: false,
        ruleMaxChars: v => v.length <= 80 || v.length + ' / 80',
        ruleFloat: v => !isNaN(parseFloat(v)) || 'Input should be a muber with one decimal',
        repDate: '',
        rowsPerPage: [ 30, 50, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
        headers: [
          { text: 'Date', align: 'left', sortable: true, value: 'date', class: 'body-1' },
          { text: 'Hours', align: 'left', sortable: true, value: 'hours', class: 'body-1' },
          { text: 'Project', align: 'left', sortable: true, value: 'project', class: 'body-1' },
          { text: 'Description', align: 'left', value: 'description', class: 'body-1' },
          { text: 'Rate', align: 'left', sortable: true, value: 'rate', class: 'body-1' },
          { text: 'Actions', align: 'center', sortable: false, value: 'actions', class: 'body-1' }
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
          return this.consultants.selected
        }
      },
      isCurrentWeek () {
        let today = moment.tz({}, 'Europe/Prague')
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
        let consultant = this.selectedConsultants
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
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
        reportedHours: state => state.reportedHours.all,
        assignedProjects: state => state.projects.all,
        rates: state => state.rates.all,
        consultants: state => state.consultants,
        selectedConsultants: state => state.consultants.selected,
        dailyWorkingHoursMax: state => state.context.dailyWorkingHoursMax,
        dailyWorkingHoursMin: state => state.context.dailyWorkingHoursMin
      })
    },

    watch: {
      dialog (val) {
        val || this.close()
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
      reportedThisDay (weekDay) {
        let rep = 0.0
        for (let i = 0; i < this.selectedReportedHours.length; i++) {
          if (moment.tz(this.selectedReportedHours[i].date, 'Europe/Prague').weekday() === weekDay) {
            rep = rep + this.selectedReportedHours[i].hours
          }
        }
        return rep
      },
      currentWeek () {
        this.$store.dispatch('context/jumpToWeek', moment.tz({}, 'Europe/Prague'))
      },
      editPreviousWeeks (itemID) {
        const editedDay = moment(this.selectedReportedHours.filter(item => item.id === itemID)[0].date)
        if (this.previousWeeksUnLock) {
          return true
        } else {
          if (editedDay.isBetween(moment.tz({}, 'Europe/Prague').startOf('isoWeek'), moment.tz({}, 'Europe/Prague').endOf('isoWeek'), null, '[]')) {
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
          let payload = {
            id: newValue.id,
            type: 'hours',
            value: parseFloat(newValue.hours)
          }
          this.$store.dispatch('reportedHours/updateAttributeValue', payload)
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
      formatDate (date) {
        if (!date) return null
        return moment(date).format('ddd, MMM Do')
      },
      addItem (item) {
        let newRecord = {}
        newRecord.id = null
        newRecord.consultant = this.selectedConsultants
        newRecord.date = this.dateFrom.format('YYYY-MM-DDTHH:mm:ssZ')
        newRecord.hours = 8
        newRecord.rate = ''
        newRecord.description = ''
        newRecord.project = ''
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      },
      duplicateItem (item) {
        if (this.editPreviousWeeks(item.id)) {
          let newRecord = Object.assign({}, item)
          newRecord.id = null
          newRecord.date = moment.tz(item.date, 'Europe/Prague').format('YYYY-MM-DDTHH:mm:ssZ')
          this.$store.dispatch('reportedHours/addRecord', newRecord)
        }
      },
      async deleteItem (item) {
        if (this.editPreviousWeeks(item.id)) {
          if (await this.$refs.confirm.open('Please confirm', 'Are you sure you want to delete the record?', { color: 'orange lighten-2' })) {
            this.$store.dispatch('reportedHours/removeRecord', parseInt(item.id, 10))
            this.$store.dispatch('context/setNotification', { text: this.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted', type: 'success' })
          } else {
            console.log('No') /* eslint-disable-line no-console */
          }
        }
      },
      previousWeek () {
        this.$store.dispatch('context/changeWeek', 'previous')
      },
      nextWeek () {
        this.$store.dispatch('context/changeWeek', 'next')
      },
      formatWeek (date) {
        let a = moment.tz(date, 'YYYY-MM-DD', 'Europe/Prague').format('MMM Do')
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
.menuable__content__active {
    transform: scale(0.875);
    transform-origin: left;
}
</style>
