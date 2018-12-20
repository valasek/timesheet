<template>
  <div>
    <v-toolbar flat>
      <v-toolbar-title>Year</v-toolbar-title>
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
              <v-toolbar-title>Month</v-toolbar-title>
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
              <v-toolbar-title>Week</v-toolbar-title>
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
      selectedReportedHours () {
        return this.reportedHours.filter(report => {
          let d = new Date(report.date)
          return (d >= this.dateFrom && d <= this.dateTo && report.consultant === this.selectedConsultants)
        })
      },
      monthlyConsultantReportedHours () {
        return this.reportedHours.filter(report => {
          return (report.consultant === this.selectedConsultants)
        })
      },
      ...mapState({
        reportedHours: state => state.reportedHours.all,
        dateFrom: state => state.reportedHours.dateFrom,
        dateTo: state => state.reportedHours.dateTo,
        selectedConsultants: state => state.consultants.selected
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
      weeklyOverview () {
        return [
          {
            text: 'Available working hours',
            value: 18 * 8
          },
          {
            text: 'Reported working hours',
            value: this.getTotals(this.selectedReportedHours, 'Off-site') + this.getTotals(this.selectedReportedHours, 'On-site')
          },
          {
            text: 'Reported working hours during holiday',
            value: this.getTotals(this.selectedReportedHours, 'Off-site Holiday')
          },
          {
            text: 'Personal days',
            value: this.getTotals(this.selectedReportedHours, 'Personal Day')
          },
          {
            text: 'Sick days',
            value: this.getTotals(this.selectedReportedHours, 'Sick Day')
          },
          {
            text: 'Vacations',
            value: this.getTotals(this.selectedReportedHours, 'Vacation')
          },
          {
            text: 'Unpaid leave',
            value: this.getTotals(this.selectedReportedHours, 'Unpaid Leave')
          },
          {
            text: 'Sick',
            value: this.getTotals(this.selectedReportedHours, 'Sick')
          },
          {
            text: 'Weekend',
            value: this.getTotals(this.selectedReportedHours, 'Off-site Weekend')
          },
          {
            text: 'Holiday',
            value: this.getTotals(this.selectedReportedHours, 'Public Holiday')
          },
          {
            text: 'Overtime',
            value: 999
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
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site') + this.getTotals(this.monthlyConsultantReportedHours, 'On-site')
          },
          {
            text: 'Reported working hours during holiday',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site Holiday')
          },
          {
            text: 'Personal days',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Personal Day')
          },
          {
            text: 'Sick days',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Sick Day')
          },
          {
            text: 'Vacations',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Vacation')
          },
          {
            text: 'Unpaid leave',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Unpaid Leave')
          },
          {
            text: 'Sick',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Sick')
          },
          {
            text: 'Weekend',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Off-site Weekend')
          },
          {
            text: 'Holiday',
            value: this.getTotals(this.monthlyConsultantReportedHours, 'Public Holiday')
          },
          {
            text: 'Overtime',
            value: 999
          }
        ]
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Overview of reported hours')
    },

    methods: {
      getTotals (hours, rate) {
        let h = hours.reduce(
          function (total, current) {
            let h = 0
            if (current.rate === rate) { h = current.hours }
            return total + h
          }, 0)
        return h || ''
      }
    }
  }
</script>

<style scoped>
</style>
