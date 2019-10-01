<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-toolbar class="q-pa-md bg-red-3 text-secondary">
      <q-toolbar-title>Administration & Settings</q-toolbar-title>
    </q-toolbar>
    <q-expansion-item
      expand-separator
      label="Manage Consultants & Projects"
    >
      <q-card>
        <q-card-section>
          <div class="row justify-around">
            <div class="column">
              <div class="column q-gutter-md">
                <div class="row q-gutter-x-md justify-around">
                  <q-input v-model="newConsultant" label="Consultant name" dense style="width: 15em" />
                  <q-btn color="primary" :disable="newConsultant.length === 0" @click="createConsultant">
                    Add Consultant
                  </q-btn>
                </div>
                <p align="center">
                  {{ defaultAllocationLabel }}
                </p>
                <div class="row q-gutter-x-md">
                  <q-table :columns="columnsConsultants" row-key="name" :data="consultants"
                           no-data-label="No consultants" :pagination.sync="consultantsPagination" :rows-per-page-options="[30,50,0]"
                           binary-state-sort bordered
                  >
                    <template v-slot:body="props">
                      <q-tr :props="props">
                        <q-td key="name" :props="props">
                          {{ props.row.name }}
                        </q-td>
                        <q-td key="allocation" :props="props">
                          <q-input :value="format_allocation(props.row.allocation)" dense
                            type="string" class="body-1" style="width: 5em"
                            debounce="500"
                            @input="onUpdateAllocation({row: props.row, newValue: $event})"
                            @focus="$event.target.select()"
                          />
                        </q-td>
                        <q-td key="actions" :props="props" class="q-gutter-x-sm">
                          <q-icon name="remove_red_eye" small :color="setColor(props.row.disabled)" size="1.5em" @click="toggleConsultant(props.row)">
                            <q-tooltip>Hide / Unhide the consultant</q-tooltip>
                          </q-icon>
                          <q-icon name="delete" small color="red" size="1.5em" @click="deleteConsultant(props.row)">
                            <q-tooltip>Permanently delete the consultant and all associated reported records</q-tooltip>
                          </q-icon>
                        </q-td>
                      </q-tr>
                    </template>
                  </q-table>
                </div>
              </div>
            </div>
            <div class="column">
              <div class="column q-gutter-md">
                <div class="row q-gutter-x-md justify-around">
                  <q-input v-model="newProject" label="Project name" dense style="width: 15em" />
                  <q-btn color="primary" :disable="newProject.length === 0" @click="createProject">
                    Add Project
                  </q-btn>
                </div>
                <div class="row justify-center">
                  <div class="col self-center">
                    {{ defaultRateLabel }}
                  </div>
                  <div class="col self-center">
                    <q-select :value="defaultRate" option-value="name" option-label="name" style="width: 10em"
                              :options="rates" dense options-dense :hide-selected="false"
                              @input="onUpdateProjectRate($event)"
                            />
                  </div>
                </div>
                <div class="row q-gutter-x-md">
                  <q-table :columns="columnsProjects" row-key="name" :data="projects"
                          no-data-label="No projects" :pagination.sync="projectsPagination" :rows-per-page-options="[30,50,0]"
                          binary-state-sort bordered
                  >
                    <template v-slot:body="props">
                      <q-tr :props="props">
                        <q-td key="name" :props="props">
                          {{ props.row.name }}
                        </q-td>
                        <q-td key="rate" :props="props">
                          <q-select :value="props.row.rate" option-value="name" option-label="name" style="width: 15em"
                            :options="rates" dense options-dense :hide-selected="false"
                            @input="onUpdateProject({row: props.row, newValue: $event})"
                          />
                        </q-td>
                        <q-td key="actions" :props="props" class="q-gutter-x-sm">
                          <q-icon name="remove_red_eye" small :color="setColor(props.row.disabled)" size="1.5em" @click="toggleProject(props.row)">
                            <q-tooltip>Hide / Unhide the project</q-tooltip>
                          </q-icon>
                          <q-icon name="delete" small color="red" size="1.5em" @click="deleteProject(props.row)">
                            <q-tooltip>Permanently delete the project and all associated reported records</q-tooltip>
                          </q-icon>
                        </q-td>
                      </q-tr>
                    </template>
                  </q-table>
                </div>
              </div>
            </div>
          </div>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="Managed Data Statistics"
    >
      <q-card>
        <q-card-section>
          <div class="row q-gutter-x-md">
            <managed-data />
          </div>
        </q-card-section>
        <q-card-section>
          <p>Timesheet internally does soft deletes and all such records are displayed in the column "# of soft-deleted records".</p>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="Backup & Restore"
    >
      <q-card>
        <q-card-section>
          <div class="font-italic">
            You can download all  data in csv format as a zip file. Demo data can be modified, production moved to another instance or exported data used for billing.<br>
            You can upload modified zipped csv files. Use the same format as downloaded archive. Import will delete and replace all existing data.
          </div>
          <q-card-section>
            <div class="row justify-left q-gutter-x-md">
              <div class="column">
                <q-btn color="primary" @click="download">
                  Download&nbsp;&nbsp;<q-icon name="cloud_download" />
                </q-btn>
              </div>
              <div class="column">
                <q-uploader
                  label="Upload"
                  accept=".zip"
                  :url="uploaderUrl"
                  auto-upload
                  style="max-width: 250px"
                />
              </div>
            </div>
          </q-card-section>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="Warning Limits"
    >
      <q-card>
        <q-card-section>
          <div class="font-italic">
            Set minimum and maximum working hours  used on page Report my work to color code total daily work hours
          </div>
          <div class="row justify-around">
            <div class="column">
              <q-input :value="dailyWorkingHoursMin" label="Minimum working hours" :rules="hoursRules"
                       type="number" min="0" max="24" step="0.5" maxlength="2"
                       class="body-1" style="width: 15em" @input="onUpdateHours({hourType: 'dailyWorkingHoursMin', hourValue: $event})"
              />
            </div>
            <div class="column">
              <q-input :value="dailyWorkingHoursMax" label="Maximum working hours" :rules="hoursRules"
                       type="number" min="0" max="24" step="0.5" maxlength="2"
                       class="body-1" style="width: 15em" @input="onUpdateHours({hourType: 'dailyWorkingHoursMax', hourValue: $event})"
              />
            </div>
          </div>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="Vacation Settings"
    >
      <q-card>
        <q-card-section>
          <div class="font-italic">
            Used on Reported Overview page to show weekly and monthly expected working hours
          </div>
        </q-card-section>
        <q-card-section>
          <div class="row justify-around">
            <div class="column">
              Vacations
            </div>
            <div class="column">
              Personal Days
            </div>
            <div class="column">
              Sick Days
            </div>
          </div>
          <div class="row justify-around">
            <div class="column">
              <q-input :value="yearlyVacationDays" label="Vacation days per year" :rules="daysRules"
                       type="number" min="0" max="40" step="1" maxlength="2"
                       class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlyVacationDays', dayValue: $event})"
              />
            </div>
            <div class="column">
              <q-input :value="yearlyPersonalDays" label="Additional vacation days per year" :rules="daysRules"
                       type="number" min="0" max="40" step="1" maxlength="2"
                       class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlyPersonalDays', dayValue: $event})"
              />
            </div>
            <div class="column">
              <q-input :value="yearlySickDays" label="Additional vacation days per year" :rules="daysRules"
                       type="number" min="0" max="40" step="1" maxlength="2"
                       class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlySickDays', dayValue: $event})"
              />
            </div>
          </div>
          <div class="row justify-around">
            <div class="column">
              <q-select :value="vacation" option-value="name" option-label="name" style="width: 15em"
                        :options="rates" dense options-dense :hide-selected="false"
                        @input="onUpdateRate({rateType: 'vacation', rateValue: $event})"
              />
            </div>
            <div class="column">
              <q-select :value="vacationPersonal" option-value="name" option-label="name" style="width: 15em"
                        :options="rates" dense options-dense :hide-selected="false"
                        @input="onUpdateRate({rateType: 'vacationPersonal', rateValue: $event})"
              />
            </div>
            <div class="column">
              <q-select :value="vacationSick" option-value="name" option-label="name" style="width: 15em"
                        :options="rates" dense options-dense :hide-selected="false"
                        @input="onUpdateRate({rateType: 'vacationSick', rateValue: $event})"
              />
            </div>
          </div>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="Application Settings"
    >
      <q-card>
        <q-card-section>
          <p class="font-italic">
            Used for weekly and monthly expected working hours
          </p>
          <q-input :value="dailyWorkingHours" label="Daily working hours" :rules="hoursRules"
                   type="number" min="0" max="24" step="0.5" maxlength="2"
                   class="body-1" style="width: 10em" @input="onUpdateHours({hourType: 'dailyWorkingHours', hourValue: $event})"
          />
          <p class="font-italic">
            Used on Reported Overview page to show working and non-working time
          </p>
          <div class="row justify-start q-col-gutter-x-md">
            <div class="column">
              <q-select :value="isWorking" style="width: 15em" label="Working rate type"
                        :options="types" dense :hide-selected="false"
                        @input="onUpdateRateType({rateTypeType: 'isWorking', rateTypeValue: $event})"
              />
            </div>
            <div class="column">
              <q-select :value="isNonWorking" style="width: 15em" label="Non-working rate type"
                        :options="types" dense :hide-selected="false"
                        @input="onUpdateRateType({rateTypeType: 'isNonWorking', rateTypeValue: $event})"
              />
            </div>
          </div>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <q-expansion-item
      expand-separator
      label="View Application Logs"
    >
      <q-card>
        <q-card-section>
          <q-select :value="logLevelSelection" :options="logLevels" option-value="id" option-label="level" emit-value
                    dense options-dense label="Select log level" style="width: 15em" class="body-1"
                    @input="getLogFile"
          />
        </q-card-section>
        <q-card-section>
          <q-scroll-area style="height: 400px; max-width: 100%; backgroundColor: rgba(0,0,0,0.02); padding: 1em;">
            <div v-for="line in logLines" :key="line.id">
              {{ line }}
            </div>
          </q-scroll-area>
        </q-card-section>
      </q-card>
    </q-expansion-item>
    <q-separator />
    <!-- Dialog to confirm delete Project and Consultant -->
    <confirm ref="confirm" />
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import api from '../api/axiosSettings'

