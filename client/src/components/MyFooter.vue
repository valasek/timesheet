<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-footer elevated>
    <q-toolbar class="glossy">
      2018 - {{ (new Date()).getFullYear() }} &copy;&nbsp;Stanislav Valasek <q-space/>Version: {{ version === "" ? "dev" : version }}, running on {{ url }}:{{ port }}
    </q-toolbar>
  </q-footer>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      url: process.env.ENV_APP_URL,
      port: process.env.ENV_APP_PORT
    }
  },

  computed: {
    ...mapState({
      version: state => state.settings.version
    })
  }

}
</script>

<style>
</style>
