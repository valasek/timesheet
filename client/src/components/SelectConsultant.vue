<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <span>
    <v-autocomplete v-model="selectedConsultant" prepend-icon="person_outline" :dense="true" :items="consultants.all" item-text="name" item-value="name" class="body-1" />
  </span>
</template>

<script>
  import { mapState } from 'vuex'

  export default {

    computed: {
      ...mapState({
        consultants: state => state.consultants,
        selectedMonth: state => state.settings.selectedMonth
      }),
      selectedConsultant: {
        set (newValue) {
          this.$store.dispatch('consultants/setSelected', newValue)
          this.$store.dispatch('reportedHours/getMonthlyData', { date: this.selectedMonth, consultant: newValue })
        },
        get () {
          return this.$store.state.consultants.selected
        }
      }
    }

  }
</script>

<style>
</style>
