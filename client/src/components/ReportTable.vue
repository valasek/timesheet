<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>
        <v-text-field v-model="search" append-icon="search" label="Search" single-line />
      </v-toolbar-title>
      <v-spacer />
      <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" dark class="mb-2">
          +
        </v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">
              {{ formTitle }}
            </span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm6 md4>
                  <v-menu v-model="repDate" :close-on-content-click="true" :nudge-right="40" lazy transition="scale-transition" offset-y full-width min-width="290px">
                    <v-text-field slot="activator" v-model="editedItem.date" label="Date" readonly />
                    <v-date-picker v-model="editedItem.date" @input="repDate = false" />
                  </v-menu>
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="editedItem.hours" label="Hours" />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-select v-model="editedItem.project" item-text="name" item-value="name" :items="assignedProjects" label="Project" />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="editedItem.description" label="Description" />
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-select v-model="editedItem.rate" item-text="name" item-value="name" :items="rates" label="Rate" />
                  <!-- <v-text-field v-model="editedItem.rate" label="Rate" /> -->
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer />
            <v-btn color="blue darken-1" flat @click="close">
              Cancel
            </v-btn>
            <v-btn color="blue darken-1" flat @click="save">
              Save
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-toolbar>
    <v-data-table :headers="headers" :items="selectedReportedHours" :search="search" class="elevation-1" :rows-per-page-items="rowsPerPage">
      <template slot="items" slot-scope="props">
        <td>{{ formatDate(props.item.date) }}</td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.hours" lazy>
            {{ props.item.hours }}
            <v-text-field slot="input" :value="props.item.hours" label="Hours" single-line
                          type="number" min="0" max="20" step="0.5" maxlength="2"
                          @input="onUpdateHours({_id: props.item._id, hours: $event})"
            />
          </v-edit-dialog>
        <!-- {{ props.item.hours }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.project" lazy>
            {{ props.item.project }}
            <v-select slot="input" :value="props.item.project" item-text="name" item-value="name"
                      :items="assignedProjects" label="Project" :dense="true" :hide-selected="true"
                      @input="onUpdateProject({_id: props.item._id, project: $event})"
            />
          </v-edit-dialog>
          <!-- {{ props.item.project }} -->
        </td>
        <td class="text-xs-left">
          <v-edit-dialog :return-value="props.item.description" lazy>
            {{ props.item.description }}
            <v-text-field slot="input" :value="props.item.description" :rules="[ruleMax100chars]" label="Description" single-line
                          counter @input="onUpdateDescription({_id: props.item._id, description: $event})"
            />
          </v-edit-dialog>
          <!-- {{ props.item.description }} -->
        </td>
        <td class="text-xs-left">
          {{ props.item.rate }}
        </td>
        <td class="justify-center layout px-0">
          <v-icon small class="mr-2" @click="editItem(props.item)">
            add
          </v-icon>
          <v-icon small @click="deleteItem(props.item)">
            delete
          </v-icon>
        </td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" color="info" icon="info">
          Change month, select consultant or no reported hours this week :(
        </v-alert>
      </template>
    </v-data-table>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    data: () => ({
      search: '',
      ruleMax100chars: v => v.length <= 100 || 'Input too long!',
      repDate: '',
      dialog: false,
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
      reported: [],
      editedIndex: -1,
      editedItem: {
        date: '',
        hours: 0,
        project: '',
        description: '',
        rate: ''
      },
      defaultItem: {
        date: (new Date()).toString(),
        hours: 8,
        project: '',
        description: '',
        rate: ''
      }
    }),

    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'New Record' : 'Edit Record'
      },
      selectedReportedHours () {
        return this.reportedHours.filter(report => {
          let d = new Date(report.date)
          return (d >= this.dateFrom && d <= this.dateTo && report.consultant === this.selectedConsultants)
        })
      },
      ...mapState({
        reportedHours: state => state.reportedHours.all,
        dateFrom: state => state.context.dateFrom,
        dateTo: state => state.context.dateTo,
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
      // inline editing functions
      descriptionSave () {
        console.log('saved') /* eslint-disable-line no-console */
      },
      onUpdateProject (newProject) {
        console.log(newProject) /* eslint-disable-line no-console */
        this.$store.dispatch('reportedHours/updateProject', newProject)
      },
      onUpdateHours (newHours) {
        console.log(newHours) /* eslint-disable-line no-console */
        this.$store.dispatch('reportedHours/updateHours', newHours)
      },
      onUpdateDescription (newDescription) {
        console.log(newDescription) /* eslint-disable-line no-console */
        this.$store.dispatch('reportedHours/updateDescription', newDescription)
      },
      descriptionCancel () {
        console.log('cancel') /* eslint-disable-line no-console */
      },
      descriptionOpen () {
        console.log('open') /* eslint-disable-line no-console */
      },
      descriptionClose () {
        console.log('close') /* eslint-disable-line no-console */
      },
      initialize () {
        console.log('Get data clicked') /* eslint-disable-line no-console */
      },

      formatDate (date) {
        if (!date) return null

        const [year, month, day] = date.slice(0, 10).split('-')
        return `${month}/${day}/${year}`
      },

      editItem (item) {
        console.log('edit') /* eslint-disable-line no-console */
        this.editedIndex = this.selectedReportedHours.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },

      deleteItem (item) {
        confirm('Are you sure you want to delete the record?') && this.$store.dispatch('reportedHours/removeRecord', item._id)
        this.$store.dispatch('context/setNotification', this.formatDate(item.date) + ', ' + item.hours + ' hrs - record deleted')
      },

      close () {
        console.log('close') /* eslint-disable-line no-console */
        this.dialog = false
        setTimeout(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      },

      save () {
        console.log('save, editIndex: ' + this.editedIndex) /* eslint-disable-line no-console */
        if (this.editedIndex > -1) {
          Object.assign(this.selectedReportedHours[this.editedIndex], this.editedItem)
        } else {
          this.selectedReportedHours.push(this.editedItem)
        }
        this.close()
      }
    }
  }
</script>

<style>
</style>
