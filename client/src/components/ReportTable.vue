<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>
        <v-text-field v-model="search" append-icon="search" label="Search" single-line />
      </v-toolbar-title>
      <v-spacer />
      <v-label>{{ reportedThisWeek }} hours reported this week</v-label>
      <v-spacer />
      <v-btn color="primary" dark class="mb-2" @click="addItem">
        new record
      </v-btn>
    </v-toolbar>
    <v-data-table :headers="headers" :items="selectedReportedHours" :search="search" :loading="loading" class="elevation-1" :rows-per-page-items="rowsPerPage">
      <template slot="items" slot-scope="props">
        <td>
          <!-- <v-menu :value="props.item.date" :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px"> -->
          <v-menu :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px" @keyup.esc="model = false">
            <v-text-field slot="activator" :value="formatDate(props.item.date)" readonly />
            <v-date-picker first-day-of-week="1" :value="props.item.date" @input="onUpdateDate({id: props.item.id, date: $event})" />
          </v-menu>
          <!-- {{ formatDate(props.item.date) }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.hours" lazy>
            {{ props.item.hours }}
            <v-text-field slot="input" :value="props.item.hours" label="Hours" single-line
                          type="number" min="0" max="20" step="0.5" maxlength="2"
                          @change="onUpdateHours({id: props.item.id, hours: $event})"
            />
          </v-edit-dialog>
        <!-- {{ props.item.hours }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.project" lazy>
            {{ props.item.project }}
            <v-select slot="input" :value="props.item.project" item-text="name" item-value="name"
                      :items="assignedProjects" label="Project" :dense="true" :hide-selected="true"
                      @change="onUpdateProject({id: props.item.id, project: $event})"
            />
          </v-edit-dialog>
          <!-- {{ props.item.project }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.description" lazy>
            {{ props.item.description }}
            <v-text-field slot="input" :value="props.item.description" :rules="[ruleMax100chars]" label="Description" single-line
                          counter @change="onUpdateDescription({id: props.item.id, description: $event})"
            />
          </v-edit-dialog>
          <!-- {{ props.item.description }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.rate" lazy>
            {{ props.item.rate }}
            <v-select slot="input" :value="props.item.rate" item-text="name" item-value="name"
                      :items="rates" label="Rate" :dense="true" :hide-selected="true"
                      @change="onUpdateRate({id: props.item.id, rate: $event})"
            />
          </v-edit-dialog>
          <!-- {{ props.item.rate }} -->
        </td>
        <td class="justify-center layout px-0">
          <v-icon small class="mr-2" @click="duplicateItem(props.item)">
            add
          </v-icon>
          <v-icon small @click="deleteItem(props.item)">
            delete
          </v-icon>
        </td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" color="info" icon="info">
          No reported hours this week :(
        </v-alert>
      </template>
    </v-data-table>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import moment from 'moment-timezone'

  export default {
    data: () => ({
      search: '',
      ruleMax100chars: v => v.length <= 100 || 'Input too long!',
      repDate: '',
      rowsPerPage: [ 30, 50, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
      headers: [
        {
          text: 'Date',
          align: 'left',
          sortable: true,
          value: 'date'
        },
        { text: 'Hours', align: 'left', value: 'hours' },
        { text: 'Project', align: 'left', value: 'project' },
        { text: 'Description', value: 'description' },
        { text: 'Rate', value: 'rate' },
        { text: 'Actions', value: 'actions', sortable: false }
      ],
      reported: []
    }),

    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'New Record' : 'Edit Record'
      },
      selectedReportedHours () {
        return this.reportedHours.filter(report => {
          let d = new Date(report.date)
          // report.date = this.formatDate(report.date)
          return (d >= this.dateFrom && d <= this.dateTo && report.consultant === this.selectedConsultants)
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
        reportedHours: state => state.reportedHours.all,
        loading: state => state.reportedHours.loading,
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
        dateMonth: state => state.context.dateMonth,
        assignedProjects: state => state.projects.all,
        rates: state => state.rates.all,
        selectedConsultants: state => state.consultants.selected
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
      onUpdateProject (newValue) {
        let payload = {
          id: newValue.id,
          type: 'project',
          value: newValue.project
        }
        this.$store.dispatch('reportedHours/updateAttributeValue', payload)
      },
      onUpdateDate (newValue) {
        let payload = {
          id: newValue.id,
          type: 'date',
          value: newValue.date
        }
        this.$store.dispatch('reportedHours/updateAttributeValue', payload)
      },
      onUpdateHours (newValue) {
        let payload = {
          id: newValue.id,
          type: 'hours',
          value: newValue.hours
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
      formatDate (date) {
        if (!date) return null
        // console.log('before slice') /* eslint-disable-line no-console */
        const [, month, day] = date.slice(0, 10).split('-')
        // console.log('after slice') /* eslint-disable-line no-console */
        let weekdays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
        let dayName = weekdays[new Date(date).getDay()]
        return `${dayName} ${month}/${day}`
      },
      addItem (item) {
        let newRecord = {}
        newRecord.id = null
        newRecord.consultant = this.selectedConsultants
        console.log('from:', this.dateFrom) /* eslint-disable-line no-console */
        newRecord.date = this.dateFrom.format('YYYY-MM-DDTHH:mm:ssZ')
        console.log('date to ISO:', newRecord.date) /* eslint-disable-line no-console */
        newRecord.hours = 8
        newRecord.rate = 'Off-site'
        newRecord.project = ''
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      },
      duplicateItem (item) {
        let newRecord = Object.assign({}, item)
        newRecord.id = null
        newRecord.date = moment(item.date).format('YYYY-MM-DDTHH:mm:ssZ')
        console.log(newRecord) /* eslint-disable-line no-console */
        this.$store.dispatch('reportedHours/addRecord', newRecord)
      },
      deleteItem (item) {
        confirm('Are you sure you want to delete the record?') && this.$store.dispatch('reportedHours/removeRecord', parseInt(item.id, 10))
        this.$store.dispatch('context/setNotification', { text: this.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted', type: 'success' })
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
</style>
