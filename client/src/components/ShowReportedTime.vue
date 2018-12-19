<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>Yearly overview</v-toolbar-title>
    </v-toolbar>
    <v-container grid-list-md text-xs-center>
      <v-layout row wrap>
        <v-flex xs4>
          <v-card>
            <v-data-table :items="vacations" hide-actions hide-headers class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
        <v-flex xs4>
          <v-card>
            <v-data-table :items="personalDays" hide-actions hide-headers class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
        <v-flex xs4>
          <v-card>
            <v-data-table :items="sickDays" hide-actions hide-headers class="elevation-1">
              <template slot="items" slot-scope="props">
                <td class="text-xs-left">
                  {{ props.item.text }}
                </td>
                <td class="text-xs-left">
                  {{ props.item.value }}
                </td>
              </template>
            </v-data-table>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-container grid-list-md text-xs-center>
      <v-layout row wrap>
        <v-flex xs6>
          <v-card>
            <v-toolbar flat>
              <v-toolbar-title>Monthly overview</v-toolbar-title>
            </v-toolbar>
          </v-card>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :items="monthlyOverview" hide-actions hide-headers class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
        <v-flex xs6>
          <v-card>
            <v-toolbar flat>
              <v-toolbar-title>Weekly overview</v-toolbar-title>
            </v-toolbar>
          </v-card>
          <v-container grid-list-md text-xs-center>
            <v-card>
              <v-data-table :items="weeklyOverview" hide-actions hide-headers class="elevation-1">
                <template slot="items" slot-scope="props">
                  <td class="text-xs-left">
                    {{ props.item.text }}
                  </td>
                  <td class="text-xs-left">
                    {{ props.item.value }}
                  </td>
                </template>
              </v-data-table>
            </v-card>
          </v-container>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: 'ShowReportedTime',

    data () {
      return {
      }
    },

    computed: {
      ...mapState({
        reportedHours: state => state.reportedHours.all
      }),
      vacations () {
        return [
          {
            text: 'Remaining vacations',
            value: 25
          },
          {
            text: 'Reported vacations',
            value: 10
          }
        ]
      },
      personalDays () {
        return [
          {
            text: 'Remaining personal days',
            value: 15
          },
          {
            text: 'Reported personal days',
            value: 17
          }
        ]
      },
      sickDays () {
        return [
          {
            text: 'Remaining sick days',
            value: 10
          },
          {
            text: 'Reported sick days',
            value: 7
          }
        ]
      },
      monthlyOverview () {
        return [
          {
            text: 'Available working hours',
            value: 18 * 8
          },
          {
            text: 'Reported working hours',
            value: this.getTotals(this.reportedHours, 'Off-site') + this.getTotals(this.reportedHours, 'On-site')
          },
          {
            text: 'Reported working hours during holiday',
            value: this.getTotals(this.reportedHours, 'Off-site Holiday')
          },
          {
            text: 'Personal days',
            value: this.getTotals(this.reportedHours, 'Personal Day')
          },
          {
            text: 'Sick days',
            value: this.getTotals(this.reportedHours, 'Sick Day')
          },
          {
            text: 'Vacations',
            value: this.getTotals(this.reportedHours, 'Vacation')
          },
          {
            text: 'Unpaid leave',
            value: this.getTotals(this.reportedHours, 'Unpaid Leave')
          },
          {
            text: 'Sick',
            value: this.getTotals(this.reportedHours, 'Sick')
          },
          {
            text: 'Weekend',
            value: this.getTotals(this.reportedHours, 'Off-site Weekend')
          },
          {
            text: 'Holiday',
            value: this.getTotals(this.reportedHours, 'Public Holiday')
          },
          {
            text: 'Overtime',
            value: 999
          }
        ]
      },
      weeklyOverview () {
        return [
          {
            text: 'Working hours',
            value: 40
          },
          {
            text: 'Worked',
            value: 39
          },
          {
            text: 'Personal days',
            value: 38
          },
          {
            text: 'Sick days',
            value: 37
          },
          {
            text: 'Vacations',
            value: 36
          },
          {
            text: 'Unpaid leave',
            value: 36
          },
          {
            text: 'Sick',
            value: 36
          },
          {
            text: 'Weekend',
            value: 36
          },
          {
            text: 'Holiday',
            value: 36
          },
          {
            text: 'Overtime',
            value: 36
          }
        ]
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Reported work')
    },

    methods: {
      getTotals (hours, rate) {
        return hours.reduce(
          function (total, current) {
            let h = 0
            if (current.rate === rate) { h = current.hours }
            return total + h
          }, 0)
      }
    }
  }
</script>

<style scoped>
</style>
