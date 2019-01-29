<template>
  <div>
    <v-btn small @click="download">
      Download all data
    </v-btn>
    <!-- <div class="headline">
      Application logs
    </div> -->
  </div>
</template>

<script>
  import api from '../api/axiosSettings'

  export default {
    name: 'Administration',
    data () {
      return {
        // logs: 'log file content\nsdf\nsdfa\nsdfa\nsdsadfsa\nsdfa\nsdfa\nsdfa\nsdfa\nssadfasddfa\nsdfa\nsdfa\nsdafa\nsdfa\nsdfa\nsdfa\nsdfa\nsdfa\nsdfa\nsdfa\nsdfa'
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Administration')
    },

    methods: {
      download () {
        api.apiClient.get('/api/download', { responseType: 'blob' })
          .then(response => {
            const url = window.URL.createObjectURL(new Blob([response.data]))
            const link = document.createElement('a')
            link.href = url
            link.setAttribute('download', 'timesheet-backup.zip')
            document.body.appendChild(link)
            link.click()
          })
      }
    }
  }
</script>

<style scoped>
</style>
