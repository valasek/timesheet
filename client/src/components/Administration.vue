<template>
  <div>
    <div class="headline">
      Backup
    </div>
    <p>
      Download all reported data to my computer as a zip file
      <v-btn small @click="download">
        Download
      </v-btn>
    </p>
    <v-divider />
    <p />
    <div class="headline">
      Application logs
    </div>
    <v-container id="scroll-target" fluid>
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
  </div>
</template>

<script>
  import api from '../api/axiosSettings'

  export default {
    name: 'Administration',
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
        let adm = this
        api.apiClient.get('/api/download/logs/' + logLevel)
          .then(response => {
            this.logLines = response.data.split(new RegExp('\r?\n', 'g')) /* eslint-disable-line no-control-regex */
          })
          .catch(function (e) {
            console.log('getLogFile failed', e) /* eslint-disable-line no-console */
            adm.$store.dispatch('context/setNotification', { text: 'Cannot retieve log file: ' + e, type: 'error' })
          })
      }
    }
  }
</script>

<style scoped>
.logs {
  padding: 10px;
}
</style>
