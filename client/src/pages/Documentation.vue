<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-markdown :src="markdown" />
  </q-page>
</template>

<script>
import api from '../api/axiosSettings'

export default {

  data () {
    return {
      markdown: 'Loading documentation ...'
    }
  },

  computed: {
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Documentation')
    this.getDocumentation()
  },

  methods: {
    getDocumentation () {
      const comp = this
      api.apiClient.get('/api/download/docs')
        .then(response => {
          comp.markdown = response.data
        })
        .catch(function (e) {
          this.$q.notify({
            message: 'Couldn\'t find documentation: ' + e.response.data,
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e, e.response) /* eslint-disable-line no-console */
        })
    }
  }
}
</script>

<style scoped>
</style>
