<!-- Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <span>
    Mv
    <q-btn>Update State Holidays</q-btn>
    holiday
  </span>
</template>

<script>
import { mapState } from 'vuex'
import api from '../api/axiosSettings'

export default {

  data () {
    return {
      country: 'CZ',
      supportedCoutries: [
        // BA Belgium
        // BR Brazil
        // CA Canada
        // CZ Czechia
        // DK Denmark (NEW!)
        // DE Germany

        // FR France
        // GB Great Britain
        // NO Norway
        // PL Poland
        // RU Russia (NEW!)
        // SK Slovakia

        // SL Sierra Leone
        // VN Vietnam (NEW!)
        // ID Indonesia
        // US United States
      ],
      year: '2020'
    }
  },

  computed: {
    ...mapState({
      consultants: state => state.consultants
    })
  },

  methods: {
    retrieveHolidays () {
      const url = 'https://holidayapi.pl/v1/holidays'
      const querystring = '?country=' + this.country + '&year=' + this.year
      // curl https://holidayapi.pl/v1/holidays -XGET -d"country=CZ&year=2020"
      api.apiClient.get(url + querystring)
        .then(response => {
        })
        .catch(function (e) {
          this.$q.notify({
            message: 'Couldn\'t upload file: ' + e.response.data,
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e, e.response) /* eslint-disable-line no-console */
        })
        // this.$store.dispatch('holidays/getMonthlyData', { date: this.selectedMonth, consultant: newValue })
    }
  }

}
</script>

<style>
</style>
