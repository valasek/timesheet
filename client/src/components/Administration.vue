<template>
  <div>
    <v-container fluid>
      <v-layout align-start justify-start row fill-height>
        <v-flex xs4>
          <div class="headline">
            Warning limits
          </div>
        </v-flex>
        <v-flex xs2 />
        <v-flex xs6>
          <div class="headline">
            Backup
          </div>
        </v-flex>
      </v-layout>
      <v-container fluid>
        <v-layout align-start justify-start row fill-height>
          <v-flex xs4>
            <v-layout align-start justify-start row fill-height>
              <v-flex xs6>
                <v-text-field :value="dailyWorkingHoursMin" label="Minimum working hours"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 60%" @input="setWorkingHoursMin"
                />
              </v-flex>
              <v-flex xs6>
                <v-text-field :value="dailyWorkingHoursMax" label="Maximum working hours"
                              type="number" min="0" max="24" step="0.5" maxlength="2"
                              class="body-1" style="width: 60%" @input="setWorkingHoursMax"
                />
              </v-flex>
            </v-layout>
          </v-flex>
          <v-flex xs2 />
          <v-flex xs6>
            <v-btn small @click="download">
              Download
            </v-btn>
            Download all reported data as a zip file<br>
          </v-flex>
        </v-layout>
      </v-container>
      <v-container id="scroll-target" fluid>
        <div class="headline">
          Application logs
        </div>
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
    </v-container>
  </div>
</template>

<script>
  import api from '../api/axiosSettings'

  export default {

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
        logLines: ['select desired log level ...']
      }
    },

    computed: {
      dailyWorkingHoursMax: {
        set (hours) {
          this.$store.dispatch('settings/setDailyWorkingHoursMax', hours)
        },
        get () {
          return this.$store.state.settings.dailyWorkingHoursMax
        }
      },
      dailyWorkingHoursMin: {
        set (hours) {
          this.$store.dispatch('settings/setDailyWorkingHoursMin', hours)
        },
        get () {
          return this.$store.state.settings.dailyWorkingHoursMin
        }
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Administration')
    },

    methods: {
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
      },
      getLogFile (logLevel) {
        this.logs = 'calling ...'
        let adminThis = this
        api.apiClient.get('/api/download/logs/' + logLevel)
          .then(response => {
            this.logLines = response.data.split(new RegExp('\r?\n', 'g')) /* eslint-disable-line no-control-regex */
          })
          .catch(function (e) {
            console.log('getLogFile failed', e) /* eslint-disable-line no-console */
            adminThis.$store.dispatch('context/setNotification', { text: 'Cannot retrieve log file: ' + e, type: 'error' })
          })
      },
      setWorkingHoursMin (hours) {
        this.$store.dispatch('settings/setDailyWorkingHoursMin', parseFloat(hours))
      },
      setWorkingHoursMax (hours) {
        this.$store.dispatch('settings/setDailyWorkingHoursMax', parseFloat(hours))
      }
    }
  }
</script>

<style scoped>
.logs {
  padding: 10px;
}
</style>
