<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-footer elevated>
    <q-toolbar class="glossy">
      2018 - {{ (new Date()).getFullYear() }} &copy;&nbsp;Stanislav Valasek <q-space/>Version: {{ version }}, running on {{ url }}
    </q-toolbar>
  </q-footer>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      url: process.env.APP_URL
    }
  },

  computed: {
    ...mapState({
      version: state => state.settings.version
    })
  },

  methods: {
  }

}
</script>

<style>
</style>
