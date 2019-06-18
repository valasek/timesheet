<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-toolbar class="q-pa-md bg-grey-3">
      <q-toolbar-title>Administration & Settings</q-toolbar-title>
    </q-toolbar>
    <q-expansion-item
      expand-separator
      label="Warning Limits">
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
      label="Vacation Settings">
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
      label="Application Settings">
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
      label="Backup & Restore">
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
                  Download&nbsp;&nbsp;<q-icon name="cloud_download"></q-icon>
                </q-btn>
              </div>
              <div class="column">
                <q-uploader
                  label="Upload"
                  accept=".zip"
                  url="http://localhost:3000/api/upload/data"
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
      label="View Application Logs">
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
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import api from '../api/axiosSettings'

export default {

  components: {
  },

  data () {
    return {
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
      rates: state => state.rates.all,
      types: state => state.rates.types,
      selectedMonth: state => state.settings.selectedMonth
    })
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Administration')
    this.$store.commit('context/SET_PAGE_ICON', 'settings')
  },

  methods: {
    upload (file) {
      const fr = new FileReader()
      fr.readAsDataURL(file)
      fr.addEventListener('load', () => {
        this.uploadedImage = fr.result
        this.uploadFile = file
        const formData = new FormData()
        formData.append('uploadFile', this.uploadFile)
        api.apiClient.post('/api/upload/data', formData)
          .then(response => {
            this.$q.notify({
              message: 'Restore successfully completed',
              color: 'teal'
            })
            this.$store.dispatch('consultants/getConsultants')
            this.$store.dispatch('projects/getProjects')
            this.$store.dispatch('rates/getRates')
            this.$store.dispatch('holidays/getHolidays')
          })
          .catch(function (e) {
            this.$q.notify({
              message: 'Restore successfully completed',
              color: 'negative',
              icon: 'report_problem'
            })
            console.log(e, e.response) /* eslint-disable-line no-console */
          })
      })
    },
    download () {
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
          this.$q.notify({
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
          this.$q.notify({
            message: 'Couldn\'t download log file: ' + e.response.data,
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e, e.response) /* eslint-disable-line no-console */
        })
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
