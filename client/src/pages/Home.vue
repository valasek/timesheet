<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <div class="q-pa-md row items-start q-gutter-md">
      <q-card
        class="my-card text-white"
        style="background: radial-gradient(circle, #35a2ff 0%, #014a88 100%)">
        <q-card-section>
          <div class="text-h6">Self-hosted web application for weekly reporting</div>
        </q-card-section>
        <q-card-section>
          Report your consulting hours on projects with further segmentation via rates.<br/>
          You can download all data in csv format.<br/>
          Data can be modified, production moved to another instance or exported data used for billing. You can upload modified files.
        </q-card-section>
      </q-card>
      <q-card
        class="my-card text-white"
        style="background: radial-gradient(circle, #35a2ff 0%, #014a88 100%)">
        <q-card-section>
          <q-btn color="primary" label="Report" icon="work_outline" />
        </q-card-section>
      </q-card>
    </div>
    <div class="q-pa-md row items-start q-gutter-md">
      <q-card
        class="my-card text-white"
        style="background: radial-gradient(circle, #35a2ff 0%, #014a88 100%)">
        <q-card-section>
          # of consultants
          # of projects
        </q-card-section>
      </q-card>
      Request New Feature

      Do you like what you see but mising an important feature? Describe your need and next business day we will let you know if and when it will be available.
      Report a Bug

      Did you find an issue? Report it and next business day we will let you know when the fix will be available.
    </div>
  </q-page>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      url: process.env.VUE_APP_URL,
      port: process.env.VUE_APP_PORT
    }
  },

  computed: {
    ...mapState({
      version: state => state.settings.version
    })
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Timesheet')
  }
}
</script>

<style scoped>
.image {
  border-radius: 10px;
}
</style>
