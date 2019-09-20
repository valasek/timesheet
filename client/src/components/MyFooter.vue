<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-footer elevated class="bg-primary text-secondary">
    <q-toolbar>
      2018 - {{ (new Date()).getFullYear() }} &copy;&nbsp;Stanislav Valasek <q-space />Version: {{ version === "" ? "dev" : version }} | Server: {{ domain }}:{{ port }}
    </q-toolbar>
  </q-footer>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      domain: '',
      port: ''
    }
  },

  computed: {
    ...mapState({
      version: state => state.settings.version
    })
  },

  created () {
    this.domain = process.env.APP_DOMAIN
    this.port = process.env.APP_PORT
  }

}
</script>

<style>
</style>
