<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <v-expansion-panel>
      <v-expansion-panel-content>
        <div slot="header">
          Warning Limits
        </div>
        <v-card>
          <v-card-text>
            <div class="font-italic">
              Used on Reported Overview page to show weekly and monthly expected working hours
            </div>
          </v-card-text>
          <v-container fluid>
            <v-layout align-start justify-start row fill-height>
              <v-flex xs5>
                <v-text-field :value="dailyWorkingHoursMin" label="Minimum working hours"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 15em" @input="onUpdateHours({hourType: 'dailyWorkingHoursMin', hourValue: $event})"
                />
              </v-flex>
              <v-flex xs5>
                <v-text-field :value="dailyWorkingHoursMax" label="Maximum working hours"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 15em" @input="onUpdateHours({hourType: 'dailyWorkingHoursMax', hourValue: $event})"
                />
              </v-flex>
              <v-flex xs2 />
            </v-layout>
          </v-container>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content>
        <div slot="header">
          Vacation Settings
        </div>
        <v-card>
          <v-card-text>
            <div class="font-italic">
              Used on Reported Overview page to show weekly and monthly expected working hours
            </div>
          </v-card-text>
          <v-container fluid>
            <v-layout align-start justify-start row fill-height text-xs-center>
              <v-flex xs4>
                Vacations
              </v-flex>
              <v-flex xs4>
                Personal Days
              </v-flex>
              <v-flex xs4>
                Sick Days
              </v-flex>
            </v-layout>
            <v-layout align-start justify-start row fill-height>
              <v-flex xs4>
                <v-text-field :value="yearlyVacationDays" label="Vacation days per year"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlyVacationDays', dayValue: $event})"
                />
              </v-flex>
              <v-flex xs4>
                <v-text-field :value="yearlyPersonalDays" label="Additional vacation days per year"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlyPersonalDays', dayValue: $event})"
                />
              </v-flex>
              <v-flex xs4>
                <v-text-field :value="yearlySickDays" label="Additional vacation days per year"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 15em" @input="onUpdateDays({dayType: 'yearlySickDays', dayValue: $event})"
                />
              </v-flex>
            </v-layout>
            <v-layout align-start justify-start row fill-height>
              <v-flex xs4>
                <v-autocomplete :value="vacation" item-text="name" item-value="name" style="width: 15em"
                                :items="rates" :dense="true" :hide-selected="false" class="body-1"
                                @change="onUpdateRate({rateType: 'vacation', rateValue: $event})"
                />
              </v-flex>
              <v-flex xs4>
                <v-autocomplete :value="vacationPersonal" item-text="name" item-value="name" style="width: 15em"
                                :items="rates" :dense="true" :hide-selected="false" class="body-1"
                                @change="onUpdateRate({rateType: 'vacationPersonal', rateValue: $event})"
                />
              </v-flex>
              <v-flex xs4>
                <v-autocomplete :value="vacationSick" item-text="name" item-value="name" style="width: 15em"
                                :items="rates" :dense="true" :hide-selected="false" class="body-1"
                                @change="onUpdateRate({rateType: 'vacationSick', rateValue: $event})"
                />
              </v-flex>
            </v-layout>
          </v-container>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content>
        <div slot="header">
          Application Settings
        </div>
        <v-card>
          <v-card-text>
            <p class="font-italic">
              Used for weekly and monthly expected working hours
            </p>
            <v-text-field :value="dailyWorkingHours" label="Daily working hours"
                          type="number" min="0" max="24" step="0.5" maxlength="2"
                          class="body-1" style="width: 10em" @input="onUpdateHours({hourType: 'dailyWorkingHours', hourValue: $event})"
            />
            <p class="font-italic">
              When changed all dates will be save using this timezone. Allready saved dates will not be updated.
            </p>
            <v-autocomplete :value="timeZone" style="width: 15em" label="Timezone"
                            :items="timeZones" :dense="true" :hide-selected="false" class="body-1"
                            @change="onUpdateTimeZone"
            />
            <p class="font-italic">
              Used on Reported Overview page to show working and non-working time
            </p>
            <v-autocomplete :value="isWorking" style="width: 15em" label="Working rate type"
                            :items="types" :dense="true" :hide-selected="false" class="body-1"
                            @change="onUpdateRateType({rateTypeType: 'isWorking', rateTypeValue: $event})"
            />
            <v-autocomplete :value="isNonWorking" style="width: 15em" label="Non-working rate type"
                            :items="types" :dense="true" :hide-selected="false" class="body-1"
                            @change="onUpdateRateType({rateTypeType: 'isNonWorking', rateTypeValue: $event})"
            />
          </v-card-text>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content>
        <div slot="header">
          Backup & Restore
        </div>
        <v-card>
          <v-container fluid>
            <v-layout align-start justify-start row fill-height>
              <v-flex xs6>
                <v-btn color="primary" @click="download">
                  Download&nbsp;&nbsp;<v-icon>cloud_download</v-icon>
                </v-btn>
                <p class="font-italic">
                  Download all data in csv format.
                </p>
              </v-flex>
              <v-flex xs6>
                <!-- eslint-disable-next-line vue/attribute-hyphenation -->
                <upload-btn color="primary" :fileChangedCallback="upload">
                  Upload
                  <template slot="icon">
                    &nbsp;&nbsp;<v-icon dark>
                      cloud_upload
                    </v-icon>
                  </template>
                </upload-btn>
                <div class="font-italic">
                  Upload zipped csv files. Use the same format as downloaded.<br>Existing data will be replaced.
                </div>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content>
        <div slot="header">
          View Application Logs
        </div>
        <v-card>
          <v-container id="scroll-target" fluid class="logs-select">
            <v-layout align-start justify-start row fill-height>
              <v-flex xs2>
                <v-select :items="logLevels" item-value="id" item-text="level" :dense="true" label="Select log level" class="body-1" @change="getLogFile" />
              </v-flex>
            </v-layout>
            <v-layout column align-left justify-top class="scroll-y elevation-5 logs" style="max-height: 600px">
              <div v-for="line in logLines" :key="line.id">
                {{ line }}
              </div>
            </v-layout>
          </v-container>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import timeZones from './timeZones'
  import api from '../api/axiosSettings'
  // import axios from 'axios'
  import UploadButton from 'vuetify-upload-button'

  export default {

    components: {
      'upload-btn': UploadButton
    },

    data () {
      return {
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
        logLines: ['select desired log level ...'],
        timeZones: timeZones,
        uploadFile: {}
      }
    },

    computed: {
      ...mapState({
        timeZone: state => state.settings.timeZone,
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
    },

    methods: {
      upload (file) {
        const fr = new FileReader()
        fr.readAsDataURL(file)
        fr.addEventListener('load', () => {
          this.uploadedImage = fr.result
          this.uploadFile = file
          const admin = this
          const formData = new FormData()
          formData.append('uploadFile', this.uploadFile)
          api.apiClient.post('/api/upload/data', formData)
            .then(response => {
              this.$store.dispatch('context/setNotification', { text: 'Restore successfully completed', type: 'success' })
              this.$store.dispatch('consultants/getConsultants')
              this.$store.dispatch('projects/getProjects')
              this.$store.dispatch('rates/getRates')
              this.$store.dispatch('holidays/getHolidays')
            })
            .catch(function (e) {
              admin.$store.dispatch('context/setNotification', { text: 'Couldn\'t upload file: ' + e.response.data, type: 'error' }, { root: true })
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
            admin.$store.dispatch('context/setNotification', { text: 'Couldn\'t download data\n' + e.toString(), type: 'error' }, { root: true })
            console.log(e) /* eslint-disable-line no-console */
          })
      },
      getLogFile (logLevel) {
        const admin = this
        api.apiClient.get('/api/download/logs/' + logLevel)
          .then(response => {
            admin.logLines = response.data.split(new RegExp('\r?\n', 'g')) /* eslint-disable-line no-control-regex */
          })
      },
      onUpdateHours (newValue) {
        const value = {
          hourType: newValue.hourType,
          hourValue: parseFloat(newValue.hourValue)
        }
        this.$store.dispatch('settings/setHours', value)
      },
      onUpdateDays (newValue) {
        const value = {
          dayType: newValue.dayType,
          dayValue: parseInt(newValue.dayValue)
        }
        this.$store.dispatch('settings/setDays', value)
      },
      onUpdateRate (newValue) {
        this.$store.dispatch('settings/setRate', newValue)
      },
      onUpdateRateType (newValue) {
        this.$store.dispatch('settings/setRateType', newValue)
      },
      onUpdateTimeZone (newValue) {
        this.$store.dispatch('settings/setTimeZone', newValue)
      }
    }
  }
</script>

<style scoped>
.logs {
  padding: 10px;
}
.logs-select {
  padding-top: 0px;
  margin-top: 0px;
}
</style>
