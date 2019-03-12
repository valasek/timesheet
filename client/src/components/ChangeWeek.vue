<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <span>
    <v-btn :disabled="isCurrentWeek" class="mb-2" @click="setToday">
      today
    </v-btn>
    <v-btn icon @click="previousWeek">
      <v-icon>
        skip_previous
      </v-icon>
    </v-btn>
    <v-btn icon @click="nextWeek">
      <v-icon>
        skip_next
      </v-icon>
    </v-btn>
    <v-label class="text-xs-center">
      {{ formatWeek(dateFrom) }} - {{ formatWeek(dateTo) }}
    </v-label>
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
        return format(date, 'MMM do')
      },
      setToday () {
        this.$store.dispatch('settings/jumpToWeek', new Date())
      }
    }

  }
</script>

<style>
</style>
