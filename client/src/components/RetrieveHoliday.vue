<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <span>
    Mv
    <v-btn>Update State Holidays</v-btn>
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
        year: '2019'
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
        const rr = this
        // curl https://holidayapi.pl/v1/holidays -XGET -d"country=CZ&year=2019"
        api.apiClient.get(url + querystring)
          .then(response => {
          })
          .catch(function (e) {
            rr.$store.dispatch('context/setNotification', { text: 'Couldn\'t upload file: ' + e.response.data, type: 'error' }, { root: true })
            console.log(e, e.response) /* eslint-disable-line no-console */
          })
        // this.$store.dispatch('holidays/getMonthlyData', { date: this.selectedMonth, consultant: newValue })
      }
    }

  }
</script>

<style>
</style>
