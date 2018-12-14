<template>
  <div />
</template>

<script>
  export default {
    name: 'Help',
    data () {
      return {
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Help')
    }
  }
</script>

<style scoped>
</style>
