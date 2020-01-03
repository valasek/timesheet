<!-- Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-table
    :columns='columnsEntityOverview'
    row-key='name'
    :data='entityOverview'
    no-data-label='Table statistics are not available'
    hide-bottom
    bordered
    :pagination.sync='entityOverviewPagination'
  />
</template>

<script>
import { mapState } from 'vuex'

export default {
  data () {
    return {
      columnsEntityOverview: [
        {
          name: 'name',
          label: 'Table',
          align: 'left',
          sortable: true,
          field: 'name',
          style: 'width: 20%'
        },
        {
          name: 'total',
          label: '# of total records',
          align: 'left',
          sortable: true,
          field: 'total',
          style: 'width: 5%'
        },
        {
          name: 'active',
          label: '# of active records',
          align: 'left',
          field: 'active',
          style: 'width: 5%'
        },
        {
          name: 'disabled',
          label: '# of disabled out of active records',
          align: 'left',
          field: 'disabled',
          style: 'width: 5%'
        },
        {
          name: 'deleted',
          label: '# of soft-deleted records',
          align: 'left',
          field: 'deleted',
          style: 'width: 5%'
        }
      ],
      entityOverviewPagination: {
        rowsPerPage: 0,
        sortBy: 'total',
        descending: true
      }
    }
  },

  created () {
    this.$store.dispatch('settings/getEntityOverview')
  },

  computed: {
    ...mapState({
      entityOverview: state => state.settings.entityOverview
    })
  },
  methods: {}
}
</script>

<style>
</style>
