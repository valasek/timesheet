<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <span>
    <q-btn :disabled="isCurrentWeek" rounded class="bg-secondary text-primary" @click="setToday">
      today
    </q-btn>
    <q-btn flat icon="skip_previous" @click="previousWeek" />
    <q-btn flat icon="skip_next" @click="nextWeek" />
    <span>
      {{ formatWeek(dateFrom) }} - {{ formatWeek(dateTo) }}
    </span>
  </span>
</template>

<script>
import { mapState } from 'vuex'
import { format } from 'date-fns'

export default {

  computed: {
    ...mapState({
      isCurrentWeek: state => state.context.isCurrentWeek,
      dateFrom: state => state.settings.dateFrom,
      dateTo: state => state.settings.dateTo
    })
  },

  methods: {
    previousWeek () {
      this.$store.dispatch('settings/changeWeek', 'previous')
    },
    nextWeek () {
      this.$store.dispatch('settings/changeWeek', 'next')
    },
    // FIXME move to filter
    formatWeek (date) {
      return format(date, 'MMM d')
    },
    setToday () {
      this.$store.dispatch('settings/jumpToWeek', new Date())
    }
  }

}
</script>

<style>
</style>
