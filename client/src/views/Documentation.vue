<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <v-expansion-panel>
      <v-expansion-panel-content>
        <div slot="header">
          <strong>Table of Content</strong>
        </div>
        <v-card>
          <v-card-text>
            <v-list id="toc" :dense="true" />
          </v-card-text>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <br>
    <vue-markdown :source="docs" :toc="true" toc-id="toc" :toc-first-level="1"
                  :toc-anchor-link="false" toc-anchor-link-symbol="" toc-class="table"
    />
  </div>
</template>

<script>
  import VueMarkdown from 'vue-markdown'
  import api from '../api/axiosSettings'

  export default {

    components: {
      VueMarkdown
    },

    data () {
      return {
        docs: 'Loading documentation ...'
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
            comp.docs = response.data
          })
          .catch(function (e) {
            comp.$store.dispatch('context/setNotification', { text: 'Couldn\'t find documentation: ' + e.response.data, type: 'error' }, { root: true })
            console.log(e, e.response) /* eslint-disable-line no-console */
          })
      }
    }
  }
</script>

<style scoped>
</style>
