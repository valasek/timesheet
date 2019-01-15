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
                  {{ formatDay(props.item.date) }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.description }}
                </td>
                <td class="text-xs-left">
                  <v-tooltip bottom>
                    <v-icon slot="activator" small class="mr-2" @click="addHolidayToReport(props.item)">
                      add
                    </v-icon>
                    <span>Add to my reported records</span>
                  </v-tooltip>
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
  import moment from 'moment'

  export default {
    name: 'Holidays',

    data () {
      return {
        headers: [
          { text: 'Date', align: 'left', value: 'date', sortable: true, class: 'body-1' },
          { text: 'Day', align: 'left', value: 'day', sortable: true, class: 'body-1' },
          { text: 'Description', align: 'left', value: 'description', sortable: false, class: 'body-1' },
          { text: 'Actions', align: 'left', value: 'action', sortable: false, class: 'body-1' }
        ]
      }
    },

    computed: {
      ...mapState({
        holidays: state => state.holidays.all,
        dailyWorkingHours: state => state.context.dailyWorkingHours,
        consultant: state => state.consultants.selected
      })
    },

    created () {
    },

    methods: {
      addHolidayToReport (item) {
        let newRecord = {
          id: null,
          date: moment(item.date).format('YYYY-MM-DDTHH:mm:ssZ'),
          hours: this.dailyWorkingHours,
          project: 'Holiday',
          description: item.description,
          rate: '_Holiday',
          consultant: this.consultant
        }
        this.$store.dispatch('reportedHours/addRecord', newRecord)
        this.$store.dispatch('context/setNotification', { text: this.formatDate(item.date) + ' added to your reports', type: 'success' })
      },
      formatDate (date) {
        if (!date) return null
        return moment(date).format('MMM Do, YYYY')
      },
      formatDay (date) {
        return moment(date).format('dddd')
      }
    }
  }
</script>

<style scoped>
</style>
