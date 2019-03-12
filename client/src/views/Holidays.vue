<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <v-toolbar dense>
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
  import { parseISO, format } from 'date-fns'

  export default {

    data () {
      return {
        headers: [
          { text: 'Date', align: 'left', value: 'date', sortable: true, class: 'body-1' },
          { text: 'Day', align: 'left', value: 'day', sortable: true, class: 'body-1' },
          { text: 'Description', align: 'left', value: 'description', sortable: false, class: 'body-1' }
        ]
      }
    },

    computed: {
      ...mapState({
        holidays: state => state.holidays.all
      })
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'State holiday')
    },

    methods: {
      formatDate (date) {
        if (!date) return null
        return format(parseISO(date), 'MMM do, yyyy')
      },
      formatDay (date) {
        if (!date) return null
        return format(parseISO(date), 'iiii')
      }
    }
  }
</script>

<style scoped>
</style>
