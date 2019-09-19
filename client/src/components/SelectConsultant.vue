<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
<!-- <q-select v-model="selectedConsultant" :options="consultants.all" option-name="id" option-label="name"
    @filter="filterFn" dense options-dense > -->
  <q-select v-model="selectedConsultant" :options="filteredConsultants" option-name="id" option-label="name" option-disable="disabled"
    @filter="filterFn" dense options-dense use-input hide-selected fill-input input-debounce="0"
    color="secondary"
  >
    <template v-slot:prepend>
      <q-icon name="person_outline" />
    </template>
    <template v-slot:no-option>
      <q-item>
        <q-item-section class="text-grey">
          No results
        </q-item-section>
      </q-item>
    </template>
  </q-select>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      filteredConsultants: this.consultants
    }
  },

  computed: {
    ...mapState({
      consultants: state => state.consultants,
      selectedMonth: state => state.settings.selectedMonth
    }),
    selectedConsultant: {
      set (newValue) {
        this.$store.dispatch('consultants/setSelected', newValue.name)
        this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: newValue.name })
      },
      get () {
        return this.$store.state.consultants.selected
      }
    }
  },
  methods: {
    filterFn (val, update, abort) {
      update(() => {
        const needle = val.toLowerCase()
        this.filteredConsultants = this.consultants.all.filter(v => v.name.toLowerCase().indexOf(needle) > -1)
      })
    }
  }

}
</script>

<style>
</style>