export default {

  components: {
    /* webpackChunkName: "core" */
    'confirm': () => import('components/Confirm'),
    'managed-data': () => import('components/ManagedData')
  },

  data () {
    return {
      // url="http://localhost:3000/api/upload/data"
      uploaderUrl: 'http://' + process.env.APP_DOMAIN + ':' + process.env.APP_PORT + '/api/upload/data',
      newConsultant: '',
      newProject: '',
      defaultAllocation: 1,
      defaultRate: 'Standard',
      columnsConsultants: [
        { name: 'name', label: 'Name', align: 'left', sortable: true, field: 'name', style: 'width: 20%' },
        { name: 'allocation', label: 'Allocation in %', align: 'left', sortable: true, field: 'allocation', style: 'width: 5%' },
        { name: 'actions', label: 'Action', align: 'left', field: 'action', style: 'width: 5%' }
      ],
      consultantsPagination: { 'rowsPerPage': 10, 'sortBy': 'name', 'descending': false },
      columnsProjects: [
        { name: 'name', label: 'Name', align: 'left', sortable: true, field: 'name', style: 'width: 20%' },
        { name: 'rate', label: 'Default Rate', align: 'left', sortable: true, field: 'rate', style: 'width: 5%' },
        { name: 'actions', label: 'Action', align: 'left', field: 'action', style: 'width: 5%' }
      ],
      projectsPagination: { 'rowsPerPage': 10, 'sortBy': 'name', 'descending': false },
      hoursRules: [
        (v) => !isNaN(parseFloat(v)) || 'Enter hours between 0 and 24',
        (v) => (parseFloat(v) <= 24) || 'Enter number between 0 and 24',
        (v) => (parseFloat(v) >= 0) || 'Enter number between 0 and 24'
      ],
      daysRules: [
        (v) => !isNaN(parseInt(v)) || 'Enter number of days between 0 and 40',
        (v) => parseInt(v) === parseFloat(v) || 'Enter full days',
        (v) => (parseInt(v) <= 40) || 'Enter number between 0 and 40',
        (v) => (parseInt(v) >= 0) || 'Enter number between 0 and 40'
      ],
      logLevelSelection: '',
      logLevels: [
        {
          id: 0,
          level: 'Information'
        },
        {
          id: 1,
          level: 'Warnings and Errors'
        }
      ],
      // logLevelsModel: null,
      logLines: ['select desired log level ...'],
      uploadFile: {}
    }
  },

  computed: {
    defaultAllocationLabel () {
      return 'Default allocation for new consultant will be ' + this.defaultAllocation * 100 + '%'
    },
    defaultRateLabel () {
      return 'Default rate for new project will be '
    },
    ...mapState({
      dailyWorkingHours: state => state.settings.dailyWorkingHours,
      dailyWorkingHoursMin: state => state.settings.dailyWorkingHoursMin,
      dailyWorkingHoursMax: state => state.settings.dailyWorkingHoursMax,
      yearlyVacationDays: state => state.settings.yearlyVacationDays,
      vacation: state => state.settings.vacation,
      yearlyPersonalDays: state => state.settings.yearlyPersonalDays,
      vacationPersonal: state => state.settings.vacationPersonal,
      yearlySickDays: state => state.settings.yearlySickDays,
      vacationSick: state => state.settings.vacationSick,
      isWorking: state => state.settings.isWorking,
      isNonWorking: state => state.settings.isNonWorking,
      reportedHours: state => state.reportedHours.consultantMonthly,
      rates: state => state.rates.all,
      types: state => state.rates.types,
      consultants: state => state.consultants.all,
      selectedConsultant: state => state.consultants.selected,
      projects: state => state.projects.all,
      selectedMonth: state => state.settings.selectedMonth
    })
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Administration')
    this.$store.commit('context/SET_PAGE_ICON', 'settings')
    this.$store.dispatch('projects/getProjects')
  },

  methods: {
    format_allocation (value) {
      return value * 100 + ' %'
    },
    upload (file) {
      const fr = new FileReader()
      fr.readAsDataURL(file)
      const admin = this
      fr.addEventListener('load', () => {
        admin.uploadedImage = fr.result
        admin.uploadFile = file
        const formData = new FormData()
        formData.append('uploadFile', admin.uploadFile)
        api.apiClient.post('/api/upload/data', formData)
          .then(response => {
            admin.$q.notify({
              message: 'Restore successfully completed',
              color: 'teal'
            })
            admin.$store.dispatch('consultants/getConsultants')
            admin.$store.dispatch('projects/getProjects')
            admin.$store.dispatch('rates/getRates')
            admin.$store.dispatch('holidays/getHolidays')
          })
          .catch(function (e) {
            admin.$q.notify({
              message: 'Restore failed, error: ' + e.toString(),
              color: 'negative',
              icon: 'report_problem'
            })
            console.log(e, e.response) /* eslint-disable-line no-console */
          })
      })
    },
    download () {
      const admin = this
      api.apiClient.get('/api/download/data', { responseType: 'blob' })
        .then(response => {
          const url = window.URL.createObjectURL(new Blob([response.data]))
          const link = document.createElement('a')
          link.href = url
          link.setAttribute('download', 'timesheet-backup.zip')
          document.body.appendChild(link)
          link.click()
        })
        .catch(function (e) {
          admin.$q.notify({
            message: 'Couldn\'t download data\n' + e.toString(),
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e) /* eslint-disable-line no-console */
        })
    },
    getLogFile (logLevel) {
      this.logLevelSelection = this.logLevels[logLevel]
      const admin = this
      api.apiClient.get('/api/download/logs/' + logLevel.toString())
        .then(response => {
          admin.logLines = response.data.split(new RegExp('\r?\n', 'g')) /* eslint-disable-line no-control-regex */
        })
        .catch(function (e) {
          admin.$q.notify({
            message: 'Couldn\'t download log file: ' + e.response.data,
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e, e.response) /* eslint-disable-line no-console */
        })
    },
    setColor (disabled) {
      if (disabled) {
        return 'orange'
      } else {
        return 'green'
      }
    },
    createProject () {
      this.$store.dispatch('projects/createProject', { id: null, name: this.newProject, rate: this.defaultRate, disabled: false })
    },
    toggleProject (project) {
      this.$store.dispatch('projects/toggleProject', parseInt(project.id, 10))
    },
    async deleteProject (project) {
      if (await this.$refs.confirm.open('Wait a sec, this cannot be undone!', 'Are you sure you want to delete the project <b>' + project.name + '</b><br/>and all relevant reported records?', 'cancel', { color: 'bg-negative' })) {
        this.$store.dispatch('projects/removeProject', { id: project.id, name: project.name })
        if (this.reportedHours.length === 0 && this.selectedConsultant !== '') {
          this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: this.selectedConsultant })
        }
      } else {
        // console.log('canceled project delete')
      }
    },
    createConsultant () {
      const newConsultant = {
        id: null,
        name: this.newConsultant,
        allocation: this.defaultAllocation,
        disabled: false
      }
      this.$store.dispatch('consultants/createConsultant', newConsultant)
    },
    toggleConsultant (consultant) {
      this.$store.dispatch('consultants/toggleConsultant', parseInt(consultant.id, 10))
    },
    async deleteConsultant (consultant) {
      if (await this.$refs.confirm.open('Wait a sec, this cannot be undone!', 'Are you sure you want to delete the consultant <b>' + consultant.name + '</b><br/>and all relevant reported records?', 'cancel', { color: 'bg-negative' })) {
        this.$store.dispatch('consultants/removeConsultant', { id: consultant.id, name: consultant.name })
        if (this.reportedHours.length === 0 && this.selectedConsultant !== '') {
          this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: this.selectedConsultant })
        }
      }
    },
    onUpdateAllocation (payload) {
      if (payload.newValue && parseInt(payload.newValue) > 0 && parseInt(payload.newValue) <= 100) {
        const newConsultant = {
          id: parseInt(payload.row.id),
          name: payload.row.name,
          allocation: parseInt(payload.newValue) / 100,
          disabled: payload.row.disabled
        }
        this.$store.dispatch('consultants/updateConsultant', newConsultant)
      }
    },
    onUpdateProject (payload) {
      const newProject = {
        id: parseInt(payload.row.id),
        name: payload.row.name,
        rate: payload.newValue.name,
        disabled: payload.row.disabled
      }
      this.$store.dispatch('projects/updateProject', newProject)
    },
    onUpdateHours (newValue) {
      if (!isNaN(parseFloat(newValue.hourValue)) && newValue.hourValue >= 0 && newValue.hourValue <= 24) {
        const value = {
          hourType: newValue.hourType,
          hourValue: parseFloat(newValue.hourValue)
        }
        this.$store.dispatch('settings/setHours', value)
      }
    },
    onUpdateDays (newValue) {
      if (parseInt(newValue.dayValue) === parseFloat(newValue.dayValue) && !isNaN(parseInt(newValue.dayValue)) && newValue.dayValue >= 0 && newValue.dayValue <= 40) {
        const value = {
          dayType: newValue.dayType,
          dayValue: parseInt(newValue.dayValue)
        }
        this.$store.dispatch('settings/setDays', value)
      }
    },
    onUpdateRate (newValue) {
      this.$store.dispatch('settings/setRate', newValue)
    },
    onUpdateRateType (newValue) {
      this.$store.dispatch('settings/setRateType', newValue)
    }
  }
}
</script>

<style scoped>
.logs {
  padding-top: 0px !important;
  padding-left: 0px !important;
}
</style>
