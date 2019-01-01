<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>State holidays</v-toolbar-title>
    </v-toolbar>
    <v-container grid-list-md text-xs-center>
      <v-layout row wrap>
        <v-flex xs12>
          <v-card>
            <v-data-table :headers="headers" :items="holidays" hide-actions class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ formatDate(props.item.date) }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.description }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: 'Holidays',

    data () {
      return {
        headers: [
          { text: 'Date', align: 'left', value: 'date', sortable: false },
          { text: 'Description', align: 'left', value: 'description', sortable: false }
        ]
      }
    },

    computed: {
      ...mapState({
        holidays: state => state.holidays.all
      })
    },

    created () {
    },

    methods: {
      formatDate (date) {
        if (!date) return null
        let months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
        let weekdays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
        const [year, , day] = date.slice(0, 10).split('-')
        let monthName = months[new Date(date).getMonth()]
        let dayName = weekdays[new Date(date).getDay()]
        return `${dayName}, ${monthName} ${day}, ${year}`
      }
    }
  }
</script>

<style scoped>
</style>
