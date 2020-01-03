<!-- Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-toolbar class="q-pa-md bg-primary text-secondary">
      <q-toolbar-title>State Holidays</q-toolbar-title>
    </q-toolbar>
    <q-layout row wrap class="q-pa-y-md">
      <q-table
        :data="holidays"
        :columns="columns"
        row-key="name"
        hide-bottom
        :filter="filter"
        :pagination="pagination"
      >
        <template v-slot:item="props">
          <div v-for="col in props" :key="col.id">
            <td>
              {{ col.date }}
            </td>
            <td>
              {{ col.day }}
            </td>
            <td>
              {{ col.description }}
            </td>
          </div>
        </template>
      </q-table>
    </q-layout>
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import { parseISO, format } from 'date-fns'

export default {

  data () {
    return {
      filter: '',
      columns: [
        {
          name: 'date',
          required: true,
          label: 'Date',
          align: 'left',
          field: 'date',
          format: val => this.formatDate(`${val}`),
          sortable: true
        },
        { name: 'day', align: 'center', label: 'Day', field: 'date', format: val => this.formatDay(`${val}`), sortable: true },
        { name: 'description', label: 'Description', field: 'description', sortable: true }
      ],
      pagination: {
        rowsPerPage: 0
      }
    }
  },

  computed: {
    ...mapState({
      holidays: state => state.holidays.all
    })
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'State holiday')
    this.$store.commit('context/SET_PAGE_ICON', 'event')
  },

  methods: {
    formatDate (date) {
      if (!date) return null
      return format(parseISO(date), 'MMM d, yyyy')
    },
    formatDay (date) {
      if (!date) return null
      return format(parseISO(date), 'iiii')
    }
  }
}
</script>

<style lang="stylus" scoped>
.q-data-table tr {
    text-align:right;
}
</style>
